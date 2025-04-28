package common

import "encoding/json"

func TypeComverter[T any](data any) (*T, error) {
	var result T
	jsondata, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsondata, &result)
	if err != nil {
		return nil ,err
	}
	return &result, nil
}