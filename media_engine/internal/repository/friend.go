package repository

import (
	"context"
	"errors"
	"rtc_media_engine/internal/models"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateFriendRequest(sender *models.User, receiver *models.User, ctx context.Context) error {
	sqlQuery := "INSERT INTO friend_requests (sender_id, receiver_id) VALUES ($1, $2)"

	_, err := r.Pool.Exec(ctx, sqlQuery, sender.UserID, receiver.UserID)
	if err != nil {
		return errors.New("Could not create new request")
	}

	return nil
}

func (r *Repository) AcceptFriendRequest(receiver *models.User, sender *models.User, ctx context.Context, tx pgx.Tx) error {
	sqlQuery := `UPDATE friend_requests
				SET status = 'accepted'
				WHERE sender_id = $1 AND receiver_id = $2`
	_, err := tx.Exec(ctx, sqlQuery, sender.UserID, receiver.UserID)
	if err != nil {
		return errors.New("Could not accept request")
	}

	return nil
}

func (r *Repository) DenyFriendRequest(receiver *models.User, sender *models.User, ctx context.Context) error {
	sqlQuery := `UPDATE friend_requests
				SET status = 'denied'
				WHERE sender_id = $1 AND receiver_id = $2`
	_, err := r.Pool.Exec(ctx, sqlQuery, sender.UserID, receiver.UserID)
	if err != nil {
		return errors.New("Could not accept request")
	}

	return nil
}

func (r *Repository) CreateNewFriend(user *models.User, friend *models.User, ctx context.Context, tx pgx.Tx) error {
	sqlQuery := `INSERT INTO friends (user_id, friend_id) VALUES ($1, $2)`
	_, err := tx.Exec(ctx, sqlQuery, user.UserID, friend.UserID)
	if err != nil {
		return errors.New("Could not create new friend")
	}

	return nil
}

func (r *Repository) GetOutgoingRequests(sender *models.User, ctx context.Context) ([]models.User, error) {
	sqlQuery := `SELECT users.id, users.first_name, users.last_name, users.email, friend_requests.created_at
				FROM users
				JOIN friend_requests ON users.id = friend_requests.receiver_id
				WHERE friend_requests.sender_id = $1 AND friend_requests.status = 'outgoing'`
	rows, err := r.Pool.Query(ctx, sqlQuery, sender.UserID)
	if err != nil {
		return nil, errors.New("Could not query for users")
	}

	defer rows.Close()

	var receivers []models.User
	for rows.Next() {
		var receiver models.User
		err = rows.Scan(&receiver.UserID, &receiver.FirstName, &receiver.LastName, &receiver.Email, &receiver.CreatedAt)
		if err != nil {
			return nil, errors.New("Could not read users")
		}

		receivers = append(receivers, receiver)
	}

	return receivers, nil
}

func (r *Repository) GetIncomingRequests(receiver *models.User, ctx context.Context) ([]models.User, error) {
	sqlQuery := `SELECT users.id, users.first_name, users.last_name, users.email, friend_requests.created_at
				FROM users
				JOIN friend_requests ON users.id = friend_requests.sender_id
				WHERE friend_requests.receiver_id = $1 AND friend_requests.status = 'outgoing'`
	rows, err := r.Pool.Query(ctx, sqlQuery, receiver.UserID)
	if err != nil {
		return nil, errors.New("Could not query for users")
	}

	defer rows.Close()

	var senders []models.User
	for rows.Next() {
		var sender models.User
		err = rows.Scan(&sender.UserID, &sender.FirstName, &sender.LastName, &sender.Email, &sender.CreatedAt)
		if err != nil {
			return nil, errors.New("Could not read users")
		}

		senders = append(senders, sender)
	}

	return senders, nil
}

func (r *Repository) OutgoingRequestExists(user *models.User, friend *models.User, ctx context.Context) bool {
	sqlQuery := `SELECT EXISTS( SELECT 1 FROM friend_requests WHERE sender_id = $1 AND receiver_id = $2);`
	row := r.Pool.QueryRow(ctx, sqlQuery, user.UserID, friend.UserID)

	var exists int
	err := row.Scan(&exists)
	if err != nil {
		return false
	}
	return exists == 1
}

func (r *Repository) IncomingRequestExists(user *models.User, friend *models.User, ctx context.Context) bool {
	sqlQuery := `SELECT EXISTS( SELECT 1 FROM friend_requests WHERE receiver_id = $1 AND sender_id = $2);`
	row := r.Pool.QueryRow(ctx, sqlQuery, user.UserID, friend.UserID)
	var exists int
	err := row.Scan(&exists)
	if err != nil {
		return false
	}
	return exists == 1
}

func (r *Repository) FriendshipExists(user *models.User, friend *models.User, ctx context.Context) bool {
	sqlQuery := `SELECT EXISTS( SELECT 1 FROM friends WHERE user_id = $1 AND friend_id = $2);`
	row := r.Pool.QueryRow(ctx, sqlQuery, user.UserID, friend.UserID)
	var exists int
	err := row.Scan(&exists)
	if err != nil {
		return false
	}
	return exists == 1
}

func (r *Repository) GetFriends(user *models.User, ctx context.Context) ([]models.User, error) {
	sqlQuery := `SELECT users.id, users.first_name, users.last_name, users.email, friends.created_at
				FROM users
				JOIN friends ON users.id = friends.friend_id
				WHERE friends.user_id = $1`
	rows, err := r.Pool.Query(ctx, sqlQuery, user.UserID)
	if err != nil {
		return nil, errors.New("Could not query for users")
	}

	defer rows.Close()

	var friends []models.User
	for rows.Next() {
		var friend models.User
		err = rows.Scan(&friend.UserID, &friend.FirstName, &friend.LastName, &friend.Email, &friend.CreatedAt)
		if err != nil {
			return nil, errors.New("Could not read users")
		}

		friends = append(friends, friend)
	}

	return friends, nil
}
