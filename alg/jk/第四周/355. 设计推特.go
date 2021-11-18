package 第四周

type Twitter struct {
	FollowedMap map[int]map[int]int
	TwitterMap  []map[int]int
}

func Constructor() Twitter {
	var TwitterMap []map[int]int
	FollowedMap := map[int]map[int]int{}

	return Twitter{
		TwitterMap:  TwitterMap,
		FollowedMap: FollowedMap,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.TwitterMap = append(this.TwitterMap, map[int]int{userId: tweetId})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	twitterArray := []int{}

	for i := len(this.TwitterMap) - 1; i >= 0; i-- {

		for authId, twitterId := range this.TwitterMap[i] {
			if userId == authId || this.FollowedMap[userId][authId] > 0 {
				twitterArray = append(twitterArray, twitterId)
			}
		}
		if len(twitterArray) == 10 {
			return twitterArray
		}
	}
	return twitterArray
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.FollowedMap[followerId] == nil {
		this.FollowedMap[followerId] = map[int]int{}
	}
	this.FollowedMap[followerId][followeeId] = 1
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.FollowedMap[followerId] != nil {
		this.FollowedMap[followerId][followeeId] = 0
	}
}
