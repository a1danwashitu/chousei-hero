package solve

type cost struct {
	Total int
	Tmp   int
	Day   int
	Level int
}

func less(x, y *cost) bool {
	if x.Day < y.Day {
		return true
	}
	if x.Tmp < y.Tmp {
		return true
	}
	if x.Total < y.Total {
		return true
	}
	return x.Level < y.Level
}

func addCost(x, y *cost) *cost {
	return &cost{
		Total: x.Total + y.Total,
		Tmp:   x.Tmp + y.Tmp,
		Day:   x.Day + y.Day,
		Level: x.Level + y.Level,
	}
}

func subCost(x, y *cost) *cost {
	return &cost{
		Total: x.Total - y.Total,
		Tmp:   x.Tmp - y.Tmp,
		Day:   x.Day - y.Day,
		Level: x.Level - y.Level,
	}
}

type costInterface struct{}

func (ci costInterface) NewCost() any {
	return &cost{}
}

func (ci costInterface) NewMaxCost() any {
	return &cost{
		Day:   1<<63 - 1,
		Tmp:   1<<63 - 1,
		Total: 1<<63 - 1,
		Level: 1<<63 - 1,
	}
}

func (ci costInterface) AddCost(x, y any) any {
	return addCost(x.(*cost), y.(*cost))
}

func (ci costInterface) SubCost(x, y any) any {
	return subCost(x.(*cost), y.(*cost))
}

func (ci costInterface) Less(x, y any) bool {
	return less(x.(*cost), y.(*cost))
}
