package main

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
)

const (
	MAX_UPLOAD_SIZE = 500 * MB
	PKG_FILE_MODE   = 0666
	PKG_DIR         = "."
)

type Resp struct {
	Code    int    `json:"code"`
	Message string `message:"message"`
	//Data    interface{} `data:"data"`
}

type BuildParams struct {
	Version string   `json:"version"`
	Apps    []string `json:"apps"`
}

var (
	SUCCESS = Resp{
		Code:    0,
		Message: "success",
	}
	UPLOAD_FAILED = Resp{
		Code:    1,
		Message: "upload faild: ",
	}
	BUILD_FAILED = Resp{
		Code:    2,
		Message: "build failed: ",
	}
	PUSH_FAILED = Resp{
		Code:    3,
		Message: "push failed: ",
	}
)
