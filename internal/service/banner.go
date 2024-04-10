package service

import (
	openapi "BannerService/internal/api/open_api_server/go"
	"BannerService/internal/repository"
	"context"
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
	//TODO implement me
	panic("implement me")
}

func (b *BannerImplementation) BannerIdDelete(
	ctx context.Context,
	id int32,
	token string,
) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BannerImplementation) BannerIdPatch(
	ctx context.Context,
	id int32,
	bannerIdPatchRequest openapi.BannerIdPatchRequest,
	token string,
) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BannerImplementation) BannerPost(
	ctx context.Context,
	bannerPostRequest openapi.BannerPostRequest,
	token string,
) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BannerImplementation) UserBannerGet(
	ctx context.Context,
	tagId int32,
	featureId int32,
	useLastRevision bool,
	token string,
) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}
