package scheduledTask

import (
	"fmt"
    "context"
	"github.com/robfig/cron/v3"
    "campusCard/dao"
    "campusCard/cache"
    "github.com/jinzhu/gorm"
    "github.com/go-redis/redis/v8"
)


var (
    C *cron.Cron
    ctx = context.Background()

)

func init() {
	updateLimit(dao.Db, cache.Rdb)
    C = cron.New()
	_, err := C.AddFunc("*/1 * * * *", func() {
    //_, err := C.AddFunc("0 0 * * *", func() {
        // 在这里定义定时执行的任务
        updateLimit(dao.Db, cache.Rdb)
        //fmt.Println("执行定时任务...")
        //printHashFromRedis(cache.Rdb, "tradeLimit")
    })
	if err != nil {
        fmt.Println("添加定时任务出错：", err)
        return
    }

    // 启动Cron调度器
    C.Start()

}

func updateLimit(db *gorm.DB, rdb *redis.Client) error {
    var accountInfos []dao.AccountInfo
    if err := db.Find(&accountInfos).Error; err != nil {
        return err
    }

    // 删除已经存在的 "tradeLimit" 哈希表
    err := rdb.Del(ctx, "tradeLimit").Err()
    if err != nil {
        return err
    }

    // 创建一个单独的哈希表，存储多个ID和对应的Limit
    hashName := "tradeLimit"
    for _, info := range accountInfos {
        err := rdb.HSet(ctx, hashName, info.ID, info.Limit).Err()
        if err != nil {
            return err
        }
    }

    return nil
}

// 从Redis中打印指定哈希表的内容
/*func printHashFromRedis(rdb *redis.Client, hashName string) {
    result := rdb.HGetAll(ctx, hashName).Val()
    fmt.Printf("Hash table %s:\n", hashName)
    for field, value := range result {
        fmt.Printf("%s: %s\n", field, value)
    }
}*/