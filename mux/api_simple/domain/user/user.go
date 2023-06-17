package user

type (
	User struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Age      int    `json:"age"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UserRegisterRequestFormat struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		Age      int    `json:"age"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	UserRegisterResponseFormat struct {
		Name      string `json:"name"`
		Address   string `json:"address"`
		Age       int    `json:"age"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)
