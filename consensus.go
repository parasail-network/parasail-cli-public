package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

type SLACheckResult struct {
	VaultAddress common.Address
	CheckRound   int64
	Results      []*DealCheckResult
}

type DealCheckResult struct {
	ID      string
	Result  byte
	SlaData []byte
}

type ConsensusEngine interface {
	Init(ctx context.Context) error
	Propose(ctx context.Context, data *SLACheckResult) error
	RegisterValidateHandler(func(ctx context.Context, data *SLACheckResult) (bool, error))
}

type EmptyEngine struct{}

func (e *EmptyEngine) Init(ctx context.Context) error {
	return nil
}
func (e *EmptyEngine) Propose(ctx context.Context, checkResult *SLACheckResult) error {
	return nil
}
func (e *EmptyEngine) RegisterValidateHandler(func(ctx context.Context, checkResult *SLACheckResult) (bool, error)) {

}
