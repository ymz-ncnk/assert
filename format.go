package assert

import "fmt"

const MsgSep = ": "

func Format(msg string, msgAndArgs ...any) string {
	if len(msgAndArgs) > 0 {
		if len(msgAndArgs) == 1 {
			msg += MsgSep + fmt.Sprint(msgAndArgs[0])
		} else {
			msg += MsgSep + fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
		}
	}
	return msg
}
