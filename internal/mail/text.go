package mail

import "fmt"

const verificationCodeText = `
Send message bla bla %d
`

func verificationCodeMsg(code int) string {
	return fmt.Sprintf(verificationCodeText, code)
}
