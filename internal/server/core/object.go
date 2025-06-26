package core

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

type Object struct { // Object record
	// Common properties

	Number        uint16   // Object's vocab number
	Flags         uint32   // flags class to which the object belongs
	Name          string   // [limits.NAME_SIZE * 2]byte
	Synonyms      []string // [2][limits.NAME_SIZE]byte
	DescMessageNo text.MsgNum
	ScanMessageNo text.MsgNum
	Sex           model.Sex // m or f or n
	MinLoc        uint16    // Lowest location for placement
	MaxLoc        uint16    // Highest location for placement
	Weight        uint16    // wt of object in strength units
	GetEvent      uint16    // event when player tries to get object
	GiveEvent     uint16    // event when player gives object
	Value         int32     // value/price on obj, -ve goes on killer's reward!
	Events        [4]uint16

	// Object specific properties

	DropEvent uint16 // event when player drops object

	// Mobile specific properties

	ShipGuns      [4]model.OldSGuns
	ShipKit       Equipment
	AttackPercent uint16 // % chance of mobile attacking
	PrefObject    uint16 // item which mobile will pay double for!
	MaxCounter    int16  // counter for movement & counter, reset level. -ve move_counter = immobile
}

type Equipment struct { // Ship's equipment record
	Hull     uint16 // Hull strength
	Shield   uint16 // Shield strength
	Engine   uint16 // Engine capacity
	Computer uint16 // Computer size
	Fuel     uint16 // Fuel
	Hold     uint16 // Cargo capacity
	Tonnage  uint16 // Overall size of ship
}

var (
	SHIP_GUN_MAG_GUN    = model.OldSGuns{Type: model.GunMagGun, Name: "Mag Gun", Damage: 2, Power: 2}
	SHIP_GUN_MISSILE    = model.OldSGuns{Type: model.GunMissile, Name: "Missile", Damage: 4, Power: 0}
	SHIP_GUN_LASER      = model.OldSGuns{Type: model.GunLaser, Name: "Laser", Damage: 5, Power: 15}
	SHIP_GUN_TWIN_LASER = model.OldSGuns{Type: model.GunTwinLaser, Name: "Twin Laser", Damage: 7, Power: 30}
	SHIP_GUN_QUAD_LASER = model.OldSGuns{Type: model.GunQuadLaser, Name: "Quad Laser", Damage: 10, Power: 30}
)
