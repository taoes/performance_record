package main

import (
	"encoding/csv"
	"github.com/astaxie/beego/logs"
	"github.com/shirou/gopsutil/cpu"
	"os"
	"strconv"
	"sync"
	"time"
)

// 定义数据缓冲文件名称
const csvFile string = "CSV_RECORD.csv"

var wg sync.WaitGroup

var execSuccess = false

func main() {
	chanExec := make(chan int, 10)
	wg.Add(2)

	go sendSearchCommand(chanExec)
	go performance(chanExec)
	logs.Debug("启动检测线程完成")
	logs.Debug("正在启动目标程序信息")
	time.Sleep(time.Second * 5)
	go execCommand()
	logs.Debug("目标程序启动完成，正在收集数据，请稍等")
	wg.Wait()
}

// 执行命令
func execCommand() {
	// 延时5分钟，模拟数据
	time.Sleep(time.Minute * 5)

	execSuccess = true
	logs.Debug("目标程序执行完成.....")
}

// 一直发送数据直到程序执行完毕
func sendSearchCommand(ints chan int) {
	var index = 0
	//程序未执行完毕，每隔一秒发送一次数据
	for !execSuccess {
		ints <- index
		time.Sleep(time.Second)
		index++
	}
	ints <- -1
	logs.Debug("程序执行完成.....")
	defer wg.Done()
}

// 计算内存以及CPU使用率
func performance(chanExec chan int) {

	// CPU 核数
	count, _ := cpu.Counts(true)

	// 计算CPU的使用率
	csvFile, _ := os.Create(csvFile)
	defer csvFile.Close()
	defer wg.Done()

	_, _ = csvFile.WriteString("\xEF\xBB\xBF")
	write := csv.NewWriter(csvFile)

	var sum float64 = 0
	var dataStr []string
	dataStr = append(dataStr, "序号")
	for coreIndex := 1; coreIndex < count+1; coreIndex++ {
		dataStr = append(dataStr, "核心"+strconv.Itoa(coreIndex)+"使用率")
	}
	dataStr = append(dataStr, "平均使用率")
	_ = write.Write(dataStr)
	write.Flush()

	for chanNumber := range chanExec {
		if chanNumber == -1 {
			break
		}
		sum = 0
		logs.Debug("第", chanNumber, "次收集CPU数据，请稍后.....")

		// CPU 采样案例代码
		data, _ := cpu.Percent(time.Second*1, true)
		dataStr = dataStr[0:0]
		dataStr = append(dataStr, strconv.Itoa(chanNumber))
		for i := 0; i < len(data); i++ {
			sum = sum + data[i]
			dataStr = append(dataStr, strconv.FormatFloat(data[i], 'f', 6, 64))
		}
		rate := sum / (float64(count))
		dataStr = append(dataStr, strconv.FormatFloat(rate, 'f', 6, 64))
		_ = write.Write(dataStr)
		write.Flush()
	}

	logs.Info("====================数据汇总完成====================")
}
