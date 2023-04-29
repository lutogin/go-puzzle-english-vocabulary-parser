package config

import (
	"go-puzzle-english-vocabulary-parser/common/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug bool `yaml:"isDebug" env-default:"true"`
	Listen  struct {
		Host string `yaml:"host" env:"HOST" env-default:"8181"`
		Port string `yaml:"port" env:"PORT" env-default:"127.0.0.1"`
	} `yaml:"listen"`
	Pe struct {
		BaseAPIPath          string `yaml:"baseApiPath" env:"BASE_API_PATH" env-default:"https://puzzle-english.com"`
		SelectorWords        string `yaml:"selectorWords" env:"SELECTOR_WORDS" env-default:".puzzle-card__word .word-wrapper"`
		SelectorTranslations string `yaml:"selectorTranslations" env:"SELECTOR_TRANSLATIONS" env-default:".puzzle-card__word .dict__video__list-table__word__translate.puzzle-text_fz_14.puzzle_mt_4"`
	}
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() { // do it once. Singleton pattern
		logger := logging.GetLogger()

		logger.Infoln("Read application's config.")
		instance = &Config{}

		if errs := cleanenv.ReadConfig("config.yml", instance); errs != nil {
			help, _ := cleanenv.GetDescription(instance, nil)

			logger.Fatalln(errs)
			logger.Fatalln(help)
		}
	})

	return instance
}
