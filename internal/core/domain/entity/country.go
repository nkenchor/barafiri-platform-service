package entity

type Country struct {
	Reference   string `bson:"reference" json:"reference"`
	Name        string `bson:"name" json:"name" validate:"required,min=2"`
	Code        string `bson:"code" json:"code" validate:"required,numeric,min=3,max=3"`
	DialingCode string `bson:"dialing_code" json:"dialing_code" validate:"required,min=2"`
	ISOCODE2    string `bson:"iso_code_2" json:"iso_code_2" validate:"required,min=2"`
	ISOCODE3    string `bson:"iso_code_3" json:"iso_code_3" validate:"required,min=2"`
	IsEnabled   bool   `bson:"is_enabled" json:"is_enabled"`
}
