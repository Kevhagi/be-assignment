package accountmanagerdto

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id                string `json:"id"`
	SupertokensUserId string `json:"supertokens_user_id"`
	Email             string `json:"email"`
}
