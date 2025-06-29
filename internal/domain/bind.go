package domain

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gotomicro/ego/core/elog"
	"loverrecipe/internal/utils"
)

func BindJson(ginCtx *gin.Context, params any) error {
	err := ginCtx.Bind(params)
	if err != nil {
		elog.Error("数据格式错误!", elog.FieldErr(err))
		if errors, ok := err.(validator.ValidationErrors); ok {
			trans, err := utils.GetTrans("zh")
			if err != nil {
				elog.Error("获取翻译失败", elog.FieldErr(err))
				return fmt.Errorf("获取翻译失败")
			}
			return buildValidateError(errors.Translate(trans))
		}

		return fmt.Errorf("数据格式错误")
	}
	return nil
}

func buildValidateError(errInfos map[string]string) error {
	var res string
	for _, v := range errInfos {
		res += fmt.Sprintf("%s;", v)
	}
	return fmt.Errorf(res)
}
