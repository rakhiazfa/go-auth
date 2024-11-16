package responses

import "github.com/rakhiazfa/vust-identity-service/pkg/utils"

type PaginationRes[T interface{}] struct {
	Items []T             `json:"items"`
	Meta  utils.Paginator `json:"meta"`
}
