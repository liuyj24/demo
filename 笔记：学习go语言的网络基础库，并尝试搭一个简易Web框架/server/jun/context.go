package jun

import (
	"encoding/json"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	engine  *Engine

	handlers []HandlerFunc
	index    int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		index:   -1,
	}
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Writer.WriteHeader(code)
	c.Writer.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}
