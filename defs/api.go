package defs

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
	State   int    `json:"state"`
}

type BuildParams struct {
	Version string   `json:"version"`
	Apps    []string `json:"apps"`
}

var (
	SUCCESS = Resp{
		Code:    0,
		Message: "success",
		Detail:  "",
	}
	UPLOAD_FAILED = Resp{
		Code:    1,
		Message: "upload faild: ",
		Detail:  "",
	}
	BUILD_FAILED = Resp{
		Code:    2,
		Message: "build failed: ",
		Detail:  "",
	}
	PUSH_FAILED = Resp{
		Code:    3,
		Message: "push failed: ",
		Detail:  "",
	}
	CLEAN_FAILED = Resp{
		Code:    4,
		Message: "clean failed: ",
		Detail:  "",
	}
)
