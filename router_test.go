package egg

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/users/:name", nil)
	r.addRoute("GET", "/nice/fast", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePath("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePath("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePath("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePath failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.hitRoute("GET", "/hi/some-strings")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.path != "/hi/:name" {
		t.Fatal("should match /hi/:name")
	}

	if ps["name"] != "some-strings" {
		t.Fatal("name should be equal to 'some-strings'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.path, ps["name"])

}
