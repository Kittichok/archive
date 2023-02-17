package services

import (
	"beefbeef/ports"
	"io/ioutil"
	"net/http"
	"strings"
)

type srv struct {
	httpClient ports.HttpClient
}

func NewBeefService(httpClient ports.HttpClient) srv {
	return srv{
		httpClient: httpClient,
	}
}

func (s *srv) Count() map[string]int {
	beefType := []string{"t-bone",
		"fatback", "pastrami", "pork", "meatloaf", "jowl", "enim",
		"bresaola", "loin", "ribs", "tenderloin", "plate", "ribeye",
		"brisket", "flank", "chuck", "shank"}
	mapBeefType := map[string]bool{}

	for _, t := range beefType {
		mapBeefType[t] = true
	}

	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := s.httpClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	txt := strings.ToLower(string(body))
	res := map[string]int{}
	for _, word := range strings.Split(txt, " ") {
		word = strings.ReplaceAll(word, ".", "")
		word = strings.ReplaceAll(word, ",", "")
		if mapBeefType[word] {
			res[word] = res[word] + 1
		}
	}
	return res
}
