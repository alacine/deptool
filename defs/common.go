package defs

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

// 当前状态
const (
	UPLOADING = iota // 正在上传
	UPLOADED         // 已上传
	BUILDING         // 正在构建
	BUILT            // 已构建
	PUSHING          // 正在推送
	PUSHED           // 已推送
	CLEANING         // 正在清理
	CLEANED          // 已清理
)
