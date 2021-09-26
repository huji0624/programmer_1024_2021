package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"sync"
	"time"
	"server/util"
)

var magic_ids map[string]string
var magic_id_scores = make(map[string]string)
var magic_formula = make(map[string][2]string)
var teams = make(map[string]string)
var diglock sync.Mutex

var gameRoundOver bool
var gameRoundOverTimer *time.Timer
var gameStartTime int

func CalFormula(formula string) (string,[]string,string){
	tokens := make([]string,0,10)
	token := ""
	for _,c := range formula{
		if (c>='0' && c<='9') || (c>='a' && c<='z'){
			token += string(c)
		}else{
			if token!=""{
				tokens = append(tokens,token)
				token = ""
			}
			if c=='+'{
				//ok
			}else if c=='-'{
				//ok
			}else if c=='*'{
				//ok
			}else if c=='/'{
				//ok
			}else if c=='('{
				//ok
			}else if c==')'{
				//ok
			}else{
				//ignore
				continue
			}
			tokens = append(tokens,string(c))
		}
	}

	if token!=""{
		tokens = append(tokens,token)
		token = ""
	}

	//tokens to nums
	ids := make([]string,0,10)
	readable := ""
	//replaceed := make([]string,0,10)
	//for _,v := range tokens{
	//	if v!="+" && v!="-" && v!="*" && v!="/"{
	//		magic,ok := magic_ids[v]
	//		if ok{
	//			readable += magic
	//			replaceed = append(replaceed,magic)
	//			ids = append(ids,v)
	//		}else{
	//			return "",nil,""
	//		}
	//	}else{
	//		readable += v
	//		replaceed = append(replaceed,v)
	//	}
	//}
	//sort.Strings(ids)

	return CalTokens(tokens),ids,readable
}

func ReverseTokens(tokens []string) []string{
	rTokens := make([]string,0,10)
	for i:=len(tokens)-1;i>=0;i--{
		rTokens = append(rTokens,tokens[i])
	}
	return rTokens
}

func CalTokens(tokens []string) string{
	//log.Println(tokens)
	if len(tokens)==1{
		return tokens[0]
	}

	if len(tokens)==3{
		return Caltwo(tokens[0],tokens[2],tokens[1])
	}

	//128+19*12-(888+(111+1*2+3)*99) = -12016
	nums := util.NewSatck()
	ops := util.NewSatck()
	for _,v := range tokens{
		if v=="*"{
			ops.Push(v)
		}else if v=="+"{

			if !ops.Empty()&&(ops.Top()=="*"||ops.Top()=="/"||ops.Top()=="+"||ops.Top()=="-"){
				num2 := nums.Pop()
				num1 := nums.Pop()
				token := Caltwo(num1,num2,ops.Pop())
				nums.Push(token)
			}

			ops.Push(v)
		}else if v=="-"{
			if !ops.Empty()&&(ops.Top()=="*"||ops.Top()=="/"||ops.Top()=="+"||ops.Top()=="-"){
				num2 := nums.Pop()
				num1 := nums.Pop()
				token := Caltwo(num1,num2,ops.Pop())
				nums.Push(token)
			}

			ops.Push(v)
		}else if v=="/"{
			ops.Push(v)
		}else if v=="("{
			ops.Push(v)
		}else if v==")"{
			newTokens := make([]string,0,10)
			for !ops.Empty() && ops.Top()!="(" {
				newTokens = append(newTokens,nums.Pop())
				newTokens = append(newTokens,ops.Pop())
			}
			newTokens = append(newTokens,nums.Pop())
			ops.Pop()

			token := CalTokens(ReverseTokens(newTokens))

			nums.Push(token)
		}else{
			if !ops.Empty()&&(ops.Top()=="*" || ops.Top()=="/"){
				num2 := v
				num1 := nums.Pop()
				token := Caltwo(num1,num2,ops.Pop())
				nums.Push(token)
			}else{
				nums.Push(v)
			}
		}
	}

	if ops.Size()==1 && nums.Size()==2{
		num2 := nums.Pop()
		num1 := nums.Pop()
		return Caltwo(num1,num2,ops.Pop())
	}

	if nums.Size()==1 && ops.Empty(){
		return nums.Top()
	}

	newTokens := make([]string,0,10)
	for !nums.Empty(){
		newTokens = append(newTokens,nums.Pop())
		if !ops.Empty(){
			newTokens = append(newTokens,ops.Pop())
		}
	}

	if len(newTokens)==len(tokens){
		return ""
	}

	return CalTokens(ReverseTokens(newTokens))
}

func Caltwo(num1 string,num2 string,op string) string{
	bi1 := big.NewInt(0)
	bi1.SetString(num1,10)
	bi2 := big.NewInt(0)
	bi2.SetString(num2,10)
	switch op {
	case "+":
		return bi1.Add(bi1,bi2).String()
	case "-":
		return bi1.Sub(bi1,bi2).String()
	case "*":
		return bi1.Mul(bi1,bi2).String()
	case "/":
		return bi1.Div(bi1,bi2).String()
	}

	return ""
}

func main() {
	form := "(102**4+-))12(9912*12)1)"
	log.Println(form)
	log.Println(CalFormula(form))
	return

	pid := os.Getpid()
	ioutil.WriteFile("./pid",[]byte(fmt.Sprintf("%d",pid)),0644)

	loadMagicIDS()

	//token => teamname map
	teams["test1"] = "test1"
	teams["test2"] = "test2"

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default() //返回默认引擎，里面有系统定义的中间件

	//跨域问题
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://192.168.1.66:9091"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.Static("/files", "./files")
	router.Static("/h5", "../h5/dist")

	router.POST("/dig",Dig)
	router.POST("/formula",Formula)

	router.GET("/reset",Reset)
	router.GET("/info",Info)

	router.Run(":80")

}

