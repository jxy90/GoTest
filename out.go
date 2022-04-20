package main

type ModelConfigs struct {
	AreaList       []AreaObject    `json:"area_list"`
	AlerterConfigs []AlerterConfig `json:"alertor_configs"`
}

type AreaObject struct {
	//区域类型(1-周界， 2-绊线， 3-双向绊线， 4-单向双绊线， 5-双向双绊线)
	AreaType string `json:"area_type"`
	// 保证areaId在同一个布控中的唯一性
	AreaId int `json:"area_id"`
	//入参规则数组的索引
	AlerterIds []int `json:"alertor_ids"`
	//interface类型同时保存单绊线和双绊线类型
	Points interface{} `json:"points"`
}

type AlerterConfig struct {
	//入参规则数组的索引
	AlerterId int `json:"alertor_id"`
	// intrusion, duty_querying, park, wander, exit, tripwire, overwall
	AlerterType []string `json:"alertor_type"`
	//最小检测目标的ratio
	TargetMin float32 `json:"target_min"`
	//最大检测目标的ratio
	TargetMax float32 `json:"target_max"`
	// person-人, car-机动车, cycle-非机动车
	TargetType []string `json:"target_type"`
	// unit:秒， 报警间隔
	CoolDownDuration int `json:"cooldown_duration"`
	//unit：秒，停留时长，"intrusion"，"park"or"wander"事件才会触发
	Duration int `json:"duration"`
	//unit:分钟， 值岗查询间隔，“duty_querying”事件才会触发
	Interval int `json:"interval,omitempty"`
	// 用于roi_counting，提供区间[a,b)，当人数在该区间内会产生报警
	ObjCountRange []int `json:"obj_count_range,omitempty"`
	// p14 roi_counting下限值
	TargetMinNumber int `json:"target_min_num"`
	// p14 roi_counting上限值
	TargetMaxNumber int `json:"target_max_num"`
	// p14聚众参数
	ObjCount int `json:"obj_count"`
	// 0-100的值，默认是50，越大表示越容易报警（物品）
	Sensitive int `json:"sensitive"`

	//和Glider确认并无该参数，因此去掉
	//TimeRange        []int    `json:"timeRange"`                 // P14聚众参数，使用默认值[0,86400]
}
