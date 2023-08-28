package config

import (
	"go-puzzle-english-vocabulary-parser/common/logging"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug bool `yaml:"isDebug" env-default:"true"`
	Pe      struct {
		BaseAPIPath          string `yaml:"baseApiPath" env:"BASE_API_PATH" env-default:"https://puzzle-english.com"`
		SelectorWords        string `yaml:"selectorWords" env:"SELECTOR_WORDS" env-default:".puzzle-card__word .word-wrapper"`
		SelectorTranslations string `yaml:"selectorTranslations" env:"SELECTOR_TRANSLATIONS" env-default:".puzzle-card__word .dict__video__list-table__word__translate.puzzle-text_fz_14.puzzle_mt_4"`
		SelectorSentencesEng string `yaml:"selectorSentencesEng" env:"SELECTOR_SENTENCES_ENG" env-default:".balloon-row__wrapper .dictionary-phrase__eng"`
		WordsPerPage         int    `yaml:"wordsPerPage" env:"WORDS_PER_PAGE" env-default:"100"`
	}
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig(logger *logging.Logger) *Config {
	once.Do(func() { // do it once. Singleton pattern
		logger := logging.GetLogger()

		logger.Infoln("Read application's config.")
		instance = &Config{}

		if _, err := os.Stat("config.yml"); os.IsNotExist(err) {
			logger.Infoln("Config file not found, will be set default values.")
		} else {
			if errs := cleanenv.ReadConfig("config.yml", instance); errs != nil {
				help, _ := cleanenv.GetDescription(instance, nil)

				logger.Fatalln(errs)
				logger.Fatalln(help)
			}
		}
	})

	return instance
}
