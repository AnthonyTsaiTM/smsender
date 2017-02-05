package store

import "github.com/minchao/smsender/smsender/model"

type StoreResult struct {
	Data interface{}
	Err  error
}

type StoreChannel chan StoreResult

type Store interface {
	Route() RouteStore
	Message() MessageStore
}

type RouteStore interface {
	GetAll() StoreChannel
	SaveAll(routes []*model.Route) StoreChannel
}

type MessageStore interface {
	Get(id string) StoreChannel
	GetByIds(ids []string) StoreChannel
	GetByProviderAndMessageId(provider, originalMessageId string) StoreChannel
	Save(message *model.MessageRecord) StoreChannel
	Search(params map[string]interface{}) StoreChannel
	Update(message *model.MessageRecord) StoreChannel
}
