package model

const (
	DBObjectSize = 464
)

type DBObject struct {
	Number        int32
	Flags         uint32
	Name          [NAME_SIZE]byte
	Desc          [OBJECT_DESC_SIZE]byte
	Scan          [OBJECT_SCAN_SIZE]byte
	_             [1]byte // padding for alignment
	Sex           byte
	_             [3]byte // padding for alignment
	CurLoc        uint16
	StartLoc      uint16
	Weight        uint16
	GetEvent      uint16
	GiveEvent     uint16
	ConsumeEvent  uint16
	Value         int32
	DropEvent     uint16
	Offset        uint16
	MaxLoc        uint16
	MinLoc        uint16
	ShipGuns      [4]DBOldSGuns
	Hull          uint16
	Shield        uint16
	Engine        uint16
	Computer      uint16
	Fuel          uint16
	Hold          uint16
	Tonnage       uint16
	AttackPercent uint16
	KillEvent     uint16
	PrefObject    int16
	MoveCounter   int16
	MaxCounter    int16
}

type DBOldSGuns struct {
	GunType uint16
	Name    [20]byte
	Damage  uint16
	Power   uint16
}
