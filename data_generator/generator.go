package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const TotalFileCount = 100
const DataCountEachFile = 10000*1
const ConcurrentCount = 10
const MagicRatio = 10000

var magicids []string = make([]string,0,1000)
var magicidsLock sync.Mutex

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int, ConcurrentCount)
	done := make(chan int, TotalFileCount)
	for i:=0;i< TotalFileCount;i++{
		ch <- i
		go generateOneFile(ch,done)
	}

	count := 0
	for true {
		<- done

		count++
		if count== TotalFileCount {
			break
		}
	}

	bs,err := json.Marshal(magicids)
	if err!=nil{
		println(err)
		return
	}

	println("Total magics : ",len(magicids))
	ioutil.WriteFile("magic_ids.json",[]byte(bs),0644)
}

func generateID() string{

	from := "0123456789abcdefghijklmnopqrsduvwxyz"
	var s [64]byte
	for i:=0;i<64;i++{
		s[i]=from[rand.Intn(36)]
	}

	return string(s[:])
}

func randomCal(bi *big.Int,n *big.Int) *big.Int{

	var magic *big.Int

	cal := rand.Intn(4)
	switch cal {
	case 0:
		magic = bi.Add(bi,n)
		break
	case 1:
		magic = bi.Sub(bi,n)
		break
	case 2:
		magic = bi.Mul(bi,n)
		break
	case 3:
		magic = bi.Mod(bi,n)
		break
	}

	return magic
}

func generateMagicNumber(bi *big.Int) (bool,string) {

	n := big.NewInt(1024)
	ifmagic := rand.Intn(MagicRatio)==0

	magic := randomCal(bi,n)

	if ifmagic{
		return true,magic.String()
	}else{
		return false,randomCal(magic,big.NewInt(rand.Int63n(10000)+1)).String()
	}
}

func stripInt(id string) *big.Int{

	s := ""
	for _,v := range id{
		if v>='0' && v<='9'{
			s+=string(v)
		}
	}

	n := big.NewInt(0)
	n.SetString(s,10)

	return n
}

func generateOneFile(c chan int,done chan int){

	magicCount := 0

	//println(time.Now().String()," generate data...")
	var b strings.Builder
	for i:=0;i< DataCountEachFile;i++{
		id := generateID()
		numInID := stripInt(id)
		ifmagic,mg := generateMagicNumber(numInID)
		if ifmagic{
			magicidsLock.Lock()
			magicids = append(magicids,id)
			magicidsLock.Unlock()
			magicCount++
		}

		tmp := fmt.Sprintf("{\"locationid\":\"%s\",\"magic\":\"%s\"}\n",id,mg)

		b.WriteString(tmp)
	}

	//println(time.Now().String()," write data...")

	num := <- c
	filename := fmt.Sprintf("data/Treasure_%d.data",num)
	ioutil.WriteFile(filename,[]byte(b.String()),0644)


	println(time.Now().String()," done.",filename," magicCount:",magicCount)
	done <- num
}
