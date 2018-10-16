package mongo

import (
	"template/helper"
	"template/helper/constant"

	"github.com/globalsign/mgo"
)

// Connect to connect Mongo DB
// this is use for open connection to mongodb
func Connect() (sess *mgo.Session, err error) {
	sess, err = mgo.Dial(constant.CREDMONGO)
	if err != nil {
		helper.CheckErr("error connect mongo DB", err)
		return
	}

	mgo.SetDebug(true)

	sess.SetMode(mgo.Monotonic, true)

	return
}
