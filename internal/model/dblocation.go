package model

const (
	DBLocationSize = 1060
)

type DBLocation struct {
	Desc    [DESC_SIZE]byte
	Events  [EVENT_SIZE]uint16
	MapFlag uint32
	MovTab  [13]uint16
	SysLoc  uint16
}
