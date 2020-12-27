package crud

type FindStatement = map[string]interface{}
type Query = map[string]FindStatement

type FindQuery struct {
	Q      Query    `json:"q"`
	Fields []string `json:"fields"`
	Joins  []string `json:"joins"`
}

// Operators
var Operators = map[string]string{
	"$eq":   "=",
	"$gt":   ">",
	"$lt":   "<",
	"$not":  "<>",
	"$like": "LIKE",
}

type OperatorValue = [2]interface{}
