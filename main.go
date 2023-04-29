package main

import (
	"fmt"
	"go-puzzle-english-vocabulary-parser/common/logging"
	"go-puzzle-english-vocabulary-parser/config"
	peclient "go-puzzle-english-vocabulary-parser/pe-cleint"
	"go-puzzle-english-vocabulary-parser/utils"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()

	fmt.Print("Enter you cookie: ")

	cookie, err := utils.GetPrompt()
	if err != nil {
		log.Fatal("Error parsing prompt:", err)
	}

	client := peclient.NewPeClient(cookie, cfg, logger)
	htmlContent, err := client.MakeRequest(1)

	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	doc.Find(cfg.Pe.SelectorWords).Each(func(i int, s *goquery.Selection) {
		fmt.Println("Word: ", s.Text())
	})
}
