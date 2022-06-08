package testutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/maxwellgithinji/odktojson/domain"
	"github.com/mitchellh/mapstructure"
)

// ReadFileContents reads the contents of a file
// and returns the contents as a string
// this is used to read the contents of any file when testing
func ReadFileContents(path string) (string, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	return string(fileContent), err
}

// DecodeTestXML decodes the contents of a XML file
func DecodeTestXML(got interface{}) error {
	var mapData map[string]interface{}

	err := mapstructure.Decode(got, &mapData)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		os.Exit(1)
	}

	jsonData, err := json.MarshalIndent(mapData, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(
		"../testUtils/question/1.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		os.Exit(1)
	}
	return nil
}

// DecodeBodyToJson decodes the contents of a JSON file
// and returns the contents as a domain.Body
// this is used to decode the contents the JSON file when testings
func DecodeBodyToJson(path string) (domain.Body, error) {
	body, err := ReadFileContents(path)
	if err != nil {
		return domain.Body{}, err
	}
	responseBody := domain.Body{}
	err = json.Unmarshal([]byte(body), &responseBody)
	if err != nil {
		fmt.Println("Error decoding json file JSON:", err)
		os.Exit(1)
	}
	return responseBody, err
}

// DecodeInstancesToJSON decodes the contents of a JSON file
// and returns the contents as a domain.Instance
// this is used to decode the contents the JSON file when testings
func DecodeInstancesToJSON(path string) (domain.InstanceData, error) {
	instance, err := ReadFileContents(path)
	if err != nil {
		return domain.InstanceData{}, err
	}
	responseInstance := domain.InstanceData{}
	err = json.Unmarshal([]byte(instance), &responseInstance)
	if err != nil {
		fmt.Println("Error decoding json file JSON:", err)
		os.Exit(1)
	}
	return responseInstance, err
}

// DecodeBindingToJSON decodes the contents of a JSON file
// and returns the contents as a domain.BindingData
// this is used to decode the contents the JSON file when testings
func DecodeBindingToJSON(path string) (domain.BindData, error) {
	instance, err := ReadFileContents(path)
	if err != nil {
		return domain.BindData{}, err
	}
	responseBinding := domain.BindData{}
	err = json.Unmarshal([]byte(instance), &responseBinding)
	if err != nil {
		fmt.Println("Error decoding json file JSON:", err)
		os.Exit(1)
	}
	return responseBinding, err
}
