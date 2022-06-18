package tail

import (
	"fmt"
	"strings"
	"testing"
)

func Test_textFormat(t *testing.T) {
	str := `>>>>>>>>\nPOST /cgi-bin/message/subscribe/send?access_token=49_6JUHEXMlajleu889IV_lHA7arJ6ZPCgXZTj38zBWC-0b-DyknHpgcI5Qov83Ttd6eU3b9Ys709fajWaZvSfuOOQKAh3oXyvEQnfztdE8KJallSVhp-3MefnjQ1kcVuTrESOynkdTZkjCXwWLYGBbAJAOWR HTTP/1.1\\r\\nHost: api.weixin.qq.com\\r\\nContent-Length: 298\\r\\nUser-Agent: GuzzleHttp/6.5.5 curl/7.77.0 PHP/7.2.34\\r\\nContent-Type: application/json\\r\\n\\r\\n{\\\"touser\\\":\\\"oW-Hn5C2ypoU9Ia-DHST_FICpz0o\\\",\\\"template_id\\\":\\\"MIOxFkx4xCD3z4U-x2AIT2fO1GhSikrWaTqmpFFk_-U\\\",\\\"page\\\":\\\"pages\\\\/commentMessage\\\\/commentMessage\\\",\\\"data\\\":{\\\"time3\\\":{\\\"value\\\":\\\"2021-09-08 06:49:22\\\"},\\\"thing5\\\":{\\\"value\\\":\\\"A小秦装饰--18020431762\\\"},\\\"thing2\\\":{\\\"value\\\":\\\"人家卖才五万 你要七万\\\"}}}\\n<<<<<<<<\\nHTTP/1.1 200 OK\\r\\nConnection: keep-alive\\r\\nContent-Type: application/json; encoding=utf-8\\r\\nDate: Sun, 19 Sep 2021 23:27:58 GMT\\r\\nContent-Length: 90\\r\\n\\r\\n{\\\"errcode\\\":43101,\\\"errmsg\\\":\\\"user refuse to accept the msg rid: 6147c77e-3cedb16e-3f05cd8f\\\"}\\n--------\\nNULL`
	fmt.Printf("%s\n", strings.Split(str, "\n"))
}
