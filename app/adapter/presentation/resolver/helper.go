package resolver

import (
	"context"

	model "github.com/yuorei/video-server/app/domain/models"
)

// getUploaderForVideo is a helper function to get uploader information by UploaderID
// This eliminates code duplication between Video and VideoPayload uploader resolvers
func (r *Resolver) getUploaderForVideo(ctx context.Context, uploaderID string) (*model.User, error) {
	domainUser, err := r.app.User.GetUser(ctx, uploaderID)
	if err != nil {
		return nil, err
	}

	gqlUser := &model.User{
		ID:              domainUser.ID,
		Name:            domainUser.Name,
		ProfileImageURL: domainUser.ProfileImageURL,
		IsSubscribed:    domainUser.IsSubscribed,
		Role:            model.Role(domainUser.Role),
	}

	return gqlUser, nil
}

// getCurrentUserID extracts user ID from the context
func (r *Resolver) getCurrentUserID(ctx context.Context) string {
	// This assumes that authentication middleware sets the user ID in the context
	// The key should match what your authentication middleware uses
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		return ""
	}
	return userID
}
