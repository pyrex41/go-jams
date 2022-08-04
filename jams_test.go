package jams

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestJams(t *testing.T) {
	names := []string{
		"bare1",
		"bare2",
		"bareAll",
		"double-quote",
		"emptyquotes",
		"example",
		"nested_example",
		"one-empty-quote",
		"quotes-never-fail",
		"str",
		"trailing-whitespaces",
		"wethpack",
	}

	for _, name := range names {

		jams_name := "test/pass/" + name + ".jams"
		jams_bytes, jamerr := os.ReadFile(jams_name)
		if jamerr != nil {
			t.Log("Failed to read JAMS file ", jamerr)
			t.Fail()
		}
		json_name := "test/pass/" + name + ".json"
		json_bytes, jsonerr := os.ReadFile(json_name)
		if jsonerr != nil {
			t.Log("Failed to read JSON file ", jsonerr)
			t.Fail()
		}
		jams_result := Parse(jams_bytes)
		var json_result interface{}
		json.Unmarshal(json_bytes, &json_result)

		if fmt.Sprint(jams_result) != fmt.Sprint(json_result) {
			t.Log("FAIL -- JAMS != JSON -- ", name)
			t.Fail()

			fmt.Println("**FAILED** ---> ", name)
			fmt.Println("JAMS:")
			fmt.Println(jams_result)
			fmt.Println("JSON:")
			fmt.Println(json_result)
			fmt.Print("\n-------------------------\n")
			fmt.Print("\n\n\n")
		}
	}

}
