package crud

type findQuery struct {
	q      map[string]map[string]string
	fields []string
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