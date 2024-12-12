package intercom

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/deepbluedot/intercom-go/interfaces"
)

type TeamRepository interface {
	list() (TeamList, error)
	findByID(id json.Number) (Team, error)
}

type TeamAPI struct {
	httpClient interfaces.HTTPClient
}

func (api TeamAPI) list() (TeamList, error) {
	teamList := TeamList{}
	data, err := api.httpClient.Get("/teams", nil)
	if err != nil {
		return teamList, err
	}
	err = json.Unmarshal(data, &teamList)
	return teamList, err
}

func (api TeamAPI) findByID(id json.Number) (Team, error) {
	team := Team{}
	strid, _ := id.Int64()
	data, err := api.httpClient.Get(fmt.Sprintf("/teams/%s", strconv.FormatInt(strid, 10)), nil)
	if err != nil {
		return team, err
	}
	err = json.Unmarshal(data, &team)
	return team, err
}
