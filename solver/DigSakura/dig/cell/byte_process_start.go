package cell

import (
	"DigSakura/dig"
	"DigSakura/dig/point"
	"DigSakura/dig/process"
	"strconv"
	"sync/atomic"
	"time"
)

// 默认工作队列数量
var byteWorkerNum = 1

// 正在工作中的队列数量
var byteWorkeringNum int32 = 0

const mb = 1024 * 1024

var limit = 240 * mb

// 等待处理数据的分区
var WaitProcessMemory = make(chan *BuffPool, 50)

// 处理完数据后的空分区，等待再次写入数据
var FreeMemory = make(chan *BuffPool, 50)

// 开启 cell 处理线程
func StartByteWorker(workNum int, readSuccessCb func(), workSuccessCb func()) {
	if workNum > 0 {
		byteWorkerNum = workNum
	}

	// 新建4个内存分片，等待写入
	for i := 0; i < dig.DataBlockNum; i++ {
		buff := make([]byte, limit)
		FreeMemory <- &BuffPool{Content: &buff, TotalLen: limit}
	}

	//var numCount int32
	byteWorkeringNum = int32(byteWorkerNum)
	println("开启脱壳处理管线", byteWorkerNum)
	for i := 0; i < byteWorkerNum; i++ {
		go func() {
			defer func() {

				atomic.AddInt32(&byteWorkeringNum, -1)
				//println("关闭脱壳", byteWorkeringNum)
				if byteWorkeringNum == 0 {
					//println("def调用次数", numCount)
					//point.DtTime(startTime, "脱壳管线 关闭,当前管道处理时长")
					workSuccessCb()
				}
			}()
			//startTime := time.Now()
			for {
				data, ok := <-WaitProcessMemory
				if ok {
					isReturn := false
					defer func() {
						//atomic.AddInt32(&numCount, 1)
						if !isReturn {
							// 处理完成，释放内存区
							FreeMemory <- data
							isReturn = true
						}
						err := recover()
						if err != nil {
							return
						}
					}()

					start := time.Now()
					process.Read(data.Content, data.BuffLen)
					point.DtTime(start, "脱壳时间")
					// 处理完成，释放内存区
					FreeMemory <- data
					isReturn = true
				} else {
					return
				}
			}
		}()
	}

	// 开启读取线程
	go func() {
		point.StartReadTime.Count = time.Now().UnixMilli()

		seq := ReadSeq()
		for i, elem := range seq {
			//println("文件编号", elem)
			println("已读取数量", i)
			ReadData(strconv.FormatInt(int64(elem), 10))
		}
		point.EndReadTime.Count = time.Now().UnixMilli()
		// 全部数据已推送到等待处理管线
		close(WaitProcessMemory)
		readSuccessCb()
	}()

}
