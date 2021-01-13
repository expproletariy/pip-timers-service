package persistence

import (
	dataV1 "github.com/expproletariy/pip-timers-service/data/version1"
	cdata "github.com/pip-services3-go/pip-services3-commons-go/data"
	mngpersist "github.com/pip-services3-go/pip-services3-mongodb-go/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type TimersMongoDBPersistence struct {
	mngpersist.IdentifiableMongoDbPersistence
}

func NewTimersMongoDBPersistence() *TimersMongoDBPersistence {
	proto := reflect.TypeOf(&dataV1.TimeSession{})
	persist := TimersMongoDBPersistence{}
	persist.IdentifiableMongoDbPersistence = *mngpersist.NewIdentifiableMongoDbPersistence(proto, "time_sessions")
	return &persist
}

func (p TimersMongoDBPersistence) composeFilter(filter *cdata.FilterParams) interface{} {
	if filter == nil {
		filter = cdata.NewEmptyFilterParams()
	}

	criteria := make([]bson.M, 0, 0)

	id := filter.GetAsString("id")
	if id != "" {
		criteria = append(criteria, bson.M{"_id": id})
	}

	userId := filter.GetAsString("user")
	if userId != "" {
		criteria = append(criteria, bson.M{"user": userId})
	}

	createdAt := filter.GetAsDateTime("created_at")
	if !createdAt.IsZero() {
		criteria = append(criteria, bson.M{"created_at": createdAt})
	}

	updatedAt := filter.GetAsDateTime("updated_at")
	if !updatedAt.IsZero() {
		criteria = append(criteria, bson.M{"updated_at": updatedAt})
	}

	status, isValidStatus := dataV1.SessionStatusFromString(filter.GetAsString("status"))
	if isValidStatus {
		criteria = append(criteria, bson.M{"status": status})
	}

	tags := filter.GetAsNullableArray("tags")
	if tags != nil && tags.Len() > 0 {
		criteria = append(criteria, bson.M{"tags": bson.D{{"$in", tags.Value()}}})
	}
	if len(criteria) > 0 {
		return bson.D{{"$and", criteria}}
	}
	return bson.M{}
}

func (p TimersMongoDBPersistence) GetPageByFilter(correlationId string, filter *cdata.FilterParams, paging *cdata.PagingParams) (page *dataV1.TimeSessionDataPage, err error) {
	tempPage, resErr := p.IdentifiableMongoDbPersistence.GetPageByFilter(correlationId, p.composeFilter(filter), paging, nil, nil)
	if resErr != nil {
		return nil, resErr
	}
	dataLen := int64(len(tempPage.Data))
	timeSessionData := make([]*dataV1.TimeSession, dataLen)
	for i, v := range tempPage.Data {
		timeSessionData[i] = v.(*dataV1.TimeSession)
	}
	page = dataV1.NewTimeSessionDataPage(&dataLen, timeSessionData)
	return page, nil
}

func (p TimersMongoDBPersistence) GetOneById(correlationId string, id string) (res *dataV1.TimeSession, err error) {
	result, err := p.IdentifiableMongoDbPersistence.GetOneById(correlationId, id)
	if result != nil {
		val, _ := result.(*dataV1.TimeSession)
		res = val
	}
	return res, err
}

func (p TimersMongoDBPersistence) Create(correlationId string, item *dataV1.TimeSession) (res *dataV1.TimeSession, err error) {
	value, err := p.IdentifiableMongoDbPersistence.Create(correlationId, item)
	if value != nil {
		val, _ := value.(*dataV1.TimeSession)
		res = val
	}
	return res, err
}

func (p TimersMongoDBPersistence) Update(correlationId string, item *dataV1.TimeSession) (res *dataV1.TimeSession, err error) {
	value, err := p.IdentifiableMongoDbPersistence.Update(correlationId, item)
	if value != nil {
		val, _ := value.(*dataV1.TimeSession)
		res = val
	}
	return res, err
}

func (p TimersMongoDBPersistence) DeleteById(correlationId string, id string) (res *dataV1.TimeSession, err error) {
	value, err := p.IdentifiableMongoDbPersistence.DeleteById(correlationId, id)
	if value != nil {
		val, _ := value.(*dataV1.TimeSession)
		res = val
	}
	return res, err
}
