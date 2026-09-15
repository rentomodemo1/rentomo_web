package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetJSON(client *http.Client, url string, out any) error {
	response, err := client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP status %d", response.StatusCode)
	}
	return json.NewDecoder(response.Body).Decode(out)
}
