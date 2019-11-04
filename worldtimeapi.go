
import (
        "log"
        "os/exec"
        "strings"
)

func main() {
        cmd := exec.Command("sh", "-c", "curl http://worldtimeapi.org/api/ip")
        stdoutStderr, err := cmd.CombinedOutput()
        if err != nil {
                log.Fatal(err)
        }
        //      log.Printf("%s\n", stdoutStderr)
        timeanddate := (strings.Contains(string(stdoutStderr), "datetime"))
        //      log.Println("count of letters ", timeanddate)
        if timeanddate == true {
                unixtime := strings.Index(string(stdoutStderr), "unixtime")
                unixtimearray := stdoutStderr[unixtime:]

                date_n := strings.Index(string(unixtimearray), "datetime")
                //      log.Println("count of letters ", date_n)
                if date_n > -1 {
                        parsedate := unixtimearray[date_n+11 : date_n+21]
                        log.Printf("oem - date is: %s", string(parsedate))
                        parsetime := unixtimearray[date_n+22 : date_n+30]
                        log.Printf("oem - time is: %s", string(parsetime))
                }       
                
        }
}       
~   