func loadMagicIDS(){
	data,err := ioutil.ReadFile("../data_generator/magic_ids.json")
	if err!=nil{
		log.Println(err.Error())
		os.Exit(-1)
	}

	err = json.Unmarshal(data,&magic_ids)
	if err!=nil{
		log.Println(err.Error())
		os.Exit(-1)
	}
	for k,_ := range magic_ids{
		magic_id_scores[k] = ""
	}

	magic_formula = make(map[string][2]string)

	log.Println("loadMagicIDS:")
	log.Println(len(magic_id_scores))

	gameRoundOver = false
	gameRoundOverTimer = nil
	gameStartTime = -1
}

func ReturnError(c *gin.Context, apierr error) {
	c.JSON(http.StatusOK, gin.H{
		"errorno": -1,
		"msg":    apierr.Error(),
	})
}

func ReturnData(c *gin.Context,errorno int, data interface{}) {
	if data==nil{
		c.JSON(http.StatusOK, gin.H{
			"errorno": errorno,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"data":   data,
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

	if params.Code!="pg"{
		ReturnError(c,errors.New("server error."))
		return
	}

	loadMagicIDS()

	ReturnData(c,0,nil)
}

type FormulaData struct {
	Token string `json:"token"`
	Formula string `json:"formula"`
}

func Formula(c *gin.Context) {
	//if gameRoundOver{
	//	ReturnError(c,errors.New("game round is over."))
	//	return
	//}

	var fd FormulaData
	err := c.BindJSON(&fd)

	if err != nil {
		ReturnError(c,err)
		return
	}

	//if fd.Token==""{
	//	ReturnError(c,errors.New("params token missing."))
	//	return
	//}
	//if teams[fd.Token]==""{
	//	ReturnError(c,errors.New("token not valide."))
	//	return
	//}
	if fd.Formula==""{
		ReturnError(c,errors.New("params formula missing."))
		return
	}

	diglock.Lock()
	defer diglock.Unlock()

	formula := fd.Formula
	ret,ids,readable := CalFormula(formula)

	//for test
	if ret=="1024"{
		ReturnData(c,0,nil)
		return
	}else{
		ReturnData(c,1,nil)
		return
	}
	//for test

	if ret=="1024"{
		idsbytes,jerr := json.Marshal(ids)
		if jerr==nil{
			idskey := string(idsbytes)
			_,ok := magic_formula[idskey]
			if !ok{
				var tmp [2]string
				tmp[0] = teams[fd.Token]
				tmp[1] = readable
				magic_formula[idskey] = tmp

				ReturnData(c,0,nil)
			}else{
				ReturnData(c,2,nil)
			}
		}else{
			ReturnData(c,1,nil)
		}
	}else{
		//formula wrong or id wrong
		ReturnData(c,1,nil)
	}
}

type DigData struct {
	Token string `json:"token"`
	Locationid string `json:"locationid"`
}

func Dig(c *gin.Context) {
	if gameRoundOver{
		ReturnError(c,errors.New("game round is over."))
		return
	}

	var dd DigData
	err := c.BindJSON(&dd)

	if err != nil {
		ReturnError(c,err)
		return
	}

	if dd.Token==""{
		ReturnError(c,errors.New("params token missing."))
		return
	}
	if teams[dd.Token]==""{
		ReturnError(c,errors.New("token not valide."))
		return
	}
	if dd.Locationid==""{
		ReturnError(c,errors.New("params Locationid missing."))
		return
	}

	diglock.Lock()
	defer diglock.Unlock()

	if gameRoundOverTimer==nil{
		gameRoundOverTimer = time.AfterFunc(time.Second*180, func() {
			gameRoundOver = true
		})
		gameStartTime = time.Now().Second()
	}

	t,ok := magic_id_scores[dd.Locationid]
	if !ok{
		//fail.not treasure
		ReturnData(c,1,nil)
	}else{
		if t==""{
			magic_id_scores[dd.Locationid] = teams[dd.Token]
			log.Printf("Team %v dig success.",dd.Token)
			//success
			ReturnData(c,0,nil)
		}else{
			//fail.alreay digged by others.
			ReturnData(c,2,nil)
		}
	}
}

type InfoData struct {
	Formulas map[string][2]string `json:"formulas"`
	Magics map[string]string `json:"magics"`
	Lefttime int `json:"lefttime"`
}

func Info(c *gin.Context) {

	diglock.Lock()
	defer diglock.Unlock()

	var data InfoData

	var result = make(map[string]string)
	for k,v := range magic_id_scores{
		if v!=""{
			result[k] = v
		}
	}

	var tmp [2]string
	tmp[0] = teams["test1"]
	tmp[1] = "(ajiasais+kasais)-kasoaks"
	jb,_ := json.Marshal([]string{"ajiasais","kasais","kasoaks"})
	magic_formula[string(jb)] = tmp

	tmp[0] = teams["test2"]
	tmp[1] = "(ajiasais+kasais)-kasoaks"
	jb2,_ := json.Marshal([]string{"ajiaxsais","kasais","kasoaks"})
	magic_formula[string(jb2)] = tmp

	tmp[0] = teams["test2"]
	tmp[1] = "(ajiasais+kasais)-kasoaks"
	jb3,_ := json.Marshal([]string{"ajiasaisss","kasaissss","kasoakssss"})
	magic_formula[string(jb3)] = tmp

	if gameStartTime!=-1{
		data.Lefttime = 180 - (time.Now().Second() - gameStartTime)
	}
	data.Magics = result
	data.Formulas = magic_formula

	ReturnData(c,0,data)
}



