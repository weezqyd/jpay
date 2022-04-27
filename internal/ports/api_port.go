package ports

type APIPort interface {
	GetAllCustomers(page uint64, country string)
}
