package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	apipb "github.com/ptt/pttweb/proto/api"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 3000, "Port to listen on")
)

type server struct {
	apipb.UnimplementedBoardServiceServer
}

func (s *server) Hotboard(context.Context, *apipb.HotboardRequest) (*apipb.HotboardReply, error) {
	return &apipb.HotboardReply{
		Boards: []*apipb.Board{
			{
				Bid:        1,
				Name:       "Gossiping",
				Title:      "◎[老八] 共匪共諜就在本能寺",
				NumUsers:   8714,
				Bclass:     "綜合",
				Attributes: 0,
			},
			{
				Name:       "Stock",
				Title:      "◎[股票] 漲停板",
				NumUsers:   6031,
				Bclass:     "學術",
				Attributes: 0,
			},
			{
				Name:       "C_Chat",
				Title:      "◎[希洽] 發文時標題請防雷",
				NumUsers:   3792,
				Bclass:     "閒談",
				Attributes: 0,
			},
			{
				Name:       "Tech_Job",
				Title:      "◎[科技] 這裡是科技板",
				NumUsers:   378,
				Bclass:     "工作",
				Attributes: 0,
			},
			{
				Name:       "TY_Research",
				Title:      "◎強極渦劇場上映中",
				NumUsers:   27,
				Bclass:     "大氣",
				Attributes: 0,
			},
		},
	}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	apipb.RegisterBoardServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
