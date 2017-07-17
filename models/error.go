

package models
// Servisten dönen hata modeli.
//
// swagger:model
type StatusResponse struct {
	// Kod
	Code    string `json:"code"`
	// Mesaj
	Message string `json:"message"`
}


// Hatalı istek. Servise gönderilen parametreler hatalı.
// swagger:response
type ErrorResponse400 struct {
	//in:body
	Body StatusResponse
}


// Yetkisiz deneme. Yetkisiz giriş hatası kullanıcı girişine izin yok.
// swagger:response
type ErrorResponse401 struct {
	//in:body
	Body StatusResponse
}


// Kayıt bulunamadı. Verilen kriterlere uygun kayıt bulunamadı.
// swagger:response
type ErrorResponse404 struct {
	//in:body
	Body StatusResponse
}



// Alt servis çağrısından hata alındı.
// Proxy hatası
// swagger:response
type ErrorResponse502 struct {
	//in:body
	Body StatusResponse
}



// Tamam
// swagger:response
type StatusResponse200 struct {
	//in:body
	Body StatusResponse
}


// Kayıt Oluşturuldu
// swagger:response
type StatusResponse201 struct {
	//in:body
	Body StatusResponse
}



// GenericStatusResponse objesini oluşturur.
func GetGenericStatusResponse(StatusCode string, Message string) StatusResponse {
	sr := StatusResponse{Code:StatusCode, Message:Message}
	return sr
}

