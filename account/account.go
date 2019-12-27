package account

const (
	AccTypeFree = iota
	AccTypeMonth
	AccTypeYear
	AccTypeVIP
)

type Account struct {
	Type int
	Key string
	RemainCount int
	StartTime string
	EndTime string
}
