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
    uri := os.Getenv("SIMCO_REDIS_URI")
    token := os.Getenv("SIMCO_REDIS_TOKEN")
    connection_str := "redis://" + ":" + token + "@" +  uri
    //fmt.Println("DEBUG: ", connection_str)
    opt, err := redis.ParseURL(connection_str)
    if err != nil {
        panic(err)
    }
    
    rdb := redis.NewClient(opt)
    fmt.Println(rdb)

    /*
    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("V: ", val)
    */
}

func main() {
    start := time.Now()
    fmt.Println("start: ", start)
    redis_client()
    end := time.Now()
    fmt.Println("end: ", end)
}
