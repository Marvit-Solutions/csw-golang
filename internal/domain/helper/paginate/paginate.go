package paginate

import "strconv"

func Paginate(pageParam string, defaultLimitStr string) (limit, offset int) {
	defaultLimit, err := strconv.Atoi(defaultLimitStr)
	if err != nil || defaultLimit < 1 {
		defaultLimit = 10
	}
	limit = defaultLimit

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	offset = (page - 1) * limit

	return limit, offset
}
