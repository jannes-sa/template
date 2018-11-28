package delsvclog

const (
	// LogicName - LogicName
	LogicName string = "svclog"
)

// Public - External Public Function
type Public interface {
	Call() int
	Receive(input int) int
}

var Function = make(map[string]Public)

// Register - Register Public Function
func Register(nmFunc string, objFunc Public) {
	Function[nmFunc] = objFunc
}
