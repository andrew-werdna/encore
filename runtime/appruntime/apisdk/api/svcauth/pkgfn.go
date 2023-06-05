package svcauth

import (
	"fmt"

	"encore.dev/appruntime/apisdk/api/transport"
)

const (
	AuthMethodMetaKey = "Svc-Auth-Method"
)

// Sign signs the request using the given authentication method.
func Sign(method ServiceAuth, req transport.Transport) error {
	if err := method.sign(req); err != nil {
		return fmt.Errorf("failed to sign request: %w", err)
	}
	req.SetMeta(AuthMethodMetaKey, method.method())

	return nil
}

// Verify verifies the authenticity of the request using the given authentication methods.
func Verify(req transport.Transport, loadedAuthMethods ...ServiceAuth) (internalCall bool, err error) {
	method, found := req.ReadMeta(AuthMethodMetaKey)
	if !found {
		// If this is not set, it means that the request is not an internal service to service call.
		return false, nil
	}

	for _, authMethod := range loadedAuthMethods {
		if authMethod.method() == method {
			if err := authMethod.verify(req); err != nil {
				return false, fmt.Errorf("failed to verify request: %w", err)
			}
			return true, nil
		}
	}

	return false, fmt.Errorf("unknown service to service authentication method: %s", method)
}
