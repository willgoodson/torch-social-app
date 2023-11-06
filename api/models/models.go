package models

type Model interface {
	To_Map() (map[string]any, error)
	Is_Populated() error
}
