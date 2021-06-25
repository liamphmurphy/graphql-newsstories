package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/murnux/newsstories/graph/generated"
	"github.com/murnux/newsstories/graph/model"
)

// CreateStory implements support for a user to add a new story.
func (r *mutationResolver) CreateStory(ctx context.Context, input model.NewStory) (*model.Story, error) {
	// duplicate story check
	for _, st := range r.stories {
		if input.StoryTitle == st.Title {
			return nil, fmt.Errorf("A story with the title %s already exists.", input.StoryTitle)
		}
	}

	// if no dup found, continue on
	story := &model.Story{
		ID:    fmt.Sprintf("T%d", rand.Int()),
		Title: input.StoryTitle,
		Body:  input.StoryBody,
	}

	r.stories = append(r.stories, story)
	return story, nil
}

// Stories returns all stories in the DB.
func (r *queryResolver) Stories(ctx context.Context) ([]*model.Story, error) {
	return r.stories, nil
}

// Story allows for querying for a specific story by the title.
func (r *queryResolver) Story(ctx context.Context, title string) (*model.Story, error) {
	// loop through DB to find the story
	for _, story := range r.stories {
		if story.Title == title {
			return story, nil
		}
	}

	return nil, fmt.Errorf("Could not find a story with the title %s", title)
}

// RandomStory returns a random story. Should only fail if no stories exist in the DB.
func (r *queryResolver) RandomStory(ctx context.Context) (*model.Story, error) {
	if len(r.stories) == 0 {
		return nil, fmt.Errorf("Cannot find a random story, as no stories exist.")
	}
	randIndex := rand.Intn(len(r.stories))
	return r.stories[randIndex], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
