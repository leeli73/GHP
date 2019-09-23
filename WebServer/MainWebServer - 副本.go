package main
import(
	"os"
	"io"
	"fmt"
	"net"
	"time"
	"bufio"
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"os/exec"
	"math/rand"
	"io/ioutil"
	"crypto/md5"
	"crypto/aes"
    "encoding/hex"
	"encoding/json"
	"crypto/cipher"
	"gitee.com/johng/gf/g"
	"github.com/axgle/mahonia"
	"gitee.com/johng/gf/g/net/ghttp"
)
type WebConfig struct{
	Root string `json:Root`
	Asset string `json:Asset`
	Defaults []string `json:Defaults`
	Port string `json:Port`
	HttpsPort string `json:HttpsPort`
	CertFile string `json:CertFile`
	KeyFile string `json:KeyFile`
	EnableAdmin string `json:EnableAdmin`
	Plug string `json:Plug`
	Domain []string `json:Domain`
}
type ServerConfig struct{
	ReStartTimeOut string `json:ReStartTimeOut`
	LogRoot string `json:LogRoot`
	LogLevel string `json:LogLevel`
	LogMaxSize string `json:LogMaxSize`
	SEPHWorkFolder string `json:SEPHWorkFolder`
	UploadMaxSize string `json:UploadMaxSize`
	AllWeb []WebConfig `json:AllWeb`
}
type EnterData struct{
	ID string
	r *ghttp.Request
}
var Config ServerConfig
var AllServer [128]*ghttp.Server
var LogTempList [2048]string
var LogTempListCount int
var AllEnterData = make(map[string]EnterData)
var ContentTypeMap = map[string]string{".tif":"image/tiff",
									".001":"application/x-001",
									".301":"application/x-301",
									".323":"text/h323",
									".906":"application/x-906",
									".907":"drawing/907",
									".a11":"application/x-a11",
									".acp":"audio/x-mei-aac",
									".ai":"application/postscript",
									".aif":"audio/aiff",
									".aifc":"audio/aiff",
									".aiff":"audio/aiff",
									".anv":"application/x-anv",
									".asa":"text/asa",
									".asf":"video/x-ms-asf",
									".asp":"text/asp",
									".asx":"video/x-ms-asf",
									".au":"audio/basic",
									".avi":"video/avi",
									".awf":"application/vnd.adobe.workflow",
									".biz":"text/xml",
									".bmp":"application/x-bmp",
									".bot":"application/x-bot",
									".c4t":"application/x-c4t",
									".c90":"application/x-c90",
									".cal":"application/x-cals",
									".cat":"application/vnd.ms-pki.seccat",
									".cdf":"application/x-netcdf",
									".cdr":"application/x-cdr",
									".cel":"application/x-cel",
									".cer":"application/x-x509-ca-cert",
									".cg4":"application/x-g4",
									".cgm":"application/x-cgm",
									".cit":"application/x-cit",
									".class":"java/*",
									".cml":"text/xml",
									".cmp":"application/x-cmp",
									".cmx":"application/x-cmx",
									".cot":"application/x-cot",
									".crl":"application/pkix-crl",
									".crt":"application/x-x509-ca-cert",
									".csi":"application/x-csi",
									".css":"text/css",
									".cut":"application/x-cut",
									".dbf":"application/x-dbf",
									".dbm":"application/x-dbm",
									".dbx":"application/x-dbx",
									".dcd":"text/xml",
									".dcx":"application/x-dcx",
									".der":"application/x-x509-ca-cert",
									".dgn":"application/x-dgn",
									".dib":"application/x-dib",
									".dll":"application/x-msdownload",
									".doc":"application/msword",
									".dot":"application/msword",
									".drw":"application/x-drw",
									".dtd":"text/xml",
									".dwf":"Model/vnd.dwf",
									".dwg":"application/x-dwg",
									".dxb":"application/x-dxb",
									".dxf":"application/x-dxf",
									".edn":"application/vnd.adobe.edn",
									".emf":"application/x-emf",
									".eml":"message/rfc822",
									".ent":"text/xml",
									".epi":"application/x-epi",
									".eps":"application/postscript",
									".etd":"application/x-ebx",
									".exe":"application/x-msdownload",
									".fax":"image/fax",
									".fdf":"application/vnd.fdf",
									".fif":"application/fractals",
									".fo":"text/xml",
									".frm":"application/x-frm",
									".g4":"application/x-g4",
									".gbr":"application/x-gbr",
									".gif":"image/gif",
									".gl2":"application/x-gl2",
									".gp4":"application/x-gp4",
									".hgl":"application/x-hgl",
									".hmr":"application/x-hmr",
									".hpg":"application/x-hpgl",
									".hpl":"application/x-hpl",
									".hqx":"application/mac-binhex40",
									".hrf":"application/x-hrf",
									".hta":"application/hta",
									".htc":"text/x-component",
									".htm":"text/html",
									".html":"text/html",
									".htt":"text/webviewhtml",
									".htx":"text/html",
									".icb":"application/x-icb",
									".ico":"image/x-icon",
									".iff":"application/x-iff",
									".ig4":"application/x-g4",
									".igs":"application/x-igs",
									".iii":"application/x-iphone",
									".img":"application/x-img",
									".ins":"application/x-internet-signup",
									".isp":"application/x-internet-signup",
									".IVF":"video/x-ivf",
									".java":"java/*",
									".jfif":"image/jpeg",
									".jpe":"image/jpeg",
									".jpeg":"image/jpeg",
									".jpg":"image/jpeg",
									".js":"application/x-javascript",
									".jsp":"text/html",
									".la1":"audio/x-liquid-file",
									".lar":"application/x-laplayer-reg",
									".latex":"application/x-latex",
									".lavs":"audio/x-liquid-secure",
									".lbm":"application/x-lbm",
									".lmsff":"audio/x-la-lms",
									".ls":"application/x-javascript",
									".ltr":"application/x-ltr",
									".m1v":"video/x-mpeg",
									".m2v":"video/x-mpeg",
									".m3u":"audio/mpegurl",
									".m4e":"video/mpeg4",
									".mac":"application/x-mac",
									".man":"application/x-troff-man",
									".math":"text/xml",
									".mdb":"application/msaccess",
									".mfp":"application/x-shockwave-flash",
									".mht":"message/rfc822",
									".mhtml":"message/rfc822",
									".mi":"application/x-mi",
									".mid":"audio/mid",
									".midi":"audio/mid",
									".mil":"application/x-mil",
									".mml":"text/xml",
									".mnd":"audio/x-musicnet-download",
									".mns":"audio/x-musicnet-stream",
									".mocha":"application/x-javascript",
									".movie":"video/x-sgi-movie",
									".mp1":"audio/mp1",
									".mp2":"audio/mp2",
									".mp2v":"video/mpeg",
									".mp3":"audio/mp3",
									".mp4":"video/mpeg4",
									".mpa":"video/x-mpg",
									".mpd":"application/vnd.ms-project",
									".mpe":"video/x-mpeg",
									".mpeg":"video/mpg",
									".mpg":"video/mpg",
									".mpga":"audio/rn-mpeg",
									".mpp":"application/vnd.ms-project",
									".mps":"video/x-mpeg",
									".mpt":"application/vnd.ms-project",
									".mpv":"video/mpg",
									".mpv2":"video/mpeg",
									".mpw":"application/vnd.ms-project",
									".mpx":"application/vnd.ms-project",
									".mtx":"text/xml",
									".mxp":"application/x-mmxp",
									".net":"image/pnetvue",
									".nrf":"application/x-nrf",
									".nws":"message/rfc822",
									".odc":"text/x-ms-odc",
									".out":"application/x-out",
									".p10":"application/pkcs10",
									".p12":"application/x-pkcs12",
									".p7b":"application/x-pkcs7-certificates",
									".p7c":"application/pkcs7-mime",
									".p7m":"application/pkcs7-mime",
									".p7r":"application/x-pkcs7-certreqresp",
									".p7s":"application/pkcs7-signature",
									".pc5":"application/x-pc5",
									".pci":"application/x-pci",
									".pcl":"application/x-pcl",
									".pcx":"application/x-pcx",
									".pdf":"application/pdf",
									".pdx":"application/vnd.adobe.pdx",
									".pfx":"application/x-pkcs12",
									".pgl":"application/x-pgl",
									".pic":"application/x-pic",
									".pko":"application/vnd.ms-pki.pko",
									".pl":"application/x-perl",
									".plg":"text/html",
									".pls":"audio/scpls",
									".plt":"application/x-plt",
									".png":"image/png",
									".pot":"application/vnd.ms-powerpoint",
									".ppa":"application/vnd.ms-powerpoint",
									".ppm":"application/x-ppm",
									".pps":"application/vnd.ms-powerpoint",
									".ppt":"application/vnd.ms-powerpoint",
									".pr":"application/x-pr",
									".prf":"application/pics-rules",
									".prn":"application/x-prn",
									".prt":"application/x-prt",
									".ps":"application/postscript",
									".ptn":"application/x-ptn",
									".pwz":"application/vnd.ms-powerpoint",
									".r3t":"text/vnd.rn-realtext3d",
									".ra":"audio/vnd.rn-realaudio",
									".ram":"audio/x-pn-realaudio",
									".ras":"application/x-ras",
									".rat":"application/rat-file",
									".rdf":"text/xml",
									".rec":"application/vnd.rn-recording",
									".red":"application/x-red",
									".rgb":"application/x-rgb",
									".rjs":"application/vnd.rn-realsystem-rjs",
									".rjt":"application/vnd.rn-realsystem-rjt",
									".rlc":"application/x-rlc",
									".rle":"application/x-rle",
									".rm":"application/vnd.rn-realmedia",
									".rmf":"application/vnd.adobe.rmf",
									".rmi":"audio/mid",
									".rmj":"application/vnd.rn-realsystem-rmj",
									".rmm":"audio/x-pn-realaudio",
									".rmp":"application/vnd.rn-rn_music_package",
									".rms":"application/vnd.rn-realmedia-secure",
									".rmvb":"application/vnd.rn-realmedia-vbr",
									".rmx":"application/vnd.rn-realsystem-rmx",
									".rnx":"application/vnd.rn-realplayer",
									".rp":"image/vnd.rn-realpix",
									".rpm":"audio/x-pn-realaudio-plugin",
									".rsml":"application/vnd.rn-rsml",
									".rt":"text/vnd.rn-realtext",
									".rtf":"application/x-rtf",
									".rv":"video/vnd.rn-realvideo",
									".sam":"application/x-sam",
									".sat":"application/x-sat",
									".sdp":"application/sdp",
									".sdw":"application/x-sdw",
									".sit":"application/x-stuffit",
									".slb":"application/x-slb",
									".sld":"application/x-sld",
									".slk":"drawing/x-slk",
									".smi":"application/smil",
									".smil":"application/smil",
									".smk":"application/x-smk",
									".snd":"audio/basic",
									".sol":"text/plain",
									".sor":"text/plain",
									".spc":"application/x-pkcs7-certificates",
									".spl":"application/futuresplash",
									".spp":"text/xml",
									".ssm":"application/streamingmedia",
									".sst":"application/vnd.ms-pki.certstore",
									".stl":"application/vnd.ms-pki.stl",
									".stm":"text/html",
									".sty":"application/x-sty",
									".svg":"text/xml",
									".swf":"application/x-shockwave-flash",
									".tdf":"application/x-tdf",
									".tg4":"application/x-tg4",
									".tga":"application/x-tga",
									".tiff":"image/tiff",
									".tld":"text/xml",
									".top":"drawing/x-top",
									".torrent":"application/x-bittorrent",
									".tsd":"text/xml",
									".txt":"text/plain",
									".uin":"application/x-icq",
									".uls":"text/iuls",
									".vcf":"text/x-vcard",
									".vda":"application/x-vda",
									".vdx":"application/vnd.visio",
									".vml":"text/xml",
									".vpg":"application/x-vpeg005",
									".vsd":"application/vnd.visio",
									".vss":"application/vnd.visio",
									".vst":"application/x-vst",
									".vsw":"application/vnd.visio",
									".vsx":"application/vnd.visio",
									".vtx":"application/vnd.visio",
									".vxml":"text/xml",
									".wav":"audio/wav",
									".wax":"audio/x-ms-wax",
									".wb1":"application/x-wb1",
									".wb2":"application/x-wb2",
									".wb3":"application/x-wb3",
									".wbmp":"image/vnd.wap.wbmp",
									".wiz":"application/msword",
									".wk3":"application/x-wk3",
									".wk4":"application/x-wk4",
									".wkq":"application/x-wkq",
									".wks":"application/x-wks",
									".wm":"video/x-ms-wm",
									".wma":"audio/x-ms-wma",
									".wmd":"application/x-ms-wmd",
									".wmf":"application/x-wmf",
									".wml":"text/vnd.wap.wml",
									".wmv":"video/x-ms-wmv",
									".wmx":"video/x-ms-wmx",
									".wmz":"application/x-ms-wmz",
									".wp6":"application/x-wp6",
									".wpd":"application/x-wpd",
									".wpg":"application/x-wpg",
									".wpl":"application/vnd.ms-wpl",
									".wq1":"application/x-wq1",
									".wr1":"application/x-wr1",
									".wri":"application/x-wri",
									".wrk":"application/x-wrk",
									".ws":"application/x-ws",
									".ws2":"application/x-ws",
									".wsc":"text/scriptlet",
									".wsdl":"text/xml",
									".wvx":"video/x-ms-wvx",
									".xdp":"application/vnd.adobe.xdp",
									".xdr":"text/xml",
									".xfd":"application/vnd.adobe.xfd",
									".xfdf":"application/vnd.adobe.xfdf",
									".xhtml":"text/html",
									".xls":"application/x-xls",
									".xlw":"application/x-xlw",
									".xml":"text/xml",
									".xpl":"audio/scpls",
									".xq":"text/xml",
									".xql":"text/xml",
									".xquery":"text/xml",
									".xsd":"text/xml",
									".xsl":"text/xml",
									".xslt":"text/xml",
									".xwd":"application/x-xwd",
									".x_b":"application/x-x_b",
									".sis":"application/vnd.symbian.install",
									".sisx":"application/vnd.symbian.install",
									".x_t":"application/x-x_t",
									".ipa":"application/vnd.iphone",
									".apk":"application/vnd.android.package-archive",
									".xap":"application/x-silverlight-app"}
