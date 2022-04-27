package core

import (
	"github.com/weezqyd/jpay/internal/ports"
	"regexp"
	"strings"
)

type Service struct {
	dbPort ports.DbPort
}

const CameroonPhoneRegex = `\(237\)\ ?[2368]\d{7,8}$`
const EthiopiaPhoneRegex = `\(251\)\ ?[1-59]\d{8}$`
const MoroccoPhoneRegex = `\(212\)\ ?[5-9]\d{8}$`
const MozambiquePhoneRegex = `\(258\)\ ?[28]\d{7,8}$`
const UgandaPhoneRegex = `\(256\)\ ?\d{9}$`

func getCountryForPhone(phone string) string {
	mapping := map[string]string{
		"Cameroon":   CameroonPhoneRegex,
		"Morocco":    MoroccoPhoneRegex,
		"Ethiopia":   EthiopiaPhoneRegex,
		"Mozambique": MozambiquePhoneRegex,
		"Uganda":     UgandaPhoneRegex,
	}
	for country, regex := range mapping {
		pattern := regexp.MustCompile(regex)
		if pattern.MatchString(phone) {
			return country
		}
	}
	return "Unknown"
}
func NewService(dbPort ports.DbPort) *Service {
	return &Service{
		dbPort: dbPort,
	}
}

//GetAllCustomers retrieves all customers paginated
func (s Service) GetAllCustomers(page uint64) ([]*Customer, error) {
	rows, err := s.dbPort.GetCustomers(page)
	if err != nil {
		return nil, err
	}
	var customers []*Customer
	for _, c := range rows {
		customer := &Customer{
			Id:    c["id"].(int64),
			Phone: c["phone"].(string),
			Name:  c["name"].(string),
		}
		customer.Country = getCountryForPhone(customer.Phone)
		customers = append(customers, customer)
	}
	
	return customers, nil
}

// GetCustomersByCountry list all customers grouped by country
func (s Service) GetCustomersByCountry(country string) ([]*Customer, error) {
	all, err := s.GetAllCustomers(0)
	if err != nil {
		return nil, err
	}
	var customers []*Customer
	for _, c := range all {
		if strings.ToLower(c.Country) == strings.ToLower(country) {
			customers = append(customers, c)
		}
	}
	return customers, nil
}
