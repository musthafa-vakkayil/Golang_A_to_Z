package main

import "fmt"

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	// Check if user exists in the friendship map
	directFriends, exists := friendships[username]
	if !exists {
		return nil
	}

	// Use a map to track direct friends and suggested friends
	seen := make(map[string]bool)

	// Mark all direct friends as seen
	for _, friend := range directFriends {
		seen[friend] = true
	}

	// Collect suggested friends (friends of direct friends)
	suggestedFriends := []string{}
	for _, friend := range directFriends {
		for _, friendOfFriend := range friendships[friend] {
			if friendOfFriend == username || seen[friendOfFriend] {
				continue
			}
			// Add to result and mark as seen
			suggestedFriends = append(suggestedFriends, friendOfFriend)
			seen[friendOfFriend] = true
		}
	}

	if len(suggestedFriends) == 0 {
		return nil
	}

	return suggestedFriends
}

func main() {
	friendships := map[string][]string{
		"Alice":   {"Bob", "Charlie"},
		"Bob":     {"Alice", "Charlie", "David"},
		"Charlie": {"Alice", "Bob", "David", "Eve"},
		"David":   {"Bob", "Charlie"},
		"Eve":     {"Charlie"},
	}

	suggestedFriends := findSuggestedFriends("Alice", friendships)
	fmt.Println(suggestedFriends)
}
