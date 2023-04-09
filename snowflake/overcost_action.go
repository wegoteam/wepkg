package snowflake

type OverCostAction struct {
	ActionType             int32
	TimeTick               int64
	WorkerId               uint16
	OverCostCountInOneTerm int32
	GenCountInOneTerm      int32
	TermIndex              int32
}

func (overcost OverCostAction) GetOverCostAction(workerId uint16, timeTick int64, actionType int32, overCostCountInOneTerm int32, genCountWhenOverCost int32, index int32) {
	overcost.ActionType = actionType
	overcost.TimeTick = timeTick
	overcost.WorkerId = workerId
	overcost.OverCostCountInOneTerm = overCostCountInOneTerm
	overcost.GenCountInOneTerm = genCountWhenOverCost
	overcost.TermIndex = index
}
