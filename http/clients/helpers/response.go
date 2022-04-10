package HttpClientHelpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetResponseStruct(resp *http.Response, dataType struct{}) (struct{}, error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dataType, errors.New(err.Error())
	}

	err = json.Unmarshal(b, &dataType)
	if err != nil {
		return dataType, errors.New(err.Error())
	}

	return dataType, nil
}
