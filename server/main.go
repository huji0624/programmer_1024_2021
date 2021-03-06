package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"server/util"
	"sort"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var magic_ids map[string]string
var magic_id_scores = make(map[string]string)

///使用过的 magic_ids
var used_magic_ids_set = make(map[string]bool)

var teams = make(map[string]string)
var diglock sync.Mutex

var gameRoundOver bool
var gameRoundOverTimer *time.Timer
var gameStartTime int

type SRcord struct {
	Team   string `json:"team"`
	Score  int    `json:"score"`
	Record string `json:"record"`
}

var scores_record []*SRcord

func CalFormula(formula string) (string, []string, string) {
	// log.Println(formula)

	tokens := make([]string, 0, 10)
	token := ""
	for _, c := range formula {
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
			token += string(c)
		} else {
			if token != "" {
				tokens = append(tokens, token)
				token = ""
			}
			if c == '+' {
				//ok
			} else if c == '-' {
				//ok
			} else if c == '*' {
				//ok
			} else if c == '/' {
				//ok
			} else if c == '(' {
				//ok
			} else if c == ')' {
				//ok
			} else {
				//ignore
				continue
			}
			tokens = append(tokens, string(c))
		}
	}

	if token != "" {
		tokens = append(tokens, token)
		token = ""
	}

	// log.Println(magic_ids)

	//tokens to nums
	ids := make([]string, 0, 10)
	readable := ""
	replaceed := make([]string, 0, 10)
	cur_used_ids := make([]string, 0, 10)
	for _, v := range tokens {
		if v != "+" && v != "-" && v != "*" && v != "/" && v != "(" && v != ")" {
			///检查 LocationId 是否使用过,	已经有元素被使用过则只记录被占用元素
			_, ok := used_magic_ids_set[v]
			if ok {
				cur_used_ids = append(cur_used_ids, v)
			}
			if len(cur_used_ids) > 0 {
				continue
			}

			///原有逻辑
			magic, ok := magic_ids[v]
			if ok {
				readable += magic
				replaceed = append(replaceed, magic)
				ids = append(ids, v)
			} else {
				log.Println("not:", v)
				return "", nil, ""
			}
		} else {
			readable += v
			replaceed = append(replaceed, v)
		}
	}

	///检查如果有元素被占用则直接返回
	if len(cur_used_ids) > 0 {
		return "repeat", cur_used_ids, ""
	}

	sort.Strings(ids)

	if len(ids) < 4 {
		return "", nil, ""
	}

	return CalTokens(replaceed), ids, readable
}

func ReverseTokens(tokens []string) []string {
	rTokens := make([]string, 0, 10)
	for i := len(tokens) - 1; i >= 0; i-- {
		rTokens = append(rTokens, tokens[i])
	}
	return rTokens
}

func CalTokens(tokens []string) string {
	// log.Println(tokens)
	if len(tokens) == 1 {
		return tokens[0]
	}

	if len(tokens) == 3 {
		return Caltwo(tokens[0], tokens[2], tokens[1])
	}

	//128+19*12-(888+(111+1*2+3)*99) = -12016
	nums := util.NewSatck()
	ops := util.NewSatck()
	for _, v := range tokens {
		if v == "*" {
			ops.Push(v)
		} else if v == "+" {

			if !ops.Empty() && (ops.Top() == "*" || ops.Top() == "/" || ops.Top() == "+" || ops.Top() == "-") {
				num2 := nums.Pop()
				num1 := nums.Pop()
				token := Caltwo(num1, num2, ops.Pop())
				nums.Push(token)
			}

			ops.Push(v)
		} else if v == "-" {
			if !ops.Empty() && (ops.Top() == "*" || ops.Top() == "/" || ops.Top() == "+" || ops.Top() == "-") {
				num2 := nums.Pop()
				num1 := nums.Pop()
				token := Caltwo(num1, num2, ops.Pop())
				nums.Push(token)
			}

			ops.Push(v)
		} else if v == "/" {
			ops.Push(v)
		} else if v == "(" {
			ops.Push(v)
		} else if v == ")" {
			newTokens := make([]string, 0, 10)
			for !ops.Empty() && ops.Top() != "(" {
				newTokens = append(newTokens, nums.Pop())
				newTokens = append(newTokens, ops.Pop())
			}
			newTokens = append(newTokens, nums.Pop())
			ops.Pop()

			token := CalTokens(ReverseTokens(newTokens))

			nums.Push(token)
		} else {
			if !ops.Empty() && (ops.Top() == "*" || ops.Top() == "/") {
				num2 := v
				num1 := nums.Pop()
				token := Caltwo(num1, num2, ops.Pop())
				nums.Push(token)
			} else {
				nums.Push(v)
			}
		}
	}

	if ops.Size() == 1 && nums.Size() == 2 {
		num2 := nums.Pop()
		num1 := nums.Pop()
		return Caltwo(num1, num2, ops.Pop())
	}

	if nums.Size() == 1 && ops.Empty() {
		return nums.Top()
	}

	newTokens := make([]string, 0, 10)
	for !nums.Empty() {
		newTokens = append(newTokens, nums.Pop())
		if !ops.Empty() {
			newTokens = append(newTokens, ops.Pop())
		}
	}

	if len(newTokens) == len(tokens) {
		return ""
	}

	return CalTokens(ReverseTokens(newTokens))
}

