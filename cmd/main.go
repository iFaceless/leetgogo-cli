package main

import (
	"context"
	"github.com/iFaceless/leetgogo-cli/pkg/command"
	"log"
)

func main() {
	ctx := context.Background()
	err := command.Execute(ctx)
	if err != nil {
		log.Fatalf("failed to execute command: %s", err)
	}
}
