package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Post struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Author         string    `json:"author"`
	ContentLimited string    `json:"contentlimited"`
	Createdat      time.Time `json:"createdat"`
}

func NewPost(id, title, content, author string) Post {
	contentlimited := content
	if len(contentlimited) > 50 {
		contentlimited = contentlimited[:50]
	}
	return Post{
		ID:             id,
		Title:          title,
		Content:        content,
		Author:         author,
		ContentLimited: contentlimited,
		Createdat:      time.Now(),
	}
}

var posts []Post = []Post{
	NewPost("1", "This is my first Post", "This is my first post. kinda exicted and nervous at the same time!", "Guest"),
}

func main() {
	fmt.Println("Gin has started !")
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, posts)
	})

	router.GET("/detail/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, post := range posts {
			if post.ID == id {
				c.JSON(http.StatusOK, post)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Post not found!"})
	})

	router.POST("/addpost", func(c *gin.Context) {
		var newPost Post
		if err := c.ShouldBindJSON(&newPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newPost.ID = fmt.Sprintf("%d", len(posts)+1)
		newPost = NewPost(newPost.ID, newPost.Title, newPost.Content, newPost.Author)
		posts = append(posts, newPost)
		c.Redirect(http.StatusSeeOther, "/")
	})

	// x := fmt.Sprintf("%d", len(posts)+1)
	// fmt.Println(x)
	// fmt.Println(reflect.TypeOf(x))
	router.Run(":8000")
}

/*

 <script>
        document.getElementById('postForm').addEventListener('submit', function (e) {
            e.preventDefault(); // Prevent the form from submitting the default way

            const title = document.getElementById('title').value;
            const content = document.getElementById('content').value;
            const author = document.getElementById('author').value;

            fetch("http://localhost:8000/addpost", {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json', // Send JSON format
                },
                body: JSON.stringify({
                    title: title,
                    content: content,
                    author: author
                })
            })
            .then(response => response.json()) // Assuming the server sends JSON data
            .then(data => {
                console.log('Success:', data);
                alert('Post added successfully!');
            })
            .catch((error) => {
                console.error('Error:', error);
                alert('Error adding post!');
            });
        });
    </script>

*/
