package server

type XServerConfig struct {
	Name string
}

//自定义server
type XServer struct {
	Name string
	Run  func()
}

//标准配置
func StdConfig(serverName string) XServerConfig {
	config := XServerConfig{
		Name: serverName,
	}
	return config
}

func (stdConfig XServerConfig) Build(fun func()) (*XServer, error) {
	server := XServer{
		Name: stdConfig.Name,
		Run:  fun,
	}
	return &server, nil
}
