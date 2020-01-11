package participants

type Participant struct {
	EmailAddress       string `json:"emailAddress" bson:"emailAddress"`
	DeviceID           string `json:"deviceID" bson:"deviceID"`
	PhoneNumber        string `json:"PhoneNumber" bson:"phoneNumber"`
	Name               string `json:"name" bson:"name"`
	RegistrationNumber string `json:"RegistrationNumber" bson:"RegistrationNumber"`
	EventName          string `json:"eventName" bson:"eventName"`
}
