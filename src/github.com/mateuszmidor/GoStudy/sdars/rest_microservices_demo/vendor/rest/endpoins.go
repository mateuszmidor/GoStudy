package rest

type HttpResource = string

const HwAddr string = "localhost:8080"
const HwCurrentStation HttpResource = "/CurrentStation/"

func MakeHwEndpoint(resource HttpResource) string {
	return "http://" + HwAddr + string(resource)
}

const UIAddr string = "localhost:8082"
const UISubscription HttpResource = "/Subscription/"
const UIStations HttpResource = "/Stations/"

func MakeUIEndpoint(resource HttpResource) string {
	return "http://" + UIAddr + string(resource)
}

const TunerAddr string = "localhost:8081"
const TunerCurrentStations HttpResource = "/CurrentStation/"
const TunerStationList HttpResource = "/Stationlist/"
const TunerSubscription HttpResource = "/Subscrption/"

func MakeTunerEndpoint(resource HttpResource) string {
	return "http://" + TunerAddr + string(resource)
}
