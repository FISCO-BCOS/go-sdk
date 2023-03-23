package golink

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/contract/kvTableTest"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/helper"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/model"
	"github.com/sirupsen/logrus"
)

func Intsert(ctx context.Context, chanID uint64, ch chan<- *model.RequestResults, totalNumber uint64, wg *sync.WaitGroup,
	request *model.Request, kVTableTestSession *kvTableTest.KVTableTestSession) {
	defer func() {
		wg.Done()
	}()
	// fmt.Printf("start goroutines id:%05d \n", chanID)
	for i := uint64(0); i < totalNumber; i++ {
		if ctx.Err() != nil {
			fmt.Printf("ctx.Err err: %v \n", ctx.Err())
			break
		}

		list := getRequestList(request)
		isSucceed, errCode, requestTime, contentLength := insertList(chanID, list, kVTableTestSession)
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

func insertList(chanID uint64, requestList []*model.Request, kVTableTestSession *kvTableTest.KVTableTestSession) (isSucceed bool, errCode int, requestTime uint64,
	contentLength int64) {
	var (
		succeed = false
		code    = model.HTTPOk
		u       = uint64(0)
		length  int64
	)
	for _, request := range requestList {
		switch request.Method {
		case "insert":
			succeed, code, u, length = insert(chanID, request, kVTableTestSession)
		case "select":
			succeed, code, u, length = selectById(chanID, request, kVTableTestSession)
		default:
			isSucceed = false
		}
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

func selectById(chanID uint64, request *model.Request, kvtabletestSession *kvTableTest.KVTableTestSession) (bool, int, uint64, int64) {
	var (
		startTime     = time.Now()
		isSucceed     = false
		errCode       = model.HTTPOk
		contentLength = int64(0)
		requestTime   uint64
	)
	id := "3cb5896a-da72-4723-98d0-de7ead800fb4"
	//select
	_, _, err := kvtabletestSession.Select(id) // call get API
	if err != nil {
		logrus.Fatal(err)
	}
	isSucceed = true
	requestTime = uint64(helper.DiffNano(startTime))
	return isSucceed, errCode, requestTime, contentLength
}

// insert
func insert(chanID uint64, request *model.Request, kvtabletestSession *kvTableTest.KVTableTestSession) (bool, int, uint64, int64) {
	var (
		startTime     = time.Now()
		isSucceed     = false
		errCode       = model.HTTPOk
		contentLength = int64(0)
		requestTime   uint64
	)

	id := "3cb5896a-da72-4723-98d0-de7ead800fb4"
	item_name := "Laptop"
	item_age := "29"
	_, receipt, err := kvtabletestSession.Insert(id, item_name, item_age) // call set API
	if err != nil {
		errCode = model.RequestErr // 请求错误
	} else {
		contentLength = int64(len(receipt.Output))
		if receipt.Status == 0 {
			isSucceed = true
		} else {
			errCode = model.RequestErr
		}
	}
	requestTime = uint64(helper.DiffNano(startTime))
	return isSucceed, errCode, requestTime, contentLength
}
