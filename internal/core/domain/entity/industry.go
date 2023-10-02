package entity

type Industry struct {
	Reference   string `bson:"reference" json:"reference"`
	Name        string `bson:"name" json:"name" validate:"required,min=2"`
	Description string `bson:"description" json:"description" validate:"required,min=10"`
	IsEnabled   bool   `bson:"is_enabled" json:"is_enabled" validate:"required"`
}
