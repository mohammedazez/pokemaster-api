package healthCheck

type (
	Service interface {
		Get() map[string]interface{}
	}
)
