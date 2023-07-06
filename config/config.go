package config

// MySQL
// @Description: MySQL配置
//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
type MySQL struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Db       string `yaml:"db" json:"db"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Charset  string `yaml:"charset" json:"charset"`
}

// Redis
// @Description: redis配置
type Redis struct {
	Addr     string `yaml:"addr" json:"addr"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	DB       int    `yaml:"db" json:"db"`
}

// Hertz
// @Description: Hertz配置
type Hertz struct {
	Address         string `yaml:"address" json:"address"`
	EnablePprof     bool   `yaml:"enable_pprof" json:"enablePprof"`
	EnableGzip      bool   `yaml:"enable_gzip" json:"enableGzip"`
	EnableAccessLog bool   `yaml:"enable_access_log "json:"enableAccessLog"`
	LogLevel        string `yaml:"log_level" json:"logLevel"`
	LogFileName     string `yaml:"log_file_name" json:"logFileName"`
	LogMaxSize      int    `yaml:"log_max_size" json:"logMaxSize"`
	LogMaxBackups   int    `yaml:"log_max_backups" json:"logMaxBackups"`
	LogMaxAge       int    `yaml:"log_max_age" json:"logMaxAge"`
}
