// @Author Bing
// @Date 2023/2/6 17:33:00
// @Desc
package handlers

import (
	"github.com/qingconglaixueit/abing_logger"
	"github.com/qingconglaixueit/wechatbot/model/wordfilter"
)

// 校验是否有敏感词
func IsWordFilter(str string) bool {
	res := wordfilter.Filter.FindAll(str)
	abing_logger.SugarLogger.Info("wordfilter : ", res)
	if len(res) > 0 {
		return true
	}
	return false
}
