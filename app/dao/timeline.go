package dao

// import (
// 	"github.com/jmoiron/sqlx"
// )

// type (
// 	timeline struct {
// 		db *sqlx.DB
// 	}
// )

// func NewTimeline(db *sqlx.DB) repository.Timeline {
// 	return &timeline{db: db}
// }

// // GetPublicTimeline : queryに応じたstatus取得
// func (r *timeline) GetPublicTimeliine(ctx context.Context, MaxID int64, SinceID int64, Limit int64) (*object.Timeline, error) {
// 	entity := new(object.Timeline)
// 	err := r.db.QueryRowxContext(ctx, "select * from statuses where id > ? AND where id < ? AND LIMIT ?", SinceID, MaxID, Limit).StructScan(entity)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, nil
// 		}

// 		return nil, fmt.Errorf("failed to find status from db: %w", err)
// 	}

// 	return entity, nil
// }
