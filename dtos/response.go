package dtos


type Response struct {
   TotalCount int64 `json:"totat_count"`
   Pages      int64 `json:"pages"`
   List       any `json:"list"`
}





