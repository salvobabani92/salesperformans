package models

import "time"

// Kullanıcı girişi geriş dönüş bilgileri
// swagger:model
type LoginResp struct {
	// in:body
	// Diğer işlemlerde kullanılacak olan bilet
	TokenVal string `json:"token_val"`
	// ilgili bilet ne zaman'a kadar geçerli.
	Expire   time.Time `json:"expire"`
}

// Kullanıcı girişi başarılı
// swagger:response
type LoginResponce struct {
	// in:body
	Body LoginResp
}



// swagger:parameters Kullanıcı-Girişi
type LoginInput struct {
	// Kullanıcı girişi için Form içinde aşağıdaki değerler gönderilmelidir.
	// in:body
	Body struct {
			 // Kullanıcı Adı
			 UserName string `json:"user_name"`
			 // Şifre
			 Password string `json:"password"`
		 }
}
