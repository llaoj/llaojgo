package e

var CodeMsg = map[int]string {
    SUCCESS : "成功",
    ERROR : "失败",
    INVALID_PARAMS : "请求参数错误",

	RESOURCE_EXIST : "资源已存在",
    RESOURCE_NOT_EXIST : "请求的资源不存在",

    AUTH_FAIL : "授权失败",
    AUTH_TOKEN_INVALID : "Token无效",
}

func Msg(code int) string {
    if msg, ok := CodeMsg[code]; ok {
        return msg
    }

    return CodeMsg[ERROR]
}