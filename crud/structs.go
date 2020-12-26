package crud

type findQuery struct {
	q      map[string]map[string]string
	fields []string
}

// Operators
var operators = map[string]string{
	"$eq":   "=",
	"$gt":   ">",
	"$lt":   "<",
	"$not":  "<>",
	"$like": "LIKE",
}
