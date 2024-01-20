package main

import (
	"context"
	"log"
	"net"

	"github.com/ken1009us/unit-conversion-service/pb"
	"github.com/ken1009us/unit-conversion-service/units"

	"google.golang.org/grpc"
)

const (
    port = ":50051"
)

type server struct {
    pb.UnimplementedUnitConversionServiceServer
    unitsConverter *units.Units
}

func (s *server) ConvertUnit(ctx context.Context, in *pb.UnitConversionRequest) (*pb.UnitConversionResponse, error) {
    log.Printf("Received: %v %s", in.GetValue(), in.FromUnit)
    result, err := s.unitsConverter.Convert(in.GetValue(), in.GetFromUnit(), in.GetToUnit())
    if err != nil {
        // Return an error response
        return &pb.UnitConversionResponse{
            ConvertedValue: 0,
            Error:          err.Error(),
        }, nil
    }
    return &pb.UnitConversionResponse{ConvertedValue: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUnitConversionServiceServer(s, &server{unitsConverter: units.NewUnits()})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
