package main

type sendReady struct {
	Cmd   string      `json:"cmd"`
	Args  interface{} `json:"args"`
	Evt   string      `json:"evt"`
	Nonce string      `json:"nonce"`
}

type connectArgs struct {
	Type string `json:"type"`
	Pid  int    `json:"pid"`
}

type connect struct {
	Cmd   string       `json:"cmd"`
	Args  *connectArgs `json:"args"`
	Nonce string       `json:"nonce"`
}

type base struct {
	Cmd   string      `json:"cmd"`
	Data  baseData    `json:"data"`
	Event string      `json:"evt"`
	Nonce interface{} `json:"nonce"`
}

type baseData struct {
	Type     string      `json:"type"`
	PID      int         `json:"pid"`
	Payloads []*payloads `json:"payloads"`
}

type payloads struct {
	User        user             `json:"user"`
	Token       string           `json:"token"`
	PaymentInfo []*paymentSource `json:"paymentSources"`
	Friends     map[string]int   `json:"relationships"`
	MediaState  struct {
		Input map[string]struct {
			Name string `json:"name"`
		} `json:"inputDevices"`
		Output map[string]struct {
			Name string `json:"name"`
		} `json:"outputDevices"`
	} `json:"mediaEngineState"`
}

type paymentSource struct {
	Address struct {
		Name          string `json:"name"`
		FirstAddress  string `json:"line1"`
		SecondAddress string `json:"line2"`
		City          string `json:"city"`
		PostalCode    string `json:"postalCode"`
		State         string `json:"state"`
		Country       string `json:"country"`
	} `json:"billingAddress"`
	Email string `json:"email"`
}

type user struct {
	Discriminator string `json:"discriminator"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone"`
}
