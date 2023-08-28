package demo

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type T struct {
	DeviceNo     string `json:"deviceNo"`
	PointName    string `json:"pointName"`
	DetectDelay  int    `json:"detectDelay"`
	SensorDataBO struct {
		StartTime string `json:"startTime"`
		AdFreq    int    `json:"adFreq"`
		MiscFreq  int    `json:"miscFreq"`
		AdData    []struct {
			Name     string    `json:"name"`
			Value    []float64 `json:"value"`
			ResValue []byte    `json:"resvalue"`
		} `json:"adData"`
		MiscData []struct {
			Name  string `json:"name"`
			Value []int  `json:"value"`
		} `json:"miscData"`
	} `json:"sensorDataBO"`
	ScadaDataBO struct {
		BeatData struct {
			AvlGrindTime                       float64 `json:"avlGrindTime"`
			ZAxisStartTrimmingPosition         float64 `json:"zAxisStartTrimmingPosition"`
			AvgTrimTime                        int     `json:"avgTrimTime"`
			LoadWaitTime                       float64 `json:"loadWaitTime"`
			GrindTotalAmount                   float64 `json:"grindTotalAmount"`
			XAxisStartMachiningPosition        float64 `json:"xAxisStartMachiningPosition"`
			UnloadWaitTime                     float64 `json:"unloadWaitTime"`
			CurrentDiameterOfGrindingWheel     float64 `json:"currentDiameterOfGrindingWheel"`
			FastForwardTime                    float64 `json:"fastForwardTime"`
			DeviceStartTime                    int64   `json:"deviceStartTime"`
			FeedAxisOriginPosition             float64 `json:"feedAxisOriginPosition"`
			CoarseGrindTwo                     float64 `json:"coarseGrindTwo"`
			DeviceEndTime                      int64   `json:"deviceEndTime"`
			FastTrendTime                      float64 `json:"fastTrendTime"`
			CoarseGrindOne                     float64 `json:"coarseGrindOne"`
			Polish                             float64 `json:"polish"`
			FineGrind                          float64 `json:"fineGrind"`
			RetractTime                        float64 `json:"retractTime"`
			GrindingWheelFinishingSerialNumber int     `json:"grindingWheelFinishingSerialNumber"`
		} `json:"beatData"`
		CraftParam struct {
			RotationSpeedOfGrindingWheelShaft                int     `json:"rotationSpeedOfGrindingWheelShaft"`
			FastTrendAmount                                  float64 `json:"fastTrendAmount"`
			StraightLineAfterConvexity                       int     `json:"straightLineAfterConvexity"`
			CenterPositionOfGrindingRack                     int     `json:"centerPositionOfGrindingRack"`
			ArcTrimmingAngle                                 int     `json:"arcTrimmingAngle"`
			CoarseGrindTwoSpeed                              float64 `json:"coarseGrindTwoSpeed"`
			DitchPositionCompensation                        int     `json:"ditchPositionCompensation"`
			TrimmingSpeed                                    int     `json:"trimmingSpeed"`
			ConvexMetric                                     int     `json:"convexMetric"`
			FastTrendSpeed                                   int     `json:"fastTrendSpeed"`
			TrimmingCompensation                             float64 `json:"trimmingCompensation"`
			FastForwardAmount                                float64 `json:"fastForwardAmount"`
			DeviceEndTime                                    string  `json:"deviceEndTime"`
			LinearVelocity                                   int     `json:"linearVelocity"`
			CoarseGrindTwoAmount                             float64 `json:"coarseGrindTwoAmount"`
			SrcReturnSpeed                                   int     `json:"srcReturnSpeed"`
			TrimJumpAmount                                   float64 `json:"trimJumpAmount"`
			CoarseRebound                                    int     `json:"coarseRebound"`
			TrimmingInterval                                 int     `json:"trimmingInterval"`
			LinearSpeedOfGrindingWheel                       int     `json:"linearSpeedOfGrindingWheel"`
			MinimumDiameterOfGrindingWheel                   int     `json:"minimumDiameterOfGrindingWheel"`
			RoughGrindingDelay                               int     `json:"roughGrindingDelay"`
			GrindingDelay                                    int     `json:"grindingDelay"`
			InitialDressingPositionOfZAxisOfNewGrindingWheel int     `json:"initialDressingPositionOfZAxisOfNewGrindingWheel"`
			LoadingAndUnloadingPosition                      int     `json:"loadingAndUnloadingPosition"`
			FineGrindAmount                                  float64 `json:"fineGrindAmount"`
			FastForwardSpeed                                 int     `json:"fastForwardSpeed"`
			GrindingWheelSerialNumber                        int     `json:"grindingWheelSerialNumber"`
			ArcTrimmingSpeed                                 int     `json:"arcTrimmingSpeed"`
			FineRebound                                      int     `json:"fineRebound"`
			CoarseGrindOneAmount                             float64 `json:"coarseGrindOneAmount"`
			DiameterOfNewGrindingWheel                       int     `json:"diameterOfNewGrindingWheel"`
			WorkpieceShaftSpeed                              int     `json:"workpieceShaftSpeed"`
			DressingAmountOfNewGrindingWheel                 float64 `json:"dressingAmountOfNewGrindingWheel"`
			FeedCompensation                                 float64 `json:"feedCompensation"`
			GrindingFrameInPlace                             int     `json:"grindingFrameInPlace"`
			StartingPositionOfXAxisOfNewGrindingWheel        int     `json:"startingPositionOfXAxisOfNewGrindingWheel"`
			FineGrindSpeed                                   float64 `json:"fineGrindSpeed"`
			OscillationSpeed                                 int     `json:"oscillationSpeed"`
			CoarseGrindOneSpeed                              float64 `json:"coarseGrindOneSpeed"`
			OscillationSelection                             string  `json:"oscillationSelection"`
			GrindingWheelTrimmingWidth                       int     `json:"grindingWheelTrimmingWidth"`
			StraightLineBeforeConvexity                      int     `json:"straightLineBeforeConvexity"`
			GrindingFrameHomePosition                        int     `json:"grindingFrameHomePosition"`
			DeviceStartTime                                  string  `json:"deviceStartTime"`
			PolishAmount                                     float64 `json:"polishAmount"`
			TrimFallback                                     int     `json:"trimFallback"`
			GrindingWheelWidth                               int     `json:"grindingWheelWidth"`
			PolishSpeed                                      float64 `json:"polishSpeed"`
			OscillatingDistance                              int     `json:"oscillatingDistance"`
		} `json:"craftParam"`
	} `json:"scadaDataBO"`
}

