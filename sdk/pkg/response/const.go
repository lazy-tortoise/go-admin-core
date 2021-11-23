/**
 * @Author: chenlei@sqqmall.com
 * @Description:
 * @Date: 23 2021/11/23
 */
package response

var (
	ResCodeOk      = "00000"
	ResClientError = "A0001" // 用户端错误 一级宏观错误码
	ResParamError  = "A0002" // 参数错误 二级宏观错误码
	ResTokenError  = "A0003" // Token error或者未携带token
	ResTokenExpire = "A0004" // token已过期
	ResForbid      = "A0005" // token已过期
	ResSystemError = "B0001" // 系统执行出错 一级宏观错误码
	ResThirdError  = "C0001" // 调用第三方服务出错 一级宏观错误码
)
