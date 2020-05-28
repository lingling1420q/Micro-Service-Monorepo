package config

func production() {
	Etcd.Addrs = []string{
		"127.0.0.1:23791",
		"127.0.0.1:23792",
		"127.0.0.1:23793",
	}
	Srv1 = srv{
		MicroName: "go.micro.srv.srv1",
		Version:   "v0.0.1",
		Address:   "0.0.0.0:10001",
		Debug:     false,
		DbURI:     "",
		DbDebug:   false,
	}
}
