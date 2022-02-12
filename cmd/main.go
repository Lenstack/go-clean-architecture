package main

import (
	"github.com/Lenstack/clean-architecture/internal/infrastructure"
)

func main() {
	logger := infrastructure.NewLogger()
	infrastructure.Load(logger)
	mongo := infrastructure.NewMongo(logger)
	infrastructure.Dispatch(logger, mongo)
}
