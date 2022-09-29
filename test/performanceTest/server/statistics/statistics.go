// Package statistics 统计数据
package statistics

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/model"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/tools"
)

var (
	// Time to output statistics
	exportStatisticsTime = 1 * time.Second
	//p  = message.NewPrinter(language.English)
	requestTimeList []uint64 // 所有请求响应时间
)

// ReceivingResults Receive the result and process it
// statistics time is nanosecond，printf time is millisecond
// concurrent
func ReceivingResults(concurrent uint64, ch <-chan *model.RequestResults, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	var stopChan = make(chan bool)
	// time
	var (
		processingTime uint64
		requestTime    uint64
		maxTime        uint64
		minTime        uint64
		successNum     uint64
		failureNum     uint64
		chanIDLen      int
		chanIDs        = make(map[uint64]bool)
		receivedBytes  int64
		mutex          = sync.RWMutex{}
	)
	statTime := uint64(time.Now().UnixNano())
	// error code
	var errCode = &sync.Map{}
	// Periodically output a calculation result
	ticker := time.NewTicker(exportStatisticsTime)
	go func() {
		for {
			select {
			case <-ticker.C:
				endTime := uint64(time.Now().UnixNano())
				mutex.Lock()
				go calculateData(concurrent, processingTime, endTime-statTime, maxTime, minTime, successNum, failureNum,
					chanIDLen, errCode, receivedBytes)
				mutex.Unlock()
			case <-stopChan:
				return
			}
		}
	}()
	header()
	for data := range ch {
		mutex.Lock()
		processingTime = processingTime + data.Time
		if maxTime <= data.Time {
			maxTime = data.Time
		}
		if minTime == 0 {
			minTime = data.Time
		} else if minTime > data.Time {
			minTime = data.Time
		}
		// request success or not
		if data.IsSucceed == true {
			successNum = successNum + 1
		} else {
			failureNum = failureNum + 1
		}
		// Statistical error code
		if value, ok := errCode.Load(data.ErrCode); ok {
			valueInt, _ := value.(int)
			errCode.Store(data.ErrCode, valueInt+1)
		} else {
			errCode.Store(data.ErrCode, 1)
		}
		receivedBytes += data.ReceivedBytes
		if _, ok := chanIDs[data.ChanID]; !ok {
			chanIDs[data.ChanID] = true
			chanIDLen = len(chanIDs)
		}
		requestTimeList = append(requestTimeList, data.Time)
		mutex.Unlock()
	}
	// All data is accepted. Stop the periodic output of statistics
	stopChan <- true
	endTime := uint64(time.Now().UnixNano())
	requestTime = endTime - statTime
	calculateData(concurrent, processingTime, requestTime, maxTime, minTime, successNum, failureNum, chanIDLen, errCode,
		receivedBytes)

	fmt.Printf("\n\n")
	fmt.Println("*************************  end stat  ****************************")
	fmt.Println("处理协程数量:", concurrent)
	// fmt.Println("处理协程数量:", concurrent, "程序处理总时长:", fmt.Sprintf("%.3f", float64(processingTime/concurrent)/1e9), "秒")
	fmt.Println("请求总数（并发数*请求数 -c * -n）:", successNum+failureNum, "总请求时间:",
		fmt.Sprintf("%.3f", float64(requestTime)/1e9),
		"秒", "successNum:", successNum, "failureNum:", failureNum)
	printTop(requestTimeList)
	fmt.Println("*************************  结果 end   ****************************")
	fmt.Printf("\n\n")
}

// printTop 排序后计算 top 90 95 99
func printTop(requestTimeList []uint64) {
	if requestTimeList == nil {
		return
	}
	all := tools.MyUint64List{}
	all = requestTimeList
	sort.Sort(all)
	fmt.Println("tp90:", fmt.Sprintf("%.3f", float64(all[int(float64(len(all))*0.90)]/1e6)))
	fmt.Println("tp95:", fmt.Sprintf("%.3f", float64(all[int(float64(len(all))*0.95)]/1e6)))
	fmt.Println("tp99:", fmt.Sprintf("%.3f", float64(all[int(float64(len(all))*0.99)]/1e6)))
}

// calculateData
func calculateData(concurrent, processingTime, requestTime, maxTime, minTime, successNum, failureNum uint64,
	chanIDLen int, errCode *sync.Map, receivedBytes int64) {
	if processingTime == 0 {
		processingTime = 1
	}
	var (
		qps              float64
		averageTime      float64
		maxTimeFloat     float64
		minTimeFloat     float64
		requestTimeFloat float64
	)
	// 平均 每个协程成功数*总协程数据/总耗时 (每秒)
	if processingTime != 0 {
		qps = float64(successNum*1e9*concurrent) / float64(processingTime)
	}
	// 平均时长 总耗时/总请求数/并发数 纳秒=>毫秒
	if successNum != 0 && concurrent != 0 {
		averageTime = float64(processingTime) / float64(successNum*1e6)
	}
	// nanosecond=>millisecond
	maxTimeFloat = float64(maxTime) / 1e6
	minTimeFloat = float64(minTime) / 1e6
	requestTimeFloat = float64(requestTime) / 1e9
	// All print in milliseconds
	table(successNum, failureNum, errCode, qps, averageTime, maxTimeFloat, minTimeFloat, requestTimeFloat, chanIDLen,
		receivedBytes)
}

// header 打印表头信息
func header() {
	fmt.Printf("\n\n")
	// 打印的时长都为毫秒 总请数
	fmt.Println("─────┬───────┬───────┬───────┬────────┬────────┬────────┬────────┬────────┬────────┬────────")
	fmt.Println(" 耗时│ 并发数│ 成功数│ 失败数│   qps  │最长耗时│最短耗时│平均耗时│下载字节│字节每秒│ 状态码")
	fmt.Println("─────┼───────┼───────┼───────┼────────┼────────┼────────┼────────┼────────┼────────┼────────")
	return
}

// table
func table(successNum, failureNum uint64, errCode *sync.Map,
	qps, averageTime, maxTimeFloat, minTimeFloat, requestTimeFloat float64, chanIDLen int, receivedBytes int64) {
	var (
		speed int64
	)
	if requestTimeFloat > 0 {
		speed = int64(float64(receivedBytes) / requestTimeFloat)
	} else {
		speed = 0
	}
	var (
		receivedBytesStr string
		speedStr         string
	)
	// Determine whether the download length is unknown
	if receivedBytes <= 0 {
		receivedBytesStr = ""
		speedStr = ""
	} else {
		receivedBytesStr = fmt.Sprintf("%d", receivedBytes)
		speedStr = fmt.Sprintf("%d", speed)
	}
	// All print in milliseconds
	result := fmt.Sprintf("%4.0fs│%7d│%7d│%7d│%8.2f│%8.2f│%8.2f│%8.2f│%8s│%8s│%v",
		requestTimeFloat, chanIDLen, successNum, failureNum, qps, maxTimeFloat, minTimeFloat, averageTime,
		receivedBytesStr, speedStr,
		printMap(errCode))
	fmt.Println(result)
	return
}

// printMap
func printMap(errCode *sync.Map) (mapStr string) {
	var (
		mapArr []string
	)
	errCode.Range(func(key, value interface{}) bool {
		mapArr = append(mapArr, fmt.Sprintf("%v:%v", key, value))
		return true
	})
	sort.Strings(mapArr)
	mapStr = strings.Join(mapArr, ";")
	return
}
