// clientlib/clientlib.go
package clientlib

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ken1009us/unit-conversion-service/pb"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

type UnitConversionClient struct {
	conn   *grpc.ClientConn
	client pb.UnitConversionServiceClient
}

func NewUnitConversionClient() *UnitConversionClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewUnitConversionServiceClient(conn)
	return &UnitConversionClient{conn: conn, client: client}
}

func (ucc *UnitConversionClient) Convert(fromUnit, toUnit string, value float64) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := ucc.client.ConvertUnit(ctx, &pb.UnitConversionRequest{FromUnit: fromUnit, ToUnit: toUnit, Value: value})
	if err != nil {
		return 0, fmt.Errorf("could not convert: %v", err)
	}

	if r.GetConvertedValue() == 0 && r.GetError() != "" {
		return 0, fmt.Errorf("conversion error: %s", r.GetError())
	}

	return r.GetConvertedValue(), nil
}

func (ucc *UnitConversionClient) Close() {
	ucc.conn.Close()
}
