package jira

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/andygrunwald/go-jira"
)

func JiraTest() {

	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace("dima.shulhin@samba.tv"),
		Password: strings.TrimSpace("11XwQ4chaIvZTsTBfhaH27C9"),
	}

	client, err := jira.NewClient(tp.Client(), strings.TrimSpace("https://sambatv.atlassian.net"))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	GetAssignedIssues(client)

}

func GetAssignedIssues(client *jira.Client) {
	req, err := client.NewRequest("GET", "/rest/api/3/search?jql=assignee=currentuser()", nil)
	if err != nil {
		return
	}

	var buf interface{}
	//var board Board
	client.Do(req, &buf)
	bolB, _ := json.Marshal(buf)
	// //fmt.Println(bolB)
	// err = json.Unmarshal(bolB, &board)
	m := map[string]interface{}{}

	err = json.Unmarshal(bolB, &m)
	if err != nil {
		panic(err)
	}
	issues := m["issues"].([]interface{})
//	myIssues := make([]Issue, len(issues))
//	for i := range issues {
//	issue := issues[0].(Issue)
		fmt.Println(issues[0])
//	}
	
}

type Board struct {
	Expand     string `json:"expand"`
	StartAt    int    `json:"startAt"`
	MaxResults int    `json:"maxResults"`
	Total      int    `json:"total"`
	Issues     []Issue
}

//Issue struct

type Issue struct {
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		Statuscategorychangedate interface{} `json:"statuscategorychangedate"`
		Issuetype                struct {
			Self        string `json:"self"`
			ID          string `json:"id"`
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"issuetype"`
		Project struct {
			Self           string `json:"self"`
			ID             string `json:"id"`
			Key            string `json:"key"`
			Name           string `json:"name"`
			ProjectTypeKey string `json:"projectTypeKey"`
		} `json:"project"`
		Created  string `json:"created"`
		Priority struct {
			Self string `json:"self"`
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"priority"`
		Versions []interface{} `json:"versions"`
		Assignee struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			AccountID    string `json:"accountId"`
			EmailAddress string `json:"emailAddress"`
			DisplayName  string `json:"displayName"`
			Active       bool   `json:"active"`
			TimeZone     string `json:"timeZone"`
			AccountType  string `json:"accountType"`
		} `json:"assignee"`
		Status struct {
			Self           string `json:"self"`
			Description    string `json:"description"`
			Name           string `json:"name"`
			ID             string `json:"id"`
			StatusCategory struct {
				Self      string `json:"self"`
				ID        int    `json:"id"`
				Key       string `json:"key"`
				ColorName string `json:"colorName"`
				Name      string `json:"name"`
			} `json:"statusCategory"`
		} `json:"status"`
		Description struct {
			Version int    `json:"version"`
			Type    string `json:"type"`
			Content string `json:"content"`
		} `json:"description"`
		Summary string `json:"summary"`
		Creator struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			AccountID    string `json:"accountId"`
			EmailAddress string `json:"emailAddress"`
			DisplayName  string `json:"displayName"`
			Active       bool   `json:"active"`
			TimeZone     string `json:"timeZone"`
			AccountType  string `json:"accountType"`
		} `json:"creator"`
		Reporter struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			AccountID    string `json:"accountId"`
			EmailAddress string `json:"emailAddress"`
			DisplayName  string `json:"displayName"`
			Active       bool   `json:"active"`
			TimeZone     string `json:"timeZone"`
			AccountType  string `json:"accountType"`
		} `json:"reporter"`
		Aggregateprogress struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
			Percent  int `json:"percent"`
		} `json:"aggregateprogress"`
		Environment interface{} `json:"environment"`
		Duedate     interface{} `json:"duedate"`
		Progress    struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
			Percent  int `json:"percent"`
		} `json:"progress"`
	} `json:"fields"`
}
