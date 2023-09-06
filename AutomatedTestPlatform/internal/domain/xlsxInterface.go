package domain

type XlsxInterface interface {
	OpenExcelSearchEngine() ([]SearchEnginePerResult, []SearchEngineResultReload, string)
	GetApiDataFromFile() ([]SearchEngineApiResult, string)
	OpenClusterExcel() ([]ClusterEnginePerResult, []ClusterEnginePerDetailResult, string)
}

type OpenExcel struct {
	Sep EngineFileDomain
}

func NewOpenExcel(Sep EngineFileDomain) *OpenExcel {
	return &OpenExcel{Sep: Sep}
}
