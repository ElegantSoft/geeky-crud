package crud

import "reflect"

// The operator is first 3 chars
func getOperatorAndValue(query map[string]string) (string, string) {
	keys := reflect.ValueOf(query).MapKeys()
	operatorKey := keys[0].String()
	operator := operators[operatorKey]
	value := query[operatorKey]
	if operatorKey == "$like" {
		value = "%" + value + "%"
	}
	return operator, value
}
