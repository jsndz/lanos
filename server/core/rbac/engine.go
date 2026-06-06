package rbac

import (
	"os"

	"go.yaml.in/yaml/v3"
)

type PermissionIndex map[string]map[string]map[string]struct{}

type Roles struct {
	Name  string  `yaml:"name"`
	Rules []Rules `yaml:"rules"`
}

type Rules struct {
	Resource string   `yaml:"resource"`
	Action   []string `yaml:"action"`
}

type RolesConfig struct {
	Roles           []Roles `yaml:"roles"`
	ResourceForRole PermissionIndex
}

func NewRolesConfig(yamlfile string) *RolesConfig {
	data, err := os.ReadFile(yamlfile)
	if err != nil {
		return nil
	}
	var config RolesConfig
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil
	}
	config.ResourceForRole = make(PermissionIndex)
	for _, role := range config.Roles {
		resources := make(map[string]map[string]struct{})

		for _, rule := range role.Rules {
			actions := make(map[string]struct{})
			for _, action := range rule.Action {
				// set for go
				actions[action] = struct{}{}
			}
			resources[rule.Resource] = actions
		}
		config.ResourceForRole[role.Name] = resources
	}
	return &config
}
func (c *RolesConfig) IsAllowed(role string, resource string, action string) bool {
	resources, ok := c.ResourceForRole[role]
	if !ok {
		return false
	}
	actions, ok := resources[resource]
	if !ok {
		return false
	}
	_, ok = actions[action]
	return ok
}
