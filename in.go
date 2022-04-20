package main

type GuardRuleInput struct {
	// 规则类型，枚举值 与接口guardRecord/conf警戒配置数据一致。如果是离线任务可以参考在线任务的该参数枚举。
	AlerterType int `json:"alertorType" binding:"required"`
	// 目标最大值，（0， 100]之间的数据
	TargetMax int `json:"targetMax" binding:"required,min=1,max=100"`
	// 目标最小值， [0, 100)之间的数据
	TargetMin int `json:"targetMin" binding:"required,min=0,max=99"`
	// 冷却时间(或 报警间隔)，两次报警产生的最小时间间隔，单位是秒
	CoolDownDuration int `json:"coolDownDuration" binding:"required,min=0"`
	// 持续时间(或 延迟报警)， 报警产生的持续时间，只有事件持续产生并达到这个时间才会产生报警，单位秒
	Duration int `json:"duration" binding:"required,min=0"`
	// 数据上报间隔，只有值岗离岗事件才使用此字段，每隔interval时间会上报一条数据，单位秒
	Interval int `json:"interval" binding:"min=0"`
	// 区域列表(两个点组成的一条线,或者多个点组成的一个区域) 双绊线的情况内部为两个元素
	AreaList []GuardRuleAreaInput `json:"areaList" binding:"required"`
	// 数据上限，少员时用到，报警的范围为[0, limit); 数据下限， 超员或聚众时用到，报警的范围为(limit, 100000), 100000为应用内部设置的最大值
	Limit int `json:"limit"`
	// 规则是否启用开关： 1启用；2关闭
	Enabled int `json:"enabled" binding:"required"`
}

type GuardRuleAreaInput struct {
	// 区域点坐标信息
	Points []PointInput `json:"point" binding:"required"`
	// 区域ID，从0开始，并且唯一
	AreaId int `json:"areaId" binding:"required,min=0"`
	// 区域类型(1-周界， 2-绊线， 3-双向绊线， 4-单向双绊线， 5-双向双绊线)
	AreaType int `json:"areaType" binding:"required"`
	// 区域规则类型(1-报警区域， 2- 屏蔽区域， 3-mask区域， 4-预警区域)
	AreaRuleType int `json:"areaRuleType" binding:"required"`
}

type PointInput struct {
	// 横坐标，归一化坐标点
	X float64 `json:"x" binding:"required,min=0,max=1"`
	// 纵坐标，归一化坐标点
	Y float64 `json:"y" binding:"required,min=0,max=1"`
}
