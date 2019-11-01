package entities

//请求数据结构
type TestQueueRequest struct {
	Name    string `json:"name" validate:"required" example:"小二"`
	Mobile  string `json:"mobile" validate:"required" example:"15906063661"`
	Email   string `json:"email" validate:"required" example:"15906063661@qq.com"`
	TraceId string `json:"trace_id" example:"1212121"`
}

//返回数据结构
type TestQueueResponse struct {
	Status int `json:"id" example:"1"`
}

//请求数据结构
type TestQueueValidatorRequest struct {
	//tips，因为组件required不管是没传值或者传 0 or "" 都通过不了，但是如果用指针类型，那么0就是0，而nil无法通过校验
	Name   string `json:"mobile" validate:"required" example:"小二"`
	Mobile string `json:"mobile" validate:"required" example:"157xxxxx"`
	Email  string `json:"mobile" validate:"required" example:"15906063661@qq.com"`
}
