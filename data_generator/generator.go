package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

const TotalFileCount = 10
const DataCountEachFile = 10000 * 30
const ConcurrentCount = 10
const MagicRatio = 10000
const OutputDir = "data"

var magicids []string = make([]string, 0, 1000)
var magicidsLock sync.Mutex

func createEnv() {
	_, err := os.Stat(OutputDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(OutputDir, os.ModePerm)
		if err != nil {
			panic("output dir create failed,please check it;")
		}
	}
}
func main() {
	createEnv()
	rand.Seed(time.Now().UnixNano())
	begin := time.Now()
	ch := make(chan int, ConcurrentCount)
	done := make(chan int, TotalFileCount)
	for i := 0; i < TotalFileCount; i++ {
		ch <- i
		go generateOneFile(ch, done)
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

func randomCal(bi *big.Int, n *big.Int) *big.Int {

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

func generateNumber() string {

	numlen := rand.Intn(18) + 2

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

	magic := randomCal(bi, n)

	if ifmagic {
		return true, magic.String()
	} else {
		return false, generateNumber()
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

func generateOneFile(c chan int, done chan int) {

	magicCount := 0
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	var b strings.Builder
	for i := 0; i < DataCountEachFile; i++ {
		id := generateID(generator)
		numInID := stripInt(id)
		ifmagic, mg := generateMagicNumber(copyBig(numInID), generator)
		if ifmagic {
			magicidsLock.Lock()
			magicids = append(magicids, id)
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
