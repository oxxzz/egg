package egg

import (
	"fmt"
	"net/http"
	"strings"
)

type HandleFunc func(*Context) error

type routerNode struct {
	isPlaceholder bool
	path          string
	part          string
	children      []*routerNode
}

// 树状结构打印节点信息
func (n *routerNode) printTree() {
	if n.part == "" {
		fmt.Println("/")
	}
	fmt.Printf("%s", "/"+n.part)
	for _, c := range n.children {
		c.printTree()
	}
	fmt.Printf("\n")
}

// insert
func (n *routerNode) insert(path string, parts []string, height int) {
	if len(parts) == height {
		n.path = path
		return
	}

	part := parts[height]
	child := n.matchPart(part)
	if child == nil {
		child = &routerNode{part: part, isPlaceholder: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(path, parts, height+1)
}

// search
func (n *routerNode) search(parts []string, height int) *routerNode {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.path == "" {
			return nil
		}

		return n
	}

	part := parts[height]
	children := n.match(part)

	for _, child := range children {
		if node := child.search(parts, height+1); node != nil {
			return node
		}
	}
	return nil
}

func (n *routerNode) matchPart(part string) *routerNode {
	for _, c := range n.children {
		if c.part == part || c.isPlaceholder {
			return c
		}
	}
	return nil
}

func (n *routerNode) match(part string) []*routerNode {
	items := make([]*routerNode, 0)
	for _, c := range n.children {
		if c.part == part || c.isPlaceholder {
			items = append(items, c)
		}
	}
	return items
}

// router
type router struct {
	root     map[string]*routerNode
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		root:     make(map[string]*routerNode),
		handlers: make(map[string]HandleFunc),
	}
}

func (r *router) handle(c *Context) {
	node, params := r.hitRoute(c.Method, c.Path)
	if node == nil {
		_ = c.JSON(http.StatusNotFound, M{"message": "404 Not Found"})
		return
	}

	c.params = params
	key := c.Method + " - " + node.path
	_ = r.handlers[key](c)
}

// only one * is allowed
func parsePath(path string) []string {
	sPath := strings.Split(strings.TrimSpace(path), "/")
	parts := make([]string, 0)
	for _, part := range sPath {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

// add route
func (r *router) addRoute(method, path string, handler HandleFunc) {
	parts := parsePath(path)
	key := method + " - " + path
	if _, ok := r.root[method]; !ok {
		r.root[method] = &routerNode{}
	}
	r.root[method].insert(path, parts, 0)
	r.handlers[key] = handler
}

// get route
func (r *router) hitRoute(method, path string) (*routerNode, map[string]string) {
	searchParts := parsePath(path)
	root, ok := r.root[method]
	if !ok {
		return nil, nil
	}

	node := root.search(searchParts, 0)
	if node == nil {
		return nil, nil
	}

	params := make(map[string]string)
	parts := parsePath(node.path)
	for i, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[i]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[i:], "/")
			break
		}
	}
	return node, params
}

func (r *router) print() {
	fmt.Println("[EGG] \t --------------------")
	for m, h := range r.handlers {
		fmt.Println(fmt.Sprintf("[EGG] %s --> %T", m, h))
	}
	fmt.Println("[EGG] \t --------------------")
}
