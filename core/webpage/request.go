package webpage

type Request struct {
	Operation string
	Encoding  string
	Headers   map[string]string
	Data      []byte
}
