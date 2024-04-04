## Prerequisite
Docker  https://www.docker.com/products/docker-desktop/

Postman https://www.postman.com/downloads/
(or may other tool used for api testing )


## Steps

1)clone repository 

2)run command docker compose up --build 



## Testing 

### API for peginated respone 
method Get 
handler /youtube/videos
query param: pg=<page number>


### API for search filter

method Get 
handler /youtube/search/videos"
query param: q="<search query >" 
