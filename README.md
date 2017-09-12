go-mattermost
=============

Simple Mattermost client for posting messages into channels using incoming webhooks.

```go
package main

import (
	"fmt"
	"log"

	mattermost "github.com/dg-i/go-mattermost"
)

func main() {
	client := mattermost.NewClient(
		"https://chat.example.net/hooks/IiFtFtWf9WyszufkeJS3vv4v75",
		"MyBot",
		"town-square")

	err := client.SendSimpleMessage(fmt.Sprintf("Hello World! :tada:"))
	if err != nil {
		log.Fatalf("Somthing went wrong: %v", err)
	}

}
```

