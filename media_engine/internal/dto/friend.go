package dto

type SendFriendRequest struct {
	Email string `json:"email"`
}

type AcceptFriendRequest struct {
	FriendId string `json:"friend_id"`
}
