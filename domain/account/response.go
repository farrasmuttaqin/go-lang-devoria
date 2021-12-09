package account

type AccountAuthenticationResponse struct {
	Token   string  `json:"token"`
	Profile Account `json:"profile"`
}

type AccountProfileResponse struct {
	Profile Account `json:"profile"`
}
