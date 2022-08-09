package proxy


type PaymentService interface {
	pay(order string) string
}

// 微信支付
type WXPay struct {
}

func (w *WXPay) pay(order string) string {
	return "从微信获取支付token"
}


/**
 * @Description: 阿里支付类
 */
 type AliPay struct {
}

func (a *AliPay) pay(order string) string {
   return "从阿里获取支付token"
}


type PaymentProxy struct {
	realPay PaymentService
}


func (p *PaymentProxy) pay(order string) string {
	fmt.Println("处理" + order)
	fmt.Println("1校验签名")
	fmt.Println("2格式化订单数据")
	fmt.Println("3参数检查")
	fmt.Println("4记录请求日志")
	token := p.realPay.pay(order)
	return "http://组装" + token + "然后跳转到第三方支付"
}
func main() {
	proxy := &PaymentProxy{
		 realPay: &AliPay{},  // 
	}
	url := proxy.pay("阿里订单")
	fmt.Println(url)
}