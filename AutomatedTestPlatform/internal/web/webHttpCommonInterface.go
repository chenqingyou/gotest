package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gotest/AutomatedTestPlatform/internal/domain"
	"gotest/AutomatedTestPlatform/internal/service"
	"net/http"
	"os"
	"strings"
)

type Common struct {
	CommSvc *service.CommServiceInterface
}

func NewCommon(commSvc *service.CommServiceInterface) *Common {
	return &Common{
		CommSvc: commSvc,
	}
}

func (comm *Common) CommonRoutes(server *gin.Engine) {
	commInterface := server.Group("/comm")
	commInterface.POST("/uploadFile", comm.UploadFile)
	commInterface.POST("/versionsInquire", comm.VersionsInquire)
	commInterface.POST("/deleteInterface", comm.DeleteInterface)
	commInterface.POST("/updateInterface", comm.UpdateInterface)
	commInterface.POST("/searchInterface", comm.SearchInterface)
}

func (comm *Common) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("uploadfile")
	if err != nil {
		// No file is provided
		RespondWithJSON(ctx, http.StatusBadRequest, -1, "No file is provided", nil)
		return
	}
	// Save the file at the specified path
	savePath := "./result/"
	if strings.Contains(file.Filename, "Sea") {
		savePath = "./result/SearchEngine/"
	} else if strings.Contains(file.Filename, "Clu") {
		savePath = "./result/ClusterEngine/"
	} else if strings.Contains(file.Filename, "Trt") {
		savePath = "./result/TrtEngine/"
	}
	// Create the target save path folder
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		mkdirErr := os.MkdirAll(savePath, 0755)
		if mkdirErr != nil {
			RespondWithJSON(ctx, http.StatusInternalServerError, -1, "Failed to create directory", nil)
			return
		}
	}
	err = ctx.SaveUploadedFile(file, savePath+file.Filename)
	// Deal with situations where the file save failed
	if err != nil {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, "Could not save file!", nil)
		return
	}
	RespondWithJSON(ctx, http.StatusOK, 0, "File uploaded successfully.", nil)
}

func (comm *Common) VersionsInquire(ctx *gin.Context) {
	// Create a new struct to store the client's request data
	var Versions struct {
		EngineType string `json:"engineType"`
		Key        string `json:"key"`
	}
	req := Versions
	// Read and parse the data
	if err := ctx.Bind(&req); err != nil {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, fmt.Sprintf("ReadJson-ERROR-%v", err), nil)
		return
	}
	versionList, errInfo := comm.CommSvc.CommConditionsInquire(ctx, req.EngineType, req.Key)
	if errInfo != "" || len(versionList) == 0 {
		RespondWithJSON(ctx, http.StatusOK, -1, errInfo, nil)
		return
	}
	RespondWithJSON(ctx, http.StatusOK, 0, errInfo, versionList)
}

func (comm *Common) DeleteInterface(ctx *gin.Context) {
	//创建新的结构体 用于存放client的请求数据
	var req struct {
		EngineType string  `json:"engineType"`
		ID         []int64 `json:"id"`
	}
	//读取数据、解析数据
	if err := ctx.Bind(&req); err != nil {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, fmt.Sprintf("ReadJson-ERROR-%v", err), nil)
		return
	}
	errInfo := comm.CommSvc.CommDeleteInterface(ctx, req.EngineType, req.ID)
	if errInfo != "" {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, errInfo, nil)
		return
	}
	RespondWithJSON(ctx, http.StatusOK, 0, "delete Successful", nil)
}

func (comm *Common) UpdateInterface(ctx *gin.Context) {
	//创建新的结构体 用于存放client的请求数据
	var req struct {
		EngineType string  `json:"engineType"`
		ID         []int64 `json:"id"`
		Key        string  `json:"key"`
	}
	//读取数据、解析数据
	if err := ctx.Bind(&req); err != nil {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, fmt.Sprintf("ReadJson-ERROR-%v", err), nil)
		return
	}
	errInfo := comm.CommSvc.CommUpdateInterface(ctx, req.EngineType, req.ID, req.Key)
	if errInfo != "" {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, errInfo, nil)
		return
	}
	RespondWithJSON(ctx, http.StatusOK, 0, "Update Successful", nil)
}

func (comm *Common) SearchInterface(ctx *gin.Context) {
	//创建新的结构体 用于存放client的请求数据
	var req struct {
		EngineType string                 `json:"engineType"`
		Conditions map[string]interface{} `json:"conditions"`
	}
	//读取数据、解析数据
	if err := ctx.Bind(&req); err != nil {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, fmt.Sprintf("ReadJson-ERROR-%v", err), nil)
		return
	}
	errInfo, searchResult := comm.CommSvc.CommInsertInterface(ctx, req.EngineType, req.Conditions)
	if errInfo != "" {
		RespondWithJSON(ctx, http.StatusBadRequest, -1, errInfo, nil)
		return
	}
	RespondWithJSON(ctx, http.StatusOK, 0, "Search Successful", searchResult)
}

func RespondWithJSON(ctx *gin.Context, httpStatus int, errorCode int, message string, data any) {
	ctx.JSON(httpStatus,
		domain.Result{
			Code: errorCode,
			Msg:  message,
			Data: data,
		})
}
