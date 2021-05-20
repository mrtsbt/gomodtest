package firstmod

type FirstModule struct {
  Property string
}

func New() FirstModule {
  return FirstModule{
    Property: "initial"
  }
}
