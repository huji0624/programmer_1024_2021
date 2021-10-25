package point

import (
	"time"
)

type PointDesc struct {
	Count   int64  // 累积调用
	Content string // 描述
}

var TreasuresNum = &PointDesc{Content: "宝藏数量"}
var StartReadTime = &PointDesc{Content: "开始读取文件时间戳"}
var EndReadTime = &PointDesc{Content: "结束读取文件时间戳"}
var RealReadTime = &PointDesc{Content: "实际读取文件时间长度"}
var AddNum = &PointDesc{Content: "+ 数量"}
var SubNum = &PointDesc{Content: "- 数量"}
var MulNum = &PointDesc{Content: "* 数量"}
var ModNum = &PointDesc{Content: "% 数量"}
var ServerScore = &PointDesc{Content: "服务判定 获得分数"}
var ServerErrorMagic = &PointDesc{Content: "服务判定 挖到错误宝藏"}
var ServerFailMagic = &PointDesc{Content: "服务判定 被其他人挖走"}
var ServerOtherErr = &PointDesc{Content: "服务判定 其他错误"}
var ServerLocalIdUse = &PointDesc{Content: "服务判定 被使用的localid"}
var FormulaNum = &PointDesc{Content: "公式数量统计"}

// 打印报告
func Report() {

	println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	println(TreasuresNum.Content, TreasuresNum.Count)
	println(AddNum.Content, AddNum.Count)
	println(SubNum.Content, SubNum.Count)
	println(MulNum.Content, MulNum.Count)
	println(ModNum.Content, ModNum.Count)
	println(StartReadTime.Content, StartReadTime.Count)
	println(EndReadTime.Content, EndReadTime.Count)
	println("读取时间长度", EndReadTime.Count-StartReadTime.Count)
	println("平均读取时长", (EndReadTime.Count-StartReadTime.Count)/128)
	println(RealReadTime.Content, RealReadTime.Count)
	println("读取等待缓存区时长", EndReadTime.Count-StartReadTime.Count-RealReadTime.Count)
	println(ServerScore.Content, ServerScore.Count)
	println(ServerErrorMagic.Content, ServerErrorMagic.Count)
	println(ServerFailMagic.Content, ServerFailMagic.Count)
	println(ServerOtherErr.Content, ServerOtherErr.Count)
	println(ServerLocalIdUse.Content, ServerLocalIdUse.Count)
	println(FormulaNum.Content, FormulaNum.Count)
	println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

// 打印区间时间
func DtTime(start time.Time, name string) {
	//return
	end := time.Now().UnixMilli()
	println(name, end-start.UnixMilli(), "毫秒")
}
