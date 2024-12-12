package intercom

import "encoding/json"

type Team struct {
	ID       json.Number `json:"id"`
	Type     string      `json:"type"`
	Name     string      `json:"name"`
	AdminIds []int       `json:"admin_ids"`
}

type TeamList struct {
	Teams []Team `json:"teams"`
}

type TeamService struct {
	Repository TeamRepository
}

func (c *TeamService) List() (TeamList, error) {
	return c.Repository.list()
}

func (c *TeamService) Find(id json.Number) (Team, error) {
	return c.Repository.findByID(id)
}
