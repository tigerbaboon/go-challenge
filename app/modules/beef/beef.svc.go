package beef

import (
	beefdto "app/app/modules/beef/dto"
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
)

type BeefService struct {
}

func newService() *BeefService {
	return &BeefService{}
}

func (svc *BeefService) Summary() (*beefdto.BeefSummaryResponse, error) {

	request, err := http.NewRequest("GET", "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to get beef summary")
	}

	var resp bytes.Buffer
	_, err = io.Copy(&resp, response.Body)
	if err != nil {
		return nil, err
	}

	respByte := resp.Bytes()
	text := strings.ToLower(string(respByte))

	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ",", "")
	words := strings.Fields(text)

	targetWords := map[string]int{
		"t-bone":   0,
		"fatback":  0,
		"pastrami": 0,
		"pork":     0,
		"meatloaf": 0,
		"jowl":     0,
		"enim":     0,
		"bresaola": 0,
	}

	for _, word := range words {
		if _, ok := targetWords[word]; ok {
			targetWords[word]++
		}
	}

	data := beefdto.BeefSummaryResponse{
		Beef: targetWords,
	}

	return &data, nil
}
