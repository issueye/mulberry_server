package service

import (
	"mulberry/app/downstream/model"
	"mulberry/app/downstream/requests"
	"mulberry/global"
)

type CertService struct{}

func NewCert() *CertService {
	return &CertService{}
}

func (s *CertService) GetByField(field string, value interface{}) (model.CertInfo, error) {
	var cert model.CertInfo
	err := global.DB.Where(field+" = ?", value).First(&cert).Error
	return cert, err
}

func (s *CertService) GetDatas(query *requests.QueryCertDatas) ([]model.CertInfo, error) {
	var certs []model.CertInfo
	err := global.DB.Where(query).Find(&certs).Error
	return certs, err
}
