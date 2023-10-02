package entity

type Category struct {
	Reference   string       `bson:"reference" json:"reference"`
	Name        string       `bson:"name" json:"name" validate:"required,min=2"`
	Description string       `bson:"description" json:"description" validate:"required,min=10"`
	Industry    IndustryInfo `bson:"industry" json:"industry" validate:"required,min=3"`
	Type        string       `bson:"type" json:"type" validate:"required"`
	IsEnabled   bool         `bson:"is_enabled" json:"is_enabled" validate:"required"`
}

type IndustryInfo struct {
	Reference string `bson:"reference" json:"reference"`
	Name      string `bson:"name" json:"name" validate:"required,min=2"`
}
