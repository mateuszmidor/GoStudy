package spconfig

import (
	"log"

	"github.com/joeshaw/envdecode"
)

// SpConfig represents Social Poll configuration loaded from env variables
type SpConfig struct {
	ConsumerKey    string `env:"SP_TWITTER_KEY,required"`
	ConsumerSecret string `env:"SP_TWITTER_SECRET,required"`
	AccessToken    string `env:"SP_TWITTER_ACCESSTOKEN,required"`
	AccessSecret   string `env:"SP_TWITTER_ACCESSSECRET,required"`
	MongoDbAddress string `env:"SP_MONGODB_ADDR,required"`
}

// GetConfig returns SocialPoll configuration or terminates app if no config found
func GetConfig() SpConfig {
	var config SpConfig
	if err := envdecode.Decode(&config); err != nil {
		log.Fatalln(err)
	}
	return config
}
