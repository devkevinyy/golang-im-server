package util

import (
	"encoding/json"
	"io/ioutil"
)

type IMConfig struct {
	IMHost string       `json:"im_host"`       // IM Server 主机地址
	IMPort int16        `json:"im_port"`       // 端口
	MaxClients int      `json:"max_clients"`   // tcp 最大连接数
	DBConfig DBConfig   `json:"db_config"`     // 数据库配置
}

type DBConfig struct {
	DBHost string      `json:"db_host"`         // 主机地址
	DBPort int         `json:"db_port"`         // 端口
	Username string    `json:"username"`        // 用户名
	Password string    `json:"password"`        // 密码
	DBName string      `json:"db_name"`         // 数据库名称
	MaxIdleConns int   `json:"max_idle_conns"`  // 连接池最大空闲连接数
	MaxOpenConns int   `json:"max_open_conns"`  // 连接池最大连接数
}


/*
	读取配置文件
*/
func ReadConfig(path string) (*IMConfig, error) {
	imConfig := new(IMConfig)
	err := imConfig.ParseConfig(path)
	if err != nil {
		return nil, err
	}
	return imConfig, nil
}

/**
	解析配置文件
 */
func (this *IMConfig) ParseConfig(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &this)
	if err != nil {
		return err
	}
	return nil
}
