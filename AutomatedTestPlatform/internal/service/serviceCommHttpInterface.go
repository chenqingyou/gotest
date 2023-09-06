package service

import (
	"context"
	"gotest/AutomatedTestPlatform/internal/repository"
)

type CommServiceInterface struct {
	CommRepo *repository.CommRepInterface
}

func NewCommServiceInterface(repComm *repository.CommRepInterface) *CommServiceInterface {
	return &CommServiceInterface{
		CommRepo: repComm,
	}
}

func (cs *CommServiceInterface) CommConditionsInquire(ctx context.Context, engineType, key string) ([]string, string) {
	return cs.CommRepo.CommConditionsInquire(ctx, engineType, key)
}

func (cs *CommServiceInterface) CommDeleteInterface(ctx context.Context, engineType string, id []int64) string {
	return cs.CommRepo.CommDeleteInterface(ctx, engineType, id)
}

func (cs *CommServiceInterface) CommUpdateInterface(ctx context.Context, engineType string, id []int64, key string) string {
	return cs.CommRepo.CommUpdateInterface(ctx, engineType, id, key)
}

func (cs *CommServiceInterface) CommInsertInterface(ctx context.Context, engineType string, conditions map[string]interface{}) (string, interface{}) {
	return cs.CommRepo.CommInsertInterface(ctx, engineType, conditions)
}
