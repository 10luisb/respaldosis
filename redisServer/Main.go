package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	pb "redisServer/grpcServerRedis"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

var ctx = context.Background()
var rdb *redis.Client

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port = ":3001"
)

type Data struct {
	Album  string
	Year   string
	Artist string
	Ranked string
}

func redisConnect() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "10.153.60.51:6379",
		Password: "",
		DB:       15,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pong)
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí de cliente: ", in.GetArtist())
	data := Data{
		Year:   in.GetYear(),
		Album:  in.GetAlbum(),
		Artist: in.GetArtist(),
		Ranked: in.GetRanked(),
	}
	insertRedis(data)
	return &pb.ReplyInfo{Info: "Hola cliente, recibí el comentario"}, nil
}

func insertRedis(rank Data) {
	array := rank.Artist + "-" + rank.Year
	ranked, _ := strconv.ParseFloat(rank.Ranked, 64)

	rdb.ZAddArgsIncr(ctx, array, redis.ZAddArgs{
		XX:      false,
		NX:      true,
		Members: []redis.Z{{Score: ranked, Member: rank.Album}},
	})

	key := array + "-" + rank.Album
	rdb.HIncrBy(ctx, key, rank.Ranked, 1)
}

func main() {
	listen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	redisConnect()

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
