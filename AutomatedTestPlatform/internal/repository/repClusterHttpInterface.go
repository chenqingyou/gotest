package repository

import (
	"gotest/AutomatedTestPlatform/internal/repository/dao"
)

type RepClusterInterface struct {
	CluDao dao.InterfaceDao
}

func NewRepClusterInterface(CluDao dao.InterfaceDao) *RepClusterInterface {
	return &RepClusterInterface{
		CluDao: CluDao,
	}
}

//func (rs *RepClusterInterface) InsertCluPerDate(ctx context.Context, EngineResult domain.ClusterEnginePerResult) string {
//	err := rs.CluDao.DBSClusterPerInsert(ctx, EngineResult)
//	if err != "" {
//		return err
//	}
//	return ""
//}
//
//func (rs *RepClusterInterface) InsertCluApiDate(ctx context.Context, EngineResult domain.SearchEngineApiResult) string {
//	err := rs.CluDao.DBSearchApiInsert(ctx, EngineResult)
//	if err != "" {
//		return err
//	}
//	return ""
//}
//
//func (rs *RepClusterInterface) InsertCluAlgoData(ctx context.Context, EngineResult domain.SearchEngineResultReload) string {
//	err := rs.CluDao.DBSearchLoadInsert(ctx, EngineResult)
//	if err != "" {
//		return err
//	}
//	return ""
//}
