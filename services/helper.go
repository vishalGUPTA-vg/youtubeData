package services


import "youtubedata/dtos"


func ListResponse(res any, total int64, limit int64) *dtos.Response {


   pageCount := limit


   if pageCount == 0 {
       pageCount = 1
   }
   pages := total / (pageCount)


   if total%pageCount > 0 || pages == 0 {
       pages += 1
   }


   return &dtos.Response{
       TotalCount: total,
       Pages:      pages,
       List:       res,
   }
}





