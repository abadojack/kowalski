## kowalski

Kowalski is a CLI utility for Twitter that let's one follow people
who tweet a certain phrase/word/hashtag in bulk.

# Example:
```bash
#Setting up environment variables. These can be found from [Twitter](https://dev.twitter.com).
export TWITTER_CONSUMER_KEY="your-consumer-key"
export TWITTER_CONSUMER_SECRET="your-consumer-secret"
export TWITTER_ACCESS_KEY="your-twitter-access-key"
export TWITTER_ACCESS_SECRET="your-twitter-access-secret"

#Running the program
./kowalsi -q "#100DaysOfCode" -c 100
```

This example follows 100 people who recently posted tweets containing with the
hashtag "#100DaysOfCode".
