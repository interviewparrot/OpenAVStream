# OpenAVStream
A websocket based fast streaming server. You can stream both audio and video.
This is developed by engineers at interviewparrot.com and its used for there video interview platform. We are trying to open source the version which we use internally. Master branch is current work in progress so might not work. Soon we will release to first working version with 

## How to use it

1. `git clone https://github.com/interviewparrot/OpenAVStream.git`
2. `cd OpenAVStream`
3. `go build -o server pkg/main/server.go`
4. Open the client.html and click start Recording. This would stream the video to server.

## Run using docker

1. `git clone https://github.com/interviewparrot/OpenAVStream.git`
2. `cd OpenAVStream`
3. `docker build -t openavstream:1.0 .`
4. `docker run -p 4040:4040 openavstream:1.0`
5. Open the client.html and click start Recording. This would stream the video to server.



## Performance benchmark
To be published

## Developers
For any queries send a email to developers@interviewparrot.com

