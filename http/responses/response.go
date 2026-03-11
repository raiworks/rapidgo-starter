package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIResponse is the standard JSON envelope for all API responses.
type APIResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	Meta       *Meta       `json:"meta,omitempty"`
	CursorMeta *CursorMeta `json:"cursor_meta,omitempty"`
}

// Meta holds pagination metadata.
type Meta struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// CursorMeta holds cursor-based pagination metadata.
type CursorMeta struct {
	NextCursor string `json:"next_cursor,omitempty"`
	PrevCursor string `json:"prev_cursor,omitempty"`
	PerPage    int    `json:"per_page"`
	HasMore    bool   `json:"has_more"`
}

// Success responds with 200 and the data wrapped in a success envelope.
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{Success: true, Data: data})
}

// Created responds with 201 and the data wrapped in a success envelope.
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, APIResponse{Success: true, Data: data})
}

// Error responds with the given status code and an error message.
func Error(c *gin.Context, status int, message string) {
	c.JSON(status, APIResponse{Success: false, Error: message})
}

// Paginated responds with 200, data, and pagination metadata.
func Paginated(c *gin.Context, data interface{}, page, perPage int, total int64) {
	totalPages := int(total) / perPage
	if int(total)%perPage != 0 {
		totalPages++
	}
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
		Meta: &Meta{
			Page:       page,
			PerPage:    perPage,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}

// CursorPaginated responds with 200, data, and cursor-based pagination metadata.
func CursorPaginated(c *gin.Context, data interface{}, nextCursor, prevCursor string, perPage int, hasMore bool) {
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    data,
		CursorMeta: &CursorMeta{
			NextCursor: nextCursor,
			PrevCursor: prevCursor,
			PerPage:    perPage,
			HasMore:    hasMore,
		},
	})
}
