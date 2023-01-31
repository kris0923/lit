package lite

type IGrooup interface {
	Get(string, Handler)
	Post(string, Handler)
	Put(string, Handler)
	Delete(string, Handler)
}

type RouterGroup struct {
	Prefix string
}
