package crud

import "reflect"

// get operator value slice
func getOperatorAndValue(query map[string]interface{}) []OperatorValue {
	keys := reflect.ValueOf(query).MapKeys()
	var operatorValueSlice []OperatorValue

	for _, operatorKey := range keys {
		operator := Operators[operatorKey.String()]
		value := query[operatorKey.String()]
		if operatorKey.String() == "$like" {
			value = "%" + value.(string) + "%"
		}
		operatorValueSlice = append(operatorValueSlice, OperatorValue{operator, value})
	}

	return operatorValueSlice
}
