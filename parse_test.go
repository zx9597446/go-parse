package parse

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var parseClient = NewClient()

func init() {
	parseClient.AppId = ""
	parseClient.AppKey = ""
	parseClient.DebugRequest = true
	rand.Seed(time.Now().UnixNano())
}

func randString() string {
	return fmt.Sprintf("abc%d", rand.Int())
}

func TestObject(t *testing.T) {
	className := "TestClass"
	o1 := NewObject()
	o1.Set("key", "value")
	err := o1.Save(parseClient, className, true)
	if err != nil {
		t.Fatal(err)
	}
	if o1.ObjectId() == "" {
		t.Fatal("null objectId")
	}
	o1.Set("updatekey", "updatevalue")
	err = o1.Update(parseClient, className)
	if err != nil {
		t.Fatal(err)
	}
	o2, err := FetchObject(parseClient, className, o1.ObjectId(), "")
	if err != nil {
		t.Fatal(err)
	}
	err = o2.Delete(parseClient, className)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDate(t *testing.T) {
	className := "TestClass"
	o1 := NewObject()
	d := FormatDate(time.Now())
	o1.Set("keydate", d)
	err := o1.Save(parseClient, className, true)
	if err != nil {
		t.Fatal(err)
	}
	err = o1.Delete(parseClient, className)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUser(t *testing.T) {
	u1 := NewUser()
	email := fmt.Sprintf("%s@email.com", randString())
	phone := fmt.Sprintf("1386818%0d%0d", rand.Intn(99), rand.Intn(99))
	t.Log(phone)
	username := randString()
	password := "password"
	r1, err := u1.Register(parseClient, username, password, email, phone)
	if err != nil {
		t.Fatal(r1, err)
	}
	r2, err := u1.Login(parseClient, username, password)
	if err != nil {
		t.Fatal(r2, err)
	}
}

func TestCloudFunction(t *testing.T) {
	//param := parse.NewParams()
	//param["key"] = "value"
	//r, err := CallFunction(parseClient, "functionName", param.Encode())
	//if err != nil {
	//t.Fatal(err, r)
	//}
}

func TestQueryObjectsBasic(t *testing.T) {
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
}

func TestQueryObjectsWhere(t *testing.T) {
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
}
