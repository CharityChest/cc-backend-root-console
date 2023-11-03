package abstracts

type Application interface {
	Boot()
	Get(key interface{}) interface{}
	Set(key interface{}, value interface{})
	Listen()
}
