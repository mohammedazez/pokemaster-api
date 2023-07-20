package request

type (
	RequestInsert struct {
		FullName string `json:"fullname" validate:"required" required:"fullname is required"`
		Email    string `json:"email"  validate:"required" required:"email is required"`
		Password string `json:"password"  validate:"required" required:"password is required"`
	}
)
