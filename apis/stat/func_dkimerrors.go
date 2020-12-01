package stat

import (
	"context"
	"github.com/jpbede/csapi-go/internal/transport"
	"github.com/jpbede/csapi-go/internal/validation"
)

func (c *client) GlobalDKIMErrors(ctx context.Context, date string) (*GlobalDKIMError, error) {
	var globalDkimErrors GlobalDKIMError

	if err := c.dkimErrorsScoped(ctx, date, ScopeGlobal, &globalDkimErrors); err != nil {
		return nil, err
	}

	return &globalDkimErrors, nil
}

func (c *client) DKIMErrorsByIPs(ctx context.Context, date string) (map[string]float64, error) {
	var dkimerrors map[string]float64

	if err := c.dkimErrorsScoped(ctx, date, ScopeIP, &dkimerrors); err != nil {
		return nil, err
	}

	return dkimerrors, nil
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
