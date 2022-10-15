package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/DostonAkhmedov/async-logger/pkg/alog"
)

var args = os.Args[1:]

func main() {
	cntThread, cntMsg := readFromCli(), readFromCli()

	cfg, err := NewConfig()
	if err != nil {
		fmt.Printf("read config error: %s", err.Error())
	}

	log := alog.New(cfg.ALog)
	log.Start()
	defer log.Stop()

	var wg sync.WaitGroup

	wg.Add(cntThread)
	for i := 0; i < cntThread; i++ {
		i := i
		go func() {
			defer wg.Done()

			for j := 0; j < cntMsg; j++ {
				log.Info(fmt.Sprintf("msg #%d-%d", i, j))
			}
		}()
	}

	wg.Wait()
}

func readFromCli() int {
	if len(args) > 0 {
		cnt, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("string parse error: %s", err.Error())
		}
		args = args[1:]

		return cnt
	}

	return 1
}
