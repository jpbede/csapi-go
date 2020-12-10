package stat

// Scope used for selecting stats
type Scope string

const (
	// ScopeGlobal information aggregated
	ScopeGlobal Scope = "global"
	// ScopeDKIMDomains information summed by dkim domain
	ScopeDKIMDomains Scope = "dkimdomain"
	// ScopeIP information summed by IP
	ScopeIP Scope = "ip"
)
