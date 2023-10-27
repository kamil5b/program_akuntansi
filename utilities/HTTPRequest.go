package utilities

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func HTTPRequest(request, url string, header http.Header, body, response any, myclient ...http.Client) error {
	if len(myclient) > 1 {
		return errors.New("client cannot more than 1")
	}
	client := http.Client{}
	if len(myclient) > 0 {
		client = myclient[0]
	}
	var bodyReq io.Reader = nil
	if body != nil {
		bodyJSON, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReq = bytes.NewBuffer(bodyJSON)
	}

	req, err := http.NewRequest(request, url, bodyReq)

	if err != nil {
		return err
	}
	req.Header = header

	res, err := client.Do(req)

	if err != nil {
		return err
	}
	//STATUS : 400
	if !(strings.EqualFold(res.Status, "200 OK") || strings.EqualFold(res.Status, "200") || strings.EqualFold(res.Status, "OK")) {
		return errors.New(res.Status)
	}
	defer res.Body.Close()
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(responseData, &response)
}
