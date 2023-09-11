package gola

type OperatorName string

type Operator struct {
	Name          OperatorName
	Documentation string
}
type Operators = []*Operator
