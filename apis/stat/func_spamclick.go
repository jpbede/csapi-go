package stat

import (
	"context"
	"github.com/jpbede/csapi-go/internal/transport"
	"github.com/jpbede/csapi-go/internal/validation"
)

func (c *client) GlobalSpamclickRate(ctx context.Context, date string) (*SpamclickRate, error) {
	var spamclick SpamclickRate

	if err := c.spamClickrateWithScope(ctx, ScopeGlobal, date, &spamclick); err != nil {
		return nil, err
	}

	return &spamclick, nil
}

func (c *client) SpamclickRateByDKIMDomains(ctx context.Context, date string) (map[string]*SpamclickRate, error) {
	var spamclick map[string]*SpamclickRate

	if err := c.spamClickrateWithScope(ctx, ScopeDKIMDomains, date, &spamclick); err != nil {
		return nil, err
	}

	return spamclick, nil
}

func (c *client) SpamclickRateByIPs(ctx context.Context, date string) (map[string]*SpamclickRate, error) {
	var spamclick map[string]*SpamclickRate

	if err := c.spamClickrateWithScope(ctx, ScopeIP, date, &spamclick); err != nil {
		return nil, err
	}

	return spamclick, nil
}

func (c *client) spamClickrateWithScope(ctx context.Context, scope Scope, date string, out interface{}) error {
	if valErr := validation.ValidateDate(date); valErr != nil {
		return valErr
	}

	var err error

	if date != "" {
		err = c.transport.Get(ctx, "/stat/spamclickrate/"+string(scope), &out, transport.WithQueryValue("date", date))
	} else {
		err = c.transport.Get(ctx, "/stat/spamclickrate/"+string(scope), &out)
	}
	if err != nil {
		return err
	}

	return nil
}
