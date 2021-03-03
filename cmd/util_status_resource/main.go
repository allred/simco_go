package main
import (
    "context"
    "github.com/go-redis/redis/v8"
    "fmt"
    "os"
    "time"
)

var ctx = context.Background()
var redis_hash = "simco:resources"

/*
func rdb_do(rdb Redis) {
    val, err := rdb.Do()
}
*/

func redis_client() (rdb *redis.Client) {
    uri := os.Getenv("SIMCO_REDIS_URI")
    token := os.Getenv("SIMCO_REDIS_TOKEN")
    connection_str := "redis://" + ":" + token + "@" +  uri
    //fmt.Println("DEBUG: ", connection_str)
    opt, err := redis.ParseURL(connection_str)
    if err != nil {
        panic(err)
    }
    
    redis_db := redis.NewClient(opt)
    fmt.Println(redis_db)

    /*
    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("V: ", val)
    */
    return redis_db
}

func main() {
    start := time.Now()
    fmt.Println("start: ", start)

    rdb := redis_client()
    fmt.Println("DEBUG: ", rdb)
    val := rdb.Do(ctx, "keys", "*")
    fmt.Println("DEBUG: ", val)

    end := time.Now()
    fmt.Println("end: ", end)
}
