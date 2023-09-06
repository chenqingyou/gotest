package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gotest/AutomatedTestPlatform/internal/domain"
	"strings"
)

type InterfaceDao interface {
	//todo 识别
	//DBSearchApiInsert(ctx context.Context, searchEngineResult domain.SearchEngineApiResult) string
	//DBSearchLoadInsert(ctx context.Context, searchEngineResult domain.SearchEngineResultReload) string
	////聚类
	//DBSClusterPerInsert(ctx context.Context, searchEngineResult domain.ClusterEnginePerResult) string
	//DBClusterPerDetailInsert(ctx context.Context, EngineResult domain.ClusterEnginePerDetailResult) string
	//DBClusterApiInsert(ctx context.Context, searchEngineResult domain.SearchEngineApiResult) string
	//DBClusterAlgoInsert(ctx context.Context, searchEngineResult domain.SearchEngineResultReload) string
	//特征
	//公共
	DBAnyInsert(ctx context.Context, EngineResult interface{}) string
	DBCommDeleteInterface(ctx context.Context, engineType string, id []int64) string
	DBCommConditionsInquire(ctx context.Context, engineType, key string) ([]string, string)
	DBCommUpdateInterface(ctx context.Context, engineType string, idList []int64, key string) string
	DBSearchCommInterface(ctx context.Context, engineType string, conditions map[string]interface{}) (errinfo string, EngineResult interface{})
}

type StructDao struct {
	db *gorm.DB
}

func NewInterfaceDao(db *gorm.DB) InterfaceDao {
	return &StructDao{
		db: db,
	}
}

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&domain.SearchEnginePerResult{}, &domain.SearchEngineResultReload{})
}

func BuildConditions(conditions map[string]interface{}, args *[]interface{}) string {
	var conditionParts []string
	for field, value := range conditions {
		//类型断言
		switch v := value.(type) {
		case []interface{}:
			placeholders := make([]string, len(v))
			for i := range v {
				*args = append(*args, v[i])
				placeholders[i] = "?"
			}
			condition := fmt.Sprintf("%s IN (%s)", field, strings.Join(placeholders, ", "))
			conditionParts = append(conditionParts, condition)
		default:
			// 处理单个值条件
			*args = append(*args, v)
			condition := fmt.Sprintf("%s = ?", field)
			conditionParts = append(conditionParts, condition)
		}
	}
	return strings.Join(conditionParts, " AND ")
}
