package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/timliudream/go-test/workerpool/basic"
	"github.com/timliudream/go-test/workerpool/model"
)

func NotPooledWork(allData []model.SimpleData) {
	start := time.Now()
	var wg sync.WaitGroup

	dataCh := make(chan model.SimpleData, 100)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range dataCh {
			wg.Add(1)
			go func(data model.SimpleData) {
				defer wg.Done()
				basic.Process(data)
			}(data)
		}
	}()

	for i, _ := range allData {
		dataCh <- allData[i]
	}

	close(dataCh)
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Took ===============> %s\n", elapsed)
}
