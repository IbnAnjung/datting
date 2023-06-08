package mysqlgorm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserSwapGormModel struct {
	UserID        int64 `gorm:"column:user_id"`
	SwappedUserId int64 `gorm:"column:swapped_user_id"`
	LikeCount     int64 `gorm:"column:like_count"`
	SkipCount     int64 `gorm:"column:skip_count"`
}

type UserSwapRepository struct {
	db    *gorm.DB
	table string
}

func NewUserSwapRepository(db *gorm.DB) UserSwapRepository {
	return UserSwapRepository{
		db:    db,
		table: "user_swaps",
	}
}

func (r UserSwapRepository) SwapToLike(ctx context.Context, userId, swappedUserId int64) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "swapped_user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"like_count": gorm.Expr("like_count + 1")}),
	}).Table(r.table).Create(map[string]interface{}{
		"user_id":         userId,
		"swapped_user_id": swappedUserId,
		"like_count":      1,
	}).Error

}
func (r UserSwapRepository) SwapToSkip(ctx context.Context, userId, swappedUserId int64) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "swapped_user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"pass_count": gorm.Expr("pass_count + 1")}),
	}).Table(r.table).Create(map[string]interface{}{
		"user_id":         userId,
		"swapped_user_id": swappedUserId,
		"pass_count":      1,
	}).Error
}
