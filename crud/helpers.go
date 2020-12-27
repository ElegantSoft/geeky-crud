package crud

import (
	"github.com/iancoleman/strcase"
	"reflect"
	"strings"
)

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

func convertSnakeToGormPascal(snake string) string {
	strSlice := strings.Split(snake, ".")
	var camelCaseStringsSlice []string
	for _, v := range strSlice {
		camelCaseStringsSlice = append(camelCaseStringsSlice, strcase.ToCamel(v))
	}
	result := strings.Join(camelCaseStringsSlice, ".")
	return result
}
