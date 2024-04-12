package service

import (
	"context"
	"github.com/Rustamchuk/Avito-Banner-Service/internal/repository"
	openapi "github.com/Rustamchuk/Avito-Banner-Service/pkg/generated/open_api_server/go"
)

type BannerImplementation struct {
	repo repository.BannerRepo
}

func NewBannerImplementation(repo repository.BannerRepo) *BannerImplementation {
	return &BannerImplementation{repo: repo}
}

func (b *BannerImplementation) BannerGet(
	ctx context.Context,
	token string,
	featureId int32,
	tagId int32,
	limit int32,
	offset int32,
) (openapi.ImplResponse, error) {
	return b.repo.BannerGet(ctx, token, featureId, tagId, limit, offset)
}

func (b *BannerImplementation) BannerIdDelete(
	ctx context.Context,
	id int32,
	token string,
) (openapi.ImplResponse, error) {
	return b.repo.BannerIdDelete(ctx, id, token)
}

func (b *BannerImplementation) BannerIdPatch(
	ctx context.Context,
	id int32,
	bannerIdPatchRequest openapi.BannerIdPatchRequest,
	token string,
) (openapi.ImplResponse, error) {
	return b.repo.BannerIdPatch(ctx, id, bannerIdPatchRequest, token)
}

func (b *BannerImplementation) BannerPost(
	ctx context.Context,
	bannerPostRequest openapi.BannerPostRequest,
	token string,
) (openapi.ImplResponse, error) {
	return b.repo.BannerPost(ctx, bannerPostRequest, token)
}

func (b *BannerImplementation) UserBannerGet(
	ctx context.Context,
	tagId int32,
	featureId int32,
	useLastRevision bool,
	token string,
) (openapi.ImplResponse, error) {
	return b.repo.UserBannerGet(ctx, tagId, featureId, useLastRevision, token)
}
