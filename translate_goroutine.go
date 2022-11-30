package main
import "fmt"
import "math/rand"
import "sync"
import "context"
import "github.com/aws/aws-sdk-go-v2/aws"
import "github.com/aws/aws-sdk-go-v2/config"
import "github.com/aws/aws-sdk-go-v2/service/translate"

func main(){
        var wg sync.WaitGroup
        for i := 0; i < 650; i++ {
                wg.Add(1) // 每启动一个协程就加 1
                go func(){
                        defer wg.Done() // 协程退出时自动减 1
                        //fmt.Printf("%d ",i)
                        var a [2]string
                        a[0] = "en"
                        a[1] = "fr"
                        var b [101]string
                        b[0] = "a big headache令人头痛的事情"
                        b[1] = "a fraction of 一部分"
                        b[2] = "Hello World"
                        b[3] = "a matter of concern 焦点"
                        b[4] = "a series of 一系列,一连串above all 首先,尤其是"
                        b[5] = "absent from不在,缺席"
                        b[6] = "abundant in富于"
                        b[7] = "account for 解释"
                        b[8] = "accuse sb] of sth.控告"
                        b[9] = "add to增加add up to"
                        b[10] = "after all 毕竟,究竟"
                        b[11] = "agree with同意"
                        b[12] = "ahead of time  schedule提前"
                        b[13] = "ahead of 在...之前ahead of time 提前"
                        b[14] = "alien to与...相反"
                        b[15] = "all at once 突然,同时"
                        b[16] = "all but 几乎;除了...都"
                        b[17] = "all of a sudden 突然"
                        b[18] = "all over again 再一次,重新"
                        b[19] = "all over 遍及"
                        b[20] = "all right 令人满意的;可以"
                        b[21] = "all the same 仍然,照样的"
                        b[22] = "all the time 一直,始终"
                        b[23] = "angry with sb = at about sth.生气,愤怒"
                        b[24] = "anxious about for忧虑,担心"
                        b[25] = "anything but 根本不"
                        b[26] = "apart from 除...外(有/无)"
                        b[27] = "appeal to 吸引，申诉，请求"
                        b[28] = "applicable to适用于"
                        b[29] = "apply to适用"
                        b[30] = "appropriate for/to适当,合适"
                        b[31] = "approximate to近似,接近"
                        b[32] = "apt at聪明,善于"
                        b[33] = "apt to易于"
                        b[34] = "around the clock夜以继日"
                        b[35] = "as a matter of fact 实际上"
                        b[36] = "as a result(of) 因此,由于"
                        b[37] = "as a rule 通常,照例"
                        b[38] = "as far as ...be concerned 就...而言"
                        b[39] = "as far as 远至,到...程度"
                        b[40] = "as follows 如下"
                        b[41] = "as for 至于,关于"
                        b[42] = "as good as 和...几乎一样"
                        b[43] = "as if 好像,防腐"
                        b[44] = "as regards 关于,至于"
                        b[45] = "as to 至于,关于"
                        b[46] = "as usual 像平常一样,照例"
                        b[47] = "as well as 除...外(也),即...又"
                        b[48] = "as well 同样,也,还"
                        b[49] = "ashamed of羞愧,害臊"
                        b[50] = "aside from 除...外(还有)"
                        b[51] = "ask for the moon异想天开"
                        b[52] = "at a loss 茫然,不知所措"
                        b[53] = "at a time 一次,每次"
                        b[54] = "at all costs 不惜一切代价"
                        b[55] = "at all events 不管怎样,无论如何"
                        b[56] = "at all times 随时,总是"
                        b[57] = "at all 丝毫(不),一点也不"
                        b[58] = "at any rate 无论如何,至少"
                        b[59] = "at best 充其量,至多"
                        b[60] = "at first sight 乍一看,初看起来"
                        b[61] = "at first 最初,起先"
                        b[62] = "at hand 在手边,在附近"
                        b[63] = "at heart 内心里,本质上"
                        b[64] = "at home 在家,在国内"
                        b[65] = "at intervals 不时,每隔..."
                        b[66] = "at large 大多数,未被捕获的"
                        b[67] = "at last 终于"
                        b[68] = "at least 至少"
                        b[69] = "at length 最终,终于"
                        b[70] = "at most 至多,不超过"
                        b[71] = "at no time 从不,决不"
                        b[72] = "at one time 曾经,一度;同时"
                        b[73] = "at present 目前,现在"
                        b[74] = "at someones disposal 任...处理"
                        b[75] = "at the cost of 以...为代价"
                        b[76] = "at the mercy of 任凭...摆布"
                        b[77] = "at the moment 此刻,目前"
                        b[78] = "at this rate 照此速度"
                        b[79] = "at times 有时,间或"
                        b[80] = "aware of意识到"
                        b[81] = "back and forth 来回地,反复地"
                        b[82] = "back of 在...后面"
                        b[83] = "back up后备,支援"
                        b[84] = "bare of几乎没有,缺乏"
                        b[85] = "be able to do能够"
                        b[86] = "be around差不多"
                        b[87] = "be available to sb.可用,可供"
                        b[88] = "be bound to一定"
                        b[89] = "be capable of doing能够"
                        b[90] = "be concerned with 关心…，涉足…"
                        b[91] = "be dying to渴望"
                        b[92] = "be fed up with受够了be tired of"
                        b[93] = "be in hospital 住院"
                        b[94] = "be in season 上市的/in peak season旺季"
                        b[95] = "be in the mood to do sth = 想做"
                        b[96] = "be pressed for time时间不够"
                        b[97] = "be tied up with忙于"
                        b[98] = "be under the weather 身体不好"
                        b[99] = "beat around the bush 拐弯没角"
                        b[100] = "beat the crowd 避开人群"
                        cfg,err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))//"us-east-1"
                        translate_client := translate.NewFromConfig(cfg)

                        result, err := translate_client.TranslateText(context.TODO(),&translate.TranslateTextInput{
                                SourceLanguageCode: aws.String(a[0]),
                                TargetLanguageCode: aws.String(a[1]),
                                Text:               aws.String(b[rand.Intn(100)]),
                        })
                        if err != nil {
                                fmt.Println("error: %v", err)
                        }
                        fmt.Println(*result.TranslatedText)

                        if err != nil {
                                fmt.Println("error: %v", err)
                        }
                }()
        }
        wg.Wait() // 等待所有协程执行完毕
}
