package restaurant

import (
	"database/sql"
)

type MenuItem struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type Order struct {
	Id                   int64   `json:"id"`
	Table                int64   `json:"table"`
	Date                 string  `json:"date"`
	TotalPrice           float64 `json:"totalPrice"`
	TotalPriceWithoutTax float64 `json:"totalPriceWithoutTax"`
	RecommendedTip       float64 `json:"recommendedTip"`
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
