package service

import (
	"context"
	"errors"
	"rtc_media_engine/internal/models"
)

func (s *Service) GetIncomingRequests(user *models.User, ctx context.Context) ([]models.User, error) {
	users, err := s.r.GetIncomingRequests(user, ctx)
	if err != nil {
		return nil, errors.New("Could not get users")
	}
	return users, nil
}

func (s *Service) GetOutgoingRequests(user *models.User, ctx context.Context) ([]models.User, error) {
	users, err := s.r.GetOutgoingRequests(user, ctx)
	if err != nil {
		return nil, errors.New("Could not get users")
	}
	return users, nil
}

func (s *Service) GetFriends(user *models.User, ctx context.Context) ([]models.User, error) {
	users, err := s.r.GetFriends(user, ctx)
	if err != nil {
		return nil, errors.New("Could not get friends")
	}
	return users, nil
}

func (s *Service) AcceptFriendRequest(user *models.User, friend *models.User, ctx context.Context) error {
	tx, err := s.r.NewTransaction(ctx)
	if err != nil {
		return errors.New("Could not connect to database")
	}

	defer tx.Rollback(ctx)

	err = s.r.AcceptFriendRequest(user, friend, ctx, tx)
	if err != nil {
		return errors.New("Could not accept friend request")
	}

	err = s.r.CreateNewFriend(user, friend, ctx, tx)
	if err != nil {
		return errors.New("Could not create new friendship")
	}

	err = s.r.CreateNewFriend(friend, user, ctx, tx)
	if err != nil {
		return errors.New("Could not create new friendship")
	}

	return tx.Commit(ctx)
}

func (s *Service) SendFriendRequest(user *models.User, friend *models.User, ctx context.Context) error {
	err := s.r.GetUserByEmail(friend, ctx)
	if err != nil {
		return errors.New("Could not retrieve user by email")
	}

	if s.r.OutgoingRequestExists(user, friend, ctx) {
		return errors.New("Request already exists")
	}

	if s.r.IncomingRequestExists(user, friend, ctx) {
		return errors.New("Outgoing user has already sent a friend request")
	}

	if s.r.FriendshipExists(user, friend, ctx) {
		return errors.New("Friendship already exists")
	}

	err = s.r.CreateFriendRequest(user, friend, ctx)
	if err != nil {
		return errors.New("Could not create friend request")
	}

	return nil
}
