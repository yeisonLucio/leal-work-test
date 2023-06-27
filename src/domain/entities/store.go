package entities

type Status string

const (
	ActiveStatus   Status = "active"
	InactiveStatus Status = "inactive"
)

type Store struct {
	ID     uint
	Name   string
	Status Status
}
