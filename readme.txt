Finpay Video Task

- Technologies Used
    - Elastisearch
        - Elastisearch is mainly used for search related task and storing video results.
        - It is a NoSQL highly scalable, available and performat.
        - We can use MYSQL, MongoDB with TEXT index. But their search performance decreased for heavy load, high amount of data.
    - Redis
        - Used to store youtube apikeys.
        - We can use any database from MYSQL to MongoDB, even Elastisearch. There are no performance issues here.

- Code Format
    - We have 2 projects
        - Video API
            - API to get search results and store new apikeys.
        - Video Parse
            - Worker to parse youtube results and store.
    - Parser is not written in API, since API is scalable and will have multiple servers running. Since we shouldn't run parser in every server, separated it.

- Testing
    - Run both projects (details are present in each project readme.txt)
    - Keyword is cricket (change in parser project .test-env) 
    - Add Youtube API Key (POST http://localhost:5000/admin/apikey?apikey=AIzaSyAVK3SDDPelxrblquSIZ5gCgMCSmNemDCE)
    - Search for videos (GET http://localhost:5000/video/search?query=bangla%20cricket)