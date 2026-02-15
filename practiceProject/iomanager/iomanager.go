package iomanager

type IOManager interface {
	Readlines() ([]string, error)
	WriteJSON(data any) error
}
