// Package server 压测启动
package server

import (
	"context"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/contract/kvTableTest"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/contract/parallelOk"
	"sync"
	"time"

	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/model"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/server/golink"
	"github.com/FISCO-BCOS/go-sdk/test/performanceTest/server/statistics"
)

// Dispose 处理函数
func Dispose(ctx context.Context, concurrency, totalNumber uint64, request *model.Request,session interface{}) {
	// 设置接收数据缓存
	ch := make(chan *model.RequestResults, 1000)
	var (
		wg          sync.WaitGroup // 发送数据完成
		wgReceiving sync.WaitGroup // 数据处理完成
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
			// 类型不支持
			wg.Done()
		}
	}
	// 等待所有的数据都发送完成
	wg.Wait()
	// 延时1毫秒 确保数据都处理完成了
	time.Sleep(1 * time.Millisecond)
	close(ch)
	// 数据全部处理完成了
	wgReceiving.Wait()
	return
}
