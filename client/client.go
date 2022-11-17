package main

import (
	"context"
	"fmt"
	"github/go_grpc/pb/person"
	"google.golang.org/grpc"
)

func main() {
	l, _ := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	client := person.NewSearchServiceClient(l)
	res, err := client.Search(context.Background(), &person.PersonReq{Name: "wys"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
