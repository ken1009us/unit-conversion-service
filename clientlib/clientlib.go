// clientlib/clientlib.go
package clientlib

import (
	"context"
	"fmt"

	pb "github.com/ken1009us/unit-conversion-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const address = "localhost:50051"

type UnitConversionClient struct {
    conn   *grpc.ClientConn
    client pb.UnitConversionServiceClient
}

func NewUnitConversionClient() (*UnitConversionClient, error) {
    conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
    if err != nil {
        return nil, fmt.Errorf("failed to connect to the server: %v", err)
    }
    client := pb.NewUnitConversionServiceClient(conn)
    return &UnitConversionClient{conn: conn, client: client}, nil
}

func (ucc *UnitConversionClient) Close() {
    ucc.conn.Close()
}

func (ucc *UnitConversionClient) Convert(ctx context.Context, fromUnit, toUnit string, value float64) (float64, error) {
    r, err := ucc.client.ConvertUnit(ctx, &pb.UnitConversionRequest{FromUnit: fromUnit, ToUnit: toUnit, Value: value})
    if err != nil {
        st, ok := status.FromError(err)
        if !ok {
            return 0, fmt.Errorf("error during conversion: %v", err)
        }
        return 0, fmt.Errorf("conversion error: %s (code: %v)", st.Message(), st.Code())
    }

    if r.GetConvertedValue() == 0 && r.GetError() != "" {
        return 0, fmt.Errorf("conversion error: %s", r.GetError())
    }

    return r.GetConvertedValue(), nil
}