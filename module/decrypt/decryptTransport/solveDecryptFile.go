package decryptTransport

import (
	"code_hash/common"
	"code_hash/component"
	"code_hash/module/decrypt/decryptBiz"
	"code_hash/module/decrypt/decryptRepository"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func DecryptFile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		time := time.Now()
		timeConvert := time.Format("01022006150405")
		fmt.Println(timeConvert)
		c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", timeConvert+fileHeader.Filename))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//s := decryptRepository.ConvertFileToString(string(timeConvert + fileHeader.Filename))

		repo := decryptRepository.PassSource(decryptRepository.ConvertFileToString(string(timeConvert + fileHeader.Filename)))
		biz := decryptBiz.NewBizResolveDecrypt(repo)
		biz.DecryptLogString(appCtx.GetCommonIV(), appCtx.GetKeyText())

		c.JSON(200, common.SimpleSuccessReponse(repo.Des))
	}
}
