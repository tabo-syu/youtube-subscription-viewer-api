package ports

type ErrorsOutputPort interface {
	OutputError() error
}
