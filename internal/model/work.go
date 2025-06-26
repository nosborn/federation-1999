package model

type Work struct {
	Pallet  Cargo  // Cargo info
	JobType int32  // Type of job - see #defines
	From    string // Name of pickup planet
	To      string // Name of delivery planet
	Status  int32  // Status of job
	Value   int32  // Value of contract in IG/ton
	Gtu     int32  // Time to complete contract
	Credits int32  // How many trader credits the job is worth

	// Union for job type-specific data
	FactryWk FactoryJob
	GenWk    GeneralJob

	Age int32 // Age of job in transportation
}
