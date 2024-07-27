package models

import "time"

// Terminal представляет торговую точку системы
type Terminal struct {
	ID                 int       `json:"id"`
	INN                string    `json:"inn"`
	CompanyName        string    `json:"company_name"`
	Address            string    `json:"address"`
	CashRegisterNumber string    `json:"cash_register_number"`
	ModuleNumber       string    `json:"module_number"`
	AssemblyNumber     string    `json:"assembly_number"`
	LastRequestDate    time.Time `json:"last_request_date"`
	DatabaseUpdateDate time.Time `json:"database_update_date"`
	Status             string    `json:"status"`
	PartnerID          int       `json:"partner_id"`
	FreeRecordBalance  int       `json:"free_record_balance"`
}

// TerminalCreateRequest представляет данные для создания торговой точки
type TerminalCreateRequest struct {
	INN                string    `json:"inn"`
	CompanyName        string    `json:"company_name"`
	Address            string    `json:"address"`
	CashRegisterNumber string    `json:"cash_register_number"`
	ModuleNumber       string    `json:"module_number"`
	AssemblyNumber     string    `json:"assembly_number"`
	LastRequestDate    time.Time `json:"last_request_date"`
	DatabaseUpdateDate time.Time `json:"database_update_date"`
	Status             string    `json:"status"`
	UserID             int       `json:"user_id"`
}

// TerminalUpdateRequest представляет данные для обновления торговой точки
type TerminalUpdateRequest struct {
	INN                string    `json:"inn,omitempty"`
	CompanyName        string    `json:"company_name,omitempty"`
	Address            string    `json:"address,omitempty"`
	CashRegisterNumber string    `json:"cash_register_number,omitempty"`
	ModuleNumber       string    `json:"module_number,omitempty"`
	AssemblyNumber     string    `json:"assembly_number,omitempty"`
	LastRequestDate    time.Time `json:"last_request_date,omitempty"`
	DatabaseUpdateDate time.Time `json:"database_update_date,omitempty"`
	Status             string    `json:"status,omitempty"`
	UserID             int       `json:"user_id"`
	FreeRecordBalance  int       `json:"free_record_balance"`
}

// TerminalResponse представляет данные торговой точки для ответа
type TerminalResponse struct {
	ID                 int       `json:"id"`
	INN                string    `json:"inn"`
	CompanyName        string    `json:"company_name"`
	Address            string    `json:"address"`
	CashRegisterNumber string    `json:"cash_register_number"`
	ModuleNumber       string    `json:"module_number"`
	AssemblyNumber     string    `json:"assembly_number"`
	LastRequestDate    time.Time `json:"last_request_date"`
	DatabaseUpdateDate time.Time `json:"database_update_date"`
	Status             string    `json:"status"`
	UserID             int       `json:"user_id"`
	FreeRecordBalance  int       `json:"free_record_balance"`
}
