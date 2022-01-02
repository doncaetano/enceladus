package router

type Method int

const (
	ALL Method = iota
	GET
	POST
	PUT
	DELETE
)

var availableMethods = []string{"ALL", "GET", "POST", "PUT", "DELETE"}

func (method Method) String() string {
	return availableMethods[method]
}
