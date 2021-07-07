package watch

import (
	"errors"
	"fmt"
	"github.com/woozhijun/flume_exporter/config"
	"net"
	"os/exec"
	"strings"
)

type flumeInfo struct {
	name string
	port string
}

func getFlumeProcess() ([]flumeInfo, error){
	result := make([]flumeInfo,0)
	cmd2 := exec.Command("bash", "-c", "ps -ef | grep Dflume.monitoring.port| grep -v \"color=auto\" |awk '{match($0,/-Dflume.monitoring.port=([0-9]+) .+flume.node.Application -n (agent_.+) -f/,a);if(a[1]&&a[2]){print a[1]\":\"a[2]}}'")
	processInfo := ""
	if output, err := cmd2.CombinedOutput(); err != nil {
		return result, err
	}else{
		processInfo = string(output)
	}
	if len(processInfo)==0{
		fmt.Println("没查到进程")
		return result,errors.New("未查询到flume进程")
	}
	//processInfo := "9600:agent_log_aly\n9601:agent_log_aly\n9602:agent_log_aly\n9627:agent_uuid_location_new\n9614:agent_geek_user_data\n9608:agent_geek_ap_data\n9624:agent_lage_data\n9622:agent_lage_data\n9613:agent_scene_data\n9606:agent_scenelog\n"
	//fmt.Println(processInfo)
	ps := strings.Split(strings.TrimSpace(processInfo), string('\n'))
	for p:= range ps{
		x := strings.Split(ps[p], ":")
		port,name := x[0],x[1]
		result = append(result, flumeInfo{
			name: name,
			port: port,
		})
	}
	return result,nil
}


func localIP() (string,error){
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "",err
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return strings.TrimSpace(ipnet.IP.String()), nil
					}
				}
			}
		}
	}
	return "",errors.New("什么也没发现")
}

func CheckFlume() ([]config.Agent,error){
	result:= make([]config.Agent,0)
	process,err := getFlumeProcess()
	if err != nil{
		return result, err
	}else{
		ip,err := localIP()
		if err != nil {
			fmt.Printf("没找到本机真实IP地址,err=%s",err.Error())
			ip = "localhost"
		}

		// 将多个相同agent名称的，放到一起
		nMap := make(map[string]config.Agent)

		for r := range process{
			url := fmt.Sprintf("http://%s:%s/metrics",ip, process[r].port)
			_,ok := nMap[process[r].name]
			if ok{
				t := nMap[process[r].name]
				t.Urls = append(t.Urls, url)
				nMap[process[r].name] = t
			}else{
				urls := make([]string, 0)
				urls = append(urls, url)
				nMap[process[r].name] = config.Agent{
					Name:    process[r].name,
					Enabled: true,
					Urls:   urls ,
				}
			}
		}
		for _,v := range nMap{
			result = append(result, v)
		}
		return result,nil
	}
}
