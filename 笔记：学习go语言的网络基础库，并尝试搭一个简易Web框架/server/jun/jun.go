package jun

import (
	"net/http"
	"strings"
)

// jun, a mini web framework

type HandlerFunc func(*Context)

type Engine struct {
	RouterGroup
	groups []*RouterGroup
	router map[string]HandlerFunc
}

func Default() *Engine {
	engine := New()
	return engine
}

func New() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			basePath: "",
		},
	}
	engine.RouterGroup.engine = engine
	engine.groups = append(engine.groups, &engine.RouterGroup)
	return engine
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := newContext(w, req)
	context.engine = engine

	// 根据路径判断这个请求属于哪个组，然后获取这个组上安装的中间件，把这些中间件放到上下文中
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.basePath) {
			context.handlers = append(context.handlers, group.Handlers...)
		}
	}

	// 把请求处理器也放到上下文中
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		context.handlers = append(context.handlers, handler)
	} else {
		context.handlers = append(context.handlers, func(context *Context) {
			http.Error(context.Writer, "404 page not found", http.StatusNotFound)
		})
	}

	//处理请求！
	context.Next()
}

func (engine *Engine) Run(addr ...string) error {
	address := resolveAddress(addr)
	return http.ListenAndServe(address, engine)
}

func (group *RouterGroup) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	if group.engine.router == nil {
		group.engine.router = make(map[string]HandlerFunc)
	}
	pattern = group.basePath + pattern
	key := method + "-" + pattern
	group.engine.router[key] = handlerFunc
}

func (group *RouterGroup) Get(pattern string, handlerFunc HandlerFunc) {
	group.addRouter("GET", pattern, handlerFunc)
}

func (group *RouterGroup) Post(pattern string, handlerFunc HandlerFunc) {
	group.addRouter("POST", pattern, handlerFunc)
}
