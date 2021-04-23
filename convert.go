package opensearch

func convertArgsToMap(args Args) (requestParams map[string]interface{}) {
	requestParams = make(map[string]interface{})
	if args.Query != "" {
		requestParams["query"] = args.Query
	}
	if args.Filter != "" {
		requestParams["filter"] = args.Filter
	}
	if args.Sort != "" {
		requestParams["sort"] = args.Sort
	}

	return
}

func convertScrollArgsToMap(args ScrollArgs) (requestParams map[string]interface{}) {
	requestParams = convertArgsToMap(args.Args)
	if args.Scroll != "" {
		requestParams["scroll"] = args.Scroll
	}
	if args.SearchType != "" {
		requestParams["search_type"] = args.SearchType
	}
	if args.ScrollId != "" {
		requestParams["scroll_id"] = args.ScrollId
	}

	return
}

func convertSearchArgsToMap(args SearchArgs) (requestParams map[string]interface{}) {
	requestParams = convertArgsToMap(args.Args)
	if args.Qp != "" {
		requestParams["qp"] = args.Qp
	}
	if args.Disable != "" {
		requestParams["disable"] = args.Disable
	}
	if args.FirstRankName != "" {
		requestParams["first_rank_name"] = args.FirstRankName
	}
	if args.SecondRankName != "" {
		requestParams["second_rank_name"] = args.SecondRankName
	}
	if args.UserId != "" {
		requestParams["user_id"] = args.UserId
	}
	if args.Abtest != "" {
		requestParams["abtest"] = args.Abtest
	}
	if args.RawQuery != "" {
		requestParams["raw_query"] = args.RawQuery
	}
	if args.ReSearch != "" {
		requestParams["re_search"] = args.ReSearch
	}
	if args.Biz != "" {
		requestParams["biz"] = args.Biz
	}
	if args.Summary != "" {
		requestParams["summary"] = args.Summary
	}
	if args.FromRequestId != "" {
		requestParams["from_request_id"] = args.FromRequestId
	}

	return
}
