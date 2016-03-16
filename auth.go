package apptweak

type Auth struct {
	token string
}

func NewAuth(t string) Auth {
	return Auth{token: t}
}
