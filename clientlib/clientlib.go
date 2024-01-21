// clientlib/clientlib.go
package clientlib

import (
	"context"
	"fmt"
	"sync"

	pb "github.com/ken1009us/unit-conversion-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:50051"

var (
	clientInstance *UnitConversionClient
	once           sync.Once
)

type UnitConversionClient struct {
	conn   *grpc.ClientConn
	client pb.UnitConversionServiceClient
}

func NewUnitConversionClient() (*UnitConversionClient, error) {
	var err error
	once.Do(func() {
		conn, connErr := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if connErr != nil {
			err = fmt.Errorf("failed to connect to the server: %v", connErr)
			return
		}
		client := pb.NewUnitConversionServiceClient(conn)
		clientInstance = &UnitConversionClient{conn: conn, client: client}
	})
	if err != nil {
		return nil, err
	}
	return clientInstance, nil
}

func (ucc *UnitConversionClient) Convert(ctx context.Context, fromUnit, toUnit string, value float64) (float64, error) {
	r, err := ucc.client.ConvertUnit(ctx, &pb.UnitConversionRequest{FromUnit: fromUnit, ToUnit: toUnit, Value: value})
	if err != nil {
		return 0, fmt.Errorf("error during conversion: %v", err)
	}

	if r.GetConvertedValue() == 0 && r.GetError() != "" {
		return 0, fmt.Errorf("conversion error: %s", r.GetError())
	}

	return r.GetConvertedValue(), nil
}

func (ucc *UnitConversionClient) Close() {
	ucc.conn.Close()
}