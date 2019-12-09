package commons

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

//ErrorMsgs implements the field and motive of error
// will be used on every return of error
type ErrorMsgs struct {
	Field  string `json:"field"`
	Motive string `json:"motive"`
}

// GetJSONTestFiles to convert any json over an interface
// based on the json path informed.
func GetJSONTestFiles(intfc interface{}, path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	json.Unmarshal(byteValue, &intfc)
	return nil
}

// CreateRequest to be executed on test based on the body set on an interface
// and on the path to be requested.
func CreateRequest(intfc interface{}, requestType, path string) (*http.Request, error) {
	if intfc != nil {
		// if a body is provided we marshal it on a json and send to the request
		b, err := json.Marshal(intfc)
		if err != nil {
			return nil, err
		}
		bodyBuffer := bytes.NewBuffer(b)

		req, err := http.NewRequest(requestType, path, bodyBuffer)
		if err != nil {
			return nil, err
		}
		return req, nil
	}

	//body not provided the request is created using the new request function
	req, err := http.NewRequest(requestType, path, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
