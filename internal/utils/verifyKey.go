package utils

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func VerifyKey(key string) error {
	client := openai.NewClient(option.WithAPIKey(key))

	ctx := context.Background()

	_, err := client.Models.List(ctx)
	if err != nil {
		return err
	}

	return nil
}
