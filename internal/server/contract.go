package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/pkg/ibgames"
)

type Contract struct {
	Owner  ibgames.AccountID
	Planet string
	Value  int32

	Pallet model.Cargo
}
