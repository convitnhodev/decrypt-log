package decryptBiz

import "fmt"

type ResolveDecrypt interface {
	GetSource() string
	SetDes(input string)
	ParseLog() [][]int
	DecryptLogSnippet(commonIVsecret string, keyTextsecret string) string
}

type bizResolveDecrypt struct {
	repo ResolveDecrypt
}

func NewBizResolveDecrypt(repo ResolveDecrypt) *bizResolveDecrypt {
	return &bizResolveDecrypt{repo}
}

func (biz *bizResolveDecrypt) DecryptLogString(commonIVSecret string, keyTextSecret string) {
	cipherList := biz.repo.ParseLog()

	fmt.Println(cipherList)

	sizeChanging := 0

	input := biz.repo.GetSource()

	for _, item := range cipherList {

		cipher := input[item[0]-sizeChanging : item[1]-sizeChanging]

		biz.repo.SetDes(cipher[3 : len(cipher)-4])

		plainText := biz.repo.DecryptLogSnippet(commonIVSecret, keyTextSecret)

		input = input[:item[0]-sizeChanging] + plainText + input[item[1]-sizeChanging:]

		sizeChanging = sizeChanging + ((item[1] - item[0]) - len(plainText))
	}

	biz.repo.SetDes(input)
}
