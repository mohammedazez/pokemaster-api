package response

type (
	Response struct {
		ID        string `json:"id"`
		FullName  string `json:"fullname"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)
