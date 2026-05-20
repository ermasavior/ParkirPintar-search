package handler

import (
	pb "parkir-pintar/services/search/gen/search/v1"
	"parkir-pintar/services/search/internal/search/usecase"
)

// SearchServer implements the gRPC SearchServiceServer interface
type SearchServer struct {
	pb.UnimplementedSearchServiceServer
	uc usecase.Search
}

// NewSearchServer creates a new SearchServer
func NewSearchServer(uc usecase.Search) *SearchServer {
	return &SearchServer{uc: uc}
}
