package pagerduty

import (
	"bytes"
	"fmt"
	"time"
)

type PagerDutyResponse struct {
	Incidents []Incident `json:"incidents"`
	Limit     int        `json:"limit"`
	Offset    int        `json:"offset"`
	Total     int        `json:"total"`
}

func (p *PagerDutyResponse) PrettyPrint() string {
	var buffer bytes.Buffer
	for _, val := range p.Incidents {
		buffer.WriteString(val.String())
		//TODO: Not windows safe
		buffer.WriteByte('\n')
	}
	return buffer.String()
}

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	HtmlUrl string `json:"html_url"`
}

func (u *User) String() string {
	return u.Name
}

type Incident struct {
	Status         string    `json:"status"`
	CreatedOn      time.Time `json:"created_on"`
	AssignedToUser User      `json:"assigned_to_user"`
}

func (i *Incident) String() string {
	return fmt.Sprintf("%s\t%s\t%s", i.CreatedOn, i.Status, i.AssignedToUser.String())
}
