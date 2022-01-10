//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweets chan *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}
		tweets <- tweet
	}
	close(tweets)
}

func consumer(tweets chan *Tweet, done chan struct{}) {
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	done <- struct{}{}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	done := make(chan struct{})
	tweets := make(chan *Tweet)

	// Producer
	go producer(stream, tweets)

	// Consumer
	go consumer(tweets, done)

	<-done

	fmt.Printf("Process took %s\n", time.Since(start))
}
