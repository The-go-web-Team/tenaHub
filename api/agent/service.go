package agent

import "github.com/TenaHub/api/entity"

type AgentService interface {
	Agent(id uint) (*entity.Agent, []error)
	Agents() ([]entity.Agent, []error)
	UpdateAgent(user *entity.Agent) (*entity.Agent, []error)
	StoreAgent(user *entity.Agent) (*entity.Agent, []error)
	DeleteAgent(id uint) (*entity.Agent, []error)
}