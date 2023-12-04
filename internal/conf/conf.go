package conf

import (
	"log"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

const (
	ConfCurrentPath string = "."
	ConfEtcPath     string = "/etc/"
	ConfFileSuffix  string = "rc"
	ConfFileType    string = "json"
)

// TODO: try a map instead of struct
type Configuration struct {
	AppShowStats           bool
	AppBackendApi          string
	AppBackendTrust        string
	AppBackendKit          string
	AppClientId            string
	KitCountPackets        bool
	KitTrustIdsPacketRate  int
	KitTrustIdsThroughput  int
	AmqpHost               string
	AmqpPort               int
	AmqpUsername           string
	AmqpPassword           string
	AmqpUseTLS             bool
	AmqpExchange           string
	AmqpContentType        string
	AmqpWarnTopic          string
	AmqpTrustTopic         string
	AmqpMtdTopic           string
	AmqpCACertFile         string
	AmqpClientCertFile     string
	AmqpClientKeyFile      string
	PcapBufSize            int
	PcapSnapLen            int
	PcapDirection          int // 0: in, 1:out, 2:inout
	PcapFilter             string
	PcapReplay             bool
	TrustThresholdsLow     int
	TrustThresholdsNeutral int
	TrustThresholdsMax     int
	TrustThresholdsPenalty int
	TrustThresholdsReward  int
	LogType                string
	LogFile                string
	LogLevel               string
	LogUseTag              bool
	LogUseColor            bool
	keyBytes               []byte
}

func (c Configuration) KeyBytes() []byte {
	return c.keyBytes
}

func New(name string) *Configuration {
	// Initial Configuration
	confDefaults()

	// Configuration File
	// viper module is meant to only be used locally to this file!
	confFilename := defaultName(name)
	viper.AddConfigPath(ConfCurrentPath)
	viper.AddConfigPath(ConfEtcPath + name)
	viper.SetConfigName(name + ConfFileSuffix)
	viper.SetConfigType(ConfFileType)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("config file '%s' was not found\n", confFilename)
			log.Printf("using default config values")
		} else {
			log.Printf("error parsing config file '%s': '%s'\n", confFilename, err)
			log.Printf("using default config values")
		}
	} else {
		log.Printf("using config file '%s'\n", viper.ConfigFileUsed())
	}

	// Environment Variables
	log.Printf("loading configuration from env")
	viper.SetEnvPrefix(name)
	viper.AutomaticEnv()

	// From this point further, only cli options can overwrite configuration
	// values. TODO: sanitize all user input
	return Update()
}

func NewDefaults() *Configuration {
	confDefaults()
	return Update()
}

func Update() *Configuration {
	return &Configuration{
		AppShowStats:           viper.GetBool("app_show_stats"),
		AppBackendApi:          onlyAlpha(viper.GetString("app_backend_api"), 32),
		AppBackendTrust:        viper.GetString("app_backend_trust"),
		AppBackendKit:          viper.GetString("app_backend_kit"),
		AppClientId:            viper.GetString("app_client_id"),
		KitCountPackets:        viper.GetBool("kit_count_packet"),
		KitTrustIdsPacketRate:  viper.GetInt("kit_trustIds_packet_rate"),
		KitTrustIdsThroughput:  viper.GetInt("kit_trustIds_throughput"),
		AmqpHost:               viper.GetString("amqp_host"),
		AmqpPort:               viper.GetInt("amqp_port"),
		AmqpUsername:           viper.GetString("amqp_username"),
		AmqpPassword:           viper.GetString("amqp_password"),
		AmqpUseTLS:             viper.GetBool("amqp_use_tls"),
		AmqpExchange:           viper.GetString("amqp_exchange"),
		AmqpContentType:        viper.GetString("amqp_content_type"),
		AmqpWarnTopic:          viper.GetString("amqp_warn_topic"),
		AmqpTrustTopic:         viper.GetString("amqp_trust_topic"),
		AmqpMtdTopic:           viper.GetString("amqp_mtd_topic"),
		AmqpCACertFile:         viper.GetString("amqp_ca_cert"),
		AmqpClientCertFile:     viper.GetString("amqp_client_cert"),
		AmqpClientKeyFile:      viper.GetString("amqp_client_key"),
		PcapBufSize:            viper.GetInt("pcap_buf_size"),
		PcapSnapLen:            viper.GetInt("pcap_snap_len"),
		PcapDirection:          viper.GetInt("pcap_direction"),
		PcapFilter:             viper.GetString("pcap_filter"),
		PcapReplay:             viper.GetBool("pcap_replay"),
		TrustThresholdsLow:     viper.GetInt("trust_thresholds_low"),
		TrustThresholdsNeutral: viper.GetInt("trust_thresholds_neutral"),
		TrustThresholdsMax:     viper.GetInt("trust_thresholds_max"),
		TrustThresholdsPenalty: viper.GetInt("trust_thresholds_penalty"),
		TrustThresholdsReward:  viper.GetInt("trust_thresholds_reward"),
		LogType:                viper.GetString("log_type"),
		LogFile:                viper.GetString("log_file"),
		LogLevel:               viper.GetString("log_level"),
		LogUseTag:              viper.GetBool("log_use_tag"),
		LogUseColor:            viper.GetBool("log_use_color"),
	}
}

