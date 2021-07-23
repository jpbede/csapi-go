package stat

import (
	"context"
	"go.bnck.me/csapi/internal/transport"
	"go.bnck.me/csapi/internal/validation"
)

func (c *client) GlobalDKIMErrors(ctx context.Context, date string) (*DKIMError, error) {
	var globalDkimErrors DKIMError

	if err := c.dkimErrorsScoped(ctx, date, ScopeGlobal, &globalDkimErrors); err != nil {
		return nil, err
	}

	return &globalDkimErrors, nil
}

func (c *client) DKIMErrorsByIPs(ctx context.Context, date string) (map[string]*DKIMError, error) {
	var dkimErrors map[string]*DKIMError

	if err := c.dkimErrorsScoped(ctx, date, ScopeIP, &dkimErrors); err != nil {
		return nil, err
	}

	return dkimErrors, nil
}

func (c *client) DKIMErrorsByDKIMDomains(ctx context.Context, date string) (map[string]*DKIMError, error) {
	var dkimErrors map[string]*DKIMError

	if err := c.dkimErrorsScoped(ctx, date, ScopeDKIMDomains, &dkimErrors); err != nil {
		return nil, err
	}

	return dkimErrors, nil
}

func (c *client) dkimErrorsScoped(ctx context.Context, date string, scope Scope, out interface{}) error {
	if valErr := validation.ValidateDate(date); valErr != nil {
		return valErr
	}

	var err error

	if date != "" {
		err = c.transport.Get(ctx, "/stat/dkimerrors/"+string(scope), &out, transport.WithQueryValue("date", date))
	} else {
		err = c.transport.Get(ctx, "/stat/dkimerrors/"+string(scope), &out)
	}
	if err != nil {
		return err
	}

	return nil
}
