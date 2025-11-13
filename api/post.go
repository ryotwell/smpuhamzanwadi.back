package api

import (
	"net/http"
	"project_sdu/model"
	"project_sdu/service"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func slugify(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))
	re := regexp.MustCompile(`[^a-z0-9\s-]`)
	input = re.ReplaceAllString(input, "")
	re = regexp.MustCompile(`[\s\-_]+`)
	input = re.ReplaceAllString(input, "-")
	input = strings.Trim(input, "-")
	return input
}

type PostAPI interface {
	CreatePost(c *gin.Context)
	GetPostByID(c *gin.Context)
	GetPostBySlug(c *gin.Context)
	GetAllPosts(c *gin.Context)
	GetPublishedPosts(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}

type postAPI struct {
	postService service.PostService
}

func NewPostAPI(postService service.PostService) *postAPI {
	return &postAPI{postService}
}

// ====================
// CREATE POST
// ====================
func (p *postAPI) CreatePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors:  map[string]string{"body": "Invalid JSON format"},
		})
		return
	}

	// Minimal validation
	errors := make(map[string]string)
	if post.Title == "" {
		errors["title"] = "title is required"
	}
	if post.Content == "" {
		errors["content"] = "content is required"
	}

	// Always (re-)generate slug from Title.
	if post.Title != "" {
		post.Slug = slugify(post.Title)
	} else {
		post.Slug = ""
	}

	if post.Slug == "" {
		errors["slug"] = "slug is required"
	}

	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Validation failed",
			Errors:  errors,
		})
		return
	}

	if err := p.postService.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to create post",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusCreated, model.SuccessResponse{
		Success: true,
		Status:  http.StatusCreated,
		Message: "Post created successfully",
		Data:    post,
	})
}

// ====================
// GET POST BY ID
// ====================
func (p *postAPI) GetPostByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid post ID",
		})
		return
	}

	post, err := p.postService.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "Post not found",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Post retrieved successfully",
		Data:    post,
	})
}

// ====================
// GET POST BY SLUG
// ====================
func (p *postAPI) GetPostBySlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Slug is required",
		})
		return
	}

	post, err := p.postService.GetPostBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "Post not found",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Post retrieved successfully",
		Data:    post,
	})
}

// ====================
// GET ALL POSTS
// ====================
func (p *postAPI) GetAllPosts(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	pageParam := c.DefaultQuery("page", "1")
	q := c.Query("q")

	limit, _ := strconv.Atoi(limitParam)
	page, _ := strconv.Atoi(pageParam)

	posts, err := p.postService.GetAllPosts(limit, page, q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve posts",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Posts retrieved successfully",
		Data:    posts,
		Meta: gin.H{
			"page":  page,
			"limit": limit,
		},
	})
}

// ====================
// GET ALL PUBLISHED POSTS
// ====================
func (p *postAPI) GetPublishedPosts(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitParam)
	offset, _ := strconv.Atoi(offsetParam)

	posts, err := p.postService.GetPublishedPosts(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve published posts",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Published posts retrieved successfully",
		Data:    posts,
	})
}

// ====================
// UPDATE POST
// ====================
func (p *postAPI) UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid post ID",
		})
		return
	}

	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid JSON format",
			Errors:  map[string]string{"body": err.Error()},
		})
		return
	}

	// Re-generate slug from Title when updating, if Title is present and non-empty.
	if post.Title != "" {
		post.Slug = slugify(post.Title)
	}

	if err := p.postService.UpdatePost(id, &post); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to update post",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Post updated successfully",
		Data:    post,
	})
}

// ====================
// DELETE POST
// ====================
func (p *postAPI) DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: "Invalid post ID",
		})
		return
	}

	if err := p.postService.DeletePost(id); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete post",
			Errors:  map[string]string{"server": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Post deleted successfully",
	})
}
