package mincostflow

type CostInterface interface {
	NewCost() any
	NewMaxCost() any
	AddCost(any, any) any
	SubCost(any, any) any
	Less(any, any) bool
}

var ci CostInterface

func Init(costInterface CostInterface) {
	ci = costInterface
}
