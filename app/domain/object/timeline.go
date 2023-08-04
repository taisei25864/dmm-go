package object

import "time"

// 本当はPasswordHashがハッシュされたパスワードであることを型で保証したい。
// ハッシュ化されたパスワード用の型を用意してstringと区別して管理すると良い。
// 今回は簡単のためstringで管理している。

type Timeline struct {
	// The internal ID of the account
	Content   string    `json:"content"`
	ID        int64     `json:"id,omitempty"`
	AccountID int64     `json:"account_id,omitempty" db:"account_id"`
	CreateAt  time.Time `json:"create_at,omitempty" db:"create_at"`
}
