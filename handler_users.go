package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	usersList, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error encountered getting user list: %w", err)
	}

	for _, user := range usersList {
		if user.Name == s.configPtr.CurrentUserName {
			fmt.Printf("* %v (current)", user.Name)
		} else {
			fmt.Printf("* %v", user.Name)
		}
	}
	return nil
}
