package wbdata

import (
	"strconv"
)

type CountriesService service

type Country struct {
	ID           string
	Name         string
	CapitalCity  string
	Iso2Code     string
	Longitude    string
	Latitude     string
	Region       Region
	IncomeLevels IncomeLevel
	LendingType  LendingType
	AdminRegion  struct {
		ID       string
		Iso2Code string
		Value    string
	}
}

func (c *CountriesService) ListCountries(pages PageParams) (PageSummary, []Country, error) {
	summary := PageSummary{}
	countries := []Country{}

	req, err := c.client.NewRequest("GET", "countries", nil)
	if err != nil {
		return PageSummary{}, nil, err
	}
	// log.Printf(`req: %+v`, req)

	params := req.URL.Query()
	if pages.Page != 0 {
		params.Add(`page`, strconv.Itoa(int(pages.Page)))
	}
	if pages.PerPage != 0 {
		params.Add(`per_page`, strconv.Itoa(int(pages.PerPage)))
	}
	req.URL.RawQuery = params.Encode()

	_, err = c.client.do(req, &[]interface{}{&summary, &countries})
	if err != nil {
		return PageSummary{}, nil, err
	}

	return summary, countries, nil
}