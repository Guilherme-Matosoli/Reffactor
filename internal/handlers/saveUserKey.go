package handlers

import (
	"encoding/json"
	"fmt"
	"os"
)

type informations struct {
	ApiKey string `json:"apiKey"`
}

func SaveUserKey(key string) {
	userInfo := &informations{
		ApiKey: key,
	}
	info, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Print("errrrroooooo44")
		fmt.Print(err)
	}

	error := os.WriteFile("./config/info.json", []byte(info), 0644)
	if error != nil {
		fmt.Print("errrrroooooo44")
		fmt.Print(err)
	}
}
