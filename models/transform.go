package models

func ToUserResponse(u *User) *UserResponse {
	if u == nil {
		return nil
	}
	return &UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Role:  string(u.Role),
	}
}
