package main
import "fmt"
import "sync"
import "context"
import "github.com/aws/aws-sdk-go-v2/aws"
import "github.com/aws/aws-sdk-go-v2/config"
import "github.com/aws/aws-sdk-go-v2/service/translate"
import "bufio"
import "log"
import "os"

func main(){
        var wg sync.WaitGroup
        cfg,err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"), config.WithRetryMaxAttempts(5))//"us-east-1"
        translate_client := translate.NewFromConfig(cfg)
        // os.Open() opens specific file in 
        // read-only mode and this return 
        // a pointer of type os.
        file, err := os.Open("sample.txt")

        if err != nil {
                log.Fatalf("failed to open")

        }

        // The bufio.NewScanner() function is called in which the
        // object os.File passed as its parameter and this returns a
        // object bufio.Scanner which is further used on the
        // bufio.Scanner.Split() method.
        scanner := bufio.NewScanner(file)

        // The bufio.ScanLines is used as an 
        // input to the method bufio.Scanner.Split()
        // and then the scanning forwards to each
        // new line using the bufio.Scanner.Scan()
        // method.
        scanner.Split(bufio.ScanLines)
        var text []string

        for scanner.Scan() {
        text = append(text, scanner.Text())
        }

        // The method os.File.Close() is called
        // on the os.File object to close the file
        file.Close()

        // and then a loop iterates through 
        // and prints each of the slice values.
        for _, each_ln := range text {
                fmt.Println(each_ln)
                wg.Add(1) // 每启动一个协程就加 1
                go func(each_ln string){
                        //defer wg.Done() // 协程退出时自动减 1
                        //fmt.Printf("%d ",i)
                        var a [2]string
                        a[0] = "en"
                        a[1] = "zh-TW"
                        
                        
                        result, err := translate_client.TranslateText(context.TODO(),&translate.TranslateTextInput{
                                SourceLanguageCode: aws.String(a[0]),
                                TargetLanguageCode: aws.String(a[1]),
                                Text:               aws.String(each_ln),
                        })
                        
                        fmt.Println(*result.TranslatedText)

                        if err != nil {
                                fmt.Println("error: %v", err)
                        }
			defer wg.Done() // 协程退出时自动减 1
                }(each_ln)
        }
        if err != nil {
                                fmt.Println("error: %v", err)
                        }
        wg.Wait() // 等待所有协程执行完毕
}
