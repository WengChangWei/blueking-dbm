package middleware

import (
	"k8s-dbs/common/api"
	"k8s-dbs/common/constant"
	"k8s-dbs/common/util"
	apierrors "k8s-dbs/errors"
	metaentity "k8s-dbs/metadata/entity"
	models "k8s-dbs/metadata/model"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func APIAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求路径
		path := c.FullPath()
		method := c.Request.Method
		// 获取当前Api名称
		apiName := constant.GetAPIURL(path)

		if (apiName == "") || (apiName != "" && method == "GET") {
			return
		}

		reqBody := ParseReqBody(c)
		// 没有请求体直接返回无权限
		if len(reqBody) == 0 {
			api.ErrorResponse(c, apierrors.NewK8sDbsError(apierrors.NotPermissionError, nil))
			c.Abort()
			return
		}

		requestMap, err := util.JSONStrToMap(string(reqBody))
		if err != nil {
			slog.Warn("failed to parse request body to map", "error", err)
			api.ErrorResponse(c, apierrors.NewK8sDbsError(apierrors.NotPermissionError, err))
			c.Abort()
			return
		}

		userName := requestMap["bk_username"]
		// 没有用户名返回无权限
		if userName == nil {
			api.ErrorResponse(c, apierrors.NewK8sDbsError(apierrors.NotPermissionError, nil))
			c.Abort()
			return
		}

		var model models.AuthUserRoleModel
		params := metaentity.AuthUserRoleQueryParams{
			UserID: userName.(string),
			RoleID: "bkdata.superuser",
		}

		err = db.
			Where(params).
			First(&model).Error
		// 记录未找到，返回无权限
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.ErrorResponse(c, apierrors.NewK8sDbsError(apierrors.NotPermissionError, nil))
			c.Abort()
			return
		}

		if err != nil {
			api.ErrorResponse(c, apierrors.NewK8sDbsError(apierrors.ServerError, err))
			c.Abort()
			return
		}

	}
}
