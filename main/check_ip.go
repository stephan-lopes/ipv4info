package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func checkIP(ip string) (string, error) {
	returnDefault := func(err error) (string, error) { return "", err }

	resp, err := http.Post(ip, "", nil)
	if err != nil {
		return returnDefault(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return returnDefault(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return returnDefault(err)
	}

	return fmt.Sprintf("%v", result["country"]), nil
}
