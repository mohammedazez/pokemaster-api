package response

import (
	domain "pokemaster-api/core/domain/user"
	"pokemaster-api/interface/api/common"
)

type (
	ResponseList struct {
		Response
	}
)

func NewResponseList(message string, data []domain.User, code int) *common.DefaultResponse {
	users := make([]ResponseList, 0)

	for _, val := range data {
		var user ResponseList

		user.ID = val.ID
		user.FullName = val.FullName
		user.Email = val.Email
		user.Password = val.Password
		user.CreatedAt = val.CreatedAt
		user.UpdatedAt = val.UpdatedAt
		users = append(users, user)
	}

	responseData := new(common.DefaultResponse)
	responseData.SetResponseData(message, users, code, true)
	return responseData
}
