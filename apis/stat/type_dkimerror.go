package stat

type GlobalDKIMError struct {
	Errors      int `json:"errors"`
	TotalVolume int `json:"total_volume"`
}
