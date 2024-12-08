package utils

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func VerifyKey(key string) bool {
	client := openai.NewClient(option.WithAPIKey(key))

	ctx := context.Background()
	_, err := client.Models.List(ctx)
	if err != nil {
		fmt.Print(err)
		return false
	}

	return true
}
