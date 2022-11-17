package main

import (
	"context"
	"github/go_grpc/pb/person"
	"google.golang.org/grpc"
	"net"
)

type personServer struct {
	person.UnimplementedSearchServiceServer
}

func (*personServer) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{Name: "我收到了" + name + "的信息"}
	return res, nil
}
func (*personServer) SearchIn(person.SearchService_SearchInServer) error {
	return nil
}
func (*personServer) SearchOut(*person.PersonReq, person.SearchService_SearchOutServer) error {
	return nil
}
func (*personServer) SearchIO(person.SearchService_SearchIOServer) error {
	return nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personServer{})
	s.Serve(l)

}
