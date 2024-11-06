package internal

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetJokes(argType, argLang, argBlacklistFlags, argRange, argAmount string) (Response, error) {
	err := validateArgs(argType, argLang, argBlacklistFlags, argRange, argAmount)
	if err != nil {
		return Response{}, err
	}

	baseUrl := "https://v2.jokeapi.dev/joke/"
	params := url.Values{}

	params.Add("type", "twopart")
	params.Add("lang", argLang)
	params.Add("amount", argAmount)

	if argBlacklistFlags != "" {
		params.Add("blacklistFlags", argBlacklistFlags)
	}

	if argRange != "" {
		params.Add("idRange", argRange)
	}

	url := baseUrl + argType + "?" + params.Encode()

	res, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("failed to get jokes: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	result, err := FormatBody(body)
	if err != nil {
		return Response{}, err
	}

	return result, nil
}

func validateArgs(argType, argLang, argBlacklistFlags, argRange, argAmount string) error {
	switch argType {
	case "Any", "Programming", "Misc", "Dark", "Pun", "Spooky", "Christmas":
		//valid argument
	default:
		return errors.New("invalid type format")
	}

	switch argLang {
	case "en", "es", "fr", "pt", "de", "cs":
		//valid argument
	default:
		return errors.New("invalid language format")
	}

	flags := strings.Split(argBlacklistFlags, ",")
	for _, flag := range flags {
		switch flag {
		case "", "nsfw", "programming", "religious", "political", "racist", "sexist", "explicit":
			//valid argument
		default:
			return errors.New("invalid blacklist flag format")
		}
	}

	_, err := strconv.Atoi(argAmount)
	if err != nil {
		return errors.New("amount must be an integer")
	}

	ranges := strings.Split(argRange, "-")
	if argRange != "" && len(ranges) != 2 {
		return errors.New("range format must be start-end")
	}

	for _, r := range ranges {
		if r != "" {
			_, err := strconv.Atoi(r)
			if err != nil {
				return errors.New("range values must be integers")
			}
		}
	}

	return nil
}
