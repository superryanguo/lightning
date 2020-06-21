package config

// MysqlConfig mysql 配置 接口
type MysqlConfig interface {
	GetURL() string
	GetPsw() string
	GetDbName() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
	GetConnMaxLifetime() int
}

// defaultMysqlConfig mysql 配置
type defaultMysqlConfig struct {
	URL               string `json:"url"`
	Psw               string `json:"psw"`
	DbName            string `json:"dbname"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
	ConnMaxLifetime   int    `json:"connMaxLifetime"`
}

// URL mysql 连接
func (m defaultMysqlConfig) GetURL() string {
	return m.URL
}

func (m defaultMysqlConfig) GetPsw() string {
	return m.Psw
}

func (m defaultMysqlConfig) GetDbName() string {
	return m.DbName
}

// Enabled 激活
func (m defaultMysqlConfig) GetEnabled() bool {
	return m.Enable
}

// 闲置连接数
func (m defaultMysqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

// 打开连接数
func (m defaultMysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}

// 连接数断开时间
func (m defaultMysqlConfig) GetConnMaxLifetime() int {
	return m.ConnMaxLifetime
}
