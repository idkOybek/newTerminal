package models

// FiscalModule представляет фискальный модуль системы
type FiscalModule struct {
	ID            int    `json:"id"`
	FactoryNumber string `json:"factory_number"`
	FiscalNumber  string `json:"fiscal_number"`
	UserID        int    `json:"user_id"`
}

// FiscalModuleCreateRequest представляет данные для создания фискального модуля
type FiscalModuleCreateRequest struct {
	FactoryNumber string `json:"factory_number"`
	FiscalNumber  string `json:"fiscal_number"`
	UserID        int    `json:"user_id"`
}

// FiscalModuleUpdateRequest представляет данные для обновления фискального модуля
type FiscalModuleUpdateRequest struct {
	FactoryNumber string `json:"factory_number,omitempty"`
	FiscalNumber  string `json:"fiscal_number,omitempty"`
	UserID        int    `json:"user_id,omitempty"`
}

// FiscalModuleResponse представляет данные фискального модуля для ответа
type FiscalModuleResponse struct {
	ID            int    `json:"id"`
	FactoryNumber string `json:"factory_number"`
	FiscalNumber  string `json:"fiscal_number"`
	UserID        int    `json:"user_id"`
}
