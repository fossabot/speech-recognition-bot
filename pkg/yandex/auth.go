package yandex

import "context"

// tokenAuth implements credentials.PerRPCCredentials
type tokenAuth struct {
	token string
}

// Return value is mapped to request headers
func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.token,
	}, nil
}

func (tokenAuth) RequireTransportSecurity() bool {
	return true
}