func InitJson(filePath string) error {
	file, err := os.Open(filePath + "format.json")
	if err != nil {
		log.Println("文件打开错误:", err)
		return err
	}
	defer file.Close()

	// 使用 bufio.Scanner 进行高效的文件读取
	scanner := bufio.NewScanner(file)
	const maxCapacity = 1024 * 1024 * 1024 * 4 // 设置较大的缓冲区大小，根据实际情况调整
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var lines []string
	startTime := time.Now() // 记录开始时间
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	endTime := time.Now() // 记录结束时间
	if err := scanner.Err(); err != nil {
		log.Println("读取文件内容错误:", err)
		return err
	}

	// 将读取的文件内容合并为一个字符串
	content := ""
	for _, line := range lines {
		content += line
	}
	//fmt.Println(content)
	elapsedTime := endTime.Sub(startTime) // 计算读取时间
	log.Println("文件读取完成，耗时:", elapsedTime)

	// 反序列化
	data := T{}
	json.Unmarshal([]byte(content), &data)
	endTime = time.Now() // 记录结束时间
	//fmt.Println(data)
	elapsedTime = endTime.Sub(startTime) // 计算读取时间
	log.Println("JSON反序列化完成，耗时:", elapsedTime)
	// 存入ES中

	//压缩data.SensorDataBO.AdData[0].Value
	//res, err := compressData(data.SensorDataBO.AdData[0].Value)
	//data.SensorDataBO.AdData[0].Value = nil
	//data.SensorDataBO.AdData[0].ResValue = res
	startWriteTime := time.Now()
	WriteToElasticsearch(data)
	endTime = time.Now()                      // 记录结束时间
	elapsedTime = endTime.Sub(startWriteTime) // 计算写入时间
	log.Println("ES写入完成，耗时:", elapsedTime)
	return nil
}

func WriteToElasticsearch(data T) error {
	// 创建 Elasticsearch 客户端
	cfg := elasticsearch.Config{
		Addresses: []string{"http://120.79.2.7:9211"}, // Elasticsearch 地址
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println("创建 Elasticsearch 客户端失败:", err)
		return err
	}

	// 将结构体数据转换为 JSON 字节
	docBytes, err := json.Marshal(data)
	//fmt.Println(string(docBytes))
	if err != nil {
		log.Println("转换为 JSON 失败:", err)
		return err
	}

	// 创建 Elasticsearch 文档请求
	req := esapi.IndexRequest{
		Index:   "ytdata",                  // Elasticsearch 索引名称
		Body:    bytes.NewReader(docBytes), // 文档内容
		Refresh: "true",                    // 可选项，刷新索引以使文档立即可用
	}

	// 执行请求
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Println("执行写入 Elasticsearch 失败:", err)
		return err
	}
	defer res.Body.Close()

	// 检查响应状态
	// 检查响应状态
	if res.IsError() {
		log.Println("写入 Elasticsearch 失败:", res.Status())

		// 读取响应体内容
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("读取响应体失败:", err)
			return fmt.Errorf("写入 Elasticsearch 失败: %s", res.Status())
		}

		log.Println("Elasticsearch 响应内容:", string(resBody))

		return fmt.Errorf("写入 Elasticsearch 失败: %s", res.Status())
	}
	log.Println("写入 Elasticsearch 成功")
	return nil
}

// 压缩数据
func compressData(data []float64) ([]byte, error) {
	// 计算差值编码
	compressed := make([]float64, len(data))
	compressed[0] = data[0]
	for i := 1; i < len(data); i++ {
		compressed[i] = data[i] - data[i-1]
	}

	// 转换为字节切片
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, compressed)
	if err != nil {
		return nil, err
	}

	// 使用gzip进行压缩
	var compressedData bytes.Buffer
	gzipWriter := gzip.NewWriter(&compressedData)
	_, err = gzipWriter.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}
	gzipWriter.Close()

	return compressedData.Bytes(), nil
}

// 解压缩数据
func decompressData(compressed []byte) ([]float64, error) {
	// 使用gzip进行解压缩
	compressedReader, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	defer compressedReader.Close()

	// 读取解压缩的数据
	var buf bytes.Buffer
	_, err = buf.ReadFrom(compressedReader)
	if err != nil {
		return nil, err
	}

	// 转换为float64切片
	decoded := make([]float64, buf.Len()/8) // 每个float64占8字节
	err = binary.Read(&buf, binary.LittleEndian, decoded)
	if err != nil {
		return nil, err
	}

	// 计算原始数值
	data := make([]float64, len(decoded))
	data[0] = decoded[0]
	for i := 1; i < len(decoded); i++ {
		data[i] = data[i-1] + decoded[i]
	}

	return data, nil
}
