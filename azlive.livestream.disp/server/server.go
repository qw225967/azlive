package server
import (
	"fmt"
	"github.com/azlive/azlive.livestream.disp/config"
	"net/http"
)

func Startserver(conf config.FileConf) error {
	http.HandleFunc("/api/v1/startpush",Startpush) //开始推流
	http.HandleFunc("/api/v1/startpull",Startpull) //开始拉流


	fmt.Printf("service.QueryESConf : http_port:%d , http_ip:%s", conf.Httpconf.Port , conf.Httpconf.Ip)
	addr := fmt.Sprintf("%v:%v",conf.Httpconf.Ip , conf.Httpconf.Port)

	http.ListenAndServe(addr, nil)



	return nil
}