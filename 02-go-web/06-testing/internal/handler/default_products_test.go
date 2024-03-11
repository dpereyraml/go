package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing-integration/internal/handler"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
)

func TestProductDefault_GetProduct_HandlerFunc(t *testing.T) {
	t.Run("success - get product", func(t *testing.T) {
		// arrange
		// - repository: map
		/* db := map[int]internal.Product{
			1: {
				ID:          1,
				Name:        "product 1",
				Quantity:    1,
				CodeValue:   "code 1",
				IsPublished: true,
				Expiration:  "01/01/2021",
				Price:       1.0,
			},
		} */
		// rp := repository.NewProductMap(db)

		// - handler
		// hd := NewProductDefault(rp)
		h := handler.NewHandlerProducts(nil, 0)
		hdFunc := h.ProductByID()

		// act
		req := httptest.NewRequest("GET", "/products/1", nil)

		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		res := httptest.NewRecorder()
		hdFunc(res, req)

		// assert
		/* expectedProduct, err := json.Marshal(handler.ProductJSON{
			ID:          1,
			Name:        "Oil - Margarine",
			Quantity:    439,
			CodeValue:   "S82254D",
			IsPublished: true,
			Expiration:  "15/12/2021",
			Price:       71.42,

				// "id": 1,
				// "name": "Oil - Margarine",
				// "quantity": 439,
				// "code_value": "S82254D",
				// "is_published": true,
				// "expiration": "15/12/2021",
				// "price": 71.42
		}) */
		expectedProduct := `{
			"id": 1,
			"name": "Oil - Margarine",
			"quantity": 439,
			"code_value": "S82254D",
			"is_published": true,
			"expiration": "15/12/2021",
			"price": 71
		}`

		// require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, expectedProduct, res.Body.String())
		// require.JSONEq(t, `{
		// 	"id": 1,
		// 	"name": "product 1",
		// 	"description": "description 1",
		// 	"completed": false
		// }`, res.Body.String())
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	})
}
