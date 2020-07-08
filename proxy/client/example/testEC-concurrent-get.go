package main

import (
	"math/rand"
	"fmt"
	"github.com/neboduus/infinicache/proxy/client"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	requestsNumber, size, addrList := client.GetArgs(os.Args)
	var wg sync.WaitGroup
	for i:=0; i<3; i++{
		wg.Add(1)
		go test2(i, &wg, addrList, requestsNumber, size)
	}
	wg.Wait()
}

func test5(i int, wg *sync.WaitGroup, addrList string, reqNumber int, size int){
	defer wg.Done()

	// parse server address
	addrArr := strings.Split(addrList, ",")

	// initial new ecRedis client
	cli := client.NewClient(10, 2, 32, 3)

	// start dial and PUT/GET
	cli.Dial(addrArr)
	val := make([]byte, size)
	rand.Read(val)
	// ToDo: Replace with a channel
	var setStats []float64

	for k:=0; k<reqNumber; k++{
		key := fmt.Sprintf("k%d", k)

		var s float64 = 0
		for l:=0; l<9; l++ {
			if _, stats, ok := cli.EcSet(key, val); !ok {
				log.Println("Failed to SET ", key)
			}else{
				log.Println("Successfull SET ", key)
				s += stats
			}
		}
		if s!=0{
			setStats = append(setStats, s)
		}
	}

	gMin, gMax, gAvg, gSd, gPercentiles := cli.GetStats(setStats)
	log.Println("SET stats ", gMin, gMax, gAvg, gSd, gPercentiles)
	return
}