[![Build Status](https://travis-ci.org/FeiniuBus/jpush-sdk-go.svg?branch=master)](https://travis-ci.org/FeiniuBus/jpush-sdk-go)

# JPush SDK for Go
[JPush](https://www.jiguang.cn/) REST API sdk for Golang

---------------------------------------

## Installation
Simple install the package to your [$GOPATH](http://code.google.com/p/go-wiki/wiki/GOPATH "GOPATH") with the [go tool](http://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get -u code.feelbus.cn/golang/jpush-sdk-go/...
```
Make sure [Git is installed](http://git-scm.com/downloads) on your machine and in your system's `PATH`.

## Example
1. push api
```golang

	//set push object
	push := jpush.NewPushPayload()
	push.Audience = new(jpush.Audience)
	push.Audience.SetAlias("78633997736779513PG99000828576232")
	push.Notification = jpush.NewNotification()
	push.Notification.SetAndroid("push test", "这是一条推送的测试")
	push.Notification.Android.AddExtra("testk", "testv")
	push.Platform = new(jpush.Platform)
	push.Platform.Android()
	push.Options = jpush.NewOptions()
	push.Options.TimeToLive = 3600

	//create push client
	client := jpush.NewPushClient(appkey, secret)

	//send push
	err := client.SendPush(push)
	if err != nil {
		println(err.Error())
	}
```
2. push schedule api(based on push api)
```golang 
	//set push object(same as object of push api)
	push := jpush.NewPushPayload()
	push.Audience = new(jpush.Audience)
	push.Audience.SetAlias("xxxx")
	push.Notification = jpush.NewNotification()
	push.Notification.SetAndroid("push schedule test", "这是一条定期推送的测试")
	push.Platform = new(jpush.Platform)
	push.Platform.Android()

	//set periodical or single object(only one can be set in a schedule task)
	start := time.Now()
	end := time.Now().Add(time.Hour * 3600)
	t := time.Now().Add(time.Second * 20)

	periodical := &jpush.TriggerPeriodicalNode{
		Start:     jpush.ScheduleDateTime{Time: &start},
		End:       jpush.ScheduleDateTime{Time: &end},
		Frequency: 1,
		TimeUnit:  "day",
		Time:      jpush.ScheduleTime{Time: &t},
	}

	// t2 := time.Now().Add(time.Hour * 24 * 365)
	// single := &jpush.TriggerSingleNode{Time: jpush.ScheduleDateTime{Time: &t2}}

	//set schedule object
	schedule := jpush.NewSchedulePayloadWithPeriodical("push_schedule_test", periodical, push)
	//schedule := jpush.NewSchedulePayloadWithSingle("push_schedule_test", single, push)

	//create push schedule client
	client := jpush.NewScheduleClient(appkey, secret)

	//create schedule task
	resp, err := client.CreateSchedule(schedule)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(resp)
```
