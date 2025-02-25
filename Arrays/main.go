package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// ?
	for i, x := range messages {
		val := tagger(x)
		messages[i].tags = val
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	// ?
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}

	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}

func main() {
	messages := []sms{
		{id: "001", content: "Urgent! Last chance to see!"},
		{id: "002", content: "Big sale on all items!"},
		// Additional messages...
	}
	taggedMessages := tagMessages(messages, tagger)

	fmt.Println(taggedMessages[0].tags)
}
