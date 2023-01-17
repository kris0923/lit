package lit

type Core struct {}

func NewCore() *Core {
	return &Core{}
}

func (*Core) ServeHTTP(req *http.Request, resp http.ResponseWriter) {
	// 
}

// 实现contxt.Context, 并封装对request和response的操作
type Context struct {
	ctx context.Context
	request *http.Request
	responseWriter http.ResponseWriter
}


func NewContext(req *http.Request, resp http.ResponseWriter) *Context{
	return &Context{}
}

func (*Context) WriterMux() {
	//TODO
}

func (*Context) BaseContext() {

}

func (*Context) Deadline() {

}

func (*Context) Done() {

}

func (*Context) Err() {

}
