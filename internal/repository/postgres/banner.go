package postgres

import (
	openapi "BannerService/internal/api/open_api_server/go"
	"BannerService/internal/model"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strings"
)

type BannerPostgres struct {
	db *sqlx.DB
}

func NewBannerPostgres(db *sqlx.DB) *BannerPostgres {
	return &BannerPostgres{db: db}
}

func (b *BannerPostgres) BannerGet(
	ctx context.Context,
	token string,
	featureId int32,
	tagId int32,
	limit int32,
	offset int32,
) (openapi.ImplResponse, error) {
	var banners []openapi.BannerGet200ResponseInner
	query := `SELECT id, content, is_active, feature_id, tag_ids, created_at, updated_at FROM banners WHERE feature_id = $1 AND $2 = ANY(tag_ids) LIMIT $3 OFFSET $4`
	err := b.db.SelectContext(ctx, &banners, query, featureId, tagId, limit, offset)
	if err != nil {
		return openapi.Response(500, nil), err
	}
	return openapi.Response(200, banners), nil
}

func (b *BannerPostgres) BannerIdDelete(
	ctx context.Context,
	id int32,
	token string,
) (openapi.ImplResponse, error) {
	query := `DELETE FROM banners WHERE id = $1`
	result, err := b.db.ExecContext(ctx, query, id)
	if err != nil {
		return openapi.Response(500, nil), err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return openapi.Response(500, nil), err
	}
	if rowsAffected == 0 {
		return openapi.Response(404, nil), nil
	}
	return openapi.Response(204, nil), nil
}

func (b *BannerPostgres) BannerIdPatch(
	ctx context.Context,
	id int32,
	bannerIdPatchRequest openapi.BannerIdPatchRequest,
	token string,
) (openapi.ImplResponse, error) {
	query := `UPDATE banners SET `
	params := []interface{}{}
	paramId := 1

	if bannerIdPatchRequest.TagIds != nil {
		query += fmt.Sprintf("tag_ids = $%d, ", paramId)
		params = append(params, pq.Array(*bannerIdPatchRequest.TagIds))
		paramId++
	}
	if bannerIdPatchRequest.FeatureId != nil {
		query += fmt.Sprintf("feature_id = $%d, ", paramId)
		params = append(params, *bannerIdPatchRequest.FeatureId)
		paramId++
	}
	if bannerIdPatchRequest.Content != nil {
		query += fmt.Sprintf("content = $%d, ", paramId)
		params = append(params, *bannerIdPatchRequest.Content)
		paramId++
	}
	if bannerIdPatchRequest.IsActive != nil {
		query += fmt.Sprintf("is_active = $%d, ", paramId)
		params = append(params, *bannerIdPatchRequest.IsActive)
		paramId++
	}

	// Удаление последней запятой и добавление условия WHERE
	query = strings.TrimSuffix(query, ", ") + fmt.Sprintf(" WHERE id = $%d", paramId)
	params = append(params, id)

	result, err := b.db.ExecContext(ctx, query, params...)
	if err != nil {
		return openapi.Response(500, nil), err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return openapi.Response(500, nil), err
	}
	if rowsAffected == 0 {
		return openapi.Response(404, nil), nil
	}
	return openapi.Response(200, nil), nil
}

func (b *BannerPostgres) BannerPost(
	ctx context.Context,
	bannerPostRequest openapi.BannerPostRequest,
	token string,
) (openapi.ImplResponse, error) {
	query := `INSERT INTO banners (tag_ids, feature_id, content, is_active) VALUES ($1, $2, $3, $4) RETURNING id`
	var newBannerId int32
	err := b.db.QueryRowContext(ctx, query, pq.Array(bannerPostRequest.TagIds), bannerPostRequest.FeatureId, bannerPostRequest.Content, bannerPostRequest.IsActive).Scan(&newBannerId)
	if err != nil {
		return openapi.Response(500, nil), err
	}
	return openapi.Response(201, map[string]int32{"id": newBannerId}), nil
}

func (b *BannerPostgres) UserBannerGet(
	ctx context.Context,
	tagId int32,
	featureId int32,
	useLastRevision bool,
	token string,
) (openapi.ImplResponse, error) {
	var banner model.Banner
	query := `SELECT id, content, is_active, feature_id, tag_ids, created_at, updated_at FROM banners WHERE feature_id = $1 AND $2 = ANY(tag_ids) ORDER BY RANDOM() LIMIT 1`
	err := b.db.GetContext(ctx, &banner, query, featureId, tagId)
	if err != nil {
		return openapi.Response(500, nil), err
	}
	return openapi.Response(200, banner), nil
}
