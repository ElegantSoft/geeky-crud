package crud

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestHelpers_GetOperatorValue(t *testing.T) {
	query := map[string]interface{}{
		"$gt": 1,
		"$lt": 5,
	}

	operatorValues := getOperatorAndValue(query)
	assert.Equal(t, operatorValues[0][0], ">")
	assert.Equal(t, operatorValues[0][1], 1)
	assert.Equal(t, operatorValues[1][0], "<")
	assert.Equal(t, operatorValues[1][1], 5)

}

func TestHelpers_GetOperatorValueWithPercentWhenUsingLike(t *testing.T) {
	query := map[string]interface{}{
		"$lt":   5,
		"$like": "test",
	}
	operatorValues := getOperatorAndValue(query)
	assert.Equal(t, operatorValues[1][1], "%test%")

}

func TestStrCase(t *testing.T) {
	tests := map[string]string{
		"posts.comments.user": "Posts.Comments.User",
		"posts":               "Posts",
		"post_comments":       "PostComments",
		"credit_cards":        "CreditCards",
		"users.credit_cards":  "Users.CreditCards",
	}
	for k, v := range tests {
		assert.Equal(t, convertSnakeToGormPascal(k), v)
	}
}
