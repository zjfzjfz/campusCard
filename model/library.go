package model

type Seat struct {
	ID       string `json:"id"`
	Area     string `json:"area"`
	Status   string `json:"status"`
}

type Reservation struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	SeatID    string `json:"seat_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Status    string `json:"status"`
}

type CheckIn struct {
	ID           string `json:"id"`
	ReservationID string `json:"reservation_id"`
	CheckInTime  string `json:"check_in_time"`
}