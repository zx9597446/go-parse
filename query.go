package parse

import (
	"encoding/json"
	"log"
)

type WhereType string

const OpLessThan = WhereType("$lt")
const OpLessThanOrEqualTo = WhereType("$lte")
const OpGreaterThan = WhereType("$gt")
const OpGreaterThanOrEqualTo = WhereType("$gte")
const OpNotEqualTo = WhereType("$nt")
const OpContainedIn = WhereType("$in")
const OpNotContainedIn = WhereType("$nin")
const OpExists = WhereType("$exists")
const OpSelect = WhereType("$select")
const OpNotSelect = WhereType("$dontSelect")
const OpAll = WhereType("$all")

type WhereQuery map[string]interface{}

func NewWhere() WhereQuery {
	return WhereQuery(make(map[string]interface{}))
}

func (w WhereQuery) Add(left string, op WhereType, right interface{}) {
	w[left] = map[string]interface{}{
		string(op): right,
	}
}

func (w WhereQuery) AddEqualTo(left string, right interface{}) {
	w[left] = right
}

func (w WhereQuery) ToJSON() string {
	bs, err := json.Marshal(w)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(bs)
}

const DefaultLimit = 100
const DefaultSkip = 0

type QueryOptions struct {
	Class   string
	Where   WhereQuery
	Limit   int
	Skip    int
	Order   string
	Keys    string
	Include string
	Count   bool
}
