package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

const (
	workerCount = 10
	pageSize    = 100
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	client := twilio.NewRestClient()

	params := &openapi.ListCallParams{}
	params.SetPageSize(pageSize)
	params.SetStatus("queued")

	ch, _ := client.Api.StreamCall(params)

	wg := &sync.WaitGroup{}
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go work(ctx, wg, i, client, ch)
	}

	wg.Wait()
}

func work(ctx context.Context, wg *sync.WaitGroup, i int, client *twilio.RestClient, ch <-chan openapi.ApiV2010Call) {
	defer wg.Done()

	params := &openapi.UpdateCallParams{}
	params.SetStatus("canceled")

	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d: %s", i, ctx.Err())
			return

		case record, ok := <-ch:
			if !ok {
				return
			}

			resp, err := client.Api.UpdateCall(*record.Sid, params)
			if err != nil {
				log.Printf("client.Api.UpdateCall(%s) failed", *record.Sid)
				return
			}
			fmt.Printf("worker %d: %s\n", i, *resp.Sid)
		}
	}
}
