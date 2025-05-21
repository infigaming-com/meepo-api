package util

import (
	"context"
	"errors"

	system "github.com/infigaming-com/meepo-api/system/service/v1"
	"github.com/shopspring/decimal"
)

type currency struct {
	Currency      string
	Enabled       bool
	Hidden        bool
	Type          string
	DecimalPlaces int32
}

type SystemCurrency struct {
	currencies  map[string]currency
	curCurrency *currency
}

func GetSystemCurrency(ctx context.Context, sys system.SystemClient, currencies []string) (*SystemCurrency, error) {

	resp, err := sys.GetCurrencies(ctx, &system.GetCurrenciesRequest{Currencies: currencies})
	if err != nil {
		return nil, err
	}

	if len(resp.Currencies) == 0 {
		return nil, errors.New("empty currencies")
	}

	var currencymap = make(map[string]currency)
	for _, curr := range resp.Currencies {
		currencymap[curr.Currency] = currency{
			Currency:      curr.Currency,
			Enabled:       curr.Enabled,
			Hidden:        curr.Hidden,
			Type:          curr.Type,
			DecimalPlaces: curr.DecimalPlaces,
		}
	}

	return &SystemCurrency{
		currencies: currencymap,
		curCurrency: &currency{
			Currency:      resp.Currencies[0].Currency,
			Enabled:       resp.Currencies[0].Enabled,
			Hidden:        resp.Currencies[0].Hidden,
			Type:          resp.Currencies[0].Type,
			DecimalPlaces: resp.Currencies[0].DecimalPlaces,
		},
	}, nil
}

func (c *SystemCurrency) SetDefaultCurrency(currency string) error {
	if cur, ok := c.currencies[currency]; ok {
		c.curCurrency = &cur
	}
	return errors.New("Currency not found")
}

func (c *SystemCurrency) ToExternalStringAbs(amount int64) string {
	return decimal.NewFromInt(amount).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt32(c.curCurrency.DecimalPlaces))).Abs().String()
}

func (c *SystemCurrency) ToExternalString(amount int64) string {
	return decimal.NewFromInt(amount).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt32(c.curCurrency.DecimalPlaces))).String()
}

func (c *SystemCurrency) ToExternalStringAbsByCurrency(amount int64, currency string) (string, error) {
	if cur, ok := c.currencies[currency]; ok {
		return decimal.NewFromInt(amount).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt32(cur.DecimalPlaces))).Abs().String(), nil
	}
	return "", errors.New("Currency not found")
}

func (c *SystemCurrency) ToExternalStringByCurrency(amount int64, currency string) (string, error) {
	if cur, ok := c.currencies[currency]; ok {
		return decimal.NewFromInt(amount).Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt32(cur.DecimalPlaces))).String(), nil
	}
	return "", errors.New("Currency not found")
}
