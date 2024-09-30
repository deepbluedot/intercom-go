package intercom

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/deepbluedot/intercom-go/interfaces"
)

// AdminRepository defines the interface for working with Admins through the API.
type AdminRepository interface {
	list() (AdminList, error)
	findByID(id json.Number) (Admin, error)
}

// AdminAPI implements AdminRepository
type AdminAPI struct {
	httpClient interfaces.HTTPClient
}

func (api AdminAPI) list() (AdminList, error) {
	adminList := AdminList{}
	data, err := api.httpClient.Get("/admins", nil)
	if err != nil {
		return adminList, err
	}
	err = json.Unmarshal(data, &adminList)
	return adminList, err
}

func (api AdminAPI) findByID(id json.Number) (Admin, error) {
	admin := Admin{}
	strid, _ := id.Int64()
	data, err := api.httpClient.Get(fmt.Sprintf("/admins/%s", strconv.FormatInt(strid, 10)), nil)
	if err != nil {
		return admin, err
	}
	err = json.Unmarshal(data, &admin)
	return admin, err
}
