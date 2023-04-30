package peclient

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"go-puzzle-english-vocabulary-parser/common/logging"
	"go-puzzle-english-vocabulary-parser/config"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type PeClient struct {
	cookie string
	config *config.Config
	logger *logging.Logger
}

func NewPeClient(cookie string, config *config.Config, logger *logging.Logger) *PeClient {
	return &PeClient{cookie: cookie, config: config, logger: logger}
}

func (pe *PeClient) getBody(page int) string {
	body := url.Values{}
	body.Add("for_dictionary_change", "true")
	body.Add("ajax_action", "ajax_pe_get_next_page_dictionary")
	body.Add("page", fmt.Sprint(page))

	return body.Encode()
}

func getClient() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 15,
	}

	return client
}

func (pe *PeClient) MakeRequest(page int) (string, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/change-my-dictionary", pe.config.Pe.BaseAPIPath),
		strings.NewReader(pe.getBody(page)),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", pe.cookie)

	client := getClient()

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Panic("Error creating gzip reader:", err)
		}
		defer reader.(*gzip.Reader).Close()
	}

	var responseBody struct {
		ListWords string `json:"listWords"`
	}

	err = json.NewDecoder(reader).Decode(&responseBody)
	if err != nil {
		return "", err
	}

	return responseBody.ListWords, nil
}
