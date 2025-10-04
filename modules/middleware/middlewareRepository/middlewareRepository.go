package middlewareRepository

type (
	MiddlewareRepositoryService interface{}

	middlewareRepository struct {
		// db *mongo.Client
	}
)

func NewMiddlewareRepository(
// db *mongo.Client,
) MiddlewareRepositoryService {
	return &middlewareRepository{}
}
