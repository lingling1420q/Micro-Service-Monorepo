package config

func production() {
	Etcd.Addrs = []string{
		"127.0.0.1:23791",
		"127.0.0.1:23792",
		"127.0.0.1:23793",
	}
	Svc1 = service{
		MicroName: "go.micro.service.svc1",
		Version:   "v0.0.1",
		Address:   "0.0.0.0:10001",
		Debug:     false,
		DbURI:     "",
		DbDebug:   false,
	}
	Svc2 = service{
		MicroName: "go.micro.service.svc2",
		Version:   "v0.0.1",
		Address:   "0.0.0.0:10002",
		Debug:     false,
		DbURI:     "",
		DbDebug:   false,
	}
}
