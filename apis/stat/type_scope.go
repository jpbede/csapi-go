package stat

type Scope string

const (
	ScopeGlobal      Scope = "global"
	ScopeDKIMDomains Scope = "dkimdomain"
	ScopeIP          Scope = "ip"
)
