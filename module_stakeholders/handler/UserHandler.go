package handler

import stakeholders "stakeholder/proto"

type UserHandler struct {
	stakeholders.UnimplementedUserServiceServer
}
