package internal

import "encoding/json"

type Joke struct {
	Category string `json:"category"`
	Type     string `json:"type"`
	Setup    string `json:"setup,omitempty"`
	Delivery string `json:"delivery,omitempty"`
	Joke     string `json:"joke,omitempty"`
	Flags    struct {
		NSFW      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Political bool `json:"political"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
		Explicit  bool `json:"explicit"`
	} `json:"flags"`
	ID   int    `json:"id"`
	Safe bool   `json:"safe"`
	Lang string `json:"lang"`
}

type Response struct {
	Error  bool   `json:"error"`
	Amount int    `json:"amount,omitempty"`
	Jokes  []Joke `json:"jokes"`
}

func FormatBody(body []byte) (Response, error) {
	var res Response
	err := json.Unmarshal(body, &res)
	if err != nil {
		return Response{}, err
	}

	if len(res.Jokes) == 0 {
		var joke Joke

		err := json.Unmarshal(body, &joke)
		if err != nil {
			return Response{}, err
		}

		res.Jokes = append(res.Jokes, joke)
		res.Amount = 1
	}

	return res, nil
}
