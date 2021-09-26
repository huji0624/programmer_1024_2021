package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var magicids = make(map[string]string)
var magiclock sync.Mutex

func main() {
	//cpuNum := runtime.NumCPU()

	begin := time.Now()
	dirPath := "../data_generator/data"
	flist, _ := ioutil.ReadDir(dirPath)

	ch := make(chan string)

	for _, f := range flist {
		fPath := dirPath + "/" + f.Name()
		go workOnFile(fPath,ch)
	}

	count := 0
	for true {
		<-ch
		count++
		if count==len(flist){
			break
		}
	}

	for k,v := range magicids{
		log.Printf("%v : %v",k,v)
	}

	println("whole duration:", time.Since(begin).String())
}

func workOnFile(fPath string,ch chan string){
	fContent, _ := ioutil.ReadFile(fPath)

	text := string(fContent)
	lines := strings.Split(text, "\n")
	for _, v := range lines {
		if v != "" {
			workOnLine(v)
		}
	}

	ch <- fPath
}

type Line struct {
	Locationid string
	Magic      string
}

func workOnLine(line string) {
	var item Line
	err := json.Unmarshal([]byte(line), &item)
	if err != nil {
		println(line)
		println(err.Error())
		os.Exit(0)
	}

	bi := stripInt(item.Locationid)
	if checkIFMagic(bi, item.Magic) {
		//magiclock.Lock()
		//magicids[item.Locationid] = item.Magic
		//magiclock.Unlock()

		go SendRequest(item.Locationid)
	}
}

func SendRequest(locationid string){
	//return
	data := make(map[string]interface{})
	data["locationid"] = locationid
	data["token"] = "test1"

	bytesData, _ := json.Marshal(data)
	resp, _ := http.Post("http://47.104.220.230/dig","application/json", bytes.NewReader(bytesData))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func stripInt(id string) *big.Int {

	s := ""
	for _, v := range id {
		if v >= '0' && v <= '9' {
			s += string(v)
		}
	}

	n := big.NewInt(0)
	n.SetString(s, 10)

	return n
}

func copyBig(id *big.Int) *big.Int {
	cal := big.NewInt(0)
	cal = cal.Set(id)
	return cal
}

func checkIFMagic(id *big.Int, mg string) bool {

	mgValue, success := big.NewInt(0).SetString(mg, 10)
	if !success {
		panic("failed decode mg")
	}
	num1024 := big.NewInt(1024)
	cal := copyBig(id)
	result := cal.Add(cal, num1024)
	if result.Cmp(mgValue) == 0 {
		return true
	}
	cal = copyBig(id)
	result = cal.Sub(cal, num1024)
	if result.Cmp(mgValue) == 0 {
		return true
	}

	cal = copyBig(id)
	result = cal.Mul(cal, num1024)
	if result.Cmp(mgValue) == 0 {
		return true
	}

	cal = copyBig(id)
	result = cal.Mod(cal, num1024)
	if result.Cmp(mgValue) == 0 {
		return true
	}

	return false
}
