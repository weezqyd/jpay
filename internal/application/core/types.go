package core

type Number interface {
	int | int64
}
type Customer struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
}
