package main

/*
 * Design a simplified version of Twitter where users can post tweets, follow/unfollow
 * another user, and is able to see the 10 most recent tweets in the user's news feed.
 *
 * Implement the Twitter class:
 *	- Twitter() Initializes your twitter object.
 *	- void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user userId.
 *	  Each call to this function will be made with a unique tweetId.
 *	- List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's news feed.
 *	  Each item in the news feed must be posted by users who the user followed or by the user themself. Tweets must be ordered from most recent to least recent.
 *	- void follow(int followerId, int followeeId) The user with ID followerId started following the user with ID followeeId.
 *	- void unfollow(int followerId, int followeeId) The user with ID followerId started unfollowing the user with ID followeeId.
 *
 * link: https://leetcode.com/problems/design-twitter/description/
 */

import "fmt"

type User struct {
	following map[int]bool
}

// maintains tweets with userID in chronological order, latest at last
type Tweet struct {
	userId, tweetId int
}

// sample data set :
//
//	{
//		userDb : {
//			1 : {2 : true},
//			2 : {1 : false},
//			3 : {1 : true, 2: true}
//		},
//		tweets : [
//			[1, 2], [2, 6], [1, 10], [3, 6] <------ latest tweet
//		]
//	}
type Twitter struct {
	userDb map[int]*User
	tweets []*Tweet
}

func Constructor() Twitter {
	return Twitter{make(map[int]*User), []*Tweet{}}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	// add user to db if not already present
	if _, ok := this.userDb[userId]; !ok {
		this.userDb[userId] = &User{make(map[int]bool)}
	}
	// add to tweets list at the end
	this.tweets = append(this.tweets, &Tweet{userId, tweetId})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	if _, ok := this.userDb[userId]; !ok {
		this.userDb[userId] = &User{make(map[int]bool)}
	}

	// get the target user to get feed for followers list
	target := this.userDb[userId]

	res, index := []int{}, len(this.tweets)-1

	// reverse traverse all tweets until we get max tweets needed, or start of list
	for len(res) < 10 && index >= 0 {
		cur := this.tweets[index]
		user, tweet := cur.userId, cur.tweetId

		// add to feed if tweet belongs to follower or target user
		if target.following[user] || user == userId {
			res = append(res, tweet)
		}

		index--
	}

	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	// add both users if not present
	if _, ok := this.userDb[followerId]; !ok {
		this.userDb[followerId] = &User{make(map[int]bool)}
	}

	if _, ok := this.userDb[followeeId]; !ok {
		this.userDb[followeeId] = &User{make(map[int]bool)}
	}

	// get the target user from db and followee ID to true
	target := this.userDb[followerId]

	target.following[followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	// add both users if not present
	if _, ok := this.userDb[followerId]; !ok {
		this.userDb[followerId] = &User{make(map[int]bool)}
	}

	if _, ok := this.userDb[followeeId]; !ok {
		this.userDb[followeeId] = &User{make(map[int]bool)}
	}

	// get the target user from db and followee ID to false
	target := this.userDb[followerId]

	target.following[followeeId] = false
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

func main() {
	twitter := Constructor()
	twitter.PostTweet(2, 5)
	twitter.PostTweet(1, 3)
	twitter.PostTweet(1, 101)
	twitter.PostTweet(2, 13)
	twitter.PostTweet(2, 10)
	twitter.PostTweet(1, 2)
	twitter.PostTweet(2, 94)
	twitter.PostTweet(2, 505)
	twitter.PostTweet(1, 333)
	twitter.PostTweet(1, 22)
	fmt.Println(twitter.GetNewsFeed(2))
	twitter.Follow(2, 1)
	fmt.Println(twitter.GetNewsFeed(2))
}
