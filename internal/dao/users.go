package dao

import (
	"github.com/Forget-C/http-structer/internal/dao/sqldb"
	"github.com/Forget-C/http-structer/internal/model"
)

var Users = new(users)

type users struct {
	sqldb.Common[*model.UserRecord]
}
