package main

import (
	"fmt"
	"observer/helpers"

	// import the http local  package
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	c.AddFunc("@every 1m", func() {
		fmt.Println("Every minute")

		res, err := helpers.SendPostJson("http://localhost:8080/api/v1/notify", `{"message":"Every minute"}`)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Response: %s\n", res)
	})

	c.Start()

}
