package response

type DashboardMobileResponse struct {
	RondaSchedule *RondaScheduleResponse `json:"ronda_schedule"`
	Announcement  *AnnouncementResponse  `json:"announcement"`
	IplBill       *IplBillResponse       `json:"ipl_bill"`
}
