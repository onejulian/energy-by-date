package formrequest

type FormRequest struct {
	Date   string `json:"date" binding:"required"`
	Period string `json:"period" binding:"required"`
}