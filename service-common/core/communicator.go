package core

import "go.kicksware.com/api/service-common/core/meta"

type InnerCommunicator interface {
	PostMessage(endpoint string, message interface{}, response interface{}, params ...*meta.RequestParams) error
	GetMessage(endpoint string, response interface{}, params ...*meta.RequestParams) error
}