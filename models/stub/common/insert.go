package common

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

//"errors"

// DBInsert ...
type DBInsert struct{}

// InsertMulti : insert multi
func (i *DBInsert) InsertMulti(o orm.Ormer, length int, v interface{}) (countRecord int64, err error) {
	if length == 0 {
		err = errors.New("Failed insert multi because there is no data")
		return 0, err
	}
	return 2, nil
}
