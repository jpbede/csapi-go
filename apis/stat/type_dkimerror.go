package stat

// GlobalDKIMError infos about global dkim errors returned by /stat/dkimerrors/global
type GlobalDKIMError struct {
	Errors      int `json:"errors"`
	TotalVolume int `json:"total_volume"`
}
