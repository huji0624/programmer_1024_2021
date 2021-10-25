package submit

import (
	"DigSakura/dig"
	"DigSakura/dig/point"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync/atomic"
	"time"
)

// 默认工作队列数量
var workerNum = 3

// 正在工作中的队列数量
var workeringNum int32 = 0

var workers = make(chan string, 10000)

// 其它错误信息
var repOtherErrData = make([]byte, 0)

// 计算错误的 localID
var errLocalIdData = make([]byte, 0)

// 开启网络线程
func StartSubmitWorker(workNum int, endCb func()) {
	if workNum > 0 {
		workerNum = workNum
	}

	workeringNum = int32(workerNum)
	println("开启提交管线", workerNum)
	for i := 0; i < workerNum; i++ {
		go func() {
			startTime := time.Now()
			netSubmitLocationId(workers)

			atomic.AddInt32(&workeringNum, -1)
			if workeringNum == 0 {
				point.DtTime(startTime, "提交管线关闭")
				endCb()
			}
		}()
	}

}
func Close() {
	ioutil.WriteFile("./bin/OtherError.log", repOtherErrData, 0666) //写入文件(字节数组)
	ioutil.WriteFile("./bin/LocalError.log", errLocalIdData, 0666)  //写入文件(字节数组)
	close(workers)
}

func SubmitLocationId(cell string) {
	workers <- cell
}

// 处理数据包线程
func netSubmitLocationId(ch chan string) {

	start := []byte(`{"token":"`)
	start = append(start, []byte(dig.Token)...)
	start = append(start, []byte(`","locationid":"`)...)
	end := []byte(`"}`)

	for {
		cell, ok := <-ch
		if ok {
			bytesData := make([]byte, 0)
			bytesData = append(bytesData, start...)
			bytesData = append(bytesData, cell...)
			bytesData = append(bytesData, end...)

			resp, err := http.Post("http://47.104.220.230/dig", "application/json",
				bytes.NewReader(bytesData))
			if err != nil {
				return
			}
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)

			//{"errorno":0}
			//0表示dig成功，获得1分；1表示该id不是宝地点的id，不得分；2表示该id的宝藏已经被挖了，不得分;-1表示其他错误;
			switch body[11] {
			case '0':
				point.ServerScore.Count++
			case '1':
				point.ServerErrorMagic.Count++
				//errLocalIdData = append(errLocalIdData, '\n')
				//errLocalIdData = append(errLocalIdData, bytesData...)
				//errLocalIdData = append(errLocalIdData, body...)
				//errLocalIdData = append(errLocalIdData, '\n')
			case '2':
				point.ServerFailMagic.Count++
				//repOtherErrData = append(repOtherErrData, '\n')
				//repOtherErrData = append(repOtherErrData, bytesData...)
				//repOtherErrData = append(repOtherErrData, '\n')
				//repOtherErrData = append(repOtherErrData, body...)
				//repOtherErrData = append(repOtherErrData, '\n')
			case '-':
				point.ServerOtherErr.Count++
				//repOtherErrData = append(repOtherErrData, '\n')
				//repOtherErrData = append(repOtherErrData, bytesData...)
				//repOtherErrData = append(repOtherErrData, body...)
				//repOtherErrData = append(repOtherErrData, '\n')
			}
		} else {
			// 通道关闭
			return
		}
	}

}

type FormulaResp struct {
	Errorno int32
	Data    []string
}

// 提交公式处理
func SubmitFormula(formula string, count int) *FormulaResp {

	point.FormulaNum.Count++

	rewards := &FormulaResp{}
	rewards.Errorno = 0
	//return rewards

	start := []byte(`{"token":"`)
	start = append(start, []byte(dig.Token)...)
	start = append(start, []byte(`","formula":"`)...)
	end := []byte(`"}`)

	bytesData := make([]byte, 0)
	bytesData = append(bytesData, start...)
	bytesData = append(bytesData, formula...)
	bytesData = append(bytesData, end...)

	resp, err := http.Post("http://47.104.220.230/formula", "application/json",
		bytes.NewReader(bytesData))
	if err != nil {
		if count > 0 {
			count--
			return SubmitFormula(formula, count)
		}
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	json.Unmarshal(body, rewards)
	return rewards
}
