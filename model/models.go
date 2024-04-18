package model

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID             int     `json:"id"`
	Fullname       string  `json:"fullname"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	Deposit_amount float64 `json:"deposit_amount"`
	Role           string  `json:"role"`
}

type Book struct {
	ID          int     `json:"id"`
	Tittle      string  `json:"tittle"`
	Category    string  `json:"category"`
	Stock       int     `json:"stock"`
	Author_name string  `json:"author_name"`
	Rent_cost   float64 `json:"rent_cost"`
}

type Cart struct {
	ID          int     `json:"id"`
	User_id     int     `json:"user_id"`
	Book_id     int     `json:"book_id"`
	Quantity    int     `json:"quantity"`
	Total_price float64 `json:"total_price"`
}

type Transaction struct {
	ID          int     `json:"id"`
	User_id     int     `json:"user_id"`
	Rental_id   int     `json:"rental_id"`
	Rent_date   string  `json:"rent_date"`
	Due_date    string  `json:"due_date"`
	Total_price float64 `json:"total_price"`
}

type Payment struct {
	ID             int     `json:"id"`
	Transaction_id int     `json:"transaction_id"`
	Total_amount   float64 `json:"total_amount"`
	Payment_date   string  `json:"payment_date"`
	Status         string  `json:"status"`
}

type ErrorMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type SuccessMessageUser struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Datas   []User `json:"usersdata,omitempty"`
	Data    *User  `json:"userdata,omitempty"`
	Token   string
}

type SuccessMessageBook struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Datas   []Book `json:"booksdata,omitempty"`
	Data    *Book  `json:"bookdata,omitempty"`
}

type SuccessMessageCart struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Datas   []Cart `json:"booksdata,omitempty"`
	Data    *Cart  `json:"bookdata,omitempty"`
}
