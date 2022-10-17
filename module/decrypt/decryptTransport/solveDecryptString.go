package decryptTransport

import (
	"code_hash/common"
	"code_hash/component"
	"code_hash/module/decrypt/decryptBiz"
	"code_hash/module/decrypt/decryptRepository"
	"code_hash/module/modelManager"
	"github.com/gin-gonic/gin"
)

func DecryptString(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data modelManager.ManagerCrud

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		repo := decryptRepository.PassSource(data.DataString)
		biz := decryptBiz.NewBizResolveDecrypt(repo)
		biz.DecryptLogString(appCtx.GetCommonIV(), appCtx.GetKeyText())

		c.JSON(200, common.SimpleSuccessReponse(repo.Des))
	}
}
