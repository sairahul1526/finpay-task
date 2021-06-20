Video parser
- It will parse videos from youtube API, based on specifed keyword and add to elasti search
- It will use multiple API Keys
    - It uses one API Key, with least score, continously until its credits got used. Then updates its score to next day.
    - It then uses next least score API Key
- For prod, it should be deployed to AWS Lambda and run every 10 sec using AWS Cloudwatch Events
- To run locally, use "make run" (you need to have docker installed)