package advertisement

import (
	advertisementRequest "advertisement_api/api/request/advertisement"
	advertisementResponse "advertisement_api/api/response/advertisement"
	"advertisement_api/internal/repository/sql"
)

type ServiceAddAdvertisement struct {
	SqlRepository sql.Repository
}

func (m *ServiceAddAdvertisement) Add(data advertisementRequest.AddAdvertisementRequest) (advertisementResponse.AddAdvertisementResponse, error) {
	err := m.SqlRepository.AddAdvertisement(data)
	return advertisementResponse.AddAdvertisementResponse{}, err
}
