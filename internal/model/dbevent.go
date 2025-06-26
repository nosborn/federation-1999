package model

const (
	DBEventSize = 176
)

type DBEvent struct {
	Type   uint16
	Desc   [EVENT_DESC_SIZE]byte
	Field1 int16
	Field2 int16
	Field3 int16
	Field4 int16
	Field7 int16
	Field8 int16
	NewLoc uint16
}
