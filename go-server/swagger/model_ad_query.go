package swagger

type AdQuery struct {
	OffsetVal *int `db:"offset_val" json:"offset" form:"offset"`

	LimitVal *int `db:"limit_val" json:"limit" form:"limit"`

	Age *int `db:"age" json:"age" form:"age"`

	Gender *string `db:"gender" json:"gender" form:"gender"`

	Country *string `db:"country" json:"country" form:"country"`

	Platform *string `db:"platform" json:"platform" form:"platform"`
}
