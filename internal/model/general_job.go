package model

type GeneralJob struct {
	WhereTo int32  // Destined for exchange or warehouse?
	Owner   string // Player who owns the goods being hauled
}
