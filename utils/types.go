package utils

type Task struct {
	ID          string      `json:"id"`
	CustomID    interface{} `json:"custom_id"`
	Name        string      `json:"name"`
	TextContent string      `json:"text_content"`
	Description string      `json:"description"`
	Status      struct {
		Status     string `json:"status"`
		Color      string `json:"color"`
		Type       string `json:"type"`
		Orderindex int    `json:"orderindex"`
	} `json:"status"`
	Orderindex  string      `json:"orderindex"`
	DateCreated string      `json:"date_created"`
	DateUpdated string      `json:"date_updated"`
	DateClosed  interface{} `json:"date_closed"`
	Archived    bool        `json:"archived"`
	Creator     struct {
		ID             int         `json:"id"`
		Username       string      `json:"username"`
		Color          string      `json:"color"`
		Email          string      `json:"email"`
		ProfilePicture interface{} `json:"profilePicture"`
	} `json:"creator"`
	Assignees []struct {
		ID             int         `json:"id"`
		Username       string      `json:"username"`
		Color          string      `json:"color"`
		Initials       string      `json:"initials"`
		Email          string      `json:"email"`
		ProfilePicture interface{} `json:"profilePicture"`
	} `json:"assignees"`
	Watchers        []interface{} `json:"watchers"`
	Checklists      []interface{} `json:"checklists"`
	Tags            []interface{} `json:"tags"`
	Parent          interface{}   `json:"parent"`
	Priority        interface{}   `json:"priority"`
	DueDate         interface{}   `json:"due_date"`
	StartDate       interface{}   `json:"start_date"`
	Points          interface{}   `json:"points"`
	TimeEstimate    interface{}   `json:"time_estimate"`
	CustomFields    []interface{} `json:"custom_fields"`
	Dependencies    []interface{} `json:"dependencies"`
	LinkedTasks     []interface{} `json:"linked_tasks"`
	TeamID          string        `json:"team_id"`
	URL             string        `json:"url"`
	PermissionLevel string        `json:"permission_level"`
	List            struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Access bool   `json:"access"`
	} `json:"list"`
	Project struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
		Access bool   `json:"access"`
	} `json:"project"`
	Folder struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Hidden bool   `json:"hidden"`
		Access bool   `json:"access"`
	} `json:"folder"`
	Space struct {
		ID string `json:"id"`
	} `json:"space"`
}

type ClickupResp struct {
	Tasks []Task `json:"tasks"`
}
