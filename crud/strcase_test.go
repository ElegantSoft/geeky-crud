package crud

import (
	"github.com/go-playground/assert/v2"
	"github.com/iancoleman/strcase"
	"strings"
	"testing"
)

func TestStrCase(t *testing.T) {
	str := "posts.comments.user"
	expectedResult := "Posts.Comments.User"

	strSlice := strings.Split(str, ".")
	var camelCaseStringsSlice []string
	for _, v := range strSlice {
		camelCaseStringsSlice = append(camelCaseStringsSlice, strcase.ToCamel(v))
	}
	result := strings.Join(camelCaseStringsSlice, ".")
	assert.Equal(t, result, expectedResult)
}
