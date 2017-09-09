package http



type Serializer interface {
	Serialize()
	Deserialize()
}