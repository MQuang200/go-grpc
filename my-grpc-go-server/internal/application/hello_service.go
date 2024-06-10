package application

type HelloService struct {
}

func (helloService *HelloService) GenerateHello(name string) string {
	return "Hello" + name
}
