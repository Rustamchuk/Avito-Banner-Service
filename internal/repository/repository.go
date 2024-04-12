package repository

import (
	"context"
	"github.com/Rustamchuk/Avito-Banner-Service/internal/repository/postgres"
	openapi "github.com/Rustamchuk/Avito-Banner-Service/pkg/generated/open_api_server/go"
	"github.com/jmoiron/sqlx"
)

type BannerRepo interface {
	BannerGet(
		ctx context.Context,
		token string,
		featureId int32,
		tagId int32,
		limit int32,
		offset int32,
	) (openapi.ImplResponse, error)

	BannerIdDelete(
		ctx context.Context,
		id int32,
		token string,
	) (openapi.ImplResponse, error)

	BannerIdPatch(
		ctx context.Context,
		id int32,
		bannerIdPatchRequest openapi.BannerIdPatchRequest,
		token string,
	) (openapi.ImplResponse, error)

	BannerPost(
		ctx context.Context,
		bannerPostRequest openapi.BannerPostRequest,
		token string,
	) (openapi.ImplResponse, error)

	UserBannerGet(
		ctx context.Context,
		tagId int32,
		featureId int32,
		useLastRevision bool,
		token string,
	) (openapi.ImplResponse, error)
}

type Repository struct {
	BannerRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{BannerRepo: postgres.NewBannerPostgres(db)}
}
