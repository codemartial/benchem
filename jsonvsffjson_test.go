package benchem_test

// Some say, drop in replacement X for a standard library pkg for more speed.
// So, is ffjson really faster than encoding/json? Let's find out.

import (
	"encoding/json"
	_ "errors"
	"github.com/pquerna/ffjson/ffjson"
	"strconv"
	"testing"
	"time"
)

type Foo struct {
	T time.Time
	M Bars
	L string
}

type Bar struct {
	I1 int32
	I2 int32
}

type Bars map[string]Bar

var foo = Foo{time.Now(), Bars{strconv.Itoa(1 << 30): Bar{(1 << 30), 0}}, "A Sample Foo Instance"}

func BenchmarkJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(foo); err != nil {
			b.Error("JSON Marshaling failed")
		}
	}
}

func BenchmarkFFJSONMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := ffjson.Marshal(foo); err != nil {
			b.Error("JSON Marshaling failed")
		}
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := Foo{M: Bars{}}
		if err := json.Unmarshal([]byte(`{"T":"2009-11-10T23:00:00Z","M":{"1073741824":{"I1":1073741824,"I2":0}},"L":"A Sample Foo Instance"}`), &f); err != nil {
			b.Error("JSON Unmarshal failed")
		}
	}
}

func BenchmarkFFJSONUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := Foo{M: Bars{}}
		if err := ffjson.Unmarshal([]byte(`{"T":"2009-11-10T23:00:00Z","M":{"1073741824":{"I1":1073741824,"I2":0}},"L":"A Sample Foo Instance"}`), &f); err != nil {
			b.Error("JSON Unmarshal failed")
		}
	}
}
