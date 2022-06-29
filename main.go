package main

import (
	"fmt"
	bootpay "github.com/bootpay/backend-go/v2"
	"time"
)

func main() {
	bootpay := bootpay.Api{}.New("5b8f6a4d396fa665fdc2b5ea", "rm6EYECr6aroQVG2ntW0A6LpWnkTgP4uQ3H18sDDUYw=", nil, "")
	fmt.Println("--------------- Certificate() End ---------------")
	GetToken(bootpay)
	ReceiptCancel(bootpay)
	GetReceipt(bootpay)
	GetBillingKey(bootpay)
	RequestSubscribe(bootpay)
	LookupBillingKey(bootpay)
	ReserveSubscribe(bootpay)
	ReserveCancel(bootpay)
	DestroyBillingKey(bootpay)
	GetUserToken(bootpay)
	ServerConfirm(bootpay)
	Certificate(bootpay)
	ShoppingStart(bootpay)
}

func GetToken(api *bootpay.Api) {
	fmt.Println("--------------- GetToken() Start ---------------")
	token, err := api.GetToken()
	fmt.Println(token)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println("--------------- GetToken() End ---------------")
}

func GetBillingKey(api *bootpay.Api) {
	fmt.Println("--------------- GetBillingKey() Start ---------------")
	payload := bootpay.BillingKeyPayload{
		SubscriptionId:  fmt.Sprintf("%+8d", (time.Now().UnixNano() / int64(time.Millisecond))),
		Pg:              "nicepay",
		OrderName:       "정기결제 테스트 아이템",
		CardNo:          "5570********1074",
		CardPw:          "**",
		CardExpireYear:  "**",
		CardExpireMonth: "**",
		CardIdentityNo:  "",
	}
	billingKey, err := api.GetBillingKey(payload)

	fmt.Println(billingKey)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println("--------------- GetBillingKey() End ---------------")
}

func GetReceipt(api *bootpay.Api) {
	receiptId := "62afc194e38c300021b345d4"
	fmt.Println("--------------- getReceipt() Start ---------------")
	verify, err := api.GetReceipt(receiptId)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	fmt.Println(verify)
	fmt.Println("--------------- GetVerify() End ---------------")
}

//lookupBillingKey

func LookupBillingKey(api *bootpay.Api) {
	receiptId := "62afccb3cf9f6d001b7d101d"
	fmt.Println("--------------- LookupBillingKey() Start ---------------")
	verify, err := api.LookupBillingKey(receiptId)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	fmt.Println(verify)
	fmt.Println("--------------- LookupBillingKey() End ---------------")
}

func ReceiptCancel(api *bootpay.Api) {
	payload := bootpay.CancelData{
		ReceiptId:      "62afc3c5cf9f6d001b7d101a",
		CancelId:       fmt.Sprintf("%+8d", (time.Now().UnixNano() / int64(time.Millisecond))),
		CancelUsername: "관리자",
		CancelMessage:  "테스트 결제 취소를 테스트",
	}
	//receiptId := "610cc0cb7b5ba40044b04530"
	//name := "관리자"
	//reason := "테스트 결제 취소를 테스트"
	fmt.Println("--------------- ReceiptCancel() Start ---------------")
	cancel, err := api.ReceiptCancel(payload)

	fmt.Println(cancel)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println("--------------- ReceiptCancel() End ---------------")
}

func RequestSubscribe(api *bootpay.Api) {
	payload := bootpay.SubscribePayload{
		BillingKey: "62afc52dcf9f6d001d7d1035",
		OrderName:  "정기결제 테스트",
		OrderId:    fmt.Sprintf("%+8d", (time.Now().UnixNano() / int64(time.Millisecond))),
		Price:      1000,
	}

	fmt.Println("--------------- requestSubscribe() Start ---------------")
	cancel, err := api.RequestSubscribe(payload)

	fmt.Println(cancel)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println("--------------- requestSubscribe() End ---------------")
}

func ReserveSubscribe(api *bootpay.Api) {
	s10 := time.Now().Add(time.Second * 100).Format("2006-01-02T15:04:05-07:00")
	payload := bootpay.SubscribePayload{
		BillingKey:       "62aff193cf9f6d001a7d10be",
		OrderName:        "정기결제 테스트",
		OrderId:          fmt.Sprintf("%+8d", (time.Now().UnixNano() / int64(time.Millisecond))),
		ReserveExecuteAt: s10,
		Price:            1000,
	}

	fmt.Println("--------------- ReserveSubscribe() Start ---------------")
	cancel, err := api.ReserveSubscribe(payload)

	fmt.Println(cancel)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println("--------------- ReserveSubscribe() End ---------------")
}

func ReserveCancel(api *bootpay.Api) {
	reserveId := "62aff2a0cf9f6d001a7d10c4"
	fmt.Println("--------------- ReserveCancel() Start ---------------")
	cancel, err := api.ReserveCancelSubscribe(reserveId)

	fmt.Println(cancel)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println("--------------- ReserveCancel() End ---------------")
}

func ServerConfirm(api *bootpay.Api) {
	receiptId := "62afda41cf9f6d001f7d105f"
	fmt.Println("--------------- ServerConfirm() Start ---------------")
	res, err := api.ServerConfirm(receiptId)

	fmt.Println(res)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println("--------------- ServerConfirm() End ---------------")
}

func GetUserToken(api *bootpay.Api) {
	payload := bootpay.EasyUserTokenPayload{
		UserId: "user_1234",
		Email:  "test1234@gmail.com",
		Name:   "홍길동",
		Gender: 0,
		Birth:  "19861014",
		Phone:  "01012345678",
	}

	fmt.Println("--------------- GetUserToken() Start ---------------")
	cancel, err := api.GetUserToken(payload)

	fmt.Println(cancel)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println("--------------- GetUserToken() End ---------------")
}

func DestroyBillingKey(api *bootpay.Api) {
	billingKey := "62afc52dcf9f6d001d7d1035"
	fmt.Println("--------------- DestroyBillingKey() Start ---------------")
	res, err := api.DestroyBillingKey(billingKey)

	fmt.Println(res)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println("--------------- DestroyBillingKey() End ---------------")
}

func Certificate(api *bootpay.Api) {
	receiptId := "6285ffa6cf9f6d0022c4346b"
	fmt.Println("--------------- Certificate() Start ---------------")
	res, err := api.Certificate(receiptId)

	fmt.Println(res)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println("--------------- Certificate() End ---------------")
}

func ShoppingStart(api *bootpay.Api) {
	shipping := bootpay.Shipping{
		ReceiptId:      "628ae7ffd01c7e001e9b6066",
		TrackingNumber: "123456",
		DeliveryCorp:   "CJ대한통운",
		User: bootpay.ShippingUser{
			Username: "홍길동",
			Phone:    "01000000000",
			Address:  "서울특별시 종로구",
			Zipcode:  "08490",
		},
	}

	fmt.Println("--------------- ShoppingStart() Start ---------------")
	res, err := api.PutShippingStart(shipping)

	fmt.Println(res)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println("--------------- ShoppingStart() End ---------------")
}
