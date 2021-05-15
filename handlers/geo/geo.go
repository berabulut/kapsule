package geo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {
	CountryCode string `json:"country_code3"`
}

func GetCountryCode(IP string) (string, error) {

	url := fmt.Sprintf("https://get.geojs.io/v1/ip/geo/%s.json", IP)
	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return "", err
	}

	resp := response{}
	err = json.Unmarshal(body, &resp)

	if err != nil {
		return "", err
	}

	return resp.CountryCode, nil
}
