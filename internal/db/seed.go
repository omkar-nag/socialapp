package db

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/omkar-nag/socialapp/internal/store"
)

var usernames = []string{"alice", "bob", "charlie", "dan", "eve", "frank", "grace", "heidi", "ivan", "judy", "mallory", "nina", "olivia", "peter", "quinn", "rachel", "sam", "trudy", "ursula", "victor"}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)

	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			fmt.Printf("Error creating user: %v\n", err)
		}
	}

	posts := generatePosts(300, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			fmt.Printf("Error creating post: %v\n", err)
		}
	}

	comments := generateComments(500, posts, users)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			fmt.Printf("Error creating comment: %v\n", err)
		}
	}

	fmt.Println("Database seeded successfully!")
	fmt.Printf("Created %d users, %d posts, and %d comments.\n", len(users), len(posts), len(comments))
}

func generateUsers(n int) []*store.User {

	users := make([]*store.User, n)

	for i := 0; i < n; i++ {
		users[i] = &store.User{
			Username:  usernames[i%len(usernames)] + fmt.Sprint(i),
			Email:     usernames[i%len(usernames)] + fmt.Sprint(i) + "@example.com",
			Password:  "password" + fmt.Sprint(i),
			CreatedAt: time.Now().Format(time.RFC3339), // Example date, you can use time.Now().Format(time.RFC3339) for current time
		}
	}

	return users
}

func generatePosts(n int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, n)

	for i := 0; i < n; i++ {
		user := users[rand.Intn(len(users))] // Randomly select a user for the post
		posts[i] = &store.Post{
			Content:   fmt.Sprintf(user.Username) + "This is a sample post content for post " + fmt.Sprint(i),
			Title:     fmt.Sprintf(user.Username) + ": Sample Post " + fmt.Sprint(i),
			UserID:    user.ID,
			Comments:  []store.Comment{},
			Tags:      []string{},
			CreatedAt: time.Now().Format(time.RFC3339),
		}

	}
	return posts
}

func generateComments(n int, posts []*store.Post, users []*store.User) []*store.Comment {
	cms := make([]*store.Comment, n)
	for i := 0; i < n; i++ {
		post := posts[rand.Intn(len(posts))] // Randomly select a post for the comment
		user := users[rand.Intn(len(users))] // Randomly select a user for the comment
		cms[i] = &store.Comment{
			Content: fmt.Sprintf(user.Username) + ": This is a sample comment content for post " + fmt.Sprint(post.ID) + " comment " + fmt.Sprint(i),
			User: &store.User{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			},
			PostID:    post.ID,
			CreatedAt: time.Now().Format(time.RFC3339),
		}
	}

	return cms
}
