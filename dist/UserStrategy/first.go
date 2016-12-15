package bitbot

import (
	"fmt"
)

func onTick() {
	// 策略采用轮询而非事件驱动是因为作者喜欢对代码100%的掌控力.
}

func main() {
	fmt.Println("Running first strategy")
}

/*

exchange	:默认主交易所对像, 添加交易平台时排列第一的交易所
exchanges	:交易所数组, 如果添加多个交易所, 可以访问此变量获取交易所对像

exchange.GetAccount();
exchanges[0].GetAccount();

Version
Log	()

Sleep(Millisecond)

LogProfit(Profit)

LogProfitReset	 清空所有收益日志, 可以带一个数字参数, 指定保留的条数
LogReset	清空所有日志, 可以带一个数字参数, 指定保留的条数
LogStatus(Msg)	此信息不保存到日志列表里, 只更新当前机器人的状态信息, 在日志上方显示, 可多次调用, 更新状态
EnableLog(IsEnable)	打开或者关闭定单和出错信息的日志记录
Chart({...})
Mail(...)
SetErrorFilter(RegEx)
GetPid
GetLastError
_G(K, V)
HttpQuery(...)
Dial(Address, Timeout)
GetCommand()


//EnableWebsocket()
*/
