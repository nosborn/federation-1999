package model

const (
	GunMagGun    = 1 + iota // Mag gun
	GunMissile              // Missile rack
	GunLaser                // Single laser
	GunTwinLaser            // Twin laser
	GunQuadLaser            // Quad laser
)

type SGuns struct { // Ship mounted weapons
	Type   int32 // Weapon type
	Damage int32 // Base damage
}
