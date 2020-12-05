package config

type FileConf struct {
	Httpconf httpconf
	Azlsconf azlsconf
}

type httpconf struct {
	Port int32
	Ip   string
}

type azlsconf struct {
	Port     []int32
	Ip       []string
	Hostname []string
}

//--------------------- url struct -----------------

type UrlInfo struct {
	Liveid    string
	Uid       string
	StartTime string
	Streamid  string
}

//-------------------   reqData  --------------------
type StartQueryRes struct {
	Dm_error  int     `json:"dm_error"`
	Error_msg string  `json:"error_msg"`
	Data      ResData `json:"data"`
}

type ResData struct {
	Hostname string `json:"hostname"`
	Ip       string `json:"ip"`
	Port     int32  `json:"port"`
}
