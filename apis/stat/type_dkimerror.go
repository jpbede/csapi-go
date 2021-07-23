package stat

// DKIMError infos about global dkim errors returned by /stat/dkimerrors
type DKIMError struct {
	Errors      int `json:"errors"`
	TotalVolume int `json:"total_volume"`
}
