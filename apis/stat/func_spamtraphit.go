package stat

import (
	"context"
	"github.com/jpbede/csapi-go/internal/transport"
	"github.com/jpbede/csapi-go/internal/validation"
)

func (c *client) GlobalSpamtrapHits(ctx context.Context, date string) (int, error) {
	var spamtraphits int

	if err := c.spamtrapHitsWithScope(ctx, ScopeGlobal, date, &spamtraphits); err != nil {
		return -1, err
	}

	return spamtraphits, nil
}

func (c *client) SpamtrapHitsByDKIMDomains(ctx context.Context, date string) (map[string]int, error) {
	var spamtraphits map[string]int

	if err := c.spamtrapHitsWithScope(ctx, ScopeDKIMDomains, date, &spamtraphits); err != nil {
		return nil, err
	}

	return spamtraphits, nil
}

func (c *client) SpamtrapHitsByIPs(ctx context.Context, date string) (map[string]int, error) {
	var spamtraphits map[string]int

	if err := c.spamtrapHitsWithScope(ctx, ScopeIP, date, &spamtraphits); err != nil {
		return nil, err
	}

	return spamtraphits, nil
}

func (c *client) spamtrapHitsWithScope(ctx context.Context, scope Scope, date string, out interface{}) error {
	if valErr := validation.ValidateDate(date); valErr != nil {
		return valErr
	}

	var err error

	if date != "" {
		err = c.transport.Get(ctx, "/stat/spamtrap/"+string(scope), &out, transport.WithQueryValue("date", date))
	} else {
		err = c.transport.Get(ctx, "/stat/spamtrap/"+string(scope), &out)
	}
	if err != nil {
		return err
	}

	return nil
}
