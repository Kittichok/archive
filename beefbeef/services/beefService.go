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
		"bresaola", "loin", "bacon", "ham", "belly", "shoulder", "ribs", "tenderloin", "chops", "ribeye",
		"brisket", "flank", "chuck"}

	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := s.httpClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	txt := strings.ToLower(string(body))

	res := map[string]int{}
	for _, t := range beefType {
		res[t] = strings.Count(txt, strings.ToLower(t))
	}
	return res
}
