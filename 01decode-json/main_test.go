package main

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}

	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}

	if post.Content != "Hello Dear Family!" {
		t.Error("Wrong content, was expecting 'Hello Dear Family!' but got", post.Content)
	}
	if post.Author.Name != "Kaoru" {
		t.Error("Wrong Auther Name,was expecting 'kaoru' but got", post.Author.Name)
	}

	for i, comment := range post.Comments {
		fmt.Println(i, comment)
		if i == 0 {
			if comment.Author != "Maiko" {
				t.Error("Wrong Comment Auther Name,was expecting 'Maiko' but got", comment.Author)
			}
			if comment.Content != "Have a great day!" {
				t.Error("Wrong Comment Content,was expecting 'Have a great day!' but got", comment.Content)
			}

		} else if i == 1 {
			if comment.Author != "Towa" {
				t.Error("Wrong Comment Auther Name,was expecting 'Towa' but got", comment.Author)
			}
			if comment.Content != "How are you today?" {
				t.Error("Wrong Comment Content,was expecting 'Have a great day!' but got", comment.Content)
			}
		}

	}
}
