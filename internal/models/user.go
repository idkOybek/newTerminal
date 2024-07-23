package models

// User представляет пользователя
type User struct {
	ID       int    `json:"id"`
	INN      string `json:"inn"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}

// UserRegistrationRequest представляет данные для регистрации пользователя
type UserRegistrationRequest struct {
	INN      string `json:"inn"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}

// UserLoginRequest представляет данные для входа пользователя
type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserLoginResponse представляет ответ на успешный вход пользователя
type UserLoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

// UserUpdateRequest представляет данные для обновления пользователя
type UserUpdateRequest struct {
	INN      string `json:"inn"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}

// UserResponse представляет данные пользователя для ответа
type UserResponse struct {
	ID       int    `json:"id"`
	INN      string `json:"inn"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
	IsAdmin  bool   `json:"is_admin"`
}
