package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Event struct {
	VideoID string                 `json:"video_id"`
	UserID  string                 `json:"user_id"`
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

func main() {

	videos := []string{
		"7108d562-2f5c-47f5-9a02-56da5e5438f8",
		"9e1861e2-1b00-4973-848f-622fe905b842",
		"b588a489-5a2d-4fc6-a142-138728c1989b",
		"bec1b03d-f737-42f8-8f8d-360798ea1a2e",
		"c5068cbf-76c8-4a09-8b7b-13ff5b58d679",
		"18ccfa6a-387f-4900-8fbd-0462678c00ec",
	}
	user := []string{
		"057e6539-733b-47f9-a034-b72a2c00c291",
		"09976c26-5d9b-499e-93f9-0dbaadea7dbd",
		"4b81bf9a-cc01-471b-bda0-46318cf74d55",
	}

	events := []string{"VIEW", "COMMENT", "SHARE", "REACT"}

	for {

		randVideo := rand.Intn(len(videos))
		randUser := rand.Intn(len(user))
		randEv := rand.Intn(len(events))

		payload := Event{
			VideoID: videos[randVideo],
			UserID:  user[randUser],
			Type:    events[randEv],
			Payload: map[string]interface{}{
				"name": "test",
			},
		}
		body, _ := json.Marshal(payload)
		resp, err := http.Post(fmt.Sprintf("http://app:8080/api/v1/ranking/event/%v", payload.VideoID), "application/json", bytes.NewBuffer(body))
		if err != nil {
			fmt.Printf("error: %v\n", err)
		} else {
			fmt.Printf("event: %+v (%v)\n", payload, resp.Status)
			resp.Body.Close()
		}
	}

}
