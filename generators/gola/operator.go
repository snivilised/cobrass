package gola

type OperatorName string

type Operator struct {
	Name          OperatorName
	Documentation string
}

type operatorCollection []*Operator

func buildOperators() operatorCollection {
	return operatorCollection{
		&Operator{
			Name:          "Within",
			Documentation: "fails validation if the option value does not lie within 'low' and 'high' (inclusive)",
		},
	}
}
