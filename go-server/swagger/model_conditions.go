/*
 * Ad-service API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Conditions struct {
	AgeStart int32 `db:"age_start" json:"ageStart,omitempty"`

	AgeEnd int32 `db:"age_end" json:"ageEnd,omitempty"`

	Country []string `db:"country" json:"country,omitempty"`

	Platform []string `db:"platform" json:"platform,omitempty"`
}
