package stat

import "context"

func (c *client) LatestDateWithData(ctx context.Context) (string, error) {
	var date string
	err := c.transport.Get(ctx, "/stat/", &date)
	if err != nil {
		return "", err
	}

	return date, nil
}
