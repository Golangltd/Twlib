package twlib_uniqueid

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
	"sync"
	"time"
)

const (
	timeTemplate      = "2006-01-02 15:04:05"
	appHostIDBits int = 13 // 最大节点数 2^13 8192
	sequenceBits  int = 10 // 最大毫秒并发数量 2^10 1024
)

var (
	mutex              sync.Mutex
	startTime          int64                                // 开始使用时间，使用之后，禁止变更，否则会出现重复ID
	lastTimestamp      int64 = -1                           // 上次生成ID的时间戳
	sequence           int64                                // 当前毫秒生成的序列号
	maxSequence        int64 = (1 << sequenceBits) - 1      // 当前毫秒生成的最大序列号
	timestampLeftShift       = appHostIDBits + sequenceBits // 时间戳的位移量
	appHostID          int64                                // 节点ID
	maxAppHostID       int64 = (1 << appHostIDBits) - 1     // 最大节点ID
)

// 初始化ID生成器
func InitGenerator(appID int64) error {
	if appHostID > maxAppHostID {
		return errors.New(fmt.Sprintf("app host id bigger than max:%d", maxAppHostID))
	}
	appHostID = appID
	startTime = stringToUnixMillionTime(timeTemplate, "2020-10-20 00:00:00") // 开始使用时间，使用之后，禁止变更，否则会出现重复ID
	return nil
}

// 获取一个本服唯一ID
func NextID() int64 {
	return createOne()
}

// 获取一个UUID(去除"-")
func NextUUID() (string, error) {
	u4, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("something went wrong:%s\n", err)
		return "", errors.New(fmt.Sprintf("something went wrong:%s", err))
	}
	str := strings.Replace(u4.String(), "-", "", -1)
	return str, nil
}

func createOne() int64 {
	id := generateUniqueID()
	if id > 0 {
		return id
	}
	return -1
}

// 生成唯一ID
func generateUniqueID() int64 {
	mutex.Lock()
	defer mutex.Unlock()
	currentTime := time.Now().UnixNano() / 1e6
	if currentTime < lastTimestamp {
		return -1
	}
	if currentTime == lastTimestamp { // 出现毫秒并发情况
		sequence = (sequence + 1) & maxSequence
		if sequence == maxSequence { // 达到当前毫秒并发最大值，那么切换到下一毫秒
			currentTime = nextMillionTime(lastTimestamp)
		}
	}
	lastTimestamp = currentTime
	return ((currentTime - startTime) << timestampLeftShift) | (appHostID << sequenceBits) | sequence
}

// 获取下一个毫秒值
func nextMillionTime(timestamp int64) int64 {
	for {
		current := time.Now().Nanosecond() / 1e6
		if int64(current) > timestamp {
			return int64(current)
		}
	}
}

func stringToUnixMillionTime(timeTemplate string, timeStr string) int64 {
	tampTime, _ := time.ParseInLocation(timeTemplate, timeStr, time.Local)
	return tampTime.UnixNano() / 1e6
}

func UnixMillionTimeToString(timeTemplate string, t int64) string {
	return time.Unix(t/1000, 0).Format(timeTemplate)
}
