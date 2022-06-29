package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	POST = "POST"
	GET  = "GET"
)

// HttpError http error
var HttpError = errors.New("http do error")

func transportError(status int) error {
	return errors.New(fmt.Sprintf("transport failed code is: %d", status))
}

func Get(url string, body io.ReadCloser, head *http.Header, param map[string]string, result interface{}) error {
	client := http.Client{}

	url = url + "?"
	for key, value := range param {
		url = url + key + "=" + value + "&"
	}
	request, err := http.NewRequest(GET, url[:len(url)-1], body)
	if err != nil {
		return HttpError
	}
	if head != nil {
		request.Header = *head
	}
	do, err := client.Do(request)
	if err != nil {
		return HttpError
	} else if do.StatusCode != 200 {
		return transportError(do.StatusCode)
	} else {
		res, err := ioutil.ReadAll(do.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(res, &result)
		if err != nil {
			return err
		}
		return nil
	}
}

func Post(url string, body io.ReadCloser, head map[string]string) (interface{}, error) {
	client := &http.Client{}
	request, err := http.NewRequest(POST, url, body)
	if err != nil {
		return nil, HttpError
	}
	if head != nil {
		for key, value := range head {
			request.Header.Add(key, value)
		}
	}
	do, err := client.Do(request)
	if err != nil {
		return nil, HttpError
	} else if do.StatusCode != 200 {
		return nil, transportError(do.StatusCode)
	} else {
		res, err := ioutil.ReadAll(do.Body)
		if err != nil {
			return nil, err
		}
		var resMap map[string]interface{}
		err = json.Unmarshal(res, &resMap)
		if err != nil {
			return nil, err
		}
		return resMap["data"], err
	}
}
