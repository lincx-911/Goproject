package main

import (
	"errors"
)

type FeedItem struct {
	ItemId int // Id

	Title string // 新闻标题

	Pics []string // 图片列表

	Author string // 作者

	Category string // 类别
}

func reorderFeedItems(inputItems []FeedItem)([]FeedItem, error) {
	result := make([]FeedItem, 10)
	n := len(inputItems)
	if len(inputItems) < 10 {
		return result, errors.New("inputItems less then 10")
	}
	// 动态规划，记录从i出发符合条件的最大个数
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	found := false
	endPoint := 0
	maxCount := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 符合第二个条件，非连续相同autor
			if inputItems[i].Author != inputItems[j].Author {
				if inputItems[i].Category == inputItems[j].Category {
					// 不符合最多连续显示两条相同category,跳过不更新
					if parent[i] != i && inputItems[i].Category == inputItems[parent[i]].Category {
						continue
					}
				}
				// 比较原来的，跟新的相邻item 数量
				if dp[j] < (dp[i] + 1) {
					dp[j] = dp[i] + 1
					parent[j] = i
				}
				if dp[j] == 10 {
					found = true
					endPoint = j
					maxCount=10
				}
			}
		}
		// 找到10个符合条件的
		if found {
			break
		}
	}
	// 找到目前满足条件的最后一个item
	if !found{
		for i := 0; i < n; i++ {
			if maxCount == 10 {
				break
			}
			if maxCount < dp[i] {
				maxCount = dp[i]
				endPoint = i
			}
		}
	}
	
	hadSelect := make(map[int]struct{}, 0)
	for i:=maxCount-1;i>=0;i--{
		hadSelect[endPoint]=struct{}{}
		result[i]=inputItems[endPoint]
		endPoint = parent[endPoint]
	}
	idx:=maxCount
	for index,_:=range inputItems{
		if _,ok:=hadSelect[index];!ok{
			result[idx]=inputItems[index]
			idx++
		}
		if idx==10{
			break
		}
	}
	return result,nil
}

// func reorderFeedItems1(inputItems []FeedItem) ([]FeedItem, error) {
// 	result := make([]FeedItem, 0)
// 	if len(inputItems) < 10 {
// 		return result, errors.New("inputItems less then 10")
// 	}
// 	idx := 0
// 	hadSelect := make(map[int]struct{}, 0)
// 	for index, item := range inputItems {
// 		if index == 0 {
// 			result = append(result, item)
// 			idx++
// 			hadSelect[item.ItemId] = struct{}{}
// 			continue
// 		}
// 		if item.Author == result[idx-1].Author {
// 			continue
// 		}
// 		if item.Category == result[idx-1].Category {
// 			if idx >= 2 && item.Category == result[idx-2].Category {
// 				continue
// 			}
// 		}
// 		result = append(result, item)
// 		hadSelect[item.ItemId] = struct{}{}
// 		idx++
// 	}
// 	if idx < 10 {
// 		for _, item := range inputItems {
// 			if _, ok := hadSelect[item.ItemId]; !ok {
// 				result = append(result, item)
// 				idx++
// 			}
// 			if idx == 10 {
// 				break
// 			}
// 		}
// 	}
// 	return result, nil
// }
