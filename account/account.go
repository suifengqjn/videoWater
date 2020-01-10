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
	VDAccount = GetAccountInfo(appId)
	return VDAccount
}

func (a *Account)DownloadAction()  {
	if a.AccType ==  AccTypeFree{
		a.lock.Lock()
		defer a.lock.Unlock()
		a.Count --
		a.Add()
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

