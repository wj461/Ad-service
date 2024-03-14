package swagger

type AdData struct {
	AdId *int `db:"ad_id" json:"ad_id"`

	Title *string `db:"title" json:"title"`

	StartAt *string `db:"start_at" json:"start_at"`

	EndAt *string `db:"end_at" json:"end_at"`

	AgeStart *int `db:"age_start" json:"age_start"`

	AgeEnd *int `db:"age_end" json:"age_end"`

	Male *bool `db:"male" json:"male"`

	Female *bool `db:"female" json:"female"`

	Country *[]string `db:"country" json:"country"`

	platform *[]string `db:"platform" json:"platform"`
}
