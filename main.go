package main


import (
	"fmt"
	"github.com/bootpay/backend-go"
	"time"
)

func main() {
	api := bootpay.Api{}.New("5b8f6a4d396fa665fdc2b5ea", "rm6EYECr6aroQVG2ntW0A6LpWnkTgP4uQ3H18sDDUYw=", nil, "")
	GetToken(api)
	GetVerify(api)
	ReceiptCancel(api)
	GetBillingKey(api)
	RequestSubscribe(api)
	RequestSubscribeReserve(api)
	RequestSubscribeReserveDelete(api)
	DeleteBillingKey(api)
	GetEasyUserToken(api)
	RequestLink(api)
	ServerSubmit(api)
	Certificate(api)
}

// 1. 토큰 발급
func GetToken(api *bootpay.Api) {
	fmt.Println("--------------- GetToken() Start ---------------")
	token, err := api.GetToken()
	fmt.Println("token : " + token.Data.Token)
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- GetToken() Start ---------------")
}

// 2. 결제 검증
func GetBillingKey(api *bootpay.Api) {
	fmt.Println("--------------- GetBillingKey() Start ---------------")
	payload := bootpay.BillingKeyPayload{
		OrderId: fmt.Sprintf("%+8d", (time.Now().UnixNano() / int64(time.Millisecond))),
		Pg: "nicepay",
		ItemName: "정기결제 테스트 아이템",
		CardNo: "5570********1074",
		CardPw: "**",
		ExpireYear: "**",
		ExpireMonth: "**",
		IdentifyNumber: "",
	}
	billingKey, err := api.GetBillingKey(payload)

	fmt.Println(billingKey)
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- GetBillingKey() End ---------------")
}

// 3. 결제 취소 (전액 취소 / 부분 취소)
func GetVerify(api *bootpay.Api) {
	receiptId := "610c96352386840036db8bef"
	fmt.Println("--------------- GetVerify() Start ---------------")
	verify, err := api.Verify(receiptId)

	fmt.Println(verify)
	fmt.Println(verify.Data.PaymentData)
	fmt.Println(verify.Data.PaymentData["o_id"])
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- GetVerify() End ---------------")
}

// 4. 빌링키 발급
func ReceiptCancel(api *bootpay.Api) {
	receiptId := "610cc0cb7b5ba40044b04530"
	name := "관리자"
	reason := "테스트 결제 취소를 테스트"
	fmt.Println("--------------- ReceiptCancel() Start ---------------")
	cancel, err := api.ReceiptCancel(receiptId, 0, name, reason, bootpay.RefundData{})

	fmt.Println(cancel)
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- ReceiptCancel() End ---------------")
}

// 4-1. 발급된 빌링키로 결제 승인 요청
func RequestSubscribe(api *bootpay.Api) {
	fmt.Println("--------------- RequestSubscribe() Start ---------------")
	payload := bootpay.SubscribePayload{
		BillingKey: "6100e8c80d681b001dd4e0d7",
		ItemName: "테스트아이템",
		Price: 1000,
		OrderId: fmt.Sprintf("%+8d", (time.Now().UnixNano() / int64(time.Millisecond))),
	}
	res, err := api.RequestSubscribe(payload)
	fmt.Println(res)
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- RequestSubscribe() End ---------------")
}

// 4-2. 발급된 빌링키로 결제 예약 요청
func RequestSubscribeReserve(api *bootpay.Api) {
	fmt.Println("--------------- RequestSubscribeReserve() Start ---------------")
	payload := bootpay.SubscribePayload{
		BillingKey: "6100e8c80d681b001dd4e0d7",
		ItemName: "테스트아이템",
		Price: 1000,
		OrderId: fmt.Sprintf("%+8d", (time.Now().UnixNano() / int64(time.Millisecond))),
		ExecuteAt: time.Now().UnixNano() / int64(time.Millisecond) / 1000 + 10,
	}
	res, err := api.ReserveSubscribe(payload)
	fmt.Println(res)
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- RequestSubscribeReserve() End ---------------")
}

// 4-2-1. 발급된 빌링키로 결제 예약 - 취소 요청
func RequestSubscribeReserveDelete(api *bootpay.Api) {
	fmt.Println("--------------- RequestSubscribeReserveDelete() Start ---------------")
	res, err := api.ReserveCancelSubscribe("6100e892019943002150fef3")
	fmt.Println(res)
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- RequestSubscribeReserveDelete() End ---------------")
}

// 4-3. 빌링키 삭제
func DeleteBillingKey(api *bootpay.Api) {
	fmt.Println("--------------- DeleteBillingKey() Start ---------------")

	res, err := api.DestroyBillingKey("6100e7ea0d681b001fd4de69")
	fmt.Println(res)
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- DeleteBillingKey() End ---------------")
}

// 5. 사용자 토큰 발급
func GetEasyUserToken(api *bootpay.Api) {
	fmt.Println("--------------- GetEasyUserToken() Start ---------------")
	userToken := bootpay.EasyUserTokenPayload{
		UserId: "1234",
	}

	res, err := api.GetUserToken(userToken)
	fmt.Println(res)
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- GetEasyUserToken() End ---------------")
}

// 6. 결제 링크 생성
func RequestLink(api *bootpay.Api) {
	payload := bootpay.Payload{
		Pg: "kcp",
		Method: "card",
		Price: 1000,
		OrderId: fmt.Sprintf("%+8d", (time.Now().UnixNano() / int64(time.Millisecond))),
		Name: "테스트 결제 상품",
	}
	fmt.Println("--------------- RequestLink() End ---------------")
	res, err := api.RequestLink(payload)

	fmt.Println(res)
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- RequestLink() End ---------------")
}

// 7. 서버 승인 요청
func ServerSubmit(api *bootpay.Api) {
	receiptId := "610cc01b238684002adb904e"
	fmt.Println("--------------- ServerSubmit() Start ---------------")
	res, err := api.ServerSubmit(receiptId)

	fmt.Println(res)
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- ServerSubmit() End ---------------")
}

// 8. 본인 인증 결과 조회
func Certificate(api *bootpay.Api) {
	receiptId := "610cc01b238684002adb904e"
	fmt.Println("--------------- Certificate() Start ---------------")
	res, err := api.Certificate(receiptId)

	fmt.Println(res)
	if err != nil {
		fmt.Println("get token error: " + err.Error())
	}
	fmt.Println("--------------- Certificate() End ---------------")
}