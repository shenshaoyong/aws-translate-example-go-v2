package main
import "fmt"
import "context"
import "github.com/aws/aws-sdk-go-v2/aws"
import "github.com/aws/aws-sdk-go-v2/config"
import "github.com/aws/aws-sdk-go-v2/service/translate"

func main(){
        fmt.Printf("test say")
        var b [3]string
        b[0] = "en"
        b[1] = "fr"
        b[2] = "Hello World"
        cfg,err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
        translate_client := translate.NewFromConfig(cfg)

        result, err := translate_client.TranslateText(context.TODO(),&translate.TranslateTextInput{
                SourceLanguageCode: aws.String(b[0]),
                TargetLanguageCode: aws.String(b[1]),
                Text:               aws.String(b[2]),
        })
        if err != nil {
                fmt.Println("error: %v", err)
        }
        fmt.Println(*result.TranslatedText)
}
