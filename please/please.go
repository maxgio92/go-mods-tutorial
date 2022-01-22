package please

import (
	"fmt"
	"errors"

	"rsc.io/quote"
)

func Please(name string) (string, error) {
	if name == "" {
		return "", errors.New("The name is empty")
	}

	message := fmt.Sprintf("Please %s, %s", name, quote.Go())

	return message, nil
}
