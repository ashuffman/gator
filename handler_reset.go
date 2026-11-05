package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't delete users: %w", err)
	}

	fmt.Println("Table reset successfully")
	return nil
}
