package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TicketError struct {
	msg string
}

func (e *TicketError) Error() string {
	return e.msg
}

type TicketClass struct {
	QuantitySold int `json:"quantity_sold"`
}

type TicketResponseData struct {
	TicketClasses []TicketClass `json:"ticket_classes"`
	Capacity      int           `json:"capacity"`
}

// type TicketResponseData struct {
// 	QuantitySold  int `json:"quantity_sold"`
// 	QuantityTotal int `json:"quantity_total"`
// }

func GetAvailableTickets(ctx context.Context) (int, error) {
	eventId := ctx.Value(ContextKey("EventId")).(string)
	eventbriteToken := ctx.Value(ContextKey("EventbritePrivateToken")).(string)

	eventInfoUrl := fmt.Sprintf("https://www.eventbriteapi.com/v3/events/%v/?token=%v&expand=ticket_classes", eventId, eventbriteToken)
	// eventInfoUrl := fmt.Sprintf("https://www.eventbriteapi.com/v3/events/%v/capacity_tier/?token=%v", eventId, eventbriteToken)

	res, err := http.Get(eventInfoUrl)

	if err != nil {
		fmt.Println("Error making the request")
		return 0, &TicketError{msg: "Error making the request"}
	}

	var ticketResponse TicketResponseData

	err = json.NewDecoder(res.Body).Decode(&ticketResponse)
	defer res.Body.Close()

	if err != nil {
		fmt.Println("Error decoding")
		return 0, &TicketError{msg: "Error decoding"}
	}

	// if len(ticketResponse.TicketClasses) == 0 {
	// 	return 0, &TicketError{msg: "no data"}
	// }

	fmt.Println(ticketResponse)
	return ticketResponse.Capacity - ticketResponse.TicketClasses[0].QuantitySold, nil

	// return ticketResponse.QuantityTotal - ticketResponse.QuantitySold, nil
}
