package model

type GetRequest struct {
	Request
}

type Request struct {
	Method      string
	HttpAddress string
	Headers     map[string]string
	Authorization
}

type Authorization struct {
	Type     string
	Username string
	Password string
}

type PostRequest struct {
	Request
	Body map[string]string
}

type DeleteRequest struct {
	Request
	Body map[string]string
}

type PutRequest struct {
	Request
	Body map[string]string
}
