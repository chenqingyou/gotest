package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wxnacy/wgo/file"
	"gotest/AutomatedTestPlatform/internal/domain"
	"gotest/AutomatedTestPlatform/internal/service"
	"net/http"
)

type SearchHttpStruct struct {
	SearchService *service.SearchServiceInterface
}

func (shi *SearchHttpStruct) RegisterRoutes(server *gin.Engine) {
	searchInterface := server.Group("/search")
	searchInterface.POST("/open/perFile", shi.SearchOpenPerFile)
	searchInterface.POST("/open/insertPers", shi.SearchInsertPer)
	searchInterface.POST("/open/insertApis", shi.SearchInsertPer)
}

func NewSearchHttpInterface(SearchService *service.SearchServiceInterface) *SearchHttpStruct {
	return &SearchHttpStruct{
		SearchService: SearchService,
	}
}

// SearchOpenPerFile 识别数据核对
func (shi *SearchHttpStruct) SearchOpenPerFile(ctx *gin.Context) {
	var req domain.EngineFileDomain
	if err := ctx.Bind(&req); err != nil {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, fmt.Sprintf("ReadJson-ERROR-%v", err), nil)
		return
	}
	if req.FileName == "" {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, "fileName is nil", nil)
		return
	}
	filePath := "./result/SearchEngine/" + req.FileName
	if !file.Exists(filePath) {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, "file is not exist", nil)
		return
	}
	resp := &domain.LoadPerResponse{Msg: "success"}
	per1Res, per2Res, errInfo := domain.NewOpenExcel(domain.EngineFileDomain{
		FileName:          filePath,
		TestDataType:      req.TestDataType,
		TestGroupType:     req.TestGroupType,
		TestEngineVersion: req.TestEngineVersion,
		TestPerson:        req.TestPerson,
		TestEngineType:    req.TestEngineType,
	}).OpenExcelSearchEngine()
	if errInfo != "" {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, errInfo, nil)
		return
	}
	resp.Count1, resp.Count2 = len(per1Res), len(per2Res)
	RespondWithJSON(ctx, http.StatusOK, 0, "success", resp)
}

// SearchInsertPer 识别性能数据插入
func (shi *SearchHttpStruct) SearchInsertPer(ctx *gin.Context) {
	req := &domain.SearchEnginePerFileInsert{}
	if err := ctx.Bind(&req); err != nil {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, fmt.Sprintf("ReadJson-ERROR-%v", err), nil)
		return
	}
	filePath := "./result/SearchEngine/" + req.FileName
	successfulNumber, failuresNumber, err := shi.SearchService.InsertSvcPerData(ctx,
		domain.NewOpenExcel(domain.EngineFileDomain{
			FileName:          filePath,
			TestDataType:      req.TestDataType,
			TestGroupType:     req.TestGroupType,
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

// SearchInsertApis 识别接口测试数据插入
func (shi *SearchHttpStruct) SearchInsertApis(ctx *gin.Context) {
	req := &domain.SearchEngineApiFileInsert{}
	if err := ctx.Bind(&req); err != nil {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, fmt.Sprintf("ReadJson-ERROR-%v", err), nil)
		return
	}
	filePath := "./result/SearchEngine/" + req.FileName
	successfulNumber, failuresNumber, err := shi.SearchService.InsertSvcApiData(ctx,
		domain.NewOpenExcel(domain.EngineFileDomain{
			FileName:          filePath,
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
