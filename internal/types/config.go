package types

type MlchainSandboxGlobalConfigurations struct {
	App struct {
		Port  int    `yaml:"port"`
		Debug bool   `yaml:"debug"`
		Key   string `yaml:"key"`
	} `yaml:"app"`
	MaxWorkers     int      `yaml:"max_workers"`
	MaxRequests    int      `yaml:"max_requests"`
	WorkerTimeout  int      `yaml:"worker_timeout"`
	PythonPath     string   `yaml:"python_path"`
	PythonLibPaths []string `yaml:"python_lib_path"`
	NodejsPath     string   `yaml:"nodejs_path"`
	EnableNetwork  bool     `yaml:"enable_network"`
	Proxy          struct {
		Socks5 string `yaml:"socks5"`
		Https  string `yaml:"https"`
		Http   string `yaml:"http"`
	} `yaml:"proxy"`
}
