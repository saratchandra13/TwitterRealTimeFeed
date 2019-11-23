# **Twitter RealTime Tweets**
  

_**Requirement**_ : RealTime Tweets to be shown to user as a streaming API response and search should be based on hashtag and account.


**Thought Process**

    1. Finding out how Twitter apis work, documentation was helpful.
    2. Finding out how we can authenticate of our application.
    3. Reading about Oauth1.0a to authenticate our client.
    4. Now we know what needs to be done, lets kickstart the project.
    5. Start with finding a framework for our application.
    6. Gin,chi,gorilla are considered and light-weight gin is considered for this project, We might differ in thoughts over here. (https://github.com/gin-gonic/gin)
    7. Now we will think of how we can design our system. Find below the design of this application.
    8. Finding any useful 3rd party lib which reduces our effort.
    9. Providing Streaming API to search the tweets based on Keyword.
    10. Aren`t we at the end? :P 
    
**Design**
    
    1. Dividing into Components.
        1. Server up and running. (main.go takes care of server run and incoming requests)
        2. Router taking care of routing the requests to desired controller.
        3. Controller taking care of the request and extracting Get Params from request.
        4. Controller calls Views to get the desired result for the request.
        5. Views using models and the 3rd party libraries to make the response ready.
        6. Views reverting back to controller with either response or error.
        7. Based on this controller responds to the user with either a stream of tweets or closes the request with empty response.
    2. Each component has it unique liability to its given task and then all of them combining into an application serving many requests.     
    3. We are using Vendoring as dependency management technique (Go Modules can also be used.)


**UseFul Links**
    
    1. https://blog.gopheracademy.com/advent-2015/vendor-folder/
    2. https://github.com/gin-gonic/gin
    3. https://developer.twitter.com/en/docs/tweets/search/api-reference/get-search-tweets
    4. https://github.com/dghubble/go-twitter/twitter
    5. https://github.com/dghubble/oauth1

**Run Server**
    
    1. Please take a pull from the master and set the gopath and goroot for this project.
    2. To Run Server - go run main.go (Will start server at port 8080)
    3. Curl to get the realtime filtered tweets. 
    
    curl -X GET \
      'http://localhost:8080/search/tweets/?source=%40imVkohli' \
      -H 'Accept: */*' \
      -H 'Accept-Encoding: gzip, deflate' \
      -H 'Cache-Control: no-cache' \
      -H 'Connection: keep-alive' \
      -H 'Host: localhost:8080' \
      -H 'Postman-Token: 57cede03-0927-47f6-9fd7-2261d3dbe5fa,df15bdef-548b-41dd-afe8-9573d0c86b1e,df937f20-bf71-4504-b365-56e47283cc7e' \
      -H 'User-Agent: PostmanRuntime/7.19.0' \
      -H 'cache-control: no-cache,no-cache'
      
![Alt text](/Account_search.png?raw=true "Tweets On Kohli")
![Alt text](/HashTag_search.png?raw=true "Tweets On #IndvsBan (Pink ball test match)")             