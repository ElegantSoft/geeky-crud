package crud

type findStatement = map[string]interface{}
type query = map[string]findStatement

type findQuery struct {
	q      map[string]findStatement
	fields []string
	join   []string
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
