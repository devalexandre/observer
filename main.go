package main

import (
	"fmt"
	"observer/http" // import the http local  package

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	c.AddFunc("@every 1m", func() {
		fmt.Println("Every minute")
		http.SendNotification()
	})

	c.Start()

}
