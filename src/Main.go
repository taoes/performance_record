package main

import (
	"encoding/csv"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"os"
	"strconv"
	"time"
)

const csvFile string = "CSV_RECORD.csv"

func main() {
	v, _ := mem.VirtualMemory()

	// 打印内存信息
	fmt.Printf("总计内存: %v M, 空闲内存:%v M, 使用率:%f%%\n", v.Total/(1024*1024), v.Free/(1024*1024), v.UsedPercent)

	// CPU信息
	info, _ := cpu.Info()
	print("CPU信息：\n", info[0].String())

	// CPU 核数
	count, _ := cpu.Counts(true);
	fmt.Println("CPU的核心数 = ", count)

	// 计算CPU的使用率
	csvFile, _ := os.Create(csvFile)
	defer csvFile.Close()

	_, _ = csvFile.WriteString("\xEF\xBB\xBF")
	write := csv.NewWriter(csvFile)
	// CPU 采样案例代码
	for j := 0; j < 5; j++ {
		data, _ := cpu.Percent(time.Second*1, true);
		var sum float64 = 0

		var dataStr []string
		dataStr = append(dataStr, strconv.Itoa(j))
		for i := 0; i < len(data); i++ {
			sum = sum + data[i]
		}
		rate := sum / (float64(count))
		dataStr = append(dataStr, strconv.FormatFloat(rate, 'E', -1, 32))
		_ = write.Write(dataStr)
		time.Sleep(time.Second * 1)
	}
	write.Flush()
	print("数据汇总完成")
}
