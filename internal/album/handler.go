package album

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	store *MemoryStore
}

func NewHandler() *Handler { return &Handler{store: NewMemoryStore()} }

func (h *Handler) Register(r *gin.Engine) {
	r.GET("/healthz", h.healthz)

	grp := r.Group("/albums")
	{
		grp.GET("", h.list)
		grp.GET("/:id", h.getByID)
		grp.POST("", h.create)
		grp.GET("/search", h.search) // /albums/search?title=Blue
	}
}

func (h *Handler) healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) list(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.store.List())
}

func (h *Handler) getByID(c *gin.Context) {
	id := c.Param("id")
	a, err := h.store.Get(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, a)
}

func (h *Handler) create(c *gin.Context) {
	var req Album
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, h.store.Add(req))
}

func (h *Handler) search(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing title"})
		return
	}
	c.IndentedJSON(http.StatusOK, h.store.SearchByTitle(title))
}
