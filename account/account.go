package account

const (
	AccTypeFree = iota
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
	if a.AccType ==  AccTypeFree{
		a.lock.Lock()
		defer a.lock.Unlock()
		err := a.addRequest()
		if err == nil {
			a.Count --
		}

	}
}
func (a *Account)CanUse() bool  {
	if a.AccType == AccTypeFree {
		if a.Count <= 0 {
			return false
		}
	}
	return true
}

