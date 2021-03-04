# Social Window

This is a small Go application which will clear all of tweets older then 6 months from your feed. Please note;
this was written as I'm learning Go so I doubt it's the cleanest way of achieving this, so I'm open to PRs ;).

I've used a generic name so the functionality can be expanded in the future.

## Requirements

You will need to manually create an application in [Twitters apps portal](https://apps.twitter.com/), and then
give it access to your profile. Once you've done this you will find, under the `Keys and Access Tokens` tab,
the following pieces of information you'll need to put into the source code; `Consumer Key (API Key)`,
`Consumer Secret (API Secret)`, `Access Token`, and `Access Token Secret`.

Finally you will need your [tweet archive](https://twitter.com/settings/account). I played around with the Twitter
search API, and, unless you pay, it's [severely limited](https://developer.twitter.com/en/docs/tweets/search/overview),
so, to get around this, the app parses the CSV file which is included by twitter in your tweet archive. To make use of
it you'll need to update the filename on the following line to point to the tweet CSV in your archive;

```go
	f, err := os.Open("D:\\tweets.csv")
```

## Shout Out

It *heavily* uses the work of [Dalton Hubble](https://github.com/dghubble/), which you can see in the go.mod
modules file.

