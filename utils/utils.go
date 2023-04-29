package utils

import (
	"bufio"
	"os"
)

func GetPrompt() (string, error) {
	in := bufio.NewReader(os.Stdin)
	prompt, err := in.ReadString('\n')
	if err != nil {
		return "", err
	}
	prompt = prompt[:len(prompt)-1] // remove last symbol (end of the string)

	return prompt, nil
}
