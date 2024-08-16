package interfacex

type Unmarshal interface {
	Unmarshal(data []byte, v any) error
}
