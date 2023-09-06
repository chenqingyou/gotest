package domain

type UploadApiResponse struct {
	Msg      string `json:"msg"`
	Result   int    `json:"result"`
	FileName string `json:"fileName"`
}

type EngineFileDomain struct {
	FileName          string `json:"fileName"`
	TestDataType      string `json:"testDataType"`
	TestGroupType     string `json:"testGroupType"`
	TestEngineVersion string `json:"testEngineVersion"`
	TestPerson        string `json:"testPerson"`
	TestEngineType    string `json:"testEngineType"`
}

type SearchEnginePerFileInsert struct {
	FileName           string `json:"fileName"`
	IsInsertAll        bool   `json:"isInsertAll"`
	InsertIdsInterface []int  `json:"insertIds_interface"`
	InsertIdsOther     []int  `json:"insertIds_other"`
	TestDataType       string `json:"testDataType"`
	TestGroupType      string `json:"testGroupType"`
	TestEngineVersion  string `json:"testEngineVersion"`
	TestPerson         string `json:"testPerson"`
	TestEngineType     string `json:"testEngineType"`
}

// SearchEnginePerResult 识别引擎性能结果结构体
type SearchEnginePerResult struct {
	Index                 int     `json:"index"`
	Id                    int64   `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	TestEngineVersion     string  `json:"testEngineVersion"`
	TestPerson            string  `json:"testPerson"`
	TestDataTime          int64   `json:"testDataTime"`
	TestCaseName          string  `json:"testCaseName"`
	TestSingleName        string  `json:"testSingleName"`
	TestDataType          string  `json:"testDataType"`
	TestGroupType         string  `json:"testGroupType"`
	TestEngineType        string  `json:"testEngineType"`
	TestCardNum           int     `json:"testCardNum"`
	Concurrent            int64   `json:"concurrent"`
	RequestCounts         int64   `json:"requestCounts"`
	AverageResponseTime   float64 `json:"averageResponseTime"`
	MediumResponseTime    float64 `json:"mediumResponseTime"`
	Percent90ResponseTime float64 `json:"percent90ResponseTime"`
	Percent95ResponseTime float64 `json:"percent95ResponseTime"`
	Percent99ResponseTime float64 `json:"percent99ResponseTime"`
	MinResponseTime       float64 `json:"minResponseTime"`
	MaxResponseTime       float64 `json:"maxResponseTime"`
	ErrorPercent          float64 `json:"errorPercent"`
	Tps                   float64 `json:"tps"`
	ErrInfo               string  `json:"errInfo"`
	PidName               string  `json:"pidName"`
	Cpu                   string  `json:"cpu"`
	Memory                string  `json:"memory"`
	Gpu                   string  `json:"gpu"`
	GpuMemery             string  `json:"gpuMemery"`
	ExtData               string  `json:"extData"`
	IsFinalResult         string  `json:"isFinalResult"`
}

// SearchEngineResultReload 识别引擎重载结构体
type SearchEngineResultReload struct {
	Index             int     `json:"index"`
	Id                int64   `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	TestEngineVersion string  `json:"testEngineVersion"`
	TestPerson        string  `json:"testPerson"`
	TestDataTime      int64   `json:"testDataTime"`
	TestCaseName      string  `json:"testCaseName"`
	TestSingleName    string  `json:"testSingleName"`
	TestDataType      string  `json:"testDataType"`
	TestGroupType     string  `json:"testGroupType"`
	TestEngineType    string  `json:"testEngineType"`
	TestCardNum       int     `json:"testCardNum"`
	TestFeatureCount  int64   `json:"testFeatureCount"`
	SpendTime         float64 `json:"spendTime"`
	Speed             float64 `json:"speed"`
	ExtData           string  `json:"extData"`
	IsFinalResult     string  `json:"isFinalResult"`
}

// LoadPerResponse 识别数据请求结构体
type LoadPerResponse struct {
	Msg    string                     `json:"msg"`
	Data1  []SearchEnginePerResult    `json:"interfaceData"`
	Data2  []SearchEngineResultReload `json:"otherData"`
	Result int                        `json:"result"`
	Count1 int                        `json:"interfaceCount"`
	Count2 int                        `json:"otherCount"`
}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type SearchEngineApiResult struct {
	Index             int    `json:"index"`
	Id                int64  `json:"id" gorm:"primaryKey,autoIncrement"`
	TestEngineVersion string `json:"testEngineVersion"`
	TestPerson        string `json:"testPerson"`
	TestDataTime      int64  `json:"testDataTime"`
	TestEngineType    string `json:"testEngineType"`
	TestTitle         string `json:"testTitle"`
	TestCaseName      string `json:"testCaseName"`
	TestResult        string `json:"testResult"`
	TestErrInfo       string `json:"testErrInfo"`
	ExtData           string `json:"extData"`
	IsFinalResult     string `json:"isFinalResult"`
}

type SearchEngineAlgoResult struct {
	Index             int    `json:"index"`
	Id                int64  `json:"id" gorm:"primaryKey,autoIncrement"`
	TestEngineVersion string `json:"testEngineVersion"`
	TestPerson        string `json:"testPerson"`
	TestDataTime      int64  `json:"testDataTime"`
	TestEngineType    string `json:"testEngineType"`
	TestTitle         string `json:"testTitle"`
	TestCaseName      string `json:"testCaseName"`
	TestResult        string `json:"testResult"`
	TestErrInfo       string `json:"testErrInfo"`
	ExtData           string `json:"extData"`
	IsFinalResult     string `json:"isFinalResult"`
}

