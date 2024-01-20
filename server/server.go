package main

import (
	"context"
	"log"
	"net"

	"github.com/ken1009us/unit-conversion-service/units"

	pb "github.com/ken1009us/unit-conversion-service/pb"

	"google.golang.org/grpc"
)

const (
    port = ":50051"
)

type server struct {
    pb.UnimplementedUnitConversionServiceServer
    converter *units.SafeConverter
}

func (s *server) ConvertUnit(ctx context.Context, in *pb.UnitConversionRequest) (*pb.UnitConversionResponse, error) {
    log.Printf("Received: %v", in.GetValue())
    result, err := s.converter.Convert(in.GetValue(), in.GetFromUnit(), in.GetToUnit())
    if err != nil {
        return nil, err
    }
    return &pb.UnitConversionResponse{ConvertedValue: result}, nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterUnitConversionServiceServer(s, &server{converter: units.NewSafeConverter()})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
