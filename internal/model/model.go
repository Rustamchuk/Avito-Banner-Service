package model

// Banner описывает сущность баннера
type Banner struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	Content   string `db:"content"` // JSON-документ баннера
	IsActive  bool   `db:"is_active"`
	FeatureID int64  `db:"feature_id"`
}

// Tag описывает сущность тега
type Tag struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

// Feature описывает сущность фичи
type Feature struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

// BannerTag описывает связь многие-ко-многим между баннерами и тегами
type BannerTag struct {
	BannerID int64 `db:"banner_id"`
	TagID    int64 `db:"tag_id"`
}

// BannerFeature описывает связь один-ко-многим между баннерами и фичами
// Эта структура может быть не нужна, если вы решите хранить FeatureID напрямую в Banner
type BannerFeature struct {
	BannerID  int64 `db:"banner_id"`
	FeatureID int64 `db:"feature_id"`
}
