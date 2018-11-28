package logic2

import (
	"template/models/logic/svclog/delsvclog"
)

type logic2 struct{}

func (l logic2) Call() int {
	if _, ok := delsvclog.Function[delsvclog.LogicName]; !ok {
		panic("function not exists")
	}
	svclog := delsvclog.Function[delsvclog.LogicName]

	calc := svclog.Receive(2)
	return calc
}

func (l logic2) Receive(input int) int {
	return input * 2
}
