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
	Address  string `yaml:"address" json:"address"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	DB       int    `yaml:"db" json:"db"`
}

// Hertz
// @Description: Hertz配置
type Hertz struct {
	Address         string `yaml:"address" json:"address"`
	EnablePprof     bool   `yaml:"enablePprof" json:"enablePprof"`
	EnableGzip      bool   `yaml:"enableGzip" json:"enableGzip"`
	EnableAccessLog bool   `yaml:"enableAccessLog "json:"enableAccessLog"`
	LogLevel        string `yaml:"logLevel" json:"logLevel"`
	LogFileName     string `yaml:"logFileName" json:"logFileName"`
	LogMaxSize      int    `yaml:"logMaxSize" json:"logMaxSize"`
	LogMaxBackups   int    `yaml:"logMaxBackups" json:"logMaxBackups"`
	LogMaxAge       int    `yaml:"logMaxAge" json:"logMaxAge"`
}

// Mongo
// @Description: Mongo配置
type Mongo struct {
	Address  string `yaml:"address" json:"address"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}
