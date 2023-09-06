package dao

import (
	"context"
	"fmt"
	"gotest/AutomatedTestPlatform/internal/domain"
	"log"
	"strings"
)

func (intDao *StructDao) DBCommDeleteInterface(ctx context.Context, engineType string, id []int64) string {
	var targetModel interface{}
	switch {
	case strings.Contains(engineType, "Search"):
		switch {
		case strings.Contains(engineType, "Per"):
			targetModel = &domain.SearchEnginePerResult{}
		case strings.Contains(engineType, "Api"):
			targetModel = &domain.SearchEngineApiResult{}
		case strings.Contains(engineType, "Algo"):
			targetModel = &domain.SearchEngineAlgoResult{}
		}
	case strings.Contains(engineType, "Cluster"):
		switch {
		case strings.Contains(engineType, "Per"):
			targetModel = &domain.ClusterEnginePerInsert{}
		case strings.Contains(engineType, "Api"):
			targetModel = &domain.ClusterEngineApiInsert{}
		case strings.Contains(engineType, "Algo"):
			targetModel = &domain.ClusterEngineAlgoInsert{}
		}
	case strings.Contains(engineType, "Trt"):
		targetModel = &domain.SearchEngineAlgoResult{}
	default:
		return fmt.Sprintf("Unsupported engineType: %s", engineType)
	}
	err := intDao.db.WithContext(ctx).Where("id IN (?)", id).Delete(targetModel).Error
	if err != nil {
		log.Printf("err [%v]\n", err)
		return fmt.Sprintf("DBCommDeleteInterface Mysql err[%v]", err)
	}
	return ""
}

func (intDao *StructDao) DBCommConditionsInquire(ctx context.Context, engineType, key string) (results []string, errInfo string) {
	var targetModel interface{}
	switch {
	case strings.Contains(engineType, "Search"):
		switch {
		case strings.Contains(engineType, "Per"):
			targetModel = &domain.SearchEnginePerResult{}
		case strings.Contains(engineType, "Api"):
			targetModel = &domain.SearchEngineApiResult{}
		case strings.Contains(engineType, "Algo"):
			targetModel = &domain.SearchEngineAlgoResult{}
		}
	case strings.Contains(engineType, "Cluster"):
		switch {
		case strings.Contains(engineType, "Per"):
			targetModel = &domain.ClusterEnginePerInsert{}
		case strings.Contains(engineType, "Api"):
			targetModel = &domain.ClusterEngineApiInsert{}
		case strings.Contains(engineType, "Algo"):
			targetModel = &domain.ClusterEngineAlgoInsert{}
		}
	case strings.Contains(engineType, "Trt"):
		targetModel = &domain.SearchEngineAlgoResult{}
	default:
		return nil, fmt.Sprintf("Unsupported engineType: %s", engineType)
	}
	err := intDao.db.WithContext(ctx).Model(targetModel).Select(fmt.Sprintf("DISTINCT(%v)", key)).Scan(&results).Error
	if err != nil {
		log.Printf("err [%v]\n", err)
		return nil, fmt.Sprintf("DBCommConditionsInquire Mysql err[%v]", err)
	}

	return results, ""
}

