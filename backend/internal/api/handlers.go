package api

import (
	"net/http"
	"strconv"

	"nexus/backend/internal/domain"
	"nexus/backend/internal/repository"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store *repository.Store
}

func NewHandler(store *repository.Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) ListBlogPosts(c echo.Context) error {
	filters := domain.BlogPostFilters{
		Cursor:   c.QueryParam("cursor"),
		Limit:    parseInt(c.QueryParam("limit")),
		Query:    c.QueryParam("q"),
		SourceID: c.QueryParam("sourceId"),
		TagID:    c.QueryParam("tagId"),
	}

	items, err := h.store.ListBlogPosts(c.Request().Context(), filters)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "list blog posts")
	}

	return c.JSON(http.StatusOK, items)
}

func (h *Handler) ListStartups(c echo.Context) error {
	filters := domain.StartupFilters{
		Limit:    parseInt(c.QueryParam("limit")),
		Offset:   parseInt(c.QueryParam("offset")),
		Query:    c.QueryParam("q"),
		Sector:   c.QueryParam("sector"),
		Location: c.QueryParam("location"),
		Stage:    c.QueryParam("stage"),
		TagID:    c.QueryParam("tagId"),
	}

	items, err := h.store.ListStartups(c.Request().Context(), filters)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "list startups")
	}

	return c.JSON(http.StatusOK, items)
}

func parseInt(value string) int {
	if value == "" {
		return 0
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return parsed
}