func Caltwo(num1 string, num2 string, op string) string {
	bi1 := big.NewInt(0)
	bi1.SetString(num1, 10)
	bi2 := big.NewInt(0)
	bi2.SetString(num2, 10)
	switch op {
	case "+":
		return bi1.Add(bi1, bi2).String()
	case "-":
		return bi1.Sub(bi1, bi2).String()
	case "*":
		return bi1.Mul(bi1, bi2).String()
	case "/":
		if bi2.String() == "0" {
			log.Println("div zero.wrong")
			return ""
		}
		return bi1.Div(bi1, bi2).String()
	}

	return ""
}

func main() {
	pid := os.Getpid()
	ioutil.WriteFile("./pid", []byte(fmt.Sprintf("%d", pid)), 0644)

	loadMagicIDS()

	loadTeams()

	scores_record = make([]*SRcord, 0, 0)

	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default() //返回默认引擎，里面有系统定义的中间件

	//跨域问题
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://192.168.1.66:9091"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.Static("/files", "./files")
	router.Static("/h5", "../h5/dist")

	router.POST("/dig", Dig)
	router.POST("/formula", Formula)

	router.GET("/reset", Reset)
	router.GET("/info", Info)

	router.Run(":80")

}

func loadTeams() {
	data, err := ioutil.ReadFile("./teams.json")
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1)
	}

	var teamsarray []string
	err = json.Unmarshal(data, &teamsarray)
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1)
	}

	log.Println("teams:")
	for _, v := range teamsarray {
		md5 := md5V3(v + "1024")
		teams[md5] = v

		fmt.Printf("%v : %v\n", v, md5)
	}
}

func md5V3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func loadMagicIDS() {
	data, err := ioutil.ReadFile("../match_control/magic_ids.json")
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1)
	}

	err = json.Unmarshal(data, &magic_ids)
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1)
	}
	for k, _ := range magic_ids {
		magic_id_scores[k] = ""
	}

	log.Println("loadMagicIDS:")
	log.Println(len(magic_id_scores))

	gameRoundOver = false
	gameRoundOverTimer = nil
	gameStartTime = -1
}

func ReturnError(c *gin.Context, apierr error) {
	c.JSON(http.StatusOK, gin.H{
		"errorno": -1,
		"msg":     apierr.Error(),
	})
}

func ReturnData(c *gin.Context, errorno int, data interface{}) {
	if data == nil {
		c.JSON(http.StatusOK, gin.H{
			"errorno": errorno,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    data,
			"errorno": errorno,
		})
	}
}

type ResetData struct {
	Code string `form:"code"`
}

func Reset(c *gin.Context) {
	var params ResetData

	c.ShouldBindQuery(&params)
	log.Println(params)

	if params.Code != "hiKLD.624" {
		ReturnError(c, errors.New("server error."))
		return
	}

	loadMagicIDS()
	scores_record = make([]*SRcord, 0, 0)
	//重置used_magic_ids_set
	used_magic_ids_set = make(map[string]bool)

	ReturnData(c, 0, nil)
}

type FormulaData struct {
	Token   string `json:"token"`
	Formula string `json:"formula"`
}

