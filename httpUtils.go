package server

type Method int

const (
	GET Method = iota
	POST
	PUT
	DELETE
)

func (m Method) String() string {
	return [...]string{"GET", "POST", "PUT", "DELETE"}[m]
}

var statusCode = map[int]string{
	200: "OK",
	400: "Bad Request",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
}
