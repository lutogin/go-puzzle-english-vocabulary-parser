package parser

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-puzzle-english-vocabulary-parser/common/logging"
	"go-puzzle-english-vocabulary-parser/config"
	peclient "go-puzzle-english-vocabulary-parser/pe-cleint"
	"os"
	"strings"
)

func Parser(cfg *config.Config, client *peclient.PeClient, logger *logging.Logger) {
	file, err := os.Create("vocabulary.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	hasMore := true
	page := 0
	processedWords := 0

	for hasMore {
		htmlContent, err := client.MakeRequest(page)
		if err != nil {
			logger.Traceln("Error getting HTML")
			panic(err)
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
		if err != nil {
			logger.Panic("Error parsing HTML:", err)
			panic(err)
		}

		var words []string
		doc.Find(cfg.Pe.SelectorWords).Each(func(i int, s *goquery.Selection) {
			words = append(words, strings.TrimSpace(s.Text()))
		})

		if len(words) < cfg.Pe.WordsPerPage {
			hasMore = false
		}

		var translations []string
		doc.Find(cfg.Pe.SelectorTranslations).Each(func(i int, s *goquery.Selection) {
			translations = append(translations, strings.TrimSpace(s.Text()))
		})

		for i, word := range words {
			err := writer.Write([]string{word, translations[i]})
			if err != nil {
				logger.Fatal(err)
				panic(err)
			}
		}
		writer.Flush()

		processedWords += len(words)
		page++
		//logger.Infof("\rProcessed %d words \n", processedWords)
		fmt.Printf("\rProcessed %d words...", processedWords)
	}

	logger.Infof("Were success exported %d words.", processedWords)
}
