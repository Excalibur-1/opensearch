package opensearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewQueryBuilderWith(t *testing.T) {
	q := NewQueryBuilder()

	assert.Equal(t, "query=content:'abc'", q.String())
}

func TestNewQueryBuilderWith_Zero(t *testing.T) {
	q := NewQueryBuilder()

	exp := ""
	assert.Equal(t, exp, q.query)

	// and now add another
	q.And("user_id", "123", true)
	exp = "query=user_id:'123'"
	assert.Equal(t, exp, q.query)
}

func TestQueryBuilder_AND(t *testing.T) {
	q := NewQueryBuilder()
	q.And("user_id", "123456", true)

	exp := "query=content:'abc' AND user_id:'123456'"
	assert.Equal(t, exp, q.query)
}

func TestQueryBuilder_AND_Empty(t *testing.T) {
	q := NewQueryBuilder()
	q.And("user_id", "", true)

	exp := "query=content:'abc'"
	assert.Equal(t, exp, q.query)
}
func TestQueryBuilder_AND_Zero(t *testing.T) {
	q := NewQueryBuilder()
	q.And("user_id", "0", true)

	// 0 should not be ignored (manually ignore in upper stream)
	exp := "query=content:'abc' AND user_id:'0'"
	assert.Equal(t, exp, q.query)
}

func TestNewConfigBuilderWith(t *testing.T) {
	b := NewConfigBuilderLimit(120, 30)

	exp := "config=start:120,hit:30"
	assert.Equal(t, exp, b.config)
}

func TestNewConfigBuilder_EMPTY(t *testing.T) {
	b := NewConfigBuilderLimit(0, 30)

	// zero should not be ignored
	exp := "config=start:0,hit:30"
	assert.Equal(t, exp, b.config)
}

func TestConcat(t *testing.T) {
	var testcases = []struct {
		queries []string

		exp string
	}{
		{[]string{"", ""}, ""},

		{[]string{"query=content:'abc'", "filter=created_at>=12345678"},
			"query=content:'abc' && filter=created_at>=12345678"},

		{[]string{"query=content:'abc'", "sort=-RANK;-created_at"},
			"query=content:'abc' && sort=-RANK;-created_at"},

		{[]string{"query=content:'abc' AND status:'2'", "", "config=start:30, hit:10"},
			"query=content:'abc' AND status:'2' && config=start:30, hit:10"},
	}

	for _, c := range testcases {
		got := Concat(c.queries...)
		assert.Equal(t, c.exp, got)
	}
}
