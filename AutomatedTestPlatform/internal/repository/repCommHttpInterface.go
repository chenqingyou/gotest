package repository

import (
	"context"
	"gotest/AutomatedTestPlatform/internal/repository/dao"
)

type CommRepInterface struct {
	DB dao.InterfaceDao
}

func NewCommRepInterface(db dao.InterfaceDao) *CommRepInterface {
	return &CommRepInterface{
		DB: db,
	}
}

func (cr *CommRepInterface) CommConditionsInquire(ctx context.Context, engineType, key string) ([]string, string) {
	return cr.DB.DBCommConditionsInquire(ctx, engineType, key)
}

func (cr *CommRepInterface) CommDeleteInterface(ctx context.Context, engineType string, id []int64) string {
	return cr.DB.DBCommDeleteInterface(ctx, engineType, id)
}

func (cr *CommRepInterface) CommUpdateInterface(ctx context.Context, engineType string, id []int64, key string) string {
	return cr.DB.DBCommUpdateInterface(ctx, engineType, id, key)
}

func (cr *CommRepInterface) CommInsertInterface(ctx context.Context, engineType string, conditions map[string]interface{}) (string, interface{}) {
	return cr.DB.DBSearchCommInterface(ctx, engineType, conditions)
}
