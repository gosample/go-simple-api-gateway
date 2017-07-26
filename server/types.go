package server

type Policy struct {
	User   string `json:"user" form:"user" query:"user"`
	Path   string `json:"path" form:"path" query:"path"`
	Method string `json:"method" form:"method" query:"method"`
}

type UserRole struct {
	User string `json:"user" form:"user" query:"user"`
	Role string `json:"role" form:"role" query:"role"`
}

type SuccessMessage struct {
	Status  int `json:"status" form:"status" query:"status"`
	Success bool `json:"success" form:"success" query:"success"`
}

type DataMessage struct {
	Status int `json:"status" form:"status" query:"status"`
	Data   interface{}`json:"data" form:"data" query:"data"`
}
