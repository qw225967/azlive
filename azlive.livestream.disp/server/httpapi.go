package server

import (
	"encoding/json"
	"fmt"
	"github.com/azlive/azlive.livestream.disp/config"
	"io"
	"io/ioutil"
	"net/http"
)

func Startpush(w http.ResponseWriter, req *http.Request) {
	var (
		BodyInfo config.UrlInfo
		ResInfo	 config.StartQueryRes
	)

	body, _ := ioutil.ReadAll(req.Body)
	umsh_err := json.Unmarshal([]byte(body), &BodyInfo)
	fmt.Printf("show the body , Streamid:%+v ",BodyInfo.Streamid)
	if umsh_err != nil {
		fmt.Printf("StartPush request Unmarshal error")
		ResInfo.Dm_error = 10000
		ResInfo.Error_msg = "failed"
		ret, _ := json.Marshal(ResInfo)
		io.WriteString(w, string(ret))
	}
	ResInfo.Dm_error = 200
	ResInfo.Error_msg = "success"
	ResInfo.Data.Ip = "127.0.0.1"
	ResInfo.Data.Port = 9090
	ResInfo.Data.Hostname = "azLive1"
	ret, _ := json.Marshal(ResInfo)
	io.WriteString(w, string(ret))
}




func Startpull(w http.ResponseWriter, req *http.Request) {
	var (
		BodyInfo config.UrlInfo
		ResInfo	 config.StartQueryRes
	)

	body, _ := ioutil.ReadAll(req.Body)
	umsh_err := json.Unmarshal([]byte(body), &BodyInfo)
	fmt.Printf("show the body , Streamid:%+v ",BodyInfo.Streamid)
	if umsh_err != nil {
		fmt.Printf("StartPull request Unmarshal error")
		ResInfo.Dm_error = 10000
		ResInfo.Error_msg = "failed"
		ret, _ := json.Marshal(ResInfo)
		io.WriteString(w, string(ret))
	}
	ResInfo.Dm_error = 200
	ResInfo.Error_msg = "success"
	ResInfo.Data.Ip = "127.0.0.1"
	ResInfo.Data.Port = 9090
	ResInfo.Data.Hostname = "azLive1"
	ret, _ := json.Marshal(ResInfo)
	io.WriteString(w, string(ret))

}