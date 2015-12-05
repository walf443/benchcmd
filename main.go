package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var times *int = flag.Int("num", 10, "times to run")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		os.Exit(1)
	}
	cmd := args[0]
	shellCmd := cmd + " > /dev/null"
	prevTime := time.Now()
	offsets := make([]int64, *times)
	for i := 0; i < *times; i++ {
		out, err := exec.Command("sh", "-c", shellCmd).CombinedOutput()
		now := time.Now()
		offset := now.Sub(prevTime)
		offsets[i] = int64(offset)
		if err != nil {
			fmt.Println(err)
			fmt.Println(out)
			os.Exit(1)
		}
		fmt.Println(offset)
		prevTime = time.Now()
	}
	fmt.Printf("run \"%s\"\n", cmd)
	report(int64(*times), offsets)
}

func report(times int64, offsets []int64) {
	fmt.Printf("count:\t%d times executed\n", times)
	fmt.Printf("avg:\t%s\n", time.Duration(average(times, offsets)))
}

func average(times int64, offsets []int64) int64 {
	result := int64(0)
	for _, v := range offsets {
		result += v
	}
	return result / times
}