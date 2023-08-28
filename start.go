package main

import (
	"fmt"
	"go-puzzle-english-vocabulary-parser/common/logging"
	"go-puzzle-english-vocabulary-parser/config"
	"go-puzzle-english-vocabulary-parser/parser"
	peclient "go-puzzle-english-vocabulary-parser/pe-cleint"
	"go-puzzle-english-vocabulary-parser/utils"
	"log"
	"slices"
)

var YNAnswers = []string{"y", "Y", "n", "N", ""}

func Start(cfg *config.Config, logger *logging.Logger) {
	logger.Infoln("App is started.")
	fmt.Print("Enter your cookie from puzzle-english: ")

	cookie, err := utils.GetPrompt()
	if err != nil {
		log.Fatal("Error parsing prompt:", err)
	}

	fmt.Print("Do you want to add example of using words (N/y): ")

	addExampleOfUsingWord, err := utils.GetPrompt()
	if err != nil {
		log.Fatal("Error parsing prompt: ", err)
	}
	if !slices.Contains(YNAnswers, addExampleOfUsingWord) {
		log.Fatal("Answer must be only 'Y/N'.")
	}

	addExampleOfUsingWordsBool := addExampleOfUsingWord == string('y') || addExampleOfUsingWord == string('Y')

	client := peclient.NewPeClient(cookie, cfg, logger)
	logger.Infoln("Puzzle english client is initialized.")

	parser.Parser(cfg, client, logger, &parser.Opts{AddExtraExamples: addExampleOfUsingWordsBool})
}
