package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"

	greetv1 "github.com/MatsuoTakuro/my-template-connect-go/gen/greet/v1"
	"github.com/MatsuoTakuro/my-template-connect-go/gen/greet/v1/greetv1connect"
	"github.com/MatsuoTakuro/my-template-connect-go/services"
)

type GreetController struct {
	service services.GreetServicer
}

func NewGreetController(s services.GreetServicer) *GreetController {
	return &GreetController{service: s}
}

func (c *GreetController) GreetHandler() (string, http.Handler) {
	return greetv1connect.NewGreetServiceHandler(c)
}

func (c *GreetController) Greet(
	ctx context.Context, req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
