package base

import (
	"fmt"
	"strings"

	"gorm.io/gorm/clause"
)

type SQLSortReq struct {
	Sort string `json:"sort" form:"sort"`
}

func (c *SQLSortReq) SortFilter() interface{} {
	order := clause.OrderByColumn{}
	if c.Sort == "" {
		order.Column.Name = "created_at"
		order.Column.Table = clause.CurrentTable
		order.Desc = true
		return order
	}
	c.Sort = strings.Trim(c.Sort, " ")
	var desc bool
	var sort string
	if strings.Contains(c.Sort, " ") {
		ss := strings.Split(c.Sort, " ")
		if len(ss) == 2 {
			if ss[1] == "desc" {
				desc = true
			}
		}
		sort = ss[0]
	} else {
		sort = c.Sort
	}
	order.Column.Name = sort
	order.Column.Table = clause.CurrentTable
	order.Desc = desc
	return order
}

type SQLSearchReq struct {
	Q string `json:"q" form:"q"`
}

func (c *SQLSearchReq) AddQ(k, v string) {
	if k != "" && v != "" {
		if c.Q != "" {
			c.Q = fmt.Sprintf("%s,%s=%s", c.Q, k, v)
		} else {
			c.Q = fmt.Sprintf("%s=%s", k, v)
		}
	}
}

func (c *SQLSearchReq) QToMap() map[string]interface{} {
	res := make(map[string]interface{})
	if c.Q != "" {
		filters := strings.Split(c.Q, ",")
		for _, f := range filters {
			_f := strings.Split(f, "=")
			if len(_f) != 2 {
				continue
			}
			res[_f[0]] = _f[1]
		}
	}
	return res
}

func (c *SQLSearchReq) SearchFilter(fuzzy bool) interface{} {
	if c.Q == "" {
		return nil
	}
	filter := clause.Where{}
	for k, v := range c.QToMap() {
		var expr clause.Expr
		s := ""
		if v == "true" || v == "false" {
			s = fmt.Sprintf(`%s=%s`, k, v)
		} else {
			if fuzzy {
				s = fmt.Sprintf(`%s like "%%%s%%"`, k, v)
			} else {
				s = fmt.Sprintf(`%s=%s`, k, v)
			}
		}
		expr = clause.Expr{
			SQL:                s,
			Vars:               nil,
			WithoutParentheses: false,
		}
		filter.Exprs = append(filter.Exprs, expr)
	}
	return filter
}

type WSearchReq struct {
	W string `json:"w" form:"w"`
}

func (c *WSearchReq) WordFilter(fuzzy bool, fields ...string) interface{} {
	if c.W == "" {
		return nil
	}

	var exprs []clause.Expression
	for _, f := range fields {
		var expr clause.Expression
		s := ""
		if fuzzy {
			s = fmt.Sprintf(`%s like "%%%s%%"`, f, c.W)
		} else {
			s = fmt.Sprintf(`%s=%s`, f, c.W)
		}
		expr = clause.Expr{
			SQL:                s,
			Vars:               nil,
			WithoutParentheses: false,
		}
		exprs = append(exprs, expr)
	}
	return clause.Where{
		Exprs: []clause.Expression{clause.Or(exprs...)},
	}
}
