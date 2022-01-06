package smartbonus

import (
	"os"
)

// Object implement all features
type SmartBonus struct {
	Store
}

// Create new instance of SmartBonus object
// On testing set to env variable with name SB_ROUTE
// Ask smartbonus team for your sbRoute
func NewSmartBonus(storeId, sbRoute, customerId string) *SmartBonus {
	if storeId == "" {
		panic("storeId is empty")
	}
	if customerId == "" {
		panic("customerId is empty")
	}
	if sbRoute == "" {
		sbRoute = os.Getenv("SB_ROUTE")
	}
	if rootPath = sbRoute; rootPath == "" {
		panic("sbRoute is empty")
	}
	return &SmartBonus{Store: Store{StoreId: storeId, CustomerId: customerId}}
}

func (s *SmartBonus) GetClient(userId string) (*Client, error) {
	return getClient(s.StoreId, userId)
}

func (s *SmartBonus) SyncTags(tags []Tag) error {
	return syncTags(s.StoreId, tags)
}

func (s *SmartBonus) SyncNomenclatures(nomes []Product) error {
	return syncNomenclatures(s.CustomerId, nomes)
}

func (s *SmartBonus) SyncNomenclaturesV2(nomes []Nomenclature) error {
	return syncNomenclaturesV2(s.StoreId, nomes)
}

func (s *SmartBonus) DiscountReceipt(receipt ReceiptDiscount) (*ReceiptResult, error) {
	return discountReceipt(s.StoreId, receipt)
}

func (s *SmartBonus) ConfirmReceipt(receipt ReceiptConfirm) (*ReceiptResult, error) {
	return confirmReceipt(s.StoreId, receipt)
}

func (s *SmartBonus) SyncReceipts(receipts []ReceiptConfirm) error {
	return syncReceipts(s.StoreId, receipts)
}

func (s *SmartBonus) DeleteReceipts(receipts []string) error {
	return deleteReceipts(s.StoreId, receipts)
}

func (s *SmartBonus) RefundReceipt(receipt ReceiptRefund) (*RefundResult, error) {
	return refundReceipt(s.StoreId, receipt)
}

func (s *SmartBonus) ConfigOrder(orderUrl, statusUrl, token string, syncNomenclatureFull, syncNomenclatureByCustomer bool) error {
	return configOrder(s.StoreId, orderUrl, statusUrl, token, syncNomenclatureFull, syncNomenclatureByCustomer)
}

func (s *SmartBonus) ChangeOrderStatus(statusBody StatusBody) error {
	return changeOrderStatus(s.StoreId, statusBody)
}