func Formula(c *gin.Context) {
	if gin.Mode() == gin.ReleaseMode {
		if gameRoundOver {
			ReturnError(c, errors.New("game round is over."))
			return
		}
	}

	var fd FormulaData
	err := c.BindJSON(&fd)

	if err != nil {
		ReturnError(c, err)
		return
	}

	if fd.Token == "" {
		ReturnError(c, errors.New("params token missing."))
		return
	}
	if teams[fd.Token] == "" {
		ReturnError(c, errors.New("token not valide."))
		return
	}
	if fd.Formula == "" {
		ReturnError(c, errors.New("params formula missing."))
		return
	}

	diglock.Lock()
	defer diglock.Unlock()

	formula := fd.Formula
	ret, ids, readable := CalFormula(formula)

	if gin.Mode() == gin.ReleaseMode {
		if ret == "1024" {
			///成功则计算,已经使用的 magic 对应的 locationId
			for _, id := range ids {
				used_magic_ids_set[id] = true
			}

			ReturnData(c, 0, nil)

			scores_record = append(scores_record, &SRcord{Team: teams[fd.Token], Score: len(ids) * len(ids), Record: readable})
			return
		} else if ret == "repeat" { //重复响应值
			ReturnData(c, 3, ids)
		} else {
			//formula wrong or id wrong
			ReturnData(c, 1, nil)
		}
	} else {
		//for test
		if ret == "1024" {
			///成功则计算,已经使用的 magic 对应的 locationId  测试服务暂不处理
			//for _, id := range ids {
			//	used_magic_ids_set[id] = true
			//}
			r := rand.Intn(100)
			if r < 2 {
				ReturnData(c, 3, ids)
				return
			} else {
				ReturnData(c, 0, nil)
				scores_record = append(scores_record, &SRcord{Team: teams[fd.Token], Score: len(ids) * len(ids), Record: readable})
				return
			}
		} else if ret == "repeat" { //重复响应值
			ReturnData(c, 3, ids)
		} else {
			ReturnData(c, 1, nil)
			return
		}
	}
}

type DigData struct {
	Token      string `json:"token"`
	Locationid string `json:"locationid"`
}

func Dig(c *gin.Context) {

	if gin.Mode() == gin.ReleaseMode {
		if gameRoundOver {
			ReturnError(c, errors.New("game round is over."))
			return
		}
	}

	var dd DigData
	err := c.BindJSON(&dd)

	if err != nil {
		ReturnError(c, err)
		return
	}

	if dd.Token == "" {
		ReturnError(c, errors.New("params token missing."))
		return
	}
	if teams[dd.Token] == "" {
		ReturnError(c, errors.New("token not valide."))
		return
	}
	if dd.Locationid == "" {
		ReturnError(c, errors.New("params Locationid missing."))
		return
	}

	diglock.Lock()
	defer diglock.Unlock()

	if gin.Mode() == gin.ReleaseMode {
		if gameRoundOverTimer == nil {
			gameRoundOverTimer = time.AfterFunc(time.Second*180, func() {
				gameRoundOver = true
				gameStartTime = -1
			})
			gameStartTime = time.Now().Second()
		}
	}

	t, ok := magic_id_scores[dd.Locationid]
	if !ok {
		//fail.not treasure
		ReturnData(c, 1, nil)
	} else {
		if gin.Mode() == gin.ReleaseMode {
			if t == "" {
				magic_id_scores[dd.Locationid] = teams[dd.Token]
				log.Printf("Team %v dig success.", dd.Token)
				//success
				ReturnData(c, 0, nil)

				scores_record = append(scores_record, &SRcord{Team: teams[dd.Token], Score: 1, Record: dd.Locationid})
			} else {
				//fail.alreay digged by others.
				ReturnData(c, 2, nil)
			}
		} else {
			r := rand.Intn(100)
			if r < 2 {
				ReturnData(c, 2, nil)
				return
			} else {
				log.Printf("Team %v dig success.", dd.Token)
				//success
				ReturnData(c, 0, nil)
				scores_record = append(scores_record, &SRcord{Team: teams[dd.Token], Score: 1, Record: dd.Locationid})
				return
			}
		}
	}
}

type Team struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type InfoData struct {
	Lefttime int       `json:"lefttime"`
	Records  []*SRcord `json:"records"`
	Teams    []*Team   `json:"teams"`
}

func Info(c *gin.Context) {

	diglock.Lock()
	defer diglock.Unlock()

	var data InfoData

	if gameStartTime != -1 {
		data.Lefttime = 180 - (time.Now().Second() - gameStartTime)
	} else {
		data.Lefttime = 0
	}

	if gin.Mode() == gin.ReleaseMode {
		data.Records = scores_record
	} else {
		if len(scores_record) > 100 {
			data.Records = scores_record[len(scores_record)-100:]
		} else {
			data.Records = scores_record
		}
	}

	score_map := make(map[string]int)
	for _, v := range scores_record {
		t, ok := score_map[v.Team]
		if !ok {
			t = 0
		}
		s := t + v.Score
		score_map[v.Team] = s
	}

	teams := make([]*Team, 0, 0)
	for k, v := range score_map {
		teams = append(teams, &Team{Name: k, Score: v})
	}

	data.Teams = teams

	ReturnData(c, 0, data)
}
