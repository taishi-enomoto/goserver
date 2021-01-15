package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

type ConfDB struct {
	Host    string `toml:"host"`    // ホスト名
	Port    int    `toml:"port"`    // ポート番号
	DbName  string `toml:"db-name"` // 接続先DB名
	Charset string `toml:"charset"` // 文字コード
	User    string `toml:"user"`    // 接続ユーザ名
	Pass    string // 接続パスワード
}

type Config struct {
	ConfDB `toml:"database"`
}

// URL設定の構造体
func ReadConfDB() (*Config, error) {

	// 設定ファイル名
	confpath := "config/db.toml"

	// 構造体を初期化
	conf := new(Config)

	// 読み込んだjson文字列をデコードし構造体にセット
	_, err := toml.DecodeFile(confpath, &conf)
	if err != nil {
		return conf, err
	}
	conf.ConfDB.Pass = os.Getenv("PASSWORD1")
	return conf, nil
}