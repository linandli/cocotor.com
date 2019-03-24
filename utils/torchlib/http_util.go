package torchlib

import (
	// "fmt"
	"time"
	"log"
	"net/url"
	"strconv"
	// "reflect"

	"github.com/linandli/requests"
)

const (
	baseUrl = `http://115.28.129.229:8080/powerMobileApp/`
	userName = `admin`
	passWord = `www.togeek.cn`
)


func Get(path string, params map[string]string, timeout time.Duration) string {
	req := requests.Requests()
	req.SetTimeout(timeout)
	resp, err := req.Get(baseUrl + path, params, requests.Auth{userName, passWord})

	if err != nil {
	   log.Println(err)
	   return ""
	}

	return resp.Text()
}

func Post(path string, params map[string]string, datas map[string]string, timeout time.Duration) string {
	req := requests.Requests()
	req.SetTimeout(timeout)
	req.Header.Set("Content-Type", "application/json")
	_params := requests.Params{}
	_params = params
	_data := requests.Datas{}
	_data = datas
	resp, err := req.Post(baseUrl + path, _params, _data, requests.Auth{userName, passWord})

	if err != nil {
	   log.Println(err)
	   return ""
	}
	
	return resp.Text()
}

func Put(path string, params map[string]string, datas map[string]string, timeout time.Duration) string {
	req := requests.Requests()
	req.SetTimeout(timeout)
	req.Header.Set("Content-Type", "application/json")
	_params := requests.Params{}
	_params = params
	_data := requests.Datas{}
	_data = datas
	resp, err := req.Put(baseUrl + path, _params, _data, requests.Auth{userName, passWord})

	if err != nil {
	   log.Println(err)
	   return ""
	}
	
	return resp.Text()
}

func Delete(path string, params map[string]string, datas map[string]string, timeout time.Duration) string {
	req := requests.Requests()
	req.SetTimeout(timeout)
	_params := requests.Params{}
	_params = params
	_data := requests.Datas{}
	_data = datas
	resp, err := req.Delete(baseUrl + path, _params, _data, requests.Auth{userName, passWord})

	if err != nil {
	   log.Println(err)
	   return ""
	}
	
	return resp.Text()
}

func TransGet(plantId int, path string, proxy string, params map[string]string, timeout time.Duration) string {
	_plantId := strconv.Itoa(plantId)
	res := map[string]string{"plants": _plantId, "uri": path}

	if params != nil {
		urlValues := url.Values{}
		for k, v := range params {
			urlValues.Set(k, v)
		}
		res["uri"] = res["uri"] + "?" + urlValues.Encode()
	}
	
	_params := requests.Params{}
	_params = res
	req := requests.Requests()
	req.SetTimeout(timeout)	
	resp, err := req.Get(baseUrl + "distribute/trans", _params, requests.Auth{userName, passWord})

	if err != nil {
	   log.Println(err)
	   return ""
	}
	return resp.Text()
}

func TransPost(plantId int, path string, proxy string, params map[string]string, datas map[string]string, timeout time.Duration) string {
	_plantId := strconv.Itoa(plantId)
	res := map[string]string{"plants": _plantId, "uri": path}

	if params != nil {
		urlValues := url.Values{}
		for k, v := range params {
			urlValues.Set(k, v)
		}
		res["uri"] = res["uri"] + "?" + urlValues.Encode()
	}
	
	_params := requests.Params{}
	_params = res
	_data := requests.Datas{}
	_data = datas
	req := requests.Requests()
	req.SetTimeout(timeout)
	req.Header.Set("Content-Type", "application/json")
	resp, err := req.Post(baseUrl + "distribute/trans", _params, _data, requests.Auth{userName, passWord})

	if err != nil {
	   log.Println(err)
	   return ""
	}
	return resp.Text()
}

func TransDelete(plantId int, path string, proxy string, params map[string]string, datas map[string]string, timeout time.Duration) string {
	_plantId := strconv.Itoa(plantId)
	res := map[string]string{"plants": _plantId, "uri": path}

	if params != nil {
		urlValues := url.Values{}
		for k, v := range params {
			urlValues.Set(k, v)
		}
		res["uri"] = res["uri"] + "?" + urlValues.Encode()
	}
	
	_params := requests.Params{}
	_params = res
	_data := requests.Datas{}
	_data = datas
	req := requests.Requests()
	req.SetTimeout(timeout)
	req.Header.Set("Content-Type", "application/json")
	resp, err := req.Delete(baseUrl + "distribute/trans", _params, _data, requests.Auth{userName, passWord})

	if err != nil {
	   log.Println(err)
	   return ""
	}
	return resp.Text()
}

func TransPut(plantId int, path string, proxy string, params map[string]string, datas map[string]string, timeout time.Duration) string {
	_plantId := strconv.Itoa(plantId)
	res := map[string]string{"plants": _plantId, "uri": path}

	if proxy == "true" {
		res["proxy"] = "true"
	}

	if params != nil {
		urlValues := url.Values{}
		for k, v := range params {
			urlValues.Set(k, v)
		}
		res["uri"] = res["uri"] + "?" + urlValues.Encode()
	}
	
	_params := requests.Params{}
	_params = res
	_data := requests.Datas{}
	_data = datas
	req := requests.Requests()
	req.SetTimeout(timeout)
	req.Header.Set("Content-Type", "application/json")
	resp, err := req.Put(baseUrl + "distribute/trans", _params, _data, requests.Auth{userName, passWord})

	if err != nil {
	   log.Println(err)
	   return ""
	}
	return resp.Text()
}

// func main() {
	// get test
	// fmt.Println(get("org/plants", nil, 60))

	// post test
	// _d := map[string]string{"identifier": "rizhao_8", "secret": "789@xyz"}
	// _p := map[string]string{"targetUri": "/powerMobileApp/resources/torchserver/index"}
	// fmt.Println(post("login", _p, _d, 60))

	//trans_get test
	// 'powerMobileApp/config/1/define/points'
	// plantId := 65
	// _p := map[string]string{"targetUri": "/powerMobileApp/resources/torchserver/index"}

	// fmt.Println(trans_get(plantId, "/powerMobileApp/config/" + strconv.Itoa(plantId) + "/config/points", nil, 60))

	//trans_post test
	// plantId := 65
	// fmt.Println(trans_post(plantId, ))
// }