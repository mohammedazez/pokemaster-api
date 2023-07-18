package healthCheck

type (
	Repository interface {
		Get() map[string]interface{}
	}
)
