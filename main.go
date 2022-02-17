package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gitlab.com/utopiops-water/test-image/tools"
)

func main() {

	httpHelper := tools.NewHttpHelper(tools.NewHttpClient())

	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				out, err, statusCode, _ := httpHelper.HttpRequest(http.MethodGet, "https://ipinfo.io/json", nil, nil, 0)
				if err != nil || statusCode != 200 {
					log.Println("status code:", statusCode)
					if err != nil {
						log.Println(err.Error())
					}
				} else {
					resp := map[string]interface{}{}
					json.Unmarshal(out, &resp)
					json.MarshalIndent(resp, "", "    ")
					log.Printf("\n" + string(out) + "\n")
					log.Println(time.Now())
					log.Println("--------------------------------------------------------------")
				}
			case <-quit:
				log.Println("Quit")
				ticker.Stop()
				return
			}
		}
	}()
	select {}
}
