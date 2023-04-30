package main

import (
	"fmt"
	"go-puzzle-english-vocabulary-parser/common/logging"
	"go-puzzle-english-vocabulary-parser/config"
	"go-puzzle-english-vocabulary-parser/parser"
	peclient "go-puzzle-english-vocabulary-parser/pe-cleint"
	"go-puzzle-english-vocabulary-parser/utils"
	"log"
)

func Start(cfg *config.Config, logger *logging.Logger) {
	logger.Infoln("App is started.")
	fmt.Print("Enter you cookie: ")

	cookie, err := utils.GetPrompt()
	if err != nil {
		log.Fatal("Error parsing prompt:", err)
	}

	client := peclient.NewPeClient(cookie, cfg, logger)
	logger.Infoln("Puzzle english client is initialized.")

	parser.Parser(cfg, client, logger)
}
