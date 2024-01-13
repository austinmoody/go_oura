package go_oura

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type PersonalInfo struct {
	ID     string  `json:"id"`
	Age    int     `json:"age"`
	Height float32 `json:"height"`
	Weight float32 `json:"weight"`
	Sex    string  `json:"biological_sex"`
	Email  string  `json:"email"`
}

type personalInfoBase PersonalInfo

func (pi *PersonalInfo) UnmarshalJSON(data []byte) error {
	var rawMap map[string]json.RawMessage
	err := json.Unmarshal(data, &rawMap)
	if err != nil {
		return err
	}

	t := reflect.TypeOf(*pi)
	requiredFields := make([]string, 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		requiredFields = append(requiredFields, jsonTag)
	}

	for _, field := range requiredFields {
		if _, ok := rawMap[field]; !ok {
			return fmt.Errorf("required field %s not found", field)
		}
	}

	var personalInfo personalInfoBase
	err = json.Unmarshal(data, &personalInfo)
	if err != nil {
		return err
	}

	*pi = PersonalInfo(personalInfo)
	return nil

}

func (c *Client) GetPersonalInfo() (PersonalInfo, *OuraError) {

	apiResponse, ouraError := c.Getter(PersonalInfoUrl, nil)

	if ouraError != nil {
		return PersonalInfo{},
			ouraError
	}

	var personalInfo PersonalInfo
	err := json.Unmarshal(*apiResponse, &personalInfo)
	if err != nil {
		return PersonalInfo{},
			&OuraError{
				Code:    0,
				Message: fmt.Sprintf("failed to process response body with error: %v", err),
			}
	}

	return personalInfo, nil
}