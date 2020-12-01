package stat

type IPReport struct {
	ComplaintAddress string      `json:"complaint_address"`
	FBLAddress       string      `json:"fbl_address"`
	IPs              IPStatusMap `json:"ips"`
}

type IPStatusMap map[string]IPHostnameMap

type IPHostnameMap map[string]string