type SearchEngineApiFileInsert struct {
	FileName          string `json:"fileName"`
	IsInsertAll       bool   `json:"isInsertAll"`
	InsertIds         []int  `json:"insertIds_interface"`
	TestEngineVersion string `json:"testEngineVersion"`
	TestPerson        string `json:"testPerson"`
	TestEngineType    string `json:"testEngineType"`
}

type ClusterEngineApiInsert struct {
	FileName          string `json:"fileName"`
	IsInsertAll       bool   `json:"isInsertAll"`
	InsertIds         []int  `json:"insertIds_interface"`
	TestEngineVersion string `json:"testEngineVersion"`
	TestPerson        string `json:"testPerson"`
	TestEngineType    string `json:"testEngineType"`
}

type ClusterEnginePerInsert struct {
	FileName          string `json:"fileName"`
	IsInsertAll       bool   `json:"isInsertAll"`
	InsertIds         []int  `json:"insertIds_interface"`
	TestEngineVersion string `json:"testEngineVersion"`
	TestPerson        string `json:"testPerson"`
	TestEngineType    string `json:"testEngineType"`
}

type ClusterEngineAlgoInsert struct {
	FileName          string `json:"fileName"`
	IsInsertAll       bool   `json:"isInsertAll"`
	InsertIds         []int  `json:"insertIds_interface"`
	TestEngineVersion string `json:"testEngineVersion"`
	TestPerson        string `json:"testPerson"`
	TestEngineType    string `json:"testEngineType"`
}

type ClusterEnginePerResult struct {
	Id                         int64   `json:"id,omitempty" gorm:"primaryKey,autoIncrement"`
	TestEngineVersion          string  `json:"testEngineVersion"`
	TestPerson                 string  `json:"testPerson"`
	TestDataTime               int64   `json:"testDataTime"`
	TestCaseName               string  `json:"testCaseName"`
	TestEngineType             string  `json:"testEngineType"`
	TestDataType               string  `json:"testDataType"`
	TestCardNum                int     `json:"testCardNum"`
	AllLoadCount               int64   `json:"allLoadCount"`               //总加载特征量
	AllLabelNum                int64   `json:"allLabelNum"`                //总实际档案量
	AllRegisterLabelNum        int64   `json:"allRegisterLabelNum"`        //总注册照档案个数
	AllRealLabelNum            int64   `json:"allRealLabelNum"`            //总实名档案个数
	AllFilteredNum             int64   `json:"allFilteredNum"`             //总被过滤特征总数
	AllUnFiledNum              int64   `json:"allUnFiledNum"`              //总未归档特征总数
	AllDuration                float64 `json:"allDuration"`                //总耗时
	AllAlgoInfoDuration        float64 `json:"allAlgoInfoDuration"`        //总算法耗时
	AllLoadInfoDuration        float64 `json:"allLoadInfoDuration"`        //总加载耗时
	AllPostProcessInfoDuration float64 `json:"allPostProcessInfoDuration"` //总后处理耗时
	Tps                        float64 `json:"tps"`                        //总后处理耗时
	Cpu                        string  `json:"cpu"`
	Memory                     string  `json:"memory"`
	Gpu                        string  `json:"gpu"`
	GpuMemory                  string  `json:"GpuMemory"`
	MongoDiskUsage             string  `json:"mongoDiskUsage"`
	MysqlDiskUsage             string  `json:"mysqlDiskUsage"`
	Ip                         int64   `json:"ip"`
	//SingleDetailData           []ClusterEnginePerDetailResult `json:"singleDetailData"` //每批次的详细数据以及汇总后的数据结果
	ExtData       string `json:"extData"` //理论和实际数据量对比
	IsFinalResult string `json:"isFinalResult"`
}

type ClusterEnginePerDetailResult struct {
	TestCaseName              string `json:"TestCaseName"`              //特征累计加载总数
	AllLoadCount              int64  `json:"allLoadCount"`              //特征累计加载总数
	AllLabelNum               int64  `json:"allLabelNum"`               //累加档案量
	AllSparseFeatureNum       int64  `json:"allSparseFeatureNum"`       //每批次加载特征总数
	Duration                  int64  `json:"duration"`                  //每批次总时长
	LoadInfoDuration          int64  `json:"loadInfoDuration"`          //每批次数据加载耗时
	AlgoInfoDuration          int64  `json:"algoInfoDuration"`          //每批次算法耗时
	PostProcessInfoDuration   int64  `json:"postProcessInfoDuration"`   //每批次预处理耗时
	ClusterPreProcDuration    int64  `json:"clusterPreProcDuration"`    //每批次后处理耗时
	ClusterTableDuration      int64  `json:"clusterTableDuration"`      //每批次存档案耗时
	ClusterReserveDuration    int64  `json:"clusterReserveDuration"`    //每批次存代表照耗时
	ClusterDetailDuration     int64  `json:"clusterDetailDuration"`     //每批次存明细耗时
	ClusterDenseTrackDuration int64  `json:"clusterDenseTrackDuration"` //每批次稠密轨迹细耗时
	PushRecogDuration         int64  `json:"pushRecogDuration"`         //每批次推识别耗时
	PushKafkaDuration         int64  `json:"pushKafkaDuration"`         //每批次推kafka耗时
}

type ClusterEnginePerFileInsert struct {
	FileName            string `json:"fileName"`
	TestEngineVersion   string `json:"testEngineVersion"`
	TestPerson          string `json:"testPerson"`
	TestDataType        string `json:"testDataType"`
	TestEngineType      string `json:"testEngineType"`
	ExtData             string `json:"extData"`
	IsFinalResult       string `json:"isFinalResult"`
	IsInsertAll         bool   `json:"isInsertAll"`
	InsertIdsPer        []int  `json:"InsertIdsPer"`
	InsertIdsPerDetails []int  `json:"InsertIdsPerDetails"`
}
