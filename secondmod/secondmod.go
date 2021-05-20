package secondmod

type SecondModule struct {
	Property string
}

func New() SecondModule {
	return SecondModule{
		Property: "initial2",
	}
}
