package account

const (
	AccTypeBase = iota
	AccTypeMonth
	AccTypeHalfYear
	AccTypeYear
	AccTypeSUPERVIP
)



var VDAccount *Account

func GetAccount(appId string) *Account  {
	VDAccount = getAccountInfo(appId)
	return VDAccount
}



func (a *Account)AddAction()  {
	if a.AccType ==  AccTypeBase{
		a.lock.Lock()
		defer a.lock.Unlock()
		err := a.addRequest()
		if err == nil {
			a.Count --
		}

	}
}
func (a *Account)CanUse() bool  {
	if a.AccType == AccTypeBase {
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
		return "年卡"
	} else if a.AccType == AccTypeSUPERVIP {
		return "终身高级VIP用户"
	} else {
		return "未知用户"
	}

}