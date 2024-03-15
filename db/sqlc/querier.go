// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	GetAllPost(ctx context.Context) ([]Post, error)
	GetCommunityIDbyName(ctx context.Context, communityName string) (uuid.UUID, error)
	GetCommunitybyID(ctx context.Context, id uuid.UUID) ([]Community, error)
	SearchCommunityName(ctx context.Context, communityName string) ([]SearchCommunityNameRow, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
	authPassword(ctx context.Context, username string) (string, error)
	authUsername(ctx context.Context, username string) (string, error)
	createAccount(ctx context.Context, arg createAccountParams) (Account, error)
	createComment(ctx context.Context, arg createCommentParams) (Comment, error)
	createCommunity(ctx context.Context, arg createCommunityParams) (Community, error)
	createPost(ctx context.Context, arg createPostParams) (Post, error)
	createReply(ctx context.Context, arg createReplyParams) (Reply, error)
	createSession(ctx context.Context, arg createSessionParams) (Session, error)
	deleteSession(ctx context.Context, username string) error
	getAccountIDbyID(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	getAccountIDbyUsername(ctx context.Context, username string) (uuid.UUID, error)
	getAccountRolebyUsername(ctx context.Context, username string) (string, error)
	getAllSessionID(ctx context.Context) ([]uuid.UUID, error)
	getFormsID(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	getSessionIDbyID(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	loginAccount(ctx context.Context, arg loginAccountParams) (Account, error)
	submitForm(ctx context.Context, arg submitFormParams) (Form, error)
}

var _ Querier = (*Queries)(nil)
