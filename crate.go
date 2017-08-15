package tackle




type Create interface {
	Put(key string, value interface{}) Create
	Get(key string) interface{}
}
