package domain

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"gotest/AutomatedTestPlatform/pkg"
	"strconv"
	"strings"
	"time"
)

// OpenExcelSearchEngine  读取2.0格式的性能测试excel
func (oe *OpenExcel) OpenExcelSearchEngine() ([]SearchEnginePerResult, []SearchEngineResultReload, string) {
	IsFinalResult := "false"
	TestDataTime := time.Now().UnixNano() / 1e6
	var res [][]string
	var per1Result []SearchEnginePerResult
	var per2Result []SearchEngineResultReload
	xlFile, err := xlsx.OpenFile(oe.Sep.FileName)
	if err != nil {
		return per1Result, per2Result, err.Error()
	} else {
		for _, sheet := range xlFile.Sheets {
			for _, row := range sheet.Rows {
				cells := []string{}
				for _, cell := range row.Cells {
					if cell.String() != "" {
						cells = append(cells, cell.String())
					} else {
						break
					}
				}
				if len(cells) > 0 {
					res = append(res, cells)
				}
			}
		}
	}
	var sea SearchEnginePerResult
	var load SearchEngineResultReload
	for _, single := range res {
		//fmt.Println(single)
		//fmt.Println(len(single))
		if len(single) == 1 || len(single) > 4 {
			if len(single) == 1 {
				sea = SearchEnginePerResult{}
				if !strings.Contains(single[0], "card") {
					fmt.Printf("case[%s] is not contains card ", single[0])
					return per1Result, per2Result, ""
				}
				sea.TestCaseName = strings.Split(single[0], "(")[0]
				cardNum, _ := strconv.Atoi(strings.Split(strings.Split(single[0], "(")[1], "card")[0])
				sea.TestCardNum = cardNum
				sea.TestSingleName = strings.Split(strings.Split(single[0], "(")[2], ")")[0]
			} else {
				if single[0] != "测试项" {
					if len(single) == 8 {
						Concurrent, _ := strconv.Atoi(strings.ReplaceAll(single[1], " ", ""))
						sea.Concurrent = int64(Concurrent)
						RequestCounts, _ := strconv.Atoi(strings.ReplaceAll(single[3], " ", ""))
						sea.RequestCounts = int64(RequestCounts)
						Tps, _ := strconv.ParseFloat(strings.ReplaceAll(single[4], " ", ""), 64)
						sea.Tps = Tps
						sea.ExtData = single[0]
						MinResponseTime, _ := strconv.ParseFloat(strings.ReplaceAll(single[6], " ", ""), 64)
						sea.MinResponseTime = MinResponseTime
						MaxResponseTime, _ := strconv.ParseFloat(strings.ReplaceAll(single[5], " ", ""), 64)
						sea.MaxResponseTime = MaxResponseTime
						AverageResponseTime, _ := strconv.ParseFloat(strings.ReplaceAll(single[7], " ", ""), 64)
						sea.AverageResponseTime = AverageResponseTime
					} else {
						sea.ExtData = single[0]
						Concurrent, _ := strconv.Atoi(strings.ReplaceAll(single[1], " ", ""))
						sea.Concurrent = int64(Concurrent)
						RequestCounts, _ := strconv.Atoi(strings.ReplaceAll(single[2], " ", ""))
						sea.RequestCounts = int64(RequestCounts)
						AverageResponseTime, _ := strconv.ParseFloat(strings.ReplaceAll(single[3], " ", ""), 64)
						sea.AverageResponseTime = AverageResponseTime
						MediumResponseTime, _ := strconv.ParseFloat(strings.ReplaceAll(single[4], " ", ""), 64)
						sea.MediumResponseTime = MediumResponseTime
						Percent90ResponseTime, _ := strconv.ParseFloat(strings.ReplaceAll(single[5], " ", ""), 64)
						sea.Percent90ResponseTime = Percent90ResponseTime
						Percent95ResponseTime, _ := strconv.ParseFloat(strings.ReplaceAll(single[6], " ", ""), 64)
						sea.Percent95ResponseTime = Percent95ResponseTime
						Percent99ResponseTime, _ := strconv.ParseFloat(strings.ReplaceAll(single[7], " ", ""), 64)
						sea.Percent99ResponseTime = Percent99ResponseTime
						MinResponseTime, _ := strconv.ParseFloat(strings.ReplaceAll(single[8], " ", ""), 64)
						sea.MinResponseTime = MinResponseTime
						MaxResponseTime, _ := strconv.ParseFloat(strings.ReplaceAll(single[9], " ", ""), 64)
						sea.MaxResponseTime = MaxResponseTime
						ErrorPercent, _ := strconv.ParseFloat(strings.ReplaceAll(single[10], " ", ""), 64)
						sea.ErrorPercent = ErrorPercent
						Tps, err := strconv.ParseFloat(strings.ReplaceAll(single[11], " ", ""), 64)
						if err != nil {
							fmt.Println("tps value range to float err!", err.Error())
						}
						sea.Tps = Tps
						sea.ErrInfo = single[12]
						if strings.Contains(single[13], "not contains") {
							sea.PidName = strings.ReplaceAll(single[13], "not contains SearchE&&xqpla&&mysq", "")
						}
						sea.Cpu = single[14]
						sea.Memory = single[15]
						sea.Gpu = single[16]
						sea.GpuMemery = single[17]
					}
					sea.TestDataType = oe.Sep.TestDataType
					sea.TestGroupType = oe.Sep.TestGroupType
					sea.TestEngineVersion = oe.Sep.TestEngineVersion
					sea.TestPerson = oe.Sep.TestPerson
					sea.IsFinalResult = IsFinalResult
					sea.TestDataTime = TestDataTime
					sea.TestEngineType = oe.Sep.TestEngineType
					per1Result = append(per1Result, sea)
					//fmt.Println(sea)
					//fmt.Println(per1Result)
				}
			}
		} else {
			if strings.Contains(single[1], "EngineReloadSpendTime") {
				load = SearchEngineResultReload{}
				load.TestSingleName = "searchEngineReload"
				load.ExtData = single[1]
			} else if strings.Contains(single[1], "reloadSpendTime") {
				load = SearchEngineResultReload{}
				load.TestSingleName = "sharedGroupReload"
				load.ExtData = single[1]
			} else if strings.Contains(single[1], "FeatureDbAutoClean") {
				load = SearchEngineResultReload{}
				load.TestSingleName = single[1]
				load.ExtData = single[1]
			} else if strings.Contains(single[1], "SetSaveSize") {
				load = SearchEngineResultReload{}
				load.TestSingleName = single[1]
				load.ExtData = single[1]
			} else {
				if !strings.Contains(single[0], "card") {
					fmt.Printf("case[%s] is not contains card ", single[0])
					return per1Result, per2Result, ""
				}
				load.TestCaseName = strings.Split(single[0], "(")[0]
				cardNum, _ := strconv.Atoi(strings.Split(strings.Split(single[0], "(")[1], "card")[0])
				load.TestCardNum = cardNum
				featureCount, _ := strconv.ParseInt(single[2], 10, 64)
				load.TestFeatureCount = featureCount
				speed, _ := strconv.ParseFloat(single[3], 64)
				load.Speed = speed
				spendTime, _ := strconv.ParseFloat(single[1], 64)
				load.SpendTime = spendTime
				load.TestDataType = oe.Sep.TestDataType
				load.TestGroupType = oe.Sep.TestGroupType
				load.TestEngineVersion = oe.Sep.TestEngineVersion
				load.TestPerson = oe.Sep.TestPerson
				load.IsFinalResult = IsFinalResult
				load.TestDataTime = TestDataTime
				load.TestEngineType = oe.Sep.TestEngineType
				per2Result = append(per2Result, load)
			}
		}
	}
	return per1Result, per2Result, ""
}

