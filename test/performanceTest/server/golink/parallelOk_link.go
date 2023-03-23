package golink

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/contract/parallelOk"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/helper"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/model"
	"github.com/sirupsen/logrus"
)

func Transfer(ctx context.Context, chanID uint64, ch chan<- *model.RequestResults, totalNumber uint64, wg *sync.WaitGroup,
	request *model.Request, parallelOkSession *parallelOk.ParallelOkSession) {
	defer func() {
		wg.Done()
	}()

	// fmt.Printf("启动协程 编号:%05d \n", chanID)
	for i := uint64(0); i < totalNumber; i++ {
		if ctx.Err() != nil {
			fmt.Printf("ctx.Err err: %v \n", ctx.Err())
			break
		}

		list := getRequestList(request)
		isSucceed, errCode, requestTime, contentLength := transferList(chanID, list, parallelOkSession)
		requestResults := &model.RequestResults{
			Time:          requestTime,
			IsSucceed:     isSucceed,
			ErrCode:       errCode,
			ReceivedBytes: contentLength,
		}
		requestResults.SetID(chanID, i)
		ch <- requestResults
	}
}

// sendList 多个接口分步压测
func transferList(chanID uint64, requestList []*model.Request, parallelOkSession *parallelOk.ParallelOkSession) (isSucceed bool, errCode int, requestTime uint64,
	contentLength int64) {
	errCode = model.HTTPOk
	for _, request := range requestList {
		succeed, code, u, length := transfer(request, parallelOkSession)
		isSucceed = succeed
		errCode = code
		requestTime = requestTime + u
		contentLength = contentLength + length
		if !succeed {
			break
		}
	}
	return
}

func transfer(request *model.Request, parallelOkSession *parallelOk.ParallelOkSession) (bool, int, uint64, int64) {
	var (
		startTime     = time.Now()
		isSucceed     = false
		errCode       = model.HTTPOk
		contentLength = int64(0)
		requestTime   uint64
	)
	from_name := "zhangsan"
	to_name := "lisi"
	num := big.NewInt(10)
	_, _, err := parallelOkSession.Transfer(from_name, to_name, num) // call set API
	if err != nil {
		logrus.Fatal(err)
	}
	isSucceed = true
	requestTime = uint64(helper.DiffNano(startTime))
	return isSucceed, errCode, requestTime, contentLength
}

func getBalance(request *model.Request, parallelOkSession *parallelOk.ParallelOkSession) (bool, int, uint64, int64) {
	var (
		startTime     = time.Now()
		isSucceed     = false
		errCode       = model.HTTPOk
		contentLength = int64(0)
		requestTime   uint64
	)
	to_name := "lisi"
	_, err := parallelOkSession.BalanceOf(to_name) // call get API
	if err != nil {
		logrus.Fatal(err)
	}
	//fmt.Printf("to_name: %v, item_num: %v \n", to_name, num.Int64() )
	requestTime = uint64(helper.DiffNano(startTime))
	return isSucceed, errCode, requestTime, contentLength
}
