package dto

type Request struct {
	Path   string
	Method string
	Body   map[string]string
	Header map[string]string
}
