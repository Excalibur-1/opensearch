package opensearch

import "fmt"

// QueryBuilder builds query string
//
// Doc: https://help.aliyun.com/document_detail/29157.html?spm=a2c4g.11186623.6.612.MPQuq5
type QueryBuilder struct {
	query string
}

// NewQueryBuilder creates a QueryBuilder with init key k and value v.
// k  and v should not be empty.
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (b *QueryBuilder) And(k, v string, isLike bool) *QueryBuilder {
	// 忽略空字符串
	if k == "" || v == "" {
		return b
	}
	if isLike {
		if b.query == "" {
			b.query = fmt.Sprintf("query=%s:'%s'", k, v)
		} else {
			b.query = fmt.Sprintf("%s AND %s:'%s'", b.query, k, v)
		}
	} else {
		if b.query == "" {
			b.query = fmt.Sprintf(`query=%s:"%s"`, k, v)
		} else {
			b.query = fmt.Sprintf(`%s AND %s:"%s"`, b.query, k, v)
		}
	}
	return b
}

func (b *QueryBuilder) Or(k, v string, isLike bool) *QueryBuilder {
	// 忽略空字符串
	if k == "" || v == "" {
		return b
	}
	if isLike {
		b.query = fmt.Sprintf("%s OR %s:'%s'", b.query, k, v)
	} else {
		b.query = fmt.Sprintf(`%s OR %s:"%s"`, b.query, k, v)
	}
	return b
}

func (b *QueryBuilder) String() string {
	return b.query
}

// ConfigBuilder builds a config string
//
// Doc: https://help.aliyun.com/document_detail/29156.html?spm=a2c4g.11186623.6.611.MPQuq5
type ConfigBuilder struct {
	config string
}

// NewConfigBuilderLimit creates a config string for pagination.
//
// start - offset
// hit - limit
func NewConfigBuilderLimit(start, hit int) *ConfigBuilder {
	c := ConfigBuilder{config: fmt.Sprintf("config=start:%d,hit:%d", start, hit)}
	return &c
}

func (c *ConfigBuilder) And(k, v string) *ConfigBuilder {
	if k == "" || v == "" {
		return c
	}
	if c.config == "" {
		c.config = fmt.Sprintf(`config=%s:%s`, k, v)
	} else {
		c.config = fmt.Sprintf(`%s,%s:%s`, c.config, k, v)
	}
	return c
}

func (c *ConfigBuilder) String() string {
	return c.config
}

// Concat joins URL params
func Concat(queries ...string) string {
	var s string
	for i, q := range queries {
		if len(q) > 0 {
			if i == 0 {
				s += q
			} else {
				s += " && " + q
			}
		}

	}
	return s
}
