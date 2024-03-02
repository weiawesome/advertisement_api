/*
The service for adding advertisement.
It has sql repository to realize interaction with database.
*/

package advertisement

import (
	advertisementRequest "advertisement_api/api/request/advertisement"
	advertisementResponse "advertisement_api/api/response/advertisement"
	"advertisement_api/internal/repository/sql"
)

// ServiceAddAdvertisement is the interface of the adding advertisement service
type ServiceAddAdvertisement interface {
	Add(data advertisementRequest.AddAdvertisementRequest) (advertisementResponse.AddAdvertisementResponse, error)
}

// addService is the service to add advertisement
type addService struct {
	SqlRepository sql.Repository // the repository to interact with sql database
}

func NewAddAdvertisementService(r *sql.Repository) ServiceAddAdvertisement {
	return &addService{SqlRepository: *r}
}

// Add is to get the content from handler and add advertisement by sql repository
func (m *addService) Add(data advertisementRequest.AddAdvertisementRequest) (advertisementResponse.AddAdvertisementResponse, error) {
	// use repository to add advertisement into sql database
	err := m.SqlRepository.AddAdvertisement(data)

	// return the result from the repository process
	return advertisementResponse.AddAdvertisementResponse{}, err
}
