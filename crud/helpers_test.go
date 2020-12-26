package crud

import "testing"

func TestHelpers_GetOperatorValue(t *testing.T) {
	query := map[string]string{
		"$eq": "test1",
	}

	operator, value := getOperatorAndValue(query)

	if operator != "=" || value != "test1" {
		t.Fatal("error when serialize query value")
	}

}
