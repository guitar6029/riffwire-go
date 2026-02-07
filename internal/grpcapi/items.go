package grpcapi

import (
	"context"

	itemsv1 "riffwire/internal/pb/items/v1"
	"riffwire/internal/store"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ItemsServer struct {
	itemsv1.UnimplementedItemsServiceServer
	store store.Store
}

// constructor for the store
func NewItemsServer(s store.Store) *ItemsServer {
	return &ItemsServer{store: s}
}

// retrieved the item
func (s *ItemsServer) GetItem(ctx context.Context, req *itemsv1.GetItemRequest) (*itemsv1.GetItemResponse, error) {
	id := int(req.GetId())
	if id <= 0 { // invalid armument
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}
	//if id found
	// get item
	item, ok := s.store.GetItemByID(id)
	if !ok {
		// not found
		return nil, status.Error(codes.NotFound, "item not found")
	}

	pbItem := &itemsv1.Item{Id: int32(item.ID),
		Name: item.Name}

	return &itemsv1.GetItemResponse{
		Item: pbItem,
	}, nil

}

func (s *ItemsServer) ListItems(ctx context.Context, req *itemsv1.ListItemsRequest) (*itemsv1.ListItemsResponse, error) {
	items := s.store.ListItems()

	pbItems := make([]*itemsv1.Item, 0, len(items))

	for _, curr := range items {
		pbItems = append(pbItems, &itemsv1.Item{
			Id:   int32(curr.ID),
			Name: curr.Name,
		})
	}

	return &itemsv1.ListItemsResponse{Items: pbItems}, nil
}
