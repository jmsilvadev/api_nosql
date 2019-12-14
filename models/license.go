package models

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

// License - JSON strucuture to manipultes License
type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LicenseResponse - JSON strucuture to manipultes LicenseResponse
type LicenseResponse struct {
	*License
	SourceIP string `json:"request_ip"`
}

var Licenseinfo = []*License{
	{Name: "JMSilvaDev © 2019 | Todos os direitos reservados"},
}

func dbGetLicense() (*License, error) {
	for _, a := range Licenseinfo {
		return a, nil
	}
	return nil, errors.New("License information not found")
}

// NewLicenseListResponse - the Response converted to ArrayList type LicenseResponse struct
func NewLicenseListResponse(licenses []*License) []render.Renderer {
	list := []render.Renderer{}
	for _, license := range licenses {
		list = append(list, NewLicenseResponse(license))
	}
	return list
}

// NewLicenseResponse - the Response converted to type LicenseResponse struct
func NewLicenseResponse(license *License) *LicenseResponse {
	resp := &LicenseResponse{License: license}
	return resp
}

// Render - pattern for managing payload encoding and decoding in License Response Struct
func (rd *LicenseResponse) Render(w http.ResponseWriter, r *http.Request) error {
	rd.SourceIP = r.RemoteAddr
	return nil
}
