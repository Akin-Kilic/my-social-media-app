package constant

const (
	Redis                = "key-%s-get-user_id:%d"
	RedisWithState       = "key-%s-get-user_id:%d-state:%t"
	RedisWithGroupID     = "key-%s-get-user_id:%d-group_id:%d"
	RedisWithKeyAndCount = "key-%s-get-user_id:%d-page:%d-count:%d"
	RedisWithPaginate    = "key-%s-get-user_id:%d-paginate:%s"
	RedisForJwt          = "token-%s-user_id:%s"
)
