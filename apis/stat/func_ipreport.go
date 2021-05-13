package stat

import (
	"context"
	"go.bnck.me/csapi/internal/transport"
	"go.bnck.me/csapi/internal/validation"
)

func (c *client) IPReport(ctx context.Context, date string) (map[string]*IPReport, error) {
	if valErr := validation.ValidateDate(date); valErr != nil {
		return nil, valErr
	}

	var ipreports map[string]*IPReport
	var err error

	if date != "" {
		err = c.transport.Get(ctx, "/stat/ipreport", &ipreports, transport.WithQueryValue("date", date))
	} else {
		err = c.transport.Get(ctx, "/stat/ipreport", &ipreports)
	}
	if err != nil {
		return nil, err
	}

	return ipreports, nil
}
