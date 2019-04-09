package app

type Config struct {
	Server  Server  `yaml:"server"`
	Storage Storage `yaml:"storage"`
	App     App     `yaml:"app"`
}

type Server struct {
	Port     string   `yaml:"port"`
	Ws       Ws       `yaml:"ws"`
	Shutdown Shutdown `yaml:"shutdown"`
	Timout   Timeout  `yaml:"timeout"`
}

type Storage struct {
	Path string `yaml:"path"`
}

type App struct {
	Report Report `yaml:"report"`
}

type Ws struct {
	Connections int    `yaml:"connections"`
	Endpoint    string `yaml:"endpoint"`
}

type Shutdown struct {
	Endpoint string `yaml:"endpoint"`
	Timeout  int    `yaml:"timeout"`
}

type Timeout struct {
	Read  int `yaml:"read"`
	Write int `yaml:"write"`
}

type Report struct {
	Time int `yaml:"time"`
}
