package util

import (
	"encoding/json"
	"reflect"
	"strings"
	"tablero/webserver/validation/errors"
)

func ToJSONString(value interface{}) (string, error) {
	bytes, error := json.Marshal(value)

	return string(bytes), error
}

func ToJSON(value string) (interface{}, error) {
	var jsonResult interface{}

	decoder := json.NewDecoder(strings.NewReader(value))
	decoder.UseNumber()

	if error := decoder.Decode(&jsonResult); error != nil {
		return nil, error

	} else {
		return jsonResult, nil
	}
}

func FromJSONTo(value string, instance interface{}) error {
	return json.Unmarshal([]byte(value), instance)
}

func SafeBodyToJson(body []byte) (map[string]interface{}, error) {
	jsonRaw, err := ToJSON(string(body))
	if err != nil {
		err := errors.ValidationApiError("validation_body", "Error parsing json from body.")
		return nil, err
	}
	return jsonRaw.(map[string]interface{}), nil
}

func SafeString(str interface{}, message string) (string, error) {
	if reflect.TypeOf(str) != reflect.TypeOf("string") || len(str.(string)) == 0 {
		err := errors.ValidationApiError("validation_string", message)
		return "", err
	}
	return str.(string), nil
}
