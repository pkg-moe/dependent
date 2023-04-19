package dependent

var serviceNeedList = map[ServiceNeed]struct{}{}

type ServiceNeed int64

func NeedSerivce(item ServiceNeed) {
	serviceNeedList[item] = struct{}{}
}

func NeedCheck(item ServiceNeed) bool {
	_, ok := serviceNeedList[item]
	return ok
}
