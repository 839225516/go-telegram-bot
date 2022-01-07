package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// TgConfig 系统配置
type TgConfig struct {
	TgToken                string `json:"tg_token"`
	TgProxy                string `json:"tg_proxy"`
	PromptMsgAfterDelTime  int    `json:"prompt_msg_after_del_time"`
	CaptchaMsgAfterDelTime int    `json:"captcha_msg_after_del_time"`
	CaptchaTimeOut         int    `json:"captcha_time_out"`
	PromptMsgTpl           string `json:"prompt_msg_tpl"`
	CaptchaMsgTpl          string `json:"captcha_msg_tpl"`
	CaptchaSuccessMsgTpl   string `json:"captcha_success_msg_tpl"`
	RuntimeRootPath        string `json:"runtime_root_path"`
	LogSavePath            string `json:"log_save_path"`
	VerifyImgPath          string `json:"verify_img_path"`
}

var TgConf TgConfig

// init 配置加载
func init() {
	//viper.AddConfigPath("./conf")
	//viper.SetConfigType("yaml")
	viper.SetConfigFile("./conf/bot.yaml")
	err := viper.ReadInConfig()

	fmt.Println(viper.ConfigFileUsed())

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&TgConf)
	if err != nil {
		panic(fmt.Errorf("Unmarshal error config file: %s \n", err))
	}
}
