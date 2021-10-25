package cell

import (
	"DigSakura/dig/point"
	"fmt"
	"os"
	"time"
)

type BuffPool struct {
	Content  *[]byte // 数据块指针
	TotalLen int     // 内存区总大小
	BuffLen  int     // 数据大小
}

func ReadData(index string) {
	var filePath = "./data/Treasure_" + index + ".data"
	// 等待内存分配

	for {
		fp, err := os.Open(filePath)
		data, ok := <-FreeMemory
		if ok {
			//println("获取到可写入分片内存")
			//startTime := time.Now().UnixMilli()
			startTime := time.Now()
			if err != nil {
				fmt.Println(err)
				return
			}
			data.BuffLen, err = fp.Read(*data.Content)
			if err != nil {
				fmt.Println(err)
			}
			fp.Close()
			//point.DtTime(startTime,"读取文件时间")
			endTime := time.Now().UnixMilli()
			point.RealReadTime.Count += (endTime - startTime.UnixMilli())
			// 推送数据块到处理管线
			WaitProcessMemory <- data
			return
		} else {
			// 通道关闭
			return
		}
	}

}
