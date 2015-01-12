# Bandit-server

Bandit-server is a [Multi-Armed Bandit](http://en.wikipedia.org/wiki/Multi-armed_bandit) api server which needs no configuration neither persistente store.

## Multi-armed what?!

A multi-armed bandit is essentially an online alternative to classical A/B testing. Whereas A/B testing is generally split into extended phases of execution and analysis, Bandit algorithms continually adjust to user feedback and optimize between experimental states. Bandits typically require very little curation and can in fact be left running indefinitely if need be.

The curious sounding name is drawn from the "one-armed bandit", an colloquialism for casino slot machines. Bandit algorithms can be thought of along similar lines as a eager slot player: if one were to play many slot machines continuously over many thousands of attempts, one would eventually be able to determine which machines were hotter than others. A multi-armed bandit is merely an algorithm that performs exactly this determination, using your user's interaction as its "arm pulls". Extracting winning patterns becomes a fluid part of interacting with the application.

John Myles White has an awesome treatise on Bandit implementations in his book [Bandit Algorithms for Website Optimization](http://shop.oreilly.com/product/0636920027393.do).

## Getting Started

1. Install bandit-server. ``go get github.com/peleteiro/bandit-server``
2. Run ```bandit-server --port=3000```
3. Play ``curl http://localhost:3000/ucb1?downloadButtonColor=black,white,blue\&downloadButtonText=default,now``
4. Reward ``curl -X PUT --data "downloadButtonColor=blue" http://localhost:3000/ucb1``

## Determining what to test

The first task at hand requires a little planning. What are some of the things in your app you've always been curious about changing, but never had empirical data to back up potential modifications? Bandits are best suited to cases where changes can be "slipped in" without the user noticing, but since the state assigned to a user will be persisted to their client, you can also change things like UI.

For our example case, we'll be changing the text and color of a download button in our website to see if either change increases user interaction with the feature. We'll be representing these states as two separate experiments (so a user will get separate assignments for color and text).

## Experiments Setup

You keep experiments on your app, there's no experiments setup on the bandit-server. When you call the api you give a set of arms for every experiment and the server uses this configuration. If you call the api changing the configuration it will adapt the data for the new configuration.

In your application you call the api ``http://server/ucb1?downloadButtonColor=black,white,blue`` to get the arm you need to show your user and the server will configure itself and the algorithms considering this configuration. If you give up testing "blue" arm, then your app calls ``http://server/ucb1?downloadButtonColor=black,white`` and, again, bandit-server will reconfigure and not losing any statistical data.

## Storage

By default bandit-server uses global memory to storage statistical data. It's the best option, even for multiple instances, if your bandit-server stays alive and you have enough hits. If your server goes down or you add a new server on your cluster, Multi-Armed Bandit algorithms adjust the result fast enough.

If you need to persist, you can use memcached: ``bandit-server --memcached=host:port``

## Writing your own client considerations

Bandit-server keeps no user state. But usually you want to keep the experiment's choices for a while. You don't want your user to see a different button color every reload.

It's your client job to keep a local cache on your app for the user session. I personally use bandit-server with one-page apps or native mobile apps, and it's really easy to keep this cache using localstore or filesystem. It should be easy to keep cache in traditional web app using sessions as well.

If you're going to write a client for bandit-server, read the [angular-bandit-client](https://github.com/peleteiro/angular-bandit-client) code and how it handles cache and failure.

# Algorithms

There's a couple of algorithms for MAB but we only have UCB1 implemented for now. We would love to have more algorithms. Fell free to send a push request.

# Clients

## Javascript

- [angular-bandit-client](https://github.com/peleteiro/angular-bandit-client)

# Sample Project

See [https://github.com/peleteiro/angular-bandit-client/tree/master/example](https://github.com/peleteiro/angular-bandit-client/tree/master/example).


# Contributing to bandit-server

We encourage you to contribute to bandit-server! Please check out the guidelines about how to proceed.

* Check out the latest master to make sure the feature hasn't been implemented or the bug hasn't been fixed yet.

* Check out the issue tracker to make sure someone already hasn't requested it and/or contributed it

* Fork the project

* Start a feature/bugfix branch

* Commit and push until you are happy with your contribution
Make sure to add tests for the feature/bugfix. This is important so I don't break it in a future version unintentionally.

* Please try not to mess with the Makefile, version, or history. If you want to have your own version, or is otherwise necessary, that is fine, but please isolate it to its own commit so I can cherry-pick around it.

# License

Bandit-server is released under the [MIT License](http://www.opensource.org/licenses/MIT).
