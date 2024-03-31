package plugins

type Operation string

const (
	Create  Operation = "create"
	Update  Operation = "update"
	Delete  Operation = "delete"
	Connect Operation = "connect"
)

type Interface interface {
	Handles(operation Operation) bool
}