func main(){
	Init()
	InitConfig()
    AllPort := make(map[string]int)
	for i:=0;i<len(Config.AllWeb);i++{
		AllServer[i] = g.Server(i)
		if len(Config.AllWeb[i].Domain) == 0{
			AllServer[i].BindHandler("/*any",WebRequestReceive)
		} else {
			for j:=0;j<len(Config.AllWeb[i].Domain);j++{
				AllServer[i].Domain(Config.AllWeb[i].Domain[j]).BindHandler("/*any",WebRequestReceive)
			}
		}
		Port,_ := strconv.Atoi(Config.AllWeb[i].Port)
		_,PortIsSet := AllPort[Config.AllWeb[i].Port]
		if !PortIsSet{
			AllPort[Config.AllWeb[i].Port] = Port
			AllServer[i].SetPort(Port)
		}
		if Config.AllWeb[i].HttpsPort != "0"{
			Ports,_ := strconv.Atoi(Config.AllWeb[i].HttpsPort)
			if Config.AllWeb[i].CertFile == "NULL" || Config.AllWeb[i].KeyFile == "NULL"{
				fmt.Println("Please Set Right Path of Cert and Key")
			} else {
				AllServer[i].EnableHTTPS(Config.AllWeb[i].CertFile,Config.AllWeb[i].KeyFile)
				AllServer[i].SetHTTPSPort(Ports)
			}
			AllServer[i].SetHTTPSPort(Ports)
		}
		if Config.AllWeb[i].EnableAdmin == "True"{
			AllServer[i].EnableAdmin()
		}
		AllServer[i].Start()
	}
	ShowServerInfo()
	go CommunicationEngine()
	g.Wait()
}
func PrintLog(log string){
	logLevel := strings.ToUpper(Config.LogLevel)
	log = "日志--"+time.Now().Format("2006-01-02 15:04:05")+"--"+log
	if logLevel == "OUTPUT"{
		LogTempList[LogTempListCount]=log
		LogTempListCount = LogTempListCount + 1
	} else if logLevel == "DEBUG" {
		fmt.Println(log)
		LogTempList[LogTempListCount]=log
		LogTempListCount = LogTempListCount + 1
	} else if logLevel == "NULL" {
	} else if logLevel == "PRINT"{
		fmt.Println(log)
	}
	if LogTempListCount > 2042{
		LogOutPut()
	}
}
func LogOutPut(){
	fileName := Config.LogRoot+"\\log.log"
	fileInfo, _ := os.Stat(fileName)
	filesize:= fileInfo.Size()
	MaxSize,_ := strconv.ParseInt(Config.LogMaxSize, 10, 64)
	if filesize > MaxSize*1000{
		fmt.Println(filesize,"  ",MaxSize*1000)
		err := os.Rename(fileName, Config.LogRoot+"\\log--"+strings.Replace(time.Now().Format("2006-01-02 15:04:05"),":","-",-1)+".log")
		if err != nil {
			fmt.Println(err)
		}
	}
	f,_ := os.OpenFile(fileName,os.O_APPEND|os.O_CREATE|os.O_RDWR,0644)
	log := ""
	for i:=0;i<=LogTempListCount;i++{
		log = log + LogTempList[i] + "\r\n"
	}
    f.WriteString(log)
	f.Close()
	LogTempListCount = 0
}
func InitConfig(){
	var Data,err = ioutil.ReadFile("Config.json")
	if err != nil{
		fmt.Println("Read Log File Error！")
		return
	}
	err = json.Unmarshal(Data,&Config)
	if err != nil{
		fmt.Println("Read Log JSON Error！Please Check if Your Charset is UTF-8!")
		os.Exit(0)
		return
	}
}
func Init(){
	LogTempListCount = 0
}
func ShowServerInfo(){
	fmt.Println("-------------Server Config---------------")
	fmt.Println("LogRoot:",Config.LogRoot)
	fmt.Println("LogMaxSize:",Config.LogMaxSize,"KB")
	fmt.Println("LogLevel:",Config.LogLevel)
	fmt.Println("ReStart TimeOut:",Config.ReStartTimeOut+"s")
	fmt.Println("SEPH Work Folder:",Config.SEPHWorkFolder)
	fmt.Println("Upload Max Size:",Config.UploadMaxSize,"M")
	fmt.Println("-----------------------------------------")
	for i:=0;i<len(Config.AllWeb);i++{
		fmt.Println("---------------WebSite ",i+1,"---------------")
		fmt.Println("Web Domain:",Config.AllWeb[i].Domain)
		fmt.Println("Web Root:",Config.AllWeb[i].Root)
		fmt.Println("Web Asset:",Config.AllWeb[i].Asset)
		fmt.Println("Defaults File:",Config.AllWeb[i].Defaults)
		fmt.Println("Listening Port:",Config.AllWeb[i].Port)
		if Config.AllWeb[i].HttpsPort == "0"{
			fmt.Println("isHTTPS: False")
		} else {
			fmt.Println("isHTTPS: True")
			fmt.Println("HTTPS Listening Port:",Config.AllWeb[i].HttpsPort)
			fmt.Println("HTTPS Cert Path:",Config.AllWeb[i].CertFile)
			fmt.Println("HTTPS Key Path:",Config.AllWeb[i].KeyFile)
		}
		fmt.Println("EnableAdmin: ",Config.AllWeb[i].EnableAdmin)
		fmt.Println("Plug:",Config.AllWeb[i].Plug)
		fmt.Println("-----------------------------------------")
	}
}
func WebRequestReceiveMonitor(ch chan int,ServerNumber int){
	i := 0
	TimeOut,_ := strconv.Atoi(Config.ReStartTimeOut)
	for{
		select {
		case <-ch:
			return
		default:
			time.Sleep(time.Second * 1)
			i = i + 1
			if i > TimeOut{
				AllServer[ServerNumber].Restart()
			}
		}
	}
}
func Host2Domain_Port(Host string)(string,string){
	var Domain string
	var Port string
	if strings.Contains(Host,":"){
		arrData := strings.Split(Host,":")
		Domain = arrData[0]
		Port = arrData[1]
	} else {
		Domain = Host
		Port = "80"
	}
	return Domain,Port
}
func GetWebRoot(Domain string,Port string)(string,int){
	for i:=0;i<len(Config.AllWeb);i++{
		for j:=0;j<len(Config.AllWeb[i].Domain);j++{
			if Domain == Config.AllWeb[i].Domain[j] && Port == Config.AllWeb[i].Port{
				return Config.AllWeb[i].Root,i
			}
		}
	}
	return "NULL",-1
}
func PathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		PrintLog("检查文件存在:"+path)
		return true
	}
	if os.IsNotExist(err) {
		PrintLog("检查文件不存在:"+path)
		return false
	}
	PrintLog("检查文件不存在:"+path)
	return false
}
func GetResponse_File(path string)(string,[]byte){
	FileType := path[strings.LastIndex(path, "."):]
	_,ok := ContentTypeMap[FileType]
	if ok {
		return ContentTypeMap[FileType],ReadFile(path)
	} else {
		return "application/octet-stream",ReadFile(path)
	}
}
func HTTP404(WebRoot string) (string,[]byte) {
	if PathExists(WebRoot+"\\404.html"){
		return "text/html",ReadFile(WebRoot+"\\404.html")
	} else {
		return "text/html",[]byte("HTTP 404")
	}
}
func ReadFile(path string) ([]byte){
	var Data,err = ioutil.ReadFile(path)
	if err != nil{
		PrintLog("读取文件失败:"+path)
		return []byte("Read File Error")
	}
	PrintLog("读取文件:"+path)
	return Data
}
func isAsset(Path string,i int)bool{
	UserPath := strings.Split(Path,"\\")[1]
	if strings.Compare(UserPath,Config.AllWeb[i].Asset) == 0{
		return true
	} else {
		return false
	}
}
func WebRequestReceive(r *ghttp.Request) {
	ch := make(chan int)
	r.ParseForm()
	var DataKey string
	go WebRequestReceiveMonitor(ch,0)
	EnterDomain,EnterPort := Host2Domain_Port(r.Host)
	WebRoot,WebNumber := GetWebRoot(EnterDomain,EnterPort)
	if WebRoot == "NULL"{
		r.Response.WriteStatus(500)
		return
	}
	var EchoData []byte
	var ContentType string
	Flag := 0
	Path := r.URL.Path
	Path = strings.Replace(Path,"/","\\",-1)
	PrintLog("访问路径:"+Path)
	if strings.Index(Path,".") == -1{
		for i:=0;i<len(Config.AllWeb[WebNumber].Defaults);i++{
			if PathExists(WebRoot+Path+"\\"+Config.AllWeb[WebNumber].Defaults[i]){
				Path = Path+"\\"+Config.AllWeb[WebNumber].Defaults[i]
				Flag = 1
				break;
			}
		}
		if Flag == 0{
			Flag = -1
		}
	}
	if Flag >= 0{
		if PathExists(WebRoot+Path){
			if !isAsset(Path,WebNumber){
				FileType := strings.ToUpper((WebRoot+Path)[strings.LastIndex(WebRoot+Path, "."):])
				if FileType == ".SEHP"{
					hash := md5.New()
					hashStr := strconv.FormatInt(time.Now().Unix()+int64(rand.Intn(100))*int64(rand.Intn(10)),10)
					hash.Write([]byte(hashStr))
					cipherStr := hash.Sum(nil)
					DataKey = hex.EncodeToString(cipherStr)
					var NowData EnterData
					NowData.ID = DataKey
					NowData.r = r
					AllEnterData[DataKey] = NowData
					EchoData = execSEHP(WebNumber,DataKey,Path,r);
					ContentType = "text/html"
					PrintLog("发现SEHP,访问路径:"+Path)
				} else {
					ContentType,EchoData = GetResponse_File(WebRoot+Path)
				}
			} else {
				ContentType = "application/octet-stream"
				_,EchoData = GetResponse_File(WebRoot+Path)
			}
			PrintLog("HTTP 200:"+Path)
			r.Response.Status = 200
		} else {
			PrintLog("HTTP 404 错误:"+Path)
			r.Response.Status = 404
			ContentType,EchoData = HTTP404(WebRoot)
		}
	} else {
		PrintLog("HTTP 404 错误:"+Path)
		r.Response.Status = 404
		ContentType,EchoData = HTTP404(WebRoot)
	}
	r.Response.Header().Set("content-type",ContentType+";charset=utf-8")
	r.Response.Header().Set("content-length",strconv.Itoa(len(EchoData)))
	r.Response.Write(EchoData)
	delete(AllEnterData,DataKey)
	ch <- 1
}
func compressStr(str string) string {
    if str == "" {
        return ""
    }
    reg := regexp.MustCompile("\\s+")
    return reg.ReplaceAllString(str, "")
}
func CommunicationEngine(){
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("SEHP Communication Engine Working: <%s> \n", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		RecvStr := string(data[:n])
		RecvData := strings.Split(RecvStr,"|")
		PrintLog("拓展通讯:"+RecvStr)
			if strings.Compare(RecvData[0],"GetFrom") == 0{
				if _, ok := AllEnterData[RecvData[1]];ok{
					Temp := AllEnterData[RecvData[1]]
					if _,ok := Temp.r.Form[RecvData[2]];ok{
						ResultStr := Temp.r.Form[RecvData[2]][0]
						_, err = listener.WriteToUDP([]byte("GetFrom|"+ResultStr), remoteAddr)
						if err != nil {
							fmt.Printf(err.Error())
						}
					} else {
						_, err = listener.WriteToUDP([]byte("GetFrom|KEYNULL"), remoteAddr)
						if err != nil {
							fmt.Printf(err.Error())
						}
					}
				} else {
					_, err = listener.WriteToUDP([]byte("GetFrom|IDNULL"), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				}
			} else if strings.Compare(RecvData[0],"GetUserAgent") == 0{
				if _, ok := AllEnterData[RecvData[1]];ok{
					Temp := AllEnterData[RecvData[1]]
					_, err = listener.WriteToUDP([]byte("GetUserAgent|"+Temp.r.UserAgent()), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				} else {
					_, err = listener.WriteToUDP([]byte("GetUserAgent|NULL"), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				}
			} else if strings.Compare(RecvData[0],"GetMethod") == 0{
				if _, ok := AllEnterData[RecvData[1]];ok{
					Temp := AllEnterData[RecvData[1]]
					_, err = listener.WriteToUDP([]byte("GetMethod|"+Temp.r.Method), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				} else {
					_, err = listener.WriteToUDP([]byte("GetMethod|NULL"), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				}
			} else if strings.Compare(RecvData[0],"GetHost") == 0{
				if _, ok := AllEnterData[RecvData[1]];ok{
					Temp := AllEnterData[RecvData[1]]
					_, err = listener.WriteToUDP([]byte("GetHost|"+Temp.r.Host), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				} else {
					_, err = listener.WriteToUDP([]byte("GetHost|NULL"), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				}
			} else if strings.Compare(RecvData[0],"GetPath") == 0{
				if _, ok := AllEnterData[RecvData[1]];ok{
					Temp := AllEnterData[RecvData[1]]
					_, err = listener.WriteToUDP([]byte("GetPath|"+Temp.r.URL.Path), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				} else {
					_, err = listener.WriteToUDP([]byte("GetPath|NULL"), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				}
			} else if strings.Compare(RecvData[0],"GetCookies") == 0{
				if _, ok := AllEnterData[RecvData[1]];ok{
					Temp := AllEnterData[RecvData[1]]
					ResultStr := "NULL"
					for _,Value := range Temp.r.Cookies(){
						if strings.Compare(RecvData[2],Value.Name)==0{
							ResultStr = Value.Value
							break
						}
					}
					_, err = listener.WriteToUDP([]byte("GetCookies|"+ResultStr), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				} else {
					_, err = listener.WriteToUDP([]byte("GetCookies|IDNULL"), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				}
			} else if strings.Compare(RecvData[0],"Delete") == 0{
				if _, ok := AllEnterData[RecvData[1]];ok{
					delete(AllEnterData,RecvData[1])
				}
			} else if strings.Compare(RecvData[0],"GetSEPHWorkFolder") == 0{
				if _, ok := AllEnterData[RecvData[1]];ok{
					_, err = listener.WriteToUDP([]byte("GetSEPHWorkFolder|"+Config.SEPHWorkFolder), remoteAddr)
					if err != nil {
						fmt.Printf(err.Error())
					}
				}
			}
	}
}
func execSEHP(WebNumber int,DataKey string,Path string,r *ghttp.Request) []byte{
	ResultCode := ""
	CodeOriginal := strings.Split(string(ReadFile(Config.AllWeb[WebNumber].Root +"\\"+ Path)),"\n")
	SEHPCount := 0
	for i:=0;i<len(CodeOriginal);i++{
		if strings.HasPrefix(CodeOriginal[i],"#!"){
			SetCode := compressStr(CodeOriginal[i][strings.Index(CodeOriginal[i],"#!")+2:strings.LastIndex(CodeOriginal[i],"!#")])
			CommendLine := strings.Split(SetCode,"=")
			if CommendLine[0]=="字符集"{
				ResultCode += "<meta charset=\""+CommendLine[1]+"\">"
			}
		} else if strings.HasPrefix(CodeOriginal[i],"#"){
		} else {
			if strings.Contains(CodeOriginal[i], "<#sehp"){
				exp := regexp.MustCompile("language=\"(.*?)\"")
				Langugae := strings.Replace(strings.Replace(exp.FindString(CodeOriginal[i]),"\"","",-1),"language=","",-1)
				if strings.Contains(CodeOriginal[i], "src="){
					exp = regexp.MustCompile("src=\"(.*?)\"")
					Source := strings.Replace(strings.Replace(exp.FindString(CodeOriginal[i]),"\"","",-1),"src=","",-1)
					for !strings.Contains(CodeOriginal[i], "#>"){
						i++
					}
					ResultCode += SEHPEngine(WebNumber,DataKey,Path,SEHPCount,Langugae,Source,nil,r)
				} else {
					i++
					var CHPXCode []string
					for !strings.Contains(CodeOriginal[i], "#>"){
						CHPXCode = append(CHPXCode,CodeOriginal[i]+"\n")
						i++
					}
					ResultCode += SEHPEngine(WebNumber,DataKey,Path,SEHPCount,Langugae,"NULL",CHPXCode,r)
				}
				SEHPCount++
			} else {
				ResultCode += CodeOriginal[i] + "\n"
			}
		}
	}
	return []byte(ResultCode)
}
func SEHPEngine(WebNumber int,DataKey string,Path string,Count int,Language string,Source string,Code []string,r *ghttp.Request) string{
	ResultCode := ""
	Path = strings.Split(strings.Replace(Path,"\\","-",-1),".")[0] + "-"+ strconv.Itoa(Count)
	if !PathExists(Config.SEPHWorkFolder+"\\Executable\\"+Path +".exe"){
		var CodeStr string
		for i:=0;i<len(Code);i++{
			CodeStr = CodeStr + Code[i]
		}
		err := ioutil.WriteFile(Config.SEPHWorkFolder+"\\Source\\"+Path+".cpp",[]byte(CodeStr), 0644)
		if err != nil{
			fmt.Println("Error:",err)
		}
		PrintLog("写出文件:"+Config.SEPHWorkFolder+"\\Source\\"+Path+".cpp")
		Param1 := make([]string,0)
		Param1 = append(Param1,"-c")
		Param1 = append(Param1,"-o")
		Param1 = append(Param1,Config.SEPHWorkFolder+"\\CompilerTemp\\"+Path+".o")
		Param1 = append(Param1,Config.SEPHWorkFolder+"\\Source\\"+Path+".cpp")
		Param1 = append(Param1,"-lws2_32")
		cmd := exec.Command(Config.SEPHWorkFolder+"\\Compiler\\MinGW64\\bin\\"+"g++",Param1...)
		cmd.Start()
		cmd.Wait()
		Param2 := make([]string,0)
		Param2 = append(Param2,Config.SEPHWorkFolder+"\\CompilerTemp\\sehp.o")
		Param2 = append(Param2,Config.SEPHWorkFolder+"\\CompilerTemp\\"+Path+".o")
		Param2 = append(Param2,"-o")
		Param2 = append(Param2,Config.SEPHWorkFolder+"\\Executable\\"+Path+".exe")
		Param2 = append(Param2,"-lws2_32")
		cmd2 := exec.Command(Config.SEPHWorkFolder+"\\Compiler\\MinGW64\\bin\\"+"g++",Param2...)
		cmd2.Start()
		cmd2.Wait()
		PrintLog("执行编译:"+Config.SEPHWorkFolder+"\\Executable\\"+Path+".exe")
	}
	var command string
	if Language == "C++"{
		command = Config.SEPHWorkFolder+"\\Executable\\"+Path +".exe"
	} else if Language == "Java"{
		command = Config.SEPHWorkFolder+"\\Executable\\"+Path +".exe"
	}
	params := make([]string,0)
	params = append(params,DataKey)
	ResultCodeArr := execCommand(command, params)
	for i:=0;i<len(ResultCodeArr);i++{
		ResultCode = ResultCode + ResultCodeArr[i] + "\n"
	}
	return ResultCode
}
func execCommand(commandName string, params []string) []string {
	cmd := exec.Command(commandName, params...)
	ResultCode := make([]string,0)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		ResultCode = append(ResultCode,string("执行Command出错"))
		return ResultCode
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		dec := mahonia.NewDecoder("UTF-8")
		ret:=dec.ConvertString(line)
		ResultCode = append(ResultCode,ret)
	}
	cmd.Wait()
	return ResultCode
}
func padding(src []byte,blocksize int) []byte {
	padnum:=blocksize-len(src)%blocksize
	pad:=bytes.Repeat([]byte{byte(padnum)},padnum)
	return append(src,pad...)
}
func unpadding(src []byte) []byte {
	n:=len(src)
	unpadnum:=int(src[n-1])
	return src[:n-unpadnum]
}
func encryptAES(src []byte,key []byte) []byte {
	block,_:=aes.NewCipher(key)
	src=padding(src,block.BlockSize())
	blockmode:=cipher.NewCBCEncrypter(block,key)
	blockmode.CryptBlocks(src,src)
	return src
}
func decryptAES(src []byte,key []byte) []byte {
	block,_:=aes.NewCipher(key)
	blockmode:=cipher.NewCBCDecrypter(block,key)
	blockmode.CryptBlocks(src,src)
	src=unpadding(src)
	return src
}