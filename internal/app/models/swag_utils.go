package models

// TokenResponse represents the response with access token and user ID
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	UserID       uint   `json:"user_id"`
}

// RefreshTokenResponse represents the response with access token and user ID
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// ErrorResponse represents an error message response
type ErrorResponse struct {
	Error string `json:"error"`
}

// DefaultResponse represents a default message response
type DefaultResponse struct {
	Message string `json:"message"`
}

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CategoryRequest struct {
	CategoryName string `json:"category_name" binding:"required"` // Название категории, обязательное поле
	ParentID     uint   `json:"parent_id,omitempty"`              // Идентификатор родительской категории, необязательное поле
	Description  string `json:"description,omitempty"`            // Описание категории, необязательное поле
}

type AddressRequest struct {
	AddressName string `json:"address_name"`
}

type AccountRequest struct {
	AccountName string `json:"account_number"`
}

type FillAccountRequest struct {
	AccountName string `json:"account_number"`
	Balance     uint   `json:"balance"`
}

type AccountsResponse struct {
	AccountName string `json:"account_name"`
}

type FeaturedProductsRequest struct {
	ProductID uint `json:"product_id"`
}

type ReviewRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Rating  uint   `json:"rating"`
}

type OrderRequest struct {
	StatusID  uint `json:"status_id"`
	AddressID uint `json:"address_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type OrderStatusRequest struct {
	StatusName  string `json:"status_name"`
	Description string `json:"description"`
}

type PaymentRequest struct {
	AccountID uint `json:"account_id"`
	OrderID   uint `json:"order_id"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type ProductRequest struct {
	StoreID       uint     `json:"store_id"`
	CategoryID    uint     `json:"category_id"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Price         uint     `json:"price"`
	Amount        uint     `json:"amount"`
	ProductImages []string `json:"product_image"`
}

type ProductResponse struct {
	StoreID       uint     `json:"store_id"`
	CategoryID    uint     `json:"category_id"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Price         uint     `json:"price"`
	Amount        uint     `json:"amount"`
	ProductImages []string `json:"product_images"`
}

type StoreRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type StoreReviewRequest struct {
	Rating  uint   `json:"rating"`
	Comment string `json:"comment"`
}

type CommentRequest struct {
	ParentID    uint   `json:"parent_id"`
	CommentText string `json:"text" binding:"required"`
}

type NewUsersPassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
