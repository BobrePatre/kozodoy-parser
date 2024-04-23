package service_access

import "net/http"

// Provider TODO: продумать межсервисное взаимодействие по http, с обеспечением
// безопасности (скорее всего продумать этот провайдер, который будет служить
// оберткой над http клиентом и менеджить токены доступа)
type Provider interface {
	DoRequest(*http.Request) (*http.Response, error)
}
