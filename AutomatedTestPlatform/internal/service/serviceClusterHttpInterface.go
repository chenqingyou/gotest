package service

import (
	"context"
	"fmt"
	"gotest/AutomatedTestPlatform/internal/domain"
	"gotest/AutomatedTestPlatform/internal/repository"
)

type ClusterServiceInterface struct {
	RepoCluster *repository.RepClusterInterface
}

func NewClusterServiceInterface(repoCluster *repository.RepClusterInterface) *ClusterServiceInterface {
	return &ClusterServiceInterface{
		RepoCluster: repoCluster,
	}
}

func (ss *ClusterServiceInterface) InsertSvcPerData(ctx context.Context, xlsxOp domain.XlsxInterface) (int64, int64, string) {
	per1Res, per2Res, errInfo := xlsxOp.OpenClusterExcel()
	if errInfo != "" {
		return 0, 0, fmt.Sprintf("errInfop[%v],len per1Res[0], len(per2Res) == 0 ", errInfo)
	}
	successfulNumber := int64(0)
	failuresNumber := int64(0)
	//数据插入
	for _, value := range per1Res {
		InsertSeaPerDateErrInfo := ss.RepoCluster.CluDao.DBAnyInsert(ctx, value)
		if InsertSeaPerDateErrInfo != "" {
			fmt.Println(errInfo)
			failuresNumber++
		} else {
			successfulNumber++
		}
	}
	for _, value := range per2Res {
		InsertLoadSearchEngineErrInfo := ss.RepoCluster.CluDao.DBAnyInsert(ctx, value)
		if InsertLoadSearchEngineErrInfo != "" {
			fmt.Println(errInfo)
			failuresNumber++
		} else {
			successfulNumber++
		}
	}

	return successfulNumber, failuresNumber, ""
}

func (ss *ClusterServiceInterface) InsertSvcApiData(ctx context.Context, xlsxOp domain.XlsxInterface) (int64, int64, string) {
	ApiRes, errInfo := xlsxOp.GetApiDataFromFile()
	if errInfo != "" {
		return 0, 0, fmt.Sprintf("errInfop[%v],len per1Res[0], len(per2Res) == 0 ", errInfo)
	}
	successfulNumber := int64(0)
	failuresNumber := int64(0)
	//数据插入
	for _, value := range ApiRes {
		InsertSeaPerDateErrInfo := ss.RepoCluster.CluDao.DBAnyInsert(ctx, value)
		if InsertSeaPerDateErrInfo != "" {
			fmt.Println(errInfo)
			failuresNumber++
		} else {
			successfulNumber++
		}
	}
	return successfulNumber, failuresNumber, ""
}