func (intDao *StructDao) DBCommUpdateInterface(ctx context.Context, engineType string, idList []int64, key string) string {
	var targetModel interface{}
	var existingValues []string
	switch {
	case strings.Contains(engineType, "Search"):
		switch {
		case strings.Contains(engineType, "Per"):
			targetModel = &domain.SearchEnginePerResult{}
		case strings.Contains(engineType, "Api"):
			targetModel = &domain.SearchEngineApiResult{}
		case strings.Contains(engineType, "Algo"):
			targetModel = &domain.SearchEngineAlgoResult{}
		}
	case strings.Contains(engineType, "Cluster"):
		switch {
		case strings.Contains(engineType, "Per"):
			targetModel = &domain.ClusterEnginePerInsert{}
		case strings.Contains(engineType, "Api"):
			targetModel = &domain.ClusterEngineApiInsert{}
		case strings.Contains(engineType, "Algo"):
			targetModel = &domain.ClusterEngineAlgoInsert{}
		}
	case strings.Contains(engineType, "Trt"):
		targetModel = &domain.SearchEngineAlgoResult{}
	default:
		return fmt.Sprintf("Unsupported engineType: %s", engineType)
	}
	// 查询现有值
	err := intDao.db.WithContext(ctx).Model(targetModel).Where("id IN (?)", idList).Pluck("is_final_result", &existingValues).Error
	if err != nil {
		log.Printf("err [%v]\n", err)
		return fmt.Sprintf("DBAnyInsert Mysql err[%v]", err)
	}
	// 比较值
	for id, existingValue := range existingValues {
		if existingValue == "true" {
			// 如果现有值与新值相同，抛出错误
			return fmt.Sprintf("数据库中已标记该类型,请先取消标记,标记数据id[%v]", id)
		}
	}
	// 执行更新
	err = intDao.db.WithContext(ctx).Model(targetModel).Where("id IN (?)", idList).Update("is_final_result", key).Error
	if err != nil {
		log.Printf("err [%v]\n", err)
		return fmt.Sprintf("DBAnyInsert Mysql err[%v]", err)
	}
	return ""
}

func (intDao *StructDao) DBSearchCommInterface(ctx context.Context, engineType string, conditions map[string]interface{}) (errinfo string, EngineResult interface{}) {
	var query string
	switch {
	case strings.Contains(engineType, "Search"):
		EngineResult = []domain.SearchEnginePerResult{}
		query = "SELECT * FROM search_engine_per_results"
		if strings.Contains(engineType, "Api") {
			EngineResult = []domain.SearchEngineApiResult{}
			query = "SELECT * FROM search_engine_api_results"
		} else if strings.Contains(engineType, "Algo") {
			EngineResult = []domain.SearchEnginePerResult{}
			query = "SELECT * FROM search_engine_algo_results"
		}
	case strings.Contains(engineType, "Cluster"):
		EngineResult = []domain.ClusterEnginePerInsert{}
		query = "SELECT * FROM cluster_engine_per_results"
		if strings.Contains(engineType, "Api") {
			EngineResult = []domain.ClusterEngineApiInsert{}
			query = "SELECT * FROM cluster_engine_api_results"
		} else if strings.Contains(engineType, "Algo") {
			EngineResult = []domain.ClusterEngineAlgoInsert{}
			query = "SELECT * FROM cluster_engine_algo_results"
		}
	}
	var args []interface{}
	//将查询语句转换成sql
	conditionsStr := BuildConditions(conditions, &args)
	if conditionsStr != "" {
		query += " WHERE " + conditionsStr
	}

	err := intDao.db.WithContext(ctx).Raw(query, args...).Scan(&EngineResult).Error
	if err != nil {
		log.Printf("err [%v]\n", err)
		return fmt.Sprintf("Search Mysql err[%v]", err), nil
	}
	return "", nil
}

func (intDao *StructDao) DBAnyInsert(ctx context.Context, EngineResult interface{}) string {
	switch v := EngineResult.(type) {
	case domain.SearchEnginePerResult:
		err := intDao.db.WithContext(ctx).Create(&v).Error
		if err != nil {
			return fmt.Sprintf("DBAnyInsert Mysql err[%v]", err)
		}
	case domain.SearchEngineApiResult:
		err := intDao.db.WithContext(ctx).Create(&v).Error
		if err != nil {
			return fmt.Sprintf("DBAnyInsert Mysql err[%v]", err)
		}
	case domain.SearchEngineAlgoResult:
		err := intDao.db.WithContext(ctx).Create(&v).Error
		if err != nil {
			return fmt.Sprintf("DBAnyInsert Mysql err[%v]", err)
		}
	case domain.ClusterEnginePerInsert:
		err := intDao.db.WithContext(ctx).Create(&v).Error
		if err != nil {
			return fmt.Sprintf("DBAnyInsert Mysql err[%v]", err)
		}
	default:
		return fmt.Sprintf("Unsupported EngineResult type: %T", EngineResult)
	}
	return ""
}
