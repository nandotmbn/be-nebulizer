package views

type NebulizerView struct {
	NebulizerId   interface{} `json:"nebulizer_id,omitempty" validate:"required"`
	NebulizerName string      `json:"nebulizer_name,omitempty" bson:"nebulizer_name,omitempty" validate:"required,min=0"`
}

type PayloadRetriveId struct {
	NebulizerName string `json:"nebulizer_name,omitempty" bson:"nebulizer_name,omitempty" validate:"required,min=0"`
	Password      string `json:"password,omitempty" validate:"required,min=3,max=255"`
}

type FinalRetriveId struct {
	NebulizerId   interface{} `json:"nebulizer_id,omitempty" bson:"_id,omitempty" validate:"required"`
	NebulizerName string      `json:"nebulizer_name,omitempty" bson:"nebulizer_name,omitempty" validate:"required,min=0"`
}
