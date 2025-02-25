package main

import (
	"time"
)

func processMessages(messages []string) []string {
	// ?
	processedMsgs := []string{}
	messagesChan := make(chan string, len(messages))
	go func() {
		for _, x := range messages {
			messagesChan <- process(x)
		}

		close(messagesChan)
	}()

	for {
		val, ok := <-messagesChan

		if !ok {
			break
		}

		processedMsgs = append(processedMsgs, val)
	}
	return processedMsgs
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
