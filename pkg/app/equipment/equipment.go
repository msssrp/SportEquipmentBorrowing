package equipment

import "go.mongodb.org/mongo-driver/bson/primitive"

type Equipment struct {
	Id                 primitive.ObjectID `json:"equipment_id" bson:"_id,omitempty"`
	Name               string             `json:"name" bson:"name"`
	Category           string             `json:"category" bson:"category"`
	Description        string             `json:"description" bson:"description"`
	Quantity_available string             `json:"quantity_available" bson:"quantity_available"`
	Condition          string             `json:"condition" bson:"condition"`
	Image_url          string             `json:"image_url" bson:"image_url"`
}
