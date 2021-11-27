package restaurant

import (
	"database/sql"
	"time"
)

type MenuItem struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type Order struct {
	Id                   int     `json:"id"`
	Table                int     `json:"table"`
	Date                 string  `json:"date"`
	TotalPrice           float64 `json:"totalPrice"`
	TotalPriceWithoutTax float64 `json:"totalPriceWithoutTax"`
	RecommendedTip       float64 `json:"recommendedTip"`
}

type OrderItem struct {
	ItemId   int64 `json:"itemId"`
	Quantity int64 `json:"quantity"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) GetMenu() ([]*MenuItem, error) {
	rows, err := s.Db.Query("SELECT * FROM menu")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menu []*MenuItem
	for rows.Next() {
		item := MenuItem{}
		err := rows.Scan(&item.Id, &item.Name, &item.Price)
		if err != nil {
			return nil, err
		}
		menu = append(menu, &item)
	}

	if menu == nil {
		menu = make([]*MenuItem, 0)
	}
	return menu, nil
}

func (s *Store) AddNewOrder(table int, items []*OrderItem) (*Order, error) {
	orderRow := s.Db.QueryRow("INSERT INTO orders (\"table\", date) VALUES ($1, $2) RETURNING id",
		table, time.Now().Format("2006-01-02 15:04:05"))

	var orderId int
	err := orderRow.Scan(&orderId)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		s.Db.QueryRow("INSERT INTO order_details (order_id, meal_id, quantity) VALUES ($1, $2, $3)", orderId, item.ItemId, item.Quantity)
	}

	var totalPrice, totalPriceWithoutTax, recommendedTip float64
	s.Db.QueryRow("SELECT * FROM get_total_price($1)", orderId).Scan(&orderId, &totalPrice, &totalPriceWithoutTax, &recommendedTip)

	order := &Order{
		Id:                   orderId,
		Table:                table,
		TotalPrice:           totalPrice,
		TotalPriceWithoutTax: totalPriceWithoutTax,
		RecommendedTip:       recommendedTip,
	}

	return order, nil
}
