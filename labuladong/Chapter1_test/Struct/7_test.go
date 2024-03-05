package Tree

import (
	"sort"
	"testing"
)

type Twitter struct {
	userMap map[int]*User
}
type User struct {
	id        int
	followees map[int]bool
	tweets    []*Tweet
}
type Tweet struct {
	id   int
	time int
}

var tweetCount int

func Constructor() Twitter {
	return Twitter{userMap: make(map[int]*User)}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.userMap[userId]; !ok {
		this.userMap[userId] = &User{
			id:        userId,
			followees: make(map[int]bool),
		}
	}
	this.userMap[userId].tweets = append(this.userMap[userId].tweets, &Tweet{
		id:   tweetId,
		time: tweetCount,
	})
	tweetCount++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	var result []int
	if user, ok := this.userMap[userId]; ok {
		// 所有关注者的所有推文
		tweets := make([]*Tweet, 0)
		for followeeId := range user.followees {
			if followee, ok := this.userMap[followeeId]; ok {
				tweets = append(tweets, followee.tweets...)
			}
		}
		tweets = append(tweets, user.tweets...)
		// 按时间顺序逆序排列
		sort.Slice(tweets, func(i, j int) bool {
			return tweets[i].time > tweets[j].time
		})
		for _, tweet := range tweets {
			result = append(result, tweet.id)
			if len(result) == 10 {
				break
			}
		}
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.userMap[followerId]; !ok {
		this.userMap[followerId] = &User{
			id:        followerId,
			followees: make(map[int]bool),
		}
	}
	if _, ok := this.userMap[followeeId]; !ok {
		this.userMap[followeeId] = &User{
			id:        followeeId,
			followees: make(map[int]bool),
		}
	}
	this.userMap[followerId].followees[followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.userMap[followerId]; ok {
		delete(this.userMap[followerId].followees, followeeId)
	}
}

func Test_7(t *testing.T) {
}
