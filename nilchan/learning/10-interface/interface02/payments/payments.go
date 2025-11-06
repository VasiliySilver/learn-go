package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentsInfo  map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

// Метод Pay;
// Принимает:
// - описание проводимой оплаты
// - ID проведенной операции
// Возвращает:
// - ID проведенной операции
func (p PaymentModule) Pay(description string, usd int) int {
	// 1 - проводит оплату
	// 2 - получает ID проведенной оплаты
	id := p.paymentMethod.Pay(usd)
	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}

	// 3 - сохранять информацию о проведенной операции
	//   - описание операции
	//   - сколько было потрачено
	//   - отмененная ли операция
	p.paymentsInfo[id] = info

	// 4 - возвращать ID проведенной оплате
	return id
}

// Метод Cancel()
// Принимает:
// - ID операции
// Возвращает:
// - ничего
func (p PaymentModule) Cancel(id int) {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}

	p.paymentMethod.Cancel(id)
	info.Cancelled = true

	p.paymentsInfo[id] = info

}

// Метод Info()
// Принимает:
// - ID операции
// Возвращает:
// - Информацию о проведенной операции
func (p PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}
	return info

}

// Метод Info()
// Возвращает:
// - Информацию о всех проведеных операциях
func (p PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))
	for k, v := range p.paymentsInfo {
		tempMap[k] = v
	}

	return tempMap

}
