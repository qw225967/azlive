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

// --------------------- publish ---------------------

type PublishParams struct {
	AppId        string         `json:"appId"`
	RoomId       string         `json:"roomId"`
	UserId       string         `json:"userId"`
	UserName     string         `json:"userName"`
	ServerId     string         `json:"serverId"`
	MediaStreams []*MediaStream `json:"mediaStreams"`
}

type MediaStream struct {
	StreamId     string `protobuf:"bytes,1,opt,name=streamId,proto3" json:"streamId,omitempty"`          // 媒体流id 全局唯一
	Ssrc         uint32 `protobuf:"varint,2,opt,name=ssrc,proto3" json:"ssrc,omitempty"`                 // 媒体流对应的ssrc 全局唯一
	MediaType    string `protobuf:"bytes,3,opt,name=mediaType,proto3" json:"mediaType,omitempty"`        // 媒体类型 audio|video|screen
	StreamDesc   string `protobuf:"bytes,4,opt,name=streamDesc,proto3" json:"streamDesc,omitempty"`      // 媒体描述,未来可能多video场景
	Muted        uint32 `protobuf:"varint,5,opt,name=muted,proto3" json:"muted,omitempty"`               // 推流状态是否静音 1 非静音 2 静音 0 默认没有设置此参数,sdk将忽略
	MutedDetails uint32 `protobuf:"varint,6,opt,name=mutedDetails,proto3" json:"mutedDetails,omitempty"` // 推流状态静音详细解释
	PayloadType  uint32 `protobuf:"varint,7,opt,name=payloadType,proto3" json:"payloadType,omitempty"`   // 媒体编码96
	MaxBitrate   uint32 `protobuf:"varint,8,opt,name=maxBitrate,proto3" json:"maxBitrate,omitempty"`     // 媒体流对应的码率
	ExtraData    string `protobuf:"bytes,9,opt,name=extraData,proto3" json:"extraData,omitempty"`        // 扩展数据
}
