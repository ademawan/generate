package domain

type AuthMerchantPanelPayload struct {
	UUID        string
	Email       string
	AccessToken string
	DeviceID    string
	MerchantID  string
}

type AuthMerchantPanelResponse struct {
	UUID        string   `json:"uuid"`
	RoleID      string   `json:"role_id"`
	Permissions []string `json:"permissions"`
	DeviceID    string   `json:"device_id"`
	MerchantID  int      `json:"merchant_id"`
	Email       string   `json:"email"`
}
type UserInfo struct {
	MerchantID  int      `json:"merchant_id"`
	RoleID      string   `json:"role_id"`
	Permissions []string `json:"permissions"`
	UUID        string   `json:"uuid"`
	IsMerchant  bool     `json:"is_merchant"`
	CreatedAt   string   `json:"created_at"`
	AccessToken string   `json:"access_token"`
	DeviceID    string   `json:"device_id"`
}
type AuthRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}
type AuthMerchantResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// /auth basic
type LoginBasic struct {
	Status       bool
	Message      string
	Token        string
	TokenId      string
	RefreshToken string
	StatusCode   int
}

type UserData struct {
	UsersId     string
	Uuid        string
	FullName    string
	FirstName   string
	LastName    string
	Gender      string
	Phone       string
	DOB         string
	Email       string
	LinkPicture string
	CreatedAt   string
	UpdatedAt   string
	IsMerchant  bool
}
