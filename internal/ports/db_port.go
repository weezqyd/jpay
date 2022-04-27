package ports

type DbPort interface {
	CloseDbConnection() error
	GetCustomers(page uint64) ([]map[string]interface{}, error)
}
