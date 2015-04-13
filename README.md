go-parse
============

a [parse.com](https://parse.com/) [REST API](https://parse.com/docs/rest#summary) client library for go

install
------------
```go get -u github.com/zx9597446/go-parse```

quick usage
-----------
1. config client:
```go
	var parseClient = parse.NewClient()
	parseClient.AppId = ""
	parseClient.AppKey = ""
	parseClient.DebugRequest = true
```

2. Object CRUD
```go
className := "TestClass"
o := parse.NewObject()
o.Set("key", "value")
fetchOnSave := true
o.Create(parseClient, className, fetchOnSave)

o.Set("newKey", "newValue")
o.Update(parseClient, className)

o2 := parse.NewObject()
oid := o.ObjectId()
o2.Fetch(parseClient, className, oid, "")

o2.Delete(parseClient, className)
```

3. call cloud function:
```go
param := parse.NewParams()
param["key"] = "value"
ret, err := parseClient.CallFunction("functionName", param.Encode())
```

4. query objects (basic)
```go
	opt := QueryOptions{
		Class:   "PS",
		Limit:   DefaultLimit,
		Skip:    0,
		Order:   "-createdAt'",
		Keys:    "",
		Include: "",
		Count:   false,
	}
	r, err := QueryObjects(parseClient, opt)
	if err != nil {
		t.Fatal(err)
	}
	objs, _ := r.GetResults()
	for _, o := range objs {
		t.Log(o.ObjectId())
	}

```

5. query objects (advance)
```go
	where := NewWhere()
	where.AddEqualTo("objectId", "abcd")
	where.Add("someKey", OpLessThanOrEqualTo, 100)
	opt := QueryOptions{
		Class:   "PS",
		Limit:   DefaultLimit,
		Skip:    0,
		Order:   "-createdAt'",
		Keys:    "",
		Include: "",
		Count:   false,
		Where:   where,
	}
	r, err := QueryObjects(parseClient, opt)
	if err != nil {
		t.Fatal(err)
	}
	objs, _ := r.GetResults()
	for _, o := range objs {
		t.Log(o.ObjectId())
	}

```

API doc
------------
see [doc](http://godoc.org/github.com/zx9597446/go-parse)

examples
-----------
see [test](http://github.com/zx9597446/go-parse/blob/master/parse_test.go)
