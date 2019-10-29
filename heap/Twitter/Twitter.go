package main

import (
	"container/heap"
	"fmt"
	"time"
)

type TweetHeap []*Tweet

func (h TweetHeap) Len() int {
	return len(h)
}

// max heap
func (h TweetHeap) Less(i, j int) bool {
	return h[i].ts > h[j].ts
}

func (h TweetHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *TweetHeap) Push(x interface{}) {
	*h = append(*h, x.(*Tweet))
}

func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

type Tweet struct {
	id   int
	ts   int64
	next *Tweet
}

func NewTweet(id int) *Tweet {
	return &Tweet{
		id, time.Now().UnixNano(), nil,
	}
}

type User struct {
	id         int
	followed   map[int]struct{}
	tweet_head *Tweet
}

func NewUser(id int) *User {
	u := &User{
		id:         id,
		followed:   make(map[int]struct{}),
		tweet_head: nil,
	}
	u.follow(id)
	return u
}

func (u *User) follow(id int) {
	u.followed[id] = struct{}{}
}

func (u *User) unfollow(id int) {
	delete(u.followed, id)
}

func (u *User) post(id int) {
	t := NewTweet(id)
	if u.tweet_head == nil {
		u.tweet_head = t
	} else {
		old := u.tweet_head
		t.next = old
		u.tweet_head = t
	}
}

type Twitter struct {
	userMap map[int]*User
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{make(map[int]*User)}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	user, ok := this.userMap[userId]
	if !ok {
		user = NewUser(userId)
		this.userMap[userId] = user
	}
	user.post(tweetId)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	my, ok := this.userMap[userId]
	if !ok {
		return []int{}
	}
	users := my.followed
	h := &TweetHeap{}
	heap.Init(h)
	for uid := range users {
		user, _ := this.userMap[uid]
		if user.tweet_head != nil {
			heap.Push(h, user.tweet_head)
		}
	}
	var res []int
	n := 0
	for h.Len() > 0 && n < 10 {
		t := heap.Pop(h).(*Tweet)
		res = append(res, t.id)
		if t.next != nil {
			heap.Push(h, t.next)
		}
		n++
	}
	return res
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.userMap[followerId]; !ok {
		u := NewUser(followerId)
		this.userMap[followerId] = u
	}
	if _, ok := this.userMap[followeeId]; !ok {
		u := NewUser(followeeId)
		this.userMap[followeeId] = u
	}
	this.userMap[followerId].follow(followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.userMap[followerId]; !ok || followeeId == followerId {
		return
	}

	this.userMap[followerId].unfollow(followeeId)
}

func main() {
	obj := Constructor();
	obj.PostTweet(1, 5);
	obj.Unfollow(1, 1);
	param_2 := obj.GetNewsFeed(1);
	fmt.Println(param_2)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
