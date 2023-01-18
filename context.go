package lite

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// 实现contxt.Context, 并封装对request和response的操作
type Context struct {
	ctx            context.Context
	request        *http.Request
	responseWriter http.ResponseWriter
}

func NewContext(req *http.Request, resp http.ResponseWriter) *Context {
	return &Context{}
}

func (*Context) WriterMux() {
	//TODO
}

func (c *Context) SetHasTimeout() {

}

func (c *Context) HasTimeout() {

}
func (c *Context) BaseContext() context.Context {
	return c.ctx
}

func (*Context) Deadline() {
	//TODO
}

func (*Context) Done() {
	//TODO
}

func (*Context) Err() {
	//TODO
}

func (*Context) Value(key interface{}) interface{} {
	//TODO
	return nil
}

func (c *Context) Query(key string) url.Values {
	//TODO
	return c.request.URL.Query()
}
func (c *Context) QueryInt(key string, def int) {

}
func (c *Context) QueryString(key string) string {
	//TODO
	return ""
}
func (c *Context) QueryArray(key string, def []string) {}
func (c *Context) QueryAll()                           {}

func (c *Context) FormInt(key string, def int) {
	//TODO
}
func (c *Context) FormString(key string, def string)  {}
func (c *Context) FormArray(key string, def []string) {}
func (c *Context) FormAll()                           {}
func (c *Context) BindJson(obj interface{})           {}

func (c *Context) Json(data interface{}) {
	//TODO
	b, err := json.Marshal(data)
	if err != nil {
		c.responseWriter.WriteHeader(500)
		c.responseWriter.Write([]byte("invalid json"))
		return
	}
	_, err = c.responseWriter.Write(b)
	if err != nil {
		c.responseWriter.WriteHeader(500)
	} else {
		c.responseWriter.WriteHeader(200)
	}
}

func (c Context) Html(statusCode int, obj interface{}, template ...string) {
	//TODO
}
func (c *Context) Text(statusCode int, text string) {}
