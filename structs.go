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
	Type     string         `json:"type"`
	PID      int            `json:"pid"`
	Payloads []*interface{} `json:"payloads"`
}

type basePayload struct {
	Type string `json:"type"`
}

type overlayInit struct {
	Type        string           `json:"type"`
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

type tokenUpdate struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type storageSync struct {
}

type draftChange struct {
	ChannelID string `json:"channelId"`
	Content   string `json:"draft"`
}

type userProfileFetch struct {
	Type string `json:"type"`
	User struct {
		ID            string      `json:"id"`
		Username      string      `json:"username"`
		Avatar        string      `json:"avatar"`
		Discriminator string      `json:"discriminator"`
		PublicFlags   int         `json:"public_flags"`
		Flags         int         `json:"flags"`
		Banner        interface{} `json:"banner"`
		Bio           interface{} `json:"bio"`
	} `json:"user"`
	ConnectedAccounts []struct {
		Type     string `json:"type"`
		ID       string `json:"id"`
		Name     string `json:"name"`
		Verified bool   `json:"verified"`
	} `json:"connected_accounts"`
	PremiumSince      interface{} `json:"premium_since"`
	PremiumGuildSince interface{} `json:"premium_guild_since"`
	MutualGuilds      []struct {
		ID   string      `json:"id"`
		Nick interface{} `json:"nick"`
	} `json:"mutual_guilds"`
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
