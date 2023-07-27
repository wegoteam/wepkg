package snowflake

// OverCostAction
// @Description: 用于记录超过预设的生成次数的情况
type OverCostAction struct {
	ActionType             int32
	TimeTick               int64
	WorkerId               uint16
	OverCostCountInOneTerm int32
	GenCountInOneTerm      int32
	TermIndex              int32
}

// GetOverCostAction
// @Description: 获取超过预设的生成次数的情况
// @receiver: overcost
// @param: workerId
// @param: timeTick
// @param: actionType
// @param: overCostCountInOneTerm
// @param: genCountWhenOverCost
// @param: index
func (overcost OverCostAction) GetOverCostAction(workerId uint16, timeTick int64, actionType, overCostCountInOneTerm, genCountWhenOverCost, index int32) {
	overcost.ActionType = actionType
	overcost.TimeTick = timeTick
	overcost.WorkerId = workerId
	overcost.OverCostCountInOneTerm = overCostCountInOneTerm
	overcost.GenCountInOneTerm = genCountWhenOverCost
	overcost.TermIndex = index
}
