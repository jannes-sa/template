package svclog

import (
	"template/structs"
	lStruct "template/structs/logic"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInsertServiceLog(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	var ctxStruct lStruct.ContextStruct
	ctxStruct.JobID = "jobid1"

	InsertServiceLog(ctxStruct, &errCode)

	Convey("TestInsertServiceLog", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
		})
	})
}

func TestGetAllServiceLog(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	rows := GetAllServiceLog(&errCode)

	Convey("TestGetAllServiceLog", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
			So(len(rows), ShouldNotBeEmpty)
		})
	})
}

func TestGetOneByJobIDServiceLog(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	var ctxStruct lStruct.ContextStruct
	ctxStruct.JobID = "jobid1"

	row := GetOneByJobIDServiceLog(ctxStruct, &errCode)

	Convey("TestGetOneByJobIDServiceLog", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
			So(row, ShouldNotBeNil)
		})
	})
}

func TestUpdateByJobIDServiceLog(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	var ctxStruct lStruct.ContextStruct
	ctxStruct.JobID = "jobid1"

	UpdateByJobIDServiceLog(ctxStruct, &errCode)

	Convey("TestUpdateByJobIDServiceLog", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
		})
	})
}

func TestUpdateReturnByJobIDServiceLog(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	var ctxStruct lStruct.ContextStruct
	ctxStruct.JobID = "jobid1"

	rows := UpdateReturnByJobIDServiceLog(ctxStruct, &errCode)

	Convey("TestUpdateReturnByJobIDServiceLog", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
			So(len(rows), ShouldNotBeEmpty)
		})
	})
}

func TestDeleteByJobIDServiceLog(t *testing.T) {
	errCode := make([]structs.TypeError, 0)
	var ctxStruct lStruct.ContextStruct
	ctxStruct.JobID = "jobid1"

	DeleteByJobIDServiceLog(ctxStruct, &errCode)
	Convey("TestDeleteByJobIDServiceLog", t, func() {
		Convey("Should Success", func() {
			So(len(errCode), ShouldEqual, 0)
		})
	})

}
