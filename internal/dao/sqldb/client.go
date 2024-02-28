package sqldb

import (
	"fmt"
	"time"

	mDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Client *gorm.DB

type ConOptions struct {
	Address         string `yaml:"address"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Database        string `yaml:"database"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
	AutoMigrate     bool   `yaml:"autoMigrateh"`
}

func Init(opt *ConOptions) {
	connect(opt)
	setOptions(opt)
	migrate(opt)
}

func connect(opt *ConOptions) {
	db, err := gorm.Open(mDriver.New(mDriver.Config{
		DSN: getDsn(cfg),
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Client = db
}

func setOptions(opt *ConOptions) {
	_db, _ := Client.DB()
	_db.SetMaxIdleConns(opt.MaxIdleConns)
	_db.SetMaxOpenConns(opt.MaxOpenConns)
	_db.SetConnMaxLifetime(time.Minute * time.Duration(opt.ConnMaxLifetime))
}

func migrate(cfg *ConOptions) {
	_ = Client.AutoMigrate(
		// some model
		&model.ApproveRecord{}, &model.PluginDetailRecord{},
	)
}

func getDsn(db *ConOptions) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", db.Username, db.Password, db.Address, db.Database)
}
