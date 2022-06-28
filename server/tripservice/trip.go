package trip

import (
	"context"
	trippb "coolcar/proto/gen/go"
)

type Service struct {
	*trippb.UnimplementedTripServiceServer
}

func (service *Service) GetTrip(c context.Context, request *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: request.Id,
		Trip: &trippb.Trip{
			Start:       "start",
			End:         "end",
			DurationSec: 100,
			FeeCent:     200,
			StartPos: &trippb.Location{
				Latitude:  120,
				Longitude: 30,
			},
			EndPos: &trippb.Location{
				Latitude:  130,
				Longitude: 30,
			},
			Status: trippb.TripStatus_IN_PROGRESS,
		},
	}, nil
}
