package helpers

import (
	"fmt"
	"net/http"
	"strings"
)

// send POST JSON request

func SendPostJson(url string, json string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(json))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	fmt.Printf("%s\n", res.Status)
	return "", nil
}
