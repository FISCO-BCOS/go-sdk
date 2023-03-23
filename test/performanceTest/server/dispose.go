package server

import (
	"context"
	"sync"
	"time"

	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/contract/kvTableTest"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/contract/parallelOk"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/model"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/server/golink"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/server/statistics"
)

// Dispose
func Dispose(ctx context.Context, concurrency, totalNumber uint64, request *model.Request, session interface{}) {
	// 设置接收数据缓存
	ch := make(chan *model.RequestResults, 1000)
	var (
		wg          sync.WaitGroup // Sending data is completed
		wgReceiving sync.WaitGroup // Data processing completed
	)
	wgReceiving.Add(1)
	go statistics.ReceivingResults(concurrency, ch, &wgReceiving)

	for i := uint64(0); i < concurrency; i++ {
		wg.Add(1)
		switch request.Form {
		case model.FormTypeKvTable:
			kvtabletestSession := session.(*kvTableTest.KVTableTestSession)
			go golink.Intsert(ctx, i, ch, totalNumber, &wg, request, kvtabletestSession)
		case model.FormParallelOk:
			parallelOkSession := session.(*parallelOk.ParallelOkSession)
			go golink.Transfer(ctx, i, ch, totalNumber, &wg, request, parallelOkSession)
		default:
			// Unsupported Media Type
			wg.Done()
		}
	}
	// Wait for all the data to be sent
	wg.Wait()
	//Delay 1 millisecond to ensure that all data is processed
	time.Sleep(1 * time.Millisecond)
	close(ch)
	// The data processing is complete
	wgReceiving.Wait()
}
