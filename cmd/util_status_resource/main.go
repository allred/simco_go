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

func redis_client() {
    //token := os.Getenv("SIMCO_REDIS_TOKEN")
    uri := os.Getenv("SIMCO_REDIS_URI")
    connection_str := "redis://" + uri
    fmt.Println("DEBUG: ", connection_str)
    //opt, err := redis.ParseURL("redis://")
    opt, err := redis.ParseURL("redis://")
    fmt.Println(opt, err)
    /*
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)
    */
}

func main() {
    start := time.Now()
    fmt.Println("start: ", start)
    redis_client()
    end := time.Now()
    fmt.Println("end: ", end)
}