func (oe *OpenExcel) GetApiDataFromFile() ([]SearchEngineApiResult, string) {
	var resultList []SearchEngineApiResult
	err, res := pkg.ReadFileStringLine("./result/SearchEngineApi/" + oe.Sep.FileName)
	if err != nil {
		return nil, err.Error()
	} else {
		for index, value := range res {
			//过滤掉第一行
			if index != 0 {
				var data SearchEngineApiResult
				fmt.Println(index, value)
				tt := strings.Split(value, ",")
				if len(tt) == 4 && len(strings.Split(oe.Sep.FileName, "-")) > 3 {
					data.TestTitle = tt[0]
					data.TestCaseName = tt[1]
					data.TestResult = tt[2]
					data.TestErrInfo = tt[3]
					data.TestEngineVersion = oe.Sep.TestEngineVersion
					data.TestPerson = oe.Sep.TestPerson
					data.IsFinalResult = "false"
					data.TestDataTime = time.Now().UnixNano() / 1e6
					data.TestEngineType = oe.Sep.TestEngineType
					resultList = append(resultList, data)
				}
			}
		}
	}
	return resultList, ""
}

// OpenClusterExcel 读取聚类性能测试csv
func (oe *OpenExcel) OpenClusterExcel() ([]ClusterEnginePerResult, []ClusterEnginePerDetailResult, string) {
	IsFinalResult := "false"
	TestDataTime := time.Now().UnixNano() / 1e6
	resEs := map[string][][]string{}
	var clusterPerDetail []ClusterEnginePerDetailResult
	xlFile, err := xlsx.OpenFile("./result/ClusterEnginePer/" + oe.Sep.FileName)
	if err != nil {
		return nil, nil, err.Error()
	} else {
		for _, sheet := range xlFile.Sheets {
			var res [][]string
			if strings.Contains(sheet.Name, "w") {
				for _, row := range sheet.Rows {
					var cells []string
					for _, cell := range row.Cells {
						if cell.String() != "" {
							cells = append(cells, cell.String())
						} else {
							break
						}
					}
					if len(cells) > 0 {
						res = append(res, cells)
					}
				}
				resEs[sheet.Name] = res
			}

		}
	}
	perdetail := ClusterEnginePerDetailResult{}
	resEssheetNameCount := 0
	var clusterPerOutput = make([]ClusterEnginePerResult, len(resEs))
	for sheetName, resESingle := range resEs {
		var ExtData []string
		for _, single := range resESingle {
			fmt.Println(single)
			if len(single) == 2 {
				ExtData = append(ExtData, single...)
			} else if len(single) == 14 {
				if single[0] == "时间" {
					continue
				}
				perdetail.TestCaseName = fmt.Sprintf("%v_%v", oe.Sep.TestEngineVersion, sheetName)
				perdetail.AllLoadCount, err = strconv.ParseInt(single[1], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				perdetail.AllLabelNum, err = strconv.ParseInt(single[2], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				perdetail.AllSparseFeatureNum, err = strconv.ParseInt(single[3], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				perdetail.Duration, err = strconv.ParseInt(single[4], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				perdetail.LoadInfoDuration, err = strconv.ParseInt(single[5], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				perdetail.AlgoInfoDuration, err = strconv.ParseInt(single[6], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				perdetail.PostProcessInfoDuration, err = strconv.ParseInt(single[7], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				perdetail.ClusterTableDuration, err = strconv.ParseInt(single[8], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				perdetail.ClusterReserveDuration, err = strconv.ParseInt(single[9], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				perdetail.ClusterDetailDuration, err = strconv.ParseInt(single[10], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				perdetail.PushRecogDuration, err = strconv.ParseInt(single[11], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				perdetail.PushKafkaDuration, err = strconv.ParseInt(single[12], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				clusterPerDetail = append(clusterPerDetail, perdetail)
			} else if len(single) == 18 {
				if single[0] == "时间" {
					continue
				}
				clusterPerOutput[resEssheetNameCount].AllLoadCount, err = strconv.ParseInt(single[1], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				clusterPerOutput[resEssheetNameCount].AllLabelNum, err = strconv.ParseInt(single[2], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				clusterPerOutput[resEssheetNameCount].AllRegisterLabelNum, err = strconv.ParseInt(single[3], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				clusterPerOutput[resEssheetNameCount].AllRealLabelNum, err = strconv.ParseInt(single[4], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				clusterPerOutput[resEssheetNameCount].AllFilteredNum, err = strconv.ParseInt(single[5], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				clusterPerOutput[resEssheetNameCount].AllUnFiledNum, err = strconv.ParseInt(single[6], 10, 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				clusterPerOutput[resEssheetNameCount].AllDuration, err = strconv.ParseFloat(single[7], 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				clusterPerOutput[resEssheetNameCount].AllAlgoInfoDuration, err = strconv.ParseFloat(single[8], 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				clusterPerOutput[resEssheetNameCount].AllLoadInfoDuration, err = strconv.ParseFloat(single[9], 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				clusterPerOutput[resEssheetNameCount].AllPostProcessInfoDuration, err = strconv.ParseFloat(single[10], 64)
				if err != nil {
					return nil, nil, err.Error()
				}

				clusterPerOutput[resEssheetNameCount].Tps, err = strconv.ParseFloat(single[11], 64)
				if err != nil {
					return nil, nil, err.Error()
				}
				clusterPerOutput[resEssheetNameCount].Memory = single[12]
				clusterPerOutput[resEssheetNameCount].Cpu = single[13]
				clusterPerOutput[resEssheetNameCount].Gpu = single[14]
				clusterPerOutput[resEssheetNameCount].GpuMemory = single[15]
				clusterPerOutput[resEssheetNameCount].MongoDiskUsage = single[16]
				clusterPerOutput[resEssheetNameCount].MysqlDiskUsage = single[17]
			}
		}
		clusterPerOutput[resEssheetNameCount].TestCardNum, err = strconv.Atoi(strings.Split(strings.Split(sheetName, "_")[2], "card")[0])
		if err != nil {
			return nil, nil, err.Error()
		}
		clusterPerOutput[resEssheetNameCount].ExtData = strings.Join(ExtData, "\n")
		clusterPerOutput[resEssheetNameCount].TestDataType = strings.Split(sheetName, "_")[0]
		clusterPerOutput[resEssheetNameCount].TestCaseName = fmt.Sprintf("%v_%v", oe.Sep.TestEngineVersion, sheetName)
		clusterPerOutput[resEssheetNameCount].TestDataTime = TestDataTime
		clusterPerOutput[resEssheetNameCount].TestEngineType = strings.Split(sheetName, "_")[3]
		clusterPerOutput[resEssheetNameCount].IsFinalResult = IsFinalResult
		resEssheetNameCount++
	}
	return clusterPerOutput, clusterPerDetail, ""
}
