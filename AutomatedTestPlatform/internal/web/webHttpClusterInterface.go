package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wxnacy/wgo/file"
	"gotest/AutomatedTestPlatform/internal/domain"
	"gotest/AutomatedTestPlatform/internal/service"
	"net/http"
)

type ClusterWebHttpStruct struct {
	ClusterService *service.ClusterServiceInterface
}

func (whc *ClusterWebHttpStruct) RegisterRoutes(server *gin.Engine) {
	searchInterface := server.Group("/cluster")
	searchInterface.POST("/open/perFile", whc.ClusterOpenPerFile)
	//searchInterface.POST("/open/insertPers", whc.ClusterInsertPer)
	//searchInterface.POST("/open/insertApis", whc.ClusterInsertApis)
}

func NewWebHttpCluster(clusterSvc *service.ClusterServiceInterface) *ClusterWebHttpStruct {
	return &ClusterWebHttpStruct{
		ClusterService: clusterSvc,
	}
}

// ClusterOpenPerFile  聚类数据核对
func (whc *ClusterWebHttpStruct) ClusterOpenPerFile(ctx *gin.Context) {
	var req domain.EngineFileDomain
	if err := ctx.Bind(&req); err != nil {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, fmt.Sprintf("ReadJson-ERROR-%v", err), nil)
		return
	}
	if req.FileName == "" {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, "fileName is nil", nil)
		return
	}
	filePath := "./result/ClusterEnginePer/" + req.FileName
	if !file.Exists(filePath) {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, "file is not exist", nil)
		return
	}
	resp := &domain.LoadPerResponse{Msg: "success"}
	ClusterPerOutput, ClusterPerDetail, errInfo := domain.NewOpenExcel(domain.EngineFileDomain{
		FileName:          filePath,
		TestDataType:      req.TestDataType,
		TestGroupType:     req.TestGroupType,
		TestEngineVersion: req.TestEngineVersion,
		TestPerson:        req.TestPerson,
		TestEngineType:    req.TestEngineType,
	}).OpenClusterExcel()
	if errInfo != "" {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, errInfo, nil)
		return
	}
	resp.Count1, resp.Count2 = len(ClusterPerOutput), len(ClusterPerDetail)
	RespondWithJSON(ctx, http.StatusOK, 0, "success", resp)
}

func (whc *ClusterWebHttpStruct) ClusterInsertPer(ctx *gin.Context) {
	req := &domain.ClusterEnginePerFileInsert{}
	if err := ctx.Bind(&req); err != nil {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, fmt.Sprintf("ReadJson-ERROR-%v", err), nil)
		return
	}
	filePath := "./result/ClusterEnginePer/" + req.FileName
	successfulNumber, failuresNumber, err := whc.ClusterService.InsertSvcPerData(ctx,
		domain.NewOpenExcel(domain.EngineFileDomain{
			FileName:          filePath,
			TestDataType:      req.TestDataType,
			TestEngineVersion: req.TestEngineVersion,
			TestPerson:        req.TestPerson,
			TestEngineType:    req.TestEngineType,
		}))
	if err != "" {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, err, nil)
		return
	}
	respMsg := map[string]int64{
		"successfulNumber": successfulNumber,
		"failuresNumber":   failuresNumber,
	}
	RespondWithJSON(ctx, http.StatusOK, 0, "successful insert", respMsg)
}
