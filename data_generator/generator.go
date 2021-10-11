package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var TotalFileCount int
var DataCountEachFile int
const ConcurrentCount = 2
const MagicRatio = 50000
const OutputDir = "data"

var magicids map[string]string = make(map[string]string)
var magicidsLock sync.Mutex

func createEnv() {
	err0 := os.RemoveAll(OutputDir)
	if err0!=nil{
		log.Println(err0)
		os.Exit(-1)
	}

	_, err := os.Stat(OutputDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(OutputDir, os.ModePerm)
		if err != nil {
			panic("output dir create failed,please check it;")
		}
	}
}
func main() {
	var randomSourceSeed int64
	flag.Int64Var(&randomSourceSeed,"s",0,"random source seed.")
	flag.IntVar(&TotalFileCount,"c",5,"total file count.")
	flag.IntVar(&DataCountEachFile,"d",100,"data count each file.unit 10000.")
	flag.Parse()

	DataCountEachFile = DataCountEachFile * 10000

	log.Println("Seed : ",randomSourceSeed)

	createEnv()

	begin := time.Now()
	ch := make(chan int, ConcurrentCount)
	done := make(chan int, TotalFileCount)
	for i := 0; i < TotalFileCount; i++ {
		ch <- i
		go generateOneFile(ch, done,randomSourceSeed+int64(i),randomSourceSeed==0)
	}

	count := 0
	for true {
		<-done

		count++
		if count == TotalFileCount {
			break
		}
	}

	bs, err := json.Marshal(magicids)
	if err != nil {
		println(err)
		return
	}

	println("")
	println("Total magics : ", len(magicids))
	ioutil.WriteFile("magic_ids.json", []byte(bs), 0644)

	println("whole duration:", time.Since(begin).String())
}

func generateID(rand *rand.Rand) string {

	from := "0123456789abcdefghijklmnopqrsduvwxyz"
	var s [64]byte
	for i := 0; i < 64; i++ {
		s[i] = from[rand.Intn(36)]
	}

	return string(s[:])
}

func randomCal(bi *big.Int, n *big.Int, rand *rand.Rand) *big.Int {

	var magic *big.Int

	cal := rand.Intn(4)
	switch cal {
	case 0:
		magic = bi.Add(bi, n)
		break
	case 1:
		magic = bi.Sub(bi, n)
		break
	case 2:
		magic = bi.Mul(bi, n)
		break
	case 3:
		magic = bi.Mod(bi, n)
		break
	}

	return magic
}

func generateNumber(rand *rand.Rand) string {

	numlen := rand.Intn(31) + 2

	sb := strings.Builder{}
	from := "0123456789"
	for i := 0; i < numlen; i++ {
		sb.WriteByte(from[rand.Intn(10)])
	}

	return sb.String()
}

func generateMagicNumber(bi *big.Int, rand *rand.Rand) (bool, string) {

	n := big.NewInt(1024)
	ifmagic := rand.Intn(MagicRatio) == 0

	magic := randomCal(bi, n, rand)

	if ifmagic {
		return true, magic.String()
	} else {
		return false, generateNumber(rand)
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

func generateOneFile(c chan int, done chan int,randSourceSeed int64,timesource bool) {

	magicCount := 0
	var source rand.Source
	if timesource{
		source = rand.NewSource(time.Now().UnixNano())
	}else{
		source = rand.NewSource(randSourceSeed)
	}
	generator := rand.New(source)

	var b strings.Builder
	for i := 0; i < DataCountEachFile; i++ {
		id := generateID(generator)
		numInID := stripInt(id)
		ifmagic, mg := generateMagicNumber(copyBig(numInID), generator)
		if ifmagic {
			magicidsLock.Lock()
			magicids[id] = mg
			magicidsLock.Unlock()
			magicCount++
		} else {
			if checkIFMagic(numInID, mg) {
				//println(i, " - wrong magic number:", id, " ", numInID.String(), " ", mg)
				continue
			}
		}

		tmp := fmt.Sprintf("{\"locationid\":\"%s\",\"magic\":\"%s\"}\n", id, mg)

		b.WriteString(tmp)
	}

	//println(time.Now().String()," write data...")

	num := <-c
	filename := fmt.Sprintf("%s/Treasure_%d.data", OutputDir, num)
	ioutil.WriteFile(filename, []byte(b.String()), 0644)

	println(time.Now().String(), " done.", filename, " magicCount:", magicCount)
	done <- num
}
