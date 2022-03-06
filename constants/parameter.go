package constants

type MetObsParameter string

const (
	MTempDry               MetObsParameter = "temp_dry"
	MTempDew               MetObsParameter = "temp_dew"
	MTempMeanPast1h        MetObsParameter = "temp_mean_past1h"
	MTempMaxPast1h         MetObsParameter = "temp_max_past1h"
	MTempMinPast1h         MetObsParameter = "temp_min_past1h"
	MTempMaxPast12h        MetObsParameter = "temp_max_past12h"
	MTempMinPast12h        MetObsParameter = "temp_min_past12h"
	MTempGrassMaxPast1h    MetObsParameter = "temp_grass_max_past1h"
	MTempGrassMeanPast1h   MetObsParameter = "temp_grass_mean_past1h"
	MTempGrassMinPast1h    MetObsParameter = "temp_grass_min_past1h"
	MTempSoil              MetObsParameter = "temp_soil"
	MTempGrass             MetObsParameter = "temp_grass"
	MTempSoilMaxPast1h     MetObsParameter = "temp_soil_max_past1h"
	MTempSoilMeanPast1h    MetObsParameter = "temp_soil_mean_past1h"
	MTempSoilMinPast1h     MetObsParameter = "temp_soil_min_past1h"
	MHumidity              MetObsParameter = "humidity"
	MHumidityPast1h        MetObsParameter = "humidity_past1h"
	MPressure              MetObsParameter = "pressure"
	MPressureAtSea         MetObsParameter = "pressure_at_sea"
	MWindDir               MetObsParameter = "wind_dir"
	MWindDirPast1h         MetObsParameter = "wind_dir_past1h"
	MWindSpeed             MetObsParameter = "wind_speed"
	MWindSpeedPast1h       MetObsParameter = "wind_speed_past1h"
	MWindGustAlwaysPast1h  MetObsParameter = "wind_gust_always_past1h"
	MWindMax               MetObsParameter = "wind_max"
	MWindMinPast1h         MetObsParameter = "wind_min_past1h"
	MWindMin               MetObsParameter = "wind_min"
	MWindMaxPer10minPast1h MetObsParameter = "wind_max_per10min_past1h"
	MPrecipPast1h          MetObsParameter = "precip_past1h"
	MPrecipPast10min       MetObsParameter = "precip_past10min"
	MPrecipPast1min        MetObsParameter = "precip_past1min"
	MPrecipPast24h         MetObsParameter = "precip_past24h"
	MPrecipDurPast10min    MetObsParameter = "precip_dur_past10min"
	MPrecipDurPast1h       MetObsParameter = "precip_dur_past1h"
	MSnowDepthMan          MetObsParameter = "snow_depth_man"
	MSnowCoverMan          MetObsParameter = "snow_cover_man"
	MVisibility            MetObsParameter = "visibility"
	MVisibMeanLast10min    MetObsParameter = "visib_mean_last10min"
	MCloudCover            MetObsParameter = "cloud_cover"
	MCloudHeight           MetObsParameter = "cloud_height"
	MWeather               MetObsParameter = "weather"
	MRadiaGlob             MetObsParameter = "radia_glob"
	MRadiaGlobPast1h       MetObsParameter = "radia_glob_past1h"
	MSunLast10minGlob      MetObsParameter = "sun_last10min_glob"
	MSunLast1hGlob         MetObsParameter = "sun_last1h_glob"
	MLeavHumDurPast10min   MetObsParameter = "leav_hum_dur_past10min"
	MLeavHumDurPast1h      MetObsParameter = "leav_hum_dur_past1h"
)

var MetObsParameters = []MetObsParameter{
	MTempDry,
	MTempDew,
	MTempMeanPast1h,
	MTempMaxPast1h,
	MTempMinPast1h,
	MTempMaxPast12h,
	MTempMinPast12h,
	MTempGrassMaxPast1h,
	MTempGrassMeanPast1h,
	MTempGrassMinPast1h,
	MTempSoil,
	MTempGrass,
	MTempSoilMaxPast1h,
	MTempSoilMeanPast1h,
	MTempSoilMinPast1h,
	MHumidity,
	MHumidityPast1h,
	MPressure,
	MPressureAtSea,
	MWindDir,
	MWindDirPast1h,
	MWindSpeed,
	MWindSpeedPast1h,
	MWindGustAlwaysPast1h,
	MWindMax,
	MWindMinPast1h,
	MWindMin,
	MWindMaxPer10minPast1h,
	MPrecipPast1h,
	MPrecipPast10min,
	MPrecipPast1min,
	MPrecipPast24h,
	MPrecipDurPast10min,
	MPrecipDurPast1h,
	MSnowDepthMan,
	MSnowCoverMan,
	MVisibility,
	MVisibMeanLast10min,
	MCloudCover,
	MCloudHeight,
	MWeather,
	MRadiaGlob,
	MRadiaGlobPast1h,
	MSunLast10minGlob,
	MSunLast1hGlob,
	MLeavHumDurPast10min,
	MLeavHumDurPast1h,
}
