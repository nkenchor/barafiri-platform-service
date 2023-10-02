package entity

type Currency struct {
	Reference   string `bson:"reference" json:"reference"`
	Name        string `bson:"name" json:"name" validate:"required,min=2"`
	Code        string `bson:"code" json:"code" validate:"required,min=3"`
	CountryCode string `bson:"country_code" json:"country_code" validate:"required,min=1,max=3"`
	IsEnabled   bool   `bson:"is_enabled" json:"is_enabled"`
}