func Save(appname string) string {
	filename := defaultName(appname)
	err := viper.SafeWriteConfigAs(filename)
	if err != nil {
		log.Fatalln(err)
	}

	return filename
}

func defaultName(name string) string {
	return name + ConfFileSuffix + "." + ConfFileType
}

func confDefaults() {
	viper.SetDefault("app_show_stats", false)
	viper.SetDefault("app_backend_api", "log")
	viper.SetDefault("app_backend_trust", "thresholds")
	viper.SetDefault("app_backend_kit", "trustIds")
	viper.SetDefault("app_client_id", "device1") // device1

	viper.SetDefault("kit_count_packets", false)
	viper.SetDefault("kit_trustIds_packet_rate", 10000)
	viper.SetDefault("kit_trustIds_throughput", 1000000)

	viper.SetDefault("amqp_host", "localhost")
	viper.SetDefault("amqp_port", "5672")
	viper.SetDefault("amqp_username", "client") // client client2
	viper.SetDefault("amqp_password", "client") // client client2
	viper.SetDefault("amqp_use_tls", false)
	viper.SetDefault("amqp_exchange", "amq.topic")
	viper.SetDefault("amqp_content_type", "text/plain")
	viper.SetDefault("amqp_warn_topic", "ids.warn.client_ids")
	viper.SetDefault("amqp_trust_topic", "ids.trust.client_ids")
	viper.SetDefault("amqp_mtd_topic", "mtd.warnings.client_ids")
	viper.SetDefault("amqp_ca_cert", "")
	viper.SetDefault("amqp_client_cert", "")
	viper.SetDefault("amqp_client_key", "")

	viper.SetDefault("pcap_buf_size", 1024)
	viper.SetDefault("pcap_snap_len", 128)
	viper.SetDefault("pcap_direction", 0)
	viper.SetDefault("pcap_filter", "tcp")
	viper.SetDefault("pcap_replay", false)

	viper.SetDefault("trust_thresholds_low", 25)
	viper.SetDefault("trust_thresholds_neutral", 50)
	viper.SetDefault("trust_thresholds_max", 100)
	viper.SetDefault("trust_thresholds_penalty", 5)
	viper.SetDefault("trust_thresholds_reward", 1)

	viper.SetDefault("log_type", "syslog") // simple, stdout, stderr, syslog, file
	viper.SetDefault("log_file", "")
	viper.SetDefault("log_level", "info")
	viper.SetDefault("log_use_tag", true)
	viper.SetDefault("log_use_color", true)
}

func firstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

func onlyAlpha(str string, length int) string {
	var nonAlpha = regexp.MustCompile(`[^a-zA-Z]+`)
	san := nonAlpha.ReplaceAllString(str, "")

	return firstN(strings.ToLower(san), length)
}
