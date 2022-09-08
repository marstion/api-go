package bjbusapi

import (
	"fmt"
	"testing"
)

var bus Bus

func TestGetToken(t *testing.T) {
	bus.GetToken()
	fmt.Printf("bus: %#v\n", bus)
}

func TestBusLineTime(t *testing.T) {
	var line BusLine = BusLine{
		LineName:         "472",
		FirstStationName: "鲁谷公交场站",
		LastStationName:  "金安桥北",
		IntoStationName:  "杨庄路西口",
	}

	bus.GetToken()
	bus.BusLineInfo(&line)
	fmt.Printf("line: %#v\n", line)
	fmt.Printf("上车站点: \n  %#v\n", line.intoStation)
	fmt.Println("沿途站点:")
	for _, v := range line.Stations {
		fmt.Printf("  station: %#v\n", v)
	}
	planList := bus.GetLineTime(&line)
	fmt.Println("当前车辆位置:")
	for _, v := range planList {
		fmt.Printf("  plan: %#v\n", v)
	}
}
