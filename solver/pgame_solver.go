package main

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
	"strings"
	"time"
)

var magics []string

func main() {
	begin := time.Now()
	dirPath := "../data_generator/data"
	flist, _ := ioutil.ReadDir(dirPath)

	readTime := 0.0
	calTime := 0.0

	for _, f := range flist {
		fPath := dirPath + "/" + f.Name()

		beforeRead := time.Now()
		fContent, _ := ioutil.ReadFile(fPath)
		readTime += time.Since(beforeRead).Seconds()

		beforeCal := time.Now()
		text := string(fContent)
		lines := strings.Split(text, "\n")
		for _, v := range lines {
			if v != "" {
				workOnLine(v)
			}
		}
		calTime += time.Since(beforeCal).Seconds()
	}

	println("read duration:", readTime)
	println("cal duration:", calTime)
	println("whole duration:", time.Since(begin).String())
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
		//println("Magic:",item.Locationid)
		magics = append(magics, item.Locationid)
	}
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
