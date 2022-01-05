package haberdasherserver

import (
	"context"
	"github.com/twitchtv/twirp"
	"haberdasher/rpc/haberdasher"
	"math/rand"
)

// Server implements the Haberdasher service
type Server struct{}

func (s *Server) MakeHat(ctx context.Context, size *haberdasher.Size) (hat *haberdasher.Hat, err error) {
	if size.Inches <= 0 {
		return nil, twirp.InvalidArgumentError("inches", "I can't make a hat that small!")
	}
	return &haberdasher.Hat{
		Inches: size.Inches,
		Color:  []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)],
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}
