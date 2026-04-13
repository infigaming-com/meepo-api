package v1

// PlayerListTicketsRequest is the request for listing player's own withdraw tickets.
type PlayerListTicketsRequest struct {
	Status   *string `json:"status,omitempty"`
	Page     *int32  `json:"page,omitempty"`
	PageSize *int32  `json:"page_size,omitempty"`
}

// PlayerTicket represents a ticket visible to the player.
type PlayerTicket struct {
	TicketId      int64  `json:"ticket_id"`
	Status        string `json:"status"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
	PaymentStatus string `json:"payment_status"`
	PlayerComment string `json:"player_comment"`
	CreatedAt     int64  `json:"created_at"`
	ReviewedAt    int64  `json:"reviewed_at"`
}

// PlayerListTicketsResponse is the response for listing player's own withdraw tickets.
type PlayerListTicketsResponse struct {
	Tickets    []*PlayerTicket `json:"tickets"`
	Page       int32           `json:"page"`
	PageSize   int32           `json:"page_size"`
	TotalCount int32           `json:"total_count"`
}

// PlayerGetTicketRequest is the request for getting a single player ticket.
type PlayerGetTicketRequest struct {
	TicketId int64 `json:"ticket_id"`
}

// PlayerGetTicketResponse is the response for getting a single player ticket.
type PlayerGetTicketResponse = PlayerTicket
