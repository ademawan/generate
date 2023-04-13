package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/forgoer/openssl"
)

func main() {
	input := []string{}
	for i := 0; i < 100000; i++ {
		input = append(input, "testingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtestingtesting")
	}

	WithGoroutine(input)
	NoGoroutine(input)
	memoryConsume()
	example1()

}

func WithGoroutine(input []string) {
	start := time.Now()

	var err error
	var wg sync.WaitGroup

	for _, data := range input {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			err = EnryptDataV2(data)
			if err != nil {
				log.Print("Failed Encrypt data fullname", err)
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()
	fmt.Println("success")
	elapsed := time.Since(start)
	log.Printf("Duration Goroutine :%s second", elapsed)
}
func NoGoroutine(input []string) {
	start := time.Now()

	var err error
	for _, data := range input {
		err = EnryptDataV2(data)
		if err != nil {
			log.Print("Failed Encrypt data fullname", err)
		}
	}
	fmt.Println("success")
	elapsed := time.Since(start)
	log.Printf("Duration NoGoroutine :%s second", elapsed)

}

func EnryptDataV2(input string) error {

	keyEncrypt := []byte(os.Getenv("KEY_ENCRYPT"))
	value := []byte(input)

	h := md5.New()
	h.Write(keyEncrypt)
	key := []byte(string(h.Sum(nil)))

	_, err := openssl.AesECBEncrypt(value, key, openssl.PKCS7_PADDING)
	if err != nil {
		return err
	}
	// en := base64.StdEncoding.EncodeToString(dst)

	return nil

}

func memoryConsume() {
	memConsumed := func() uint64 {

		runtime.GC()

		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys

	}

	var c <-chan interface{}

	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c
	} //1

	const numGoroutines = 1e4 //2
	wg.Add(numGoroutines)
	before := memConsumed() //3

	for i := numGoroutines; i > 0; i-- {
		go noop()
	}

	after := memConsumed() //4

	fmt.Printf("%.3fkb \n", float64(after-before)/numGoroutines/1000)

}

func example1() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) { //1
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1) //2
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup

		wg.Add(count + 1)
		beginTestTime := time.Now()

		go producer(&wg, mutex)

		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}

		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Reader\tRWmutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}

}
