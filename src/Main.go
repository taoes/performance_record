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

func main() {
	println("hello，world");
	v, _ := mem.VirtualMemory()

	// 打印内存信息
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total/(1024*1024), v.Free/(1024*1024), v.UsedPercent)

	// CPU信息
	info, _ := cpu.Info()
	fmt.Println(info)

	// CPU 核数
	count, _ := cpu.Counts(true);
	fmt.Println("cpu's count = ", count)

	// 计算CPU的使用率
	csvFile, _ := os.Create("test.csv") //创建文件
	defer csvFile.Close()
	_, _ = csvFile.WriteString("\xEF\xBB\xBF")
	write := csv.NewWriter(csvFile)
	// CPU 采样
	for j := 0; j < 50; j++ {
		data, _ := cpu.Percent(time.Second*1, true);
		var sum float64 = 0

		var dataStr []string
		dataStr = append(dataStr, strconv.Itoa(j))
		for i := 0; i < len(data); i++ {
			sum = sum + data[i]
		}
		print("index = ", j)
		println()
		rate := sum / (float64(count))
		dataStr = append(dataStr, strconv.FormatFloat(rate, 'E', -1, 32))
		_ = write.Write(dataStr)
		time.Sleep(time.Second * 1)
	}
	write.Flush()
}
