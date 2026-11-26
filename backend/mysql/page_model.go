package mysql

// PageModel 通用分页struct
type PageModel struct {
	Pagesize    int64    //每页显示多少条
	Page        int64    //页码 页码从1开始 默认不分页
	TotalCount  int64    //总条数
	SelectCount bool     //是否查询总数，false ：不查询；true：查询 默认不查询
	OrderBy     string   // 排序信息 默认不排序
	Select      []string //查询的字段
}

/*
 * 获取分页总数
 */
func (pageModel *PageModel) GetMaxPage() int64 {
	if pageModel.TotalCount == 0 || pageModel.Pagesize <= 0 {
		return 0
	}
	if pageModel.TotalCount%pageModel.Pagesize == 0 {
		return pageModel.TotalCount / pageModel.Pagesize
	} else {
		return pageModel.TotalCount/pageModel.Pagesize + 1
	}
}

/*
 *  获取分页查询起始位置
 */
func (pageModel *PageModel) GetStart() int64 {
	if pageModel.Page > 0 {
		return (pageModel.Page - 1) * pageModel.Pagesize
	}
	return 0
}

/*
 *  获取分页查询结束位置
 */
func (pageModel *PageModel) GetEnd() int64 {
	if pageModel.Page > 0 {
		return (pageModel.Page)*pageModel.Pagesize - 1
	}
	if pageModel.Pagesize > 0 {
		return pageModel.Pagesize - 1
	}
	return 0
}
