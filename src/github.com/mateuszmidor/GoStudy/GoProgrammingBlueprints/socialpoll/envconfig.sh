# export MongoDB & NSQ addresses that are set in docker-compose.yaml
export SP_MONGODB_ADDR=mongodb              # twittervotes, counter & api connect here
export SP_NSQD_ADDR=nsqd:4150               # twittervotes connects here (producer)
export SP_NSQLOOKUP_ADDR=nsqlookupd:4161    # counter connects here (consumer)

# export twitter api
export SP_TWITTER_KEY=LMxhXyLvOnXoRrJg7MA9pEizi
export SP_TWITTER_SECRET=WauAaw6Yf8JM7sS3mjvx3eRBzPRIjKdIZrYXFLkoddUNmz6ksV
export SP_TWITTER_ACCESSTOKEN=1081938944877580288-Cf4UbkmzgTYeRLp37qv8uMa27IVROe
export SP_TWITTER_ACCESSSECRET=66zCuYeWgxfv6y84JWXS47YwN2dfPa8a08x2PdA1J5Uru