package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"
)

var (
	db *sql.DB
)

func newConfig() *Config {
	return &Config{
		App:   newDefaultAPP(),
		MySQL: newDefaultMySQL(),
	}
}

// Config 应用配置
type Config struct {
	App   *App   `toml:"app"`
	MySQL *Mysql `toml:"mysql"`
}

func (c *Config) InitGlobal() error {
	global = c
	return nil
}

type App struct {
	Name      string `toml:"name"`
	OutputDir string `toml:"output_dir"` // 输出文件目录
	ThreadNum int    `toml:"thread_num"` // 线程数
	BatchSize int    `toml:"batch_size"` // 每次查询的记录数
}

func newDefaultAPP() *App {
	return &App{
		Name: "cmdb",
	}
}

type Mysql struct {
	Host        string `toml:"host"`
	Port        string `toml:"port"`
	UserName    string `toml:"username"`
	Password    string `toml:"password"`
	Database    string `toml:"database"`
	Table       string `toml:"table"`
	Columns     string `toml:"columns"`
	MaxOpenConn int    `toml:"max_open_conn"`
	MaxIdleConn int    `toml:"max_idle_conn"`
	MaxLifeTime int    `toml:"max_life_time"`
	MaxIdleTime int    `toml:"max_idle_time"`
	lock        sync.Mutex
}

func newDefaultMySQL() *Mysql {
	return &Mysql{
		Database:    "cmdb-g7",
		Host:        "127.0.0.1",
		Port:        "3306",
		MaxOpenConn: 200,
		MaxIdleConn: 100,
	}
}

// getDBConn use to get db connection pool
func (m *Mysql) getDBConn() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	if m.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	if m.MaxIdleConn != 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}
func (m *Mysql) GetDB() (*sql.DB, error) {
	// 加载全局数据量单例
	m.lock.Lock()
	defer m.lock.Unlock()
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}
