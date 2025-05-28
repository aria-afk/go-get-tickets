package queries

import (
	"github.com/aria-afk/go-get-tickets/queriesv2/tickets"
	"github.com/aria-afk/go-get-tickets/queriesv2/users"
)

var (
	GetAllUsers   = users.GetAllUsers
	GetAllTickets = tickets.GetAllTickets
)
