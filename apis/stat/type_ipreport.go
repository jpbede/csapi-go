package stat

// IPReport information returned by /stat/ipreport
type IPReport struct {
	ComplaintAddress string      `json:"complaint_address"`
	FBLAddress       string      `json:"fbl_address"`
	IPs              IPStatusMap `json:"ips"`
}

// IPStatusMap map between status and a list of ip-hostnames
type IPStatusMap map[string]IPHostnameMap

// IPHostnameMap map of ip to hostname
type IPHostnameMap map[string]string
