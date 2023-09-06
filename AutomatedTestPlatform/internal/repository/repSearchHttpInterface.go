package repository

import (
	"gotest/AutomatedTestPlatform/internal/repository/dao"
)

type RepSearchInterface struct {
	SeaDao dao.InterfaceDao
}

func NewRepSearchInterface(SeaDao dao.InterfaceDao) *RepSearchInterface {
	return &RepSearchInterface{
		SeaDao: SeaDao,
	}
}

//func (rs *RepSearchInterface) InsertSeaPerDate(ctx context.Context, SearchEngineResult domain.SearchEnginePerResult) string {
//	err := rs.SeaDao.DBAnyInsert(ctx, SearchEngineResult)
//	if err != "" {
//		return err
//	}
//	return ""
//}
//
////func (rs *RepSearchInterface) InsertSeaApiDate(ctx context.Context, SearchEngineResult domain.SearchEngineApiResult) string {
//	err := rs.SeaDao.DBSearchApiInsert(ctx, SearchEngineResult)
//	if err != "" {
//		return err
//	}
//	return ""
//}
//
//func (rs *RepSearchInterface) InsertLoadSearchEngine(ctx context.Context, SearchEngineResult domain.SearchEngineResultReload) string {
//	err := rs.SeaDao.DBSearchLoadInsert(ctx, SearchEngineResult)
//	if err != "" {
//		return err
//	}
//	return ""
//}
