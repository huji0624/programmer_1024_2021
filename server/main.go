package main

import (
	"encoding/json"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

var magic_ids []string
var magic_id_scores = make(map[string]string)
var teams = make(map[string]string)
var diglock sync.Mutex

func main() {

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
	for _,v := range magic_ids{
		magic_id_scores[v] = ""
	}

	//token => teamname map
	teams["test1"] = "test1"
	teams["test2"] = "test2"

	gin.SetMode(gin.DebugMode)
	router := gin.Default() //返回默认引擎，里面有系统定义的中间件

	//跨域问题
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080", "http://192.168.1.66:9091"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.Static("/files", "./files")

	router.POST("",nil)
	router.GET("",nil)

	router.POST("/dig",Dig)
	router.GET("/info",Info)

	router.Run(":80")

}

func ReturnError(c *gin.Context, apierr error) {
	log.Println("error:",apierr.Error())

	c.JSON(http.StatusOK, gin.H{
		"errorno": -1,
		"msg":    apierr.Error(),
	})
}

func ReturnData(c *gin.Context,errorno int, data interface{}) {
	log.Println("return data:", data)

	c.JSON(http.StatusOK, gin.H{
		"data":   data,
		"errorno": errorno,
	})
}

type DigData struct {
	Token string `json:"token"`
	Locationid string `json:"locationid"`
}

func Dig(c *gin.Context) {
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

	t,ok := magic_id_scores[dd.Locationid]
	if !ok{
		//fail.not treasure
		ReturnData(c,1,nil)
	}else{
		if t==""{
			magic_id_scores[dd.Locationid] = teams[dd.Token]
			//success
			ReturnData(c,0,nil)
		}else{
			//fail.alreay digged by others.
			ReturnData(c,2,nil)
		}
	}
}

type InfoData struct {
	Total int `json:"total"`
	Result map[string]int `json:"result"`
}

func Info(c *gin.Context) {
	var data InfoData

	var result = make(map[string]int)

	for _,v := range magic_id_scores{
		if v!=""{
			result[v] = result[v] + 1
		}
	}

	data.Result = result
	data.Total = len(magic_id_scores)

	ReturnData(c,0,data)
}



