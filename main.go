package main

import (
	"go-puzzle-english-vocabulary-parser/common/logging"
	"go-puzzle-english-vocabulary-parser/config"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()

	Start(cfg, logger)
}
