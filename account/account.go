package account

const (
	AccTypeBase = iota
	AccTypeMonth
	AccTypeHalfYear
	AccTypeYear
	AccTypeSUPERVIP
)
/*
	AccTypeBase 基础套餐 ￥1 1天20个 有效期一个月
	AccTypeMonth 月卡 ￥15 1天200个
	AccTypeHalfYear 半年卡 ￥80 1天1000个
	AccTypeYear 年卡 ￥180 数量不限
	AccTypeSUPERVIP 永久 ￥500 数量不限
*/


var VDAccount *Account

func GetAccount(appId string) *Account  {
	VDAccount = getAccountInfo(appId)
	return VDAccount
}



func (a *Account)AddAction()  {
	if a.AccType < AccTypeYear{
		a.lock.Lock()
		defer a.lock.Unlock()
		err := a.addRequest()
		if err == nil {
			a.Count --
		}
	}
}
func (a *Account)CanUse() bool  {
	if a.AccType < AccTypeYear {
		if a.Count <= 0 {
			return false
		}
	}
	return true
}

func (a *Account)TYPE()string  {
	if a.AccType == AccTypeBase {
		return "基础版"
	} else if a.AccType == AccTypeMonth {
		return "月卡用户"
	} else if a.AccType == AccTypeHalfYear {
		return "半年卡用户"
	} else if a.AccType == AccTypeYear {
		return "年卡用户"
	} else if a.AccType == AccTypeSUPERVIP {
		return "终身高级VIP用户"
	} else {
		return "未知用户"
	}

}