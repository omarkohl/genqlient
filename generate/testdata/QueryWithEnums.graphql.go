package test

// Code generated by github.com/Khan/genql, DO NOT EDIT.

import (
	"context"

	"github.com/Khan/genql/graphql"
)

type QueryWithEnumsResponse struct {
	User *QueryWithEnumsUser `json:"user"`
}

type QueryWithEnumsUser struct {
	Roles []QueryWithEnumsUserRolesRole `json:"roles"`
}

type QueryWithEnumsUserRolesRole string

const (
	QueryWithEnumsUserRolesRoleStudent QueryWithEnumsUserRolesRole = "STUDENT"
	QueryWithEnumsUserRolesRoleTeacher QueryWithEnumsUserRolesRole = "TEACHER"
)

func QueryWithEnums(client *graphql.Client) (*QueryWithEnumsResponse, error) {
	var retval QueryWithEnumsResponse
	err := client.MakeRequest(context.Background(), `
query QueryWithEnums {
	user {
		roles
	}
}
`, &retval, nil)
	return &retval, err
}