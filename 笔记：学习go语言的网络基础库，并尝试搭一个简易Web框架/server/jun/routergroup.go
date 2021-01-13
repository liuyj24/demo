package jun

type RouterGroup struct {
	basePath string
	engine   *Engine
	Handlers []HandlerFunc
}

func (group *RouterGroup) Group(relativePath string) *RouterGroup {
	newGroup := &RouterGroup{
		basePath: relativePath,
		engine:   group.engine,
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) *RouterGroup {
	group.Handlers = append(group.Handlers, middlewares...)
	return group
}
