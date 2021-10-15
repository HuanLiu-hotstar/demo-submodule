package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"demo-submodule/graph/generated"
	"demo-submodule/graph/model"
)

func (r *queryResolver) GetPlaybackUrls(ctx context.Context, playbackRequestBody model.PlaybackRequestBody) (model.PlaybackResponse, error) {
	return model.PlaybackSuccessResponse{}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
