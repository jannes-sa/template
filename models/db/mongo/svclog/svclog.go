package svclog

import (
	"template/helper"
	"template/helper/constant"
	"template/helper/constant/tablename"
	db "template/models/db/mongo"
	dbStruct "template/structs/db"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// SvcLog - Logic Struct DB
type SvcLog struct{}

func init() {
	var d SvcLog
	d.Index()
}

// GetColl - Get Collection service_log
func (d *SvcLog) GetColl() (sess *mgo.Session, coll *mgo.Collection, err error) {
	sess, err = db.Connect()
	if err != nil {
		helper.CheckErr("Failed get collection service_log", err)
		return
	}

	coll = sess.DB(constant.GOAPP).C(tablename.ServiceLog)

	return
}

// Index - Create Index
func (d *SvcLog) Index() (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	index := mgo.Index{
		Key:        []string{jobIDStr},
		Unique:     true,  // Prevent two documents from having the same index key
		DropDups:   false, // Drop documents with the same index key as a previously indexed one
		Background: false, // Build index in background and return immediately
		Sparse:     false, // Only index documents containing the Key fields
	}

	err = coll.EnsureIndex(index)

	return
}

// GetAllServiceLog - GetAllServiceLog GetAll
func (d *SvcLog) GetAllServiceLog() (rows []dbStruct.ServiceLog, err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Find(bson.M{}).All(&rows)

	return
}

// GetOneServiceLog - GetOneServiceLog
func (d *SvcLog) GetOneServiceLog() (row dbStruct.ServiceLog, err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Find(bson.M{jobIDStr: row.JobID}).One(&row)

	return
}

// UpdateServiceLog - UpdateServiceLog
func (d *SvcLog) UpdateServiceLog() (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	selector := bson.M{jobIDStr: "xxxxxx"}
	update := bson.M{
		"$set": bson.M{
			"res": "yyyy",
		},
	}

	err = coll.Update(selector, update)

	return
}

// InsertServiceLog - InsertServiceLog
func (d *SvcLog) InsertServiceLog(v interface{}) (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}

	err = coll.Insert(v)

	return
}

// RemoveServiceLog - RemoveServiceLog
func (d *SvcLog) RemoveServiceLog() (err error) {
	sess, coll, err := d.GetColl()
	defer sess.Close()
	if err != nil {
		return
	}
	selector := bson.M{jobIDStr: "xxxxxx"}
	err = coll.Remove(selector)

	return
}
