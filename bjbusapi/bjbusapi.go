package bjbusapi

import (
	"fmt"

	"github.com/marstion/api-go/jsonvalue"
	"github.com/marstion/api-go/request"
)

var j jsonvalue.J
var header map[string][]string = map[string][]string{
	"Referer":    {`https://www.bjbus.com/api/index.php`},
	"User-Agent": {`Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1`},
}

type Bus struct {
	ApiBus_Data ApiBus_Data
}

type Plan struct {
	Eta         float32 `json:"eta"`         // 剩余时间，单位：秒
	Distance    float32 `json:"distance"`    // 距离， 单位：米
	StationLeft int     `json:"stationLeft"` // 还剩 x 站
	Index       int     `json:"index"`       // 索引
}

func (p Plan) RemainTimeMin() float32 {
	return p.Eta / 60
}

func (p Plan) RemainDistanceKM() float32 {
	return p.Distance / 1000
}

type ApiBus_Data struct {
	EtaToken string `json:"etaToken"`
	// menuList
}

type BusLine struct {
	// 线路: 472
	LineName string
	// 线路内部ID: 000000058311620
	LineId string `json:"lineId"`
	// 起点站: "金安桥北"
	FirstStationName string `json:"firstStationName"`
	// 终点站: 鲁谷公交场站
	LastStationName string `json:"lastStationName"`
	// 运营时间区间, 例: 6:00-23:00
	ServiceTime string `json:"serviceTime"`
	// 夜车 0: 为白车
	IsNight string `json:"isNight"`
	// 上车站:
	IntoStationName string
	intoStation     *BusStation
	// 途径站点
	Stations []*BusStation
}

type BusStation struct {
	LineId     string `json:"lineId"`
	StationId  string `json:"stationId"`
	StopName   string `json:"stopName"`
	StopNumber string `json:"stopNumber"`
}

// 获取token
func (b *Bus) GetToken() {
	url := bjBusDoamin + "/api/api_bus.php"
	response := request.Request("GET", url, header, "")

	j.Unmarshal2Self(response)
	// 取出 json 数据中的data字段, 并解析到 b.ApiBus_Data
	j.Get("data").Unmarshal1Self(&b.ApiBus_Data)
}

// 查询线路到站时间
func (b *Bus) GetLineTime(line *BusLine) (planList []*Plan) {
	url := fmt.Sprintf(bjBusDoamin+"/api/api_etartime.php?conditionstr=%s-%s&token=%s", line.LineId, line.intoStation.StationId, b.ApiBus_Data.EtaToken)
	response := request.Request("GET", url, header, "")

	j.Unmarshal2Self(response)
	for _, v := range j.Get("data", 0, "datas", "trip").Array() {
		var plan Plan
		v.Unmarshal1Self(&plan)
		planList = append(planList, &plan)
	}
	return
}

func (b *Bus) BusLineInfo(line *BusLine) {
	b._info(line)
	b._stations(line)
}

func (bus *Bus) _info(line *BusLine) {
	// url := bjBusDoamin + "/api/api_etaline_list.php?hidden_MapTool=busex2.BusInfo&what=" + line.LineName + "&city=%25u5317%25u4EAC&pageindex=1&pagesize=30&fromuser=bjbus&datasource=bjbus&clientid=9db0f8fcb62eb46c&webapp=mobilewebapp"
	url := fmt.Sprintf(bjBusDoamin+"/api/api_etaline_list.php?hidden_MapTool=busex2.BusInfo&what=%s&city=%25u5317%25u4EAC&pageindex=1&pagesize=30&fromuser=bjbus&datasource=bjbus&clientid=9db0f8fcb62eb46c&webapp=mobilewebapp", line.LineName)
	response := request.Request("GET", url, header, "")

	j.Unmarshal2Self(response)
	for _, v := range j.Get("response", "resultset", "data", "feature").Array() {
		var l *BusLine
		v.Unmarshal1Self(&l)
		if l.LineName == line.LineName &&
			l.FirstStationName == line.FirstStationName &&
			l.LastStationName == line.LastStationName {

			v.Unmarshal1Self(&line)
			break
		}
	}
}

func (bus *Bus) _stations(line *BusLine) {
	url := fmt.Sprintf("https://www.bjbus.com/api/api_etastation.php?lineId=%s&pageNum=1&token=%s", line.LineId, bus.ApiBus_Data.EtaToken)
	response := request.Request("GET", url, header, "")

	line.Stations = make([]*BusStation, 0)
	j.Unmarshal2Self(response)
	for _, v := range j.Get("data").Array() {
		var bs BusStation
		v.Unmarshal1Self(&bs)
		line.Stations = append(line.Stations, &bs)
		if bs.StopName == line.IntoStationName {
			line.intoStation = &bs
		}
	}
}
