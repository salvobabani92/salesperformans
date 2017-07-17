package models


// swagger:model
type QuickRegister struct {
	// Firma Adı
	// required: true
	Company_Name          string `json:"company_name"`
	// Vergi Numarası
	// required: true
	Company_VatNumber     string `json:"company_vat_number"`
	// Vergi Dairesi
	// required: true
	Company_Taxarea       string `json:"company_taxarea"`
	// Firma Adresi
	Company_Address       string `json:"company_address"`
	// Firma Şehir
	Company_City          string `json:"company_city"`
	// Firma İlgili Kişisi
	// required: true
	Company_ContactPerson string `json:"company_contactperson"`
	// Firma İlgili Kişi Görevi
	Contact_Function      string `json:"contact_function"`
	// Firma İlgili Kişi Email Adresi
	// required: true
	Company_ContactEmail  string `json:"company_contact_email"`
	// Firma İlgili Kişi Telefonu
	Company_ContactPhone  string `json:"company_contact_phone"`
	// Firma Posta Kodu
	Company_PostCode      string `json:"company_post_code"`
	// Kullanıcı Email Adresi
	// required: true
	Email                 string `json:"email"`
	// Kullanıcının İsmi
	// required: true
	Name                  string `json:"name"`
	// Kullanıcının Soyadı
	// required: true
	LastName              string `json:"last_name"`
	// Kullanıcının Şifresi
	// required: true
	Password              string `json:"password"`
	// Kullanıcı Dil Kodu
	LanguageCode          string `json:"language_code"`
}

// swagger:parameters register
type QuickRegisterInput struct {
	// Kullanıcıya ait bilgiler
	// in:body
	QuickRegister QuickRegister
}


