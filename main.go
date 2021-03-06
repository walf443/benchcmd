package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"os/exec"
	"time"
)

var times *int = flag.Int("n", 10, "number of times to run")
var summary *bool = flag.Bool("summary", false, "only output summary result")
var shell *string = flag.String("shell", "sh", "shell name to execute")

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
	fmt.Printf("run \"%s\"\n", cmd)
	for i := 0; i < *times; i++ {
		out, err := exec.Command(*shell, "-c", shellCmd).CombinedOutput()
		now := time.Now()
		offset := now.Sub(prevTime)
		offsets[i] = int64(offset)
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(out))
			os.Exit(1)
		}
		if !*summary {
			fmt.Println(offset)
		}
		prevTime = time.Now()
	}
	report(int64(*times), offsets)
}

func report(times int64, offsets []int64) {
	fmt.Printf("count:\t%d times executed\n", times)
	fmt.Printf("avg:\t%s\n", time.Duration(average(times, offsets)))
	fmt.Printf("stdev:\t%s\n", time.Duration(stdev(times, offsets)))
}

func average(times int64, offsets []int64) int64 {
	result := int64(0)
	for _, v := range offsets {
		result += v
	}
	return result / times
}

func stdev(times int64, offsets []int64) float64 {
	result := int64(0)

	avg := average(times, offsets)
	for _, v := range offsets {
		diff := v - avg
		result += diff * diff
	}
	return math.Sqrt(float64(result / (times + int64(1))))
}
