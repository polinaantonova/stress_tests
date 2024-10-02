package stressTest

import (
	"context"
	"errors"
	"github.com/caio/go-tdigest"
	"log"
	"math/rand"
	"postgres/internal/sources"
	"sync"
	"sync/atomic"
	"time"
)

func StressTest(table sources.PersonSource, ammos []int, countMs time.Duration, allowedErrorNumber int32) error {
	var sentRequestsCounter int32
	var errorsCounter int32
	var mu sync.Mutex

	timeStats, err := tdigest.New()
	if err != nil {
		err = errors.New(err.Error())
		return err
	}

	freqLimitTicker := time.NewTicker(countMs * time.Microsecond)
	defer freqLimitTicker.Stop()

	printStatTicker := time.NewTicker(1 * time.Second)
	defer printStatTicker.Stop()

	programTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	defer wg.Wait()

	errmap := make(map[string]int, 10)
	var errmapMutex sync.Mutex
	defer func() {
		wg.Wait()
		log.Println("====================================")
		for msg, count := range errmap {
			log.Printf("%v\t%v\n", msg, count)
		}
		log.Println("====================================")
	}()

	for {
		select {
		case <-programTimeout.Done():
			log.Println("program timeout reached or error limit exceeded")
			return nil
		case <-freqLimitTicker.C:
			wg.Add(1)

			go func() {
				defer wg.Done()

				ctx, queryCancel := context.WithTimeout(programTimeout, 1*time.Second)
				defer queryCancel()

				ammo := ammos[rand.Intn(len(ammos))]

				start := time.Now()
				err := table.PerformQuery(ammo, ctx)
				atomic.AddInt32(&sentRequestsCounter, 1)
				if err != nil {
					atomic.AddInt32(&errorsCounter, 1)
					errmapMutex.Lock()
					errmap[err.Error()]++
					errmapMutex.Unlock()
					//log.Printf("Query error: %v", err)
					return
				}

				timeSpent := time.Since(start).Milliseconds()

				mu.Lock()
				err = timeStats.Add(float64(timeSpent))
				mu.Unlock()
				if err != nil {
					panic("tdigest not working")
				}
			}()

		case <-printStatTicker.C:
			wg.Add(1)

			go func() {
				defer wg.Done()

				if atomic.LoadInt32(&errorsCounter) > allowedErrorNumber {
					log.Println("Error limit exceeded, stopping test.")
					cancel()
					return
				}
				log.Printf("Requests sent: %v", atomic.SwapInt32(&sentRequestsCounter, 0))
				log.Printf("Errors: %d\n", atomic.SwapInt32(&errorsCounter, 0))
				printPercentiles(timeStats)
			}()
		}
	}
}

func printPercentiles(t *tdigest.TDigest) {
	log.Printf("Percentiles 50/95/99: %v %v %v ms\n", t.Quantile(0.5), t.Quantile(0.95), t.Quantile(0.99))
}
