package version1

type TimeSessionDataPage struct {
	Total *int64         `json:"total" bson:"total"`
	Data  []*TimeSession `json:"data" bson:"data"`
}

func NewEmptyTimeSessionDataPage() *TimeSessionDataPage {
	return &TimeSessionDataPage{}
}

func NewTimeSessionDataPage(total *int64, data []*TimeSession) *TimeSessionDataPage {
	return &TimeSessionDataPage{Total: total, Data: data}
}
