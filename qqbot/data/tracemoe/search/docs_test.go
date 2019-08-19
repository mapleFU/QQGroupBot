package search

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_SearchResult(t *testing.T) {
	fileBytes, err := ioutil.ReadFile("test.json")
	if err != nil {
		t.Error(err.Error())
	}
	jsonData := string(fileBytes)
	fmt.Println(jsonData)
}
