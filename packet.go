package tackle

type Packet interface {
	Get(k string) interface{}
	Put(k string, v interface{})
}

type packet struct {
	data map[string]interface{}
}

func NewPacket() Packet {
	return packet{data: make(map[string]interface{})}
}

func (p packet) Get(k string) interface{} {
	return p.data[k]
}

func (p packet) Put(k string, v interface{}) {
	p.data[k] = v
}
