package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				timeInt, err := strconv.Atoi(t.Format("030405"))
				if err != nil {
					panic(err)
				}
				var output strings.Builder
				fmt.Fprintf(&output, "%d", timeInt)
				if timeInt%15 == 0 {
					fmt.Fprintf(&output, "%v", " FizzBuzz")
				} else if timeInt%3 == 0 {
					fmt.Fprintf(&output, "%v", " Fizz")
				} else if timeInt%5 == 0 {
					fmt.Fprintf(&output, "%v", " Buzz")
				}

				fmt.Printf("%v\n", output.String())
			}
		}
	}()
	time.Sleep(30 * time.Second)
	ticker.Stop()
	done <- true
}
