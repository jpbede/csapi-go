package stat

import (
	"context"
)

// Client represents the functions implemented by this API
type Client interface {
	LatestDateWithData(ctx context.Context) (string, error)

	// Spamclickrate endpoints
	GlobalSpamclickRate(ctx context.Context, date string) (*SpamclickRate, error)
	SpamclickRateByDKIMDomains(ctx context.Context, date string) (map[string]*SpamclickRate, error)
	SpamclickRateByIPs(ctx context.Context, date string) (map[string]*SpamclickRate, error)

	// Spamtrap endpoints
	GlobalSpamtrapHits(ctx context.Context, date string) (int, error)
	SpamtrapHitsByDKIMDomains(ctx context.Context, date string) (map[string]int, error)
	SpamtrapHitsByIPs(ctx context.Context, date string) (map[string]int, error)

	// DKIM errors endpoints
	GlobalDKIMErrors(ctx context.Context, date string) (*GlobalDKIMError, error)
	DKIMErrorsByIPs(ctx context.Context, date string) (map[string]float64, error)

	// ip report endpoint
	IPReport(ctx context.Context, date string) (map[string]*IPReport, error)
}
