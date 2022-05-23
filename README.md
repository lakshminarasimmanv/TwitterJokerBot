# Twitter Joker Bot

### **Setup:**

1. Sign up for a [Twitter Developer Account](https://developer.twitter.com/en) and Create an App.

2. Make sure to setup **Elevated Access** to and enable **Read and Write** permissions.

3. Required External Packages:

   - [github.com/dghubble/go-twitter/twitter](https://github.com/dghubble/go-twitter)
   - [github.com/dghubble/oauth1](https://github.com/dghubble/oauth1)
   - [github.com/icelain/jokeapi](https://github.com/Sv443/JokeAPI)

4. To Install External Packages, run:

`go get -u <Package Name>` 

Example:

`go get -u github.com/dghubble/go-twitter/twitter`

5. Set your Twitter API Keys as Environment Variables.

```
export CONKEY=<Your Consumer Key>

export CONSEC=<Your Consumer Secret Key>

export ACTO=<Your Access Key>

export ACTOSEC=<Your Access Secret Key>
```

6. To run the code, run:

`go run .`
