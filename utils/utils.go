package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
	"path"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/basic/config"
)

const (
	RECODE_OK         = "0"
	RECODE_DBERR      = "4001"
	RECODE_NODATA     = "4002"
	RECODE_DATAEXIST  = "4003"
	RECODE_DATAERR    = "4004"
	RECODE_SESSIONERR = "4101"
	RECODE_LOGINERR   = "4102"
	RECODE_PARAMERR   = "4103"
	RECODE_USERERR    = "4104"
	RECODE_ROLEERR    = "4105"
	RECODE_PWDERR     = "4106"
	RECODE_SMSERR     = "4017"
	RECODE_REQERR     = "4201"
	RECODE_IPERR      = "4202"
	RECODE_THIRDERR   = "4301"
	RECODE_IOERR      = "4302"
	RECODE_SERVERERR  = "4500"
	RECODE_UNKNOWERR  = "4501"
)

var recodeText = map[string]string{
	RECODE_OK:         "成功",
	RECODE_DBERR:      "数据库查询错误",
	RECODE_NODATA:     "无数据",
	RECODE_DATAEXIST:  "数据已存在",
	RECODE_DATAERR:    "数据错误",
	RECODE_SESSIONERR: "用户未登录",
	RECODE_LOGINERR:   "用户登录失败",
	RECODE_PARAMERR:   "参数错误",
	RECODE_USERERR:    "用户不存在或未激活",
	RECODE_ROLEERR:    "用户身份错误",
	RECODE_PWDERR:     "密码错误",
	RECODE_REQERR:     "非法请求或请求次数受限",
	RECODE_IPERR:      "IP受限",
	RECODE_THIRDERR:   "第三方系统错误",
	RECODE_IOERR:      "文件读写错误",
	RECODE_SERVERERR:  "内部错误",
	RECODE_UNKNOWERR:  "未知错误",
	RECODE_SMSERR:     "验证码错误",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
func AddDomain2Url(url string) (domain_url string) {
	domain_url = "http://" + config.GetMisconfig().GetImageAddr() + "/" + url

	return domain_url
}
func Sha256Encode(value string) string {
	encoder := sha256.New()
	encoder.Write([]byte(value))
	hash := encoder.Sum(nil)
	result := hex.EncodeToString(hash)
	return string(result)
}
func SendEmail(emailTo string, code string) error {
	defer func() {
		if err := recover(); err != nil {
			log.Info("SendEmail Failed")
		} else {
			log.Info("SendEmail done")
		}
	}()
	config := `{"username":"` + config.GetMisconfig().GetMailUser() + `","password":"` + config.GetMisconfig().GetMailPass() + `","host":"smtp.163.com","port":25}`
	log.Info("SendEmail Config:", config)
	temail := NewEMail(config)
	temail.To = []string{emailTo}
	temail.From = "Ligthning"
	temail.Subject = "Lightning Register Code"
	temail.HTML = "Welcome to Lightning, your register code is:" + code
	_ = temail.Send()
	//err := temail.Send()
	return nil //TODO: use right sending mail account
	//return err
}

func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}
func UploadByBuffer(filebuffer []byte, fileExt string) (string, error) {
	//TODO: we do nothing currenlty to upload the avastart to picture server
	log.Info("UploadAvastar do nothing currently!")
	return "DummyFile", nil
	var (
		err        error
		sftpClient *sftp.Client
	)

	sftpClient, err = connect()
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	var remoteDir = "/home/vsftpd/sher/"
	var fileName = Md5String(time.Now().String())
	var remoteFileName = fileName + "." + fileExt
	fmt.Println("remoteFileName:", remoteFileName)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()
	dstFile.Write(filebuffer)
	return remoteFileName, nil
}

func connect() (*sftp.Client, error) {
	var (
		sftpClient *sftp.Client
		err        error
	)
	pemBytes, err := ioutil.ReadFile("../conf/id_rsa")
	if err != nil {
		log.Fatal(err)
	}
	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		log.Fatalf("parse key failed:%v", err)
	}
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 30 * time.Second,
	}
	ssh_addr := "127.0.0.1" // ssh server ip
	conn, err := ssh.Dial("tcp", ssh_addr+":22", config)
	if err != nil {
		log.Fatalf("dial failed:%v", err)
	}
	if sftpClient, err = sftp.NewClient(conn); err != nil {
		return nil, err
	}
	return sftpClient, nil
}
