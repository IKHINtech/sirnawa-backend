package response

type DashboardMobileResponse struct {
	RondaSchedule *RondaScheduleResponse `json:"ronda_schedule"`
	Announcecment *AnnouncementResponse  `json:"announcement"`
}
