package server

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func Startpush(w http.ResponseWriter, req *http.Request) {
	traceid := req.Header.Get("X-TraceId")
	body, _ := ioutil.ReadAll(req.Body)
	path := "http://127.0.0.1:89/v1/mp/stream/startpublish?appkey=storemanager"
	ret, err := httpPostJson(path, traceid, body)
	if err != nil {
		io.WriteString(w, string(ret))
	}

	io.WriteString(w, string(ret))
}




func Startpull(w http.ResponseWriter, req *http.Request) {
	traceid := req.Header.Get("X-TraceId")
	body, _ := ioutil.ReadAll(req.Body)
	path := "http://127.0.0.1:89/v1/mp/stream/startplay?appkey=storemanager"
	ret, err := httpPostJson(path, traceid, body)
	if err != nil {
		io.WriteString(w, string(ret))
	}

	io.WriteString(w, string(ret))

}

func httpPostJson(path , traceid string, data []byte) ([]byte, error){

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-TraceId", traceid)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		return []byte(""), err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}