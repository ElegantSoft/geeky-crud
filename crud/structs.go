package crud

type FindStatement = map[string]interface{}
type Query = map[string]FindStatement

type FindQuery struct {
	Q      Query    `json:"q"`
	Fields []string `json:"fields"`
	Joins  []string `json:"joins"`
}

var OperatorsKeys = struct {
	Equal          string
	Greater        string
	Lesser         string
	NotEqual       string
	Like           string
	In             string
	GreaterOrEqual string
	LesserOrEqual  string
}{
	Equal:          "$eq",
	Greater:        "$gt",
	Lesser:         "$lt",
	NotEqual:       "$not",
	Like:           "$like",
	In:             "$in",
	GreaterOrEqual: "$gte",
	LesserOrEqual:  "$lte",
}

// Operators
var Operators = map[string]string{
	OperatorsKeys.Equal:          "=",
	OperatorsKeys.Greater:        ">",
	OperatorsKeys.Lesser:         "<",
	OperatorsKeys.NotEqual:       "<>",
	OperatorsKeys.Like:           "LIKE",
	OperatorsKeys.In:             "IN",
	OperatorsKeys.GreaterOrEqual: ">=",
	OperatorsKeys.LesserOrEqual:  "<=",
}

type OperatorValue = [2]interface{}
