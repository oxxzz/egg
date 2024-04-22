package egg

import (
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type M map[string]interface{}

type Context struct {
	writer  http.ResponseWriter
	request *http.Request

	Path       string
	Method     string
	UUID       string
	StatusCode int

	params map[string]string
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		writer:  w,
		request: r,

		UUID:   uuid.NewString(),
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context) FormValue(name string) string {
	return c.request.FormValue(name)
}

func (c *Context) QueryValue(name string) string {
	return c.request.URL.Query().Get(name)
}

func (c *Context) PathValue(key string) string {
	v, _ := c.params[key]
	return v
}

func (c *Context) HeaderValue(key string) string {
	return c.request.Header.Get(key)
}

func (c *Context) TraceId() string {
	return c.UUID
}

func (c *Context) SetHeader(key string, value string) *Context {
	c.writer.Header().Set(key, value)
	return c
}

func (c *Context) Status(code int) *Context {
	c.StatusCode = code
	c.writer.WriteHeader(code)
	return c
}

func (c *Context) JSON(code int, data interface{}) error {
	c.SetHeader("Content-Type", "application/json; charset=utf-8")
	c.Status(code)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.NewEncoder(c.writer).Encode(data)
}

func (c *Context) HTML(code int, html string) error {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	_, err := c.writer.Write([]byte(html))
	return err
}

func (c *Context) String(code int, s string) error {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, err := c.writer.Write([]byte(s))
	return err
}

func (c *Context) Bytes(code int, b []byte) error {
	c.Status(code)
	_, err := c.writer.Write(b)
	return err
}

func (c *Context) Stream(code int, name string, data []byte) error {
	c.SetHeader("Content-Type", "application/octet-stream")
	c.SetHeader("Content-Disposition", "attachment; filename="+name)
	c.SetHeader("Content-Length", string(rune(len(data))))
	c.Status(code)

	_, err := c.writer.Write(data)
	return err
}
