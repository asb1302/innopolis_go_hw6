package main

import (
	"fmt"
	"maps"
	"sync"
	"time"
)

func RunProcessor(wg *sync.WaitGroup, mu *sync.Mutex, prices []map[string]float64) {
	go func() {
		defer wg.Done()
		mu.Lock()
		for _, price := range prices {
			for key, value := range price {
				price[key] = value + 1
			}
			fmt.Println(price)
		}
		mu.Unlock()
	}()
}

func RunWriter() <-chan map[string]float64 {
	var prices = make(chan map[string]float64)
	go func() {
		var currentPrice = map[string]float64{
			"inst1": 1.1,
			"inst2": 2.1,
			"inst3": 3.1,
			"inst4": 4.1,
		}
		for i := 1; i < 5; i++ {
			// Создаем новую мапу, чтобы избежать гонок данных
			for key, value := range currentPrice {
				currentPrice[key] = value + 1
			}
			prices <- maps.Clone(currentPrice)

			time.Sleep(time.Second)
		}
		close(prices)
	}()
	return prices
}
func main() {
	p := RunWriter()
	var prices []map[string]float64

	for price := range p {
		prices = append(prices, price)
	}

	var mutex sync.Mutex

	for _, price := range prices {
		fmt.Println(price)
	}

	var wg sync.WaitGroup
	wg.Add(3)

	RunProcessor(&wg, &mutex, prices)
	RunProcessor(&wg, &mutex, prices)
	RunProcessor(&wg, &mutex, prices)

	wg.Wait()
}
