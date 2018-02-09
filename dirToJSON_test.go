package dirToJSON

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var expected file
	var actual file
	expectedJSON, err := ioutil.ReadFile("expected.json")
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(expectedJSON, &expected)
	if err != nil {
		t.Fatal(err)
	}

	res, err := Parse("test")
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(res, &actual)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal(err)
	}
}
