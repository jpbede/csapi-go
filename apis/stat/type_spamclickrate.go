package stat

// SpamclickRate returned by /stat/spamclickrate
type SpamclickRate struct {
	Min float64 `json:"min"`
	Avg float64 `json:"avg"`
	Max float64 `json:"max"`
}
