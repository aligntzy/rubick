package location

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type TaobaoLocation struct {
	Data struct {
		Area      string      `json:"area"`
		Country   string      `json:"country"`
		IspID     string      `json:"isp_id"`
		QueryIP   string      `json:"queryIp"`
		City      string      `json:"city"`
		IP        string      `json:"ip"`
		Isp       string      `json:"isp"`
		County    string      `json:"county"`
		RegionID  string      `json:"region_id"`
		AreaID    string      `json:"area_id"`
		CountyID  interface{} `json:"county_id"`
		Region    string      `json:"region"`
		CountryID string      `json:"country_id"`
		CityID    string      `json:"city_id"`
	} `json:"data"`
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func Location(ip string) (*TaobaoLocation, error) {
	resp, err := http.Get(fmt.Sprintf("https://ip.taobao.com/service/getIpInfo.php?ip=%s", ip))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var loc TaobaoLocation
	err = json.Unmarshal(b, &loc)
	if err != nil {
		return nil, err
	}

	if loc.Code != 0 {
		return nil, errors.New(loc.Msg)
	}

	return &loc, nil
}
