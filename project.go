package freeagent

type ProjectStatus string

const (
	ProjectStatusActive    ProjectStatus = "active"
	ProjectStatusCompleted ProjectStatus = "completed"
	ProjectStatusCancelled ProjectStatus = "cancelled"
	ProjectStatusHidden    ProjectStatus = "hidden"
)

type ProjectBillingPeriod string

const (
	ProjectBillingPeriodHour ProjectBillingPeriod = "hour"
	ProjectBillingPeriodDay  ProjectBillingPeriod = "day"
)

type ProjectBudgetUnit string

const (
	ProjectBudgetUnitHours    ProjectBudgetUnit = "Hours"
	ProjectBudgetUnitDays     ProjectBudgetUnit = "Days"
	ProjectBudgetUnitMonetary ProjectBudgetUnit = "Monetary"
)

type Project struct {
	URL                                string               `json:"url,omitempty"`
	Contact                            string               `json:"contact"`
	Name                               string               `json:"name"`
	Status                             ProjectStatus        `json:"status"`
	ContractPOPreference               string               `json:"contract_po_preference"`
	UsesProjectInvoiceSequence         bool                 `json:"uses_project_invoice_sequence"`
	Currency                           string               `json:"currency"`
	Budget                             string               `json:"budget"`
	BudgetUnits                        ProjectBudgetUnit    `json:"budget_units"`
	HoursPerDay                        string               `json:"hours_per_day"`
	NormalBillingRate                  string               `json:"normal_billing_rate"`
	BillingPeriod                      ProjectBillingPeriod `json:"billing_period"`
	IsIR35                             bool                 `json:"is_ir35"`
	StartsOn                           string               `json:"starts_on"`
	EndsOn                             string               `json:"ends_on"`
	IncludeUnbilledTimeInProfitability bool                 `json:"include_unbilled_time_in_profitability"`
	CreatedAt                          string               `json:"created_at"`
	UpdatedAt                          string               `json:"updated_at"`
}

type projectDTO struct {
	Project *Project `json:"project"`
}

func (c *Client) PostProject(project *Project) (*Project, error) {
	request := &projectDTO{project}
	response := &projectDTO{}
	err := c.post("/projects", request, response)
	if err != nil {
		return nil, err
	}

	return response.Project, nil
}

func (c *Client) GetProject(id string) (*Project, error) {
	result := &projectDTO{}
	err := c.get("/projects/"+id, result)
	if err != nil {
		return nil, err
	}

	return result.Project, nil
}

type ProjectView string

const (
	ProjectViewActive    ProjectView = "active"
	ProjectViewCompleted ProjectView = "completed"
	ProjectViewCancelled ProjectView = "cancelled"
	ProjectViewHidden    ProjectView = "hidden"
)

type ProjectSort string

const (
	ProjectSortName               ProjectSort = "name"
	ProjectSortContactName        ProjectSort = "contact_name"
	ProjectSortContactDisplayName ProjectSort = "contact_display_name"
	ProjectSortCreatedAt          ProjectSort = "created_at"
	ProjectSortUpdatedAt          ProjectSort = "updated_at"
)

type ProjectQuery struct {
	UpdatedSince string
	View         ProjectView
	Sort         ProjectSort
	Descending   bool
}

type projectsDTO struct {
	Projects []*Project `json:"projects"`
}

func (c *Client) GetProjects(q *ProjectQuery) ([]*Project, error) {
	result := &projectsDTO{}
	err := c.get("/projects", result)
	if err != nil {
		return nil, err
	}

	return result.Projects, nil
}
