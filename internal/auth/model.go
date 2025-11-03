// User, session structs

package auth

import "time"

type UserCredential struct { 
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RefreshSession struct { 
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Token string `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

