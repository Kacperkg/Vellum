package dto

type CreateAccountRequest struct {
	Name string
}

type UpdateAccountRequest struct {
	Name string
}

type AccountResponse struct {
	ID   string
	Name string
}
