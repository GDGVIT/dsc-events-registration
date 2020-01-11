package participants

type Participant struct {
	EmailAddress       string `json:"emailAddress" bson:"emailAddress" validate:"nonzero"`
	DeviceID           string `json:"deviceID,omitempty" bson:"deviceID,omitempty"`
	PhoneNumber        string `json:"PhoneNumber" bson:"phoneNumber" validate:"min=7,max=22,nonzero"`
	Name               string `json:"name" bson:"name" validate:"nonzero"`
	RegistrationNumber string `json:"RegistrationNumber" bson:"RegistrationNumber" validate:"nonzero"`
	EventName          string `json:"eventName" bson:"eventName" validate:"nonzero"`
}
