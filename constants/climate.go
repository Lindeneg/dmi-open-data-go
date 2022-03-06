package constants

type ClimateDataParameter string

const (
	CMeanTemp                  ClimateDataParameter = "mean_temp"
	CMeanDailyMaxTemp          ClimateDataParameter = "mean_daily_max_temp"
	CMaxTempWDate              ClimateDataParameter = "max_temp_w_date"
	CMaxTemp_12h               ClimateDataParameter = "max_temp_12h"
	CNoIceDays                 ClimateDataParameter = "no_ice_days"
	CNoSummerDays              ClimateDataParameter = "no_summer_days"
	CMeanDailyMinTemp          ClimateDataParameter = "mean_daily_min_temp"
	CMinTemp                   ClimateDataParameter = "min_temp"
	CTempGrass                 ClimateDataParameter = "temp_grass"
	CMinTemperature_12h        ClimateDataParameter = "min_temperature_12h"
	CNoColdDays                ClimateDataParameter = "no_cold_days"
	CNoFrostDays               ClimateDataParameter = "no_frost_days"
	CNoTropicalNights          ClimateDataParameter = "no_tropical_nights"
	CAccHeatingDegreeDays_17   ClimateDataParameter = "acc_heating_degree_days_17"
	CAccHeatingDegreeDays_19   ClimateDataParameter = "acc_heating_degree_days_19"
	CMeanRelativeHum           ClimateDataParameter = "mean_relative_hum"
	CMaxRelativeHum            ClimateDataParameter = "max_relative_hum"
	CMinRelativeHum            ClimateDataParameter = "min_relative_hum"
	CMeanVapourPressure        ClimateDataParameter = "mean_vapour_pressure"
	CMeanWindSpeed             ClimateDataParameter = "mean_wind_speed"
	CMaxWindSpeed_10min        ClimateDataParameter = "max_wind_speed_10min"
	CMaxWindSpeed_3sec         ClimateDataParameter = "max_wind_speed_3sec"
	CNoWindyDays               ClimateDataParameter = "no_windy_days"
	CNoStormyDays              ClimateDataParameter = "no_stormy_days"
	CNoDaysWStorm              ClimateDataParameter = "no_days_w_storm"
	CNoDaysWHurricane          ClimateDataParameter = "no_days_w_hurricane"
	CMeanWindDirMin0           ClimateDataParameter = "mean_wind_dir_min0"
	CMeanWindDir               ClimateDataParameter = "mean_wind_dir"
	CMeanPressure              ClimateDataParameter = "mean_pressure"
	CMaxPressure               ClimateDataParameter = "max_pressure"
	CMinPressure               ClimateDataParameter = "min_pressure"
	CBrightSunshine            ClimateDataParameter = "bright_sunshine"
	CAccPrecip                 ClimateDataParameter = "acc_precip"
	CMaxPrecip_24h             ClimateDataParameter = "max_precip_24h"
	CAccPrecipPast12h          ClimateDataParameter = "acc_precip_past12h"
	CNoDaysAccPrecip_01        ClimateDataParameter = "no_days_acc_precip_01"
	CNoDaysAccPrecip_1         ClimateDataParameter = "no_days_acc_precip_1"
	CNoDaysAccPrecip_10        ClimateDataParameter = "no_days_acc_precip_10"
	CAccPrecipPast24h          ClimateDataParameter = "acc_precip_past24h"
	CMaxPrecip_30m             ClimateDataParameter = "max_precip_30m"
	CNoDaysSnowCover           ClimateDataParameter = "no_days_snow_cover"
	CMeanCloudCover            ClimateDataParameter = "mean_cloud_cover"
	CNoClearDays               ClimateDataParameter = "no_clear_days"
	CNoCloudyDays              ClimateDataParameter = "no_cloudy_days"
	CMaxSnowDepth              ClimateDataParameter = "max_snow_depth"
	CSnowDepth                 ClimateDataParameter = "snow_depth"
	CSnowCover                 ClimateDataParameter = "snow_cover"
	CTempSoil_10               ClimateDataParameter = "temp_soil_10"
	CTempSoil_30               ClimateDataParameter = "temp_soil_30"
	CLeafMoisture              ClimateDataParameter = "leaf_moisture"
	CVapourPressureDeficitMean ClimateDataParameter = "vapour_pressure_deficit_mean"
)

var ClimateDataParameters = []ClimateDataParameter{
	CMeanTemp,
	CMeanDailyMaxTemp,
	CMaxTempWDate,
	CMaxTemp_12h,
	CNoIceDays,
	CNoSummerDays,
	CMeanDailyMinTemp,
	CMinTemp,
	CTempGrass,
	CMinTemperature_12h,
	CNoColdDays,
	CNoFrostDays,
	CNoTropicalNights,
	CAccHeatingDegreeDays_17,
	CAccHeatingDegreeDays_19,
	CMeanRelativeHum,
	CMaxRelativeHum,
	CMinRelativeHum,
	CMeanVapourPressure,
	CMeanWindSpeed,
	CMaxWindSpeed_10min,
	CMaxWindSpeed_3sec,
	CNoWindyDays,
	CNoStormyDays,
	CNoDaysWStorm,
	CNoDaysWHurricane,
	CMeanWindDirMin0,
	CMeanWindDir,
	CMeanPressure,
	CMaxPressure,
	CMinPressure,
	CBrightSunshine,
	CAccPrecip,
	CMaxPrecip_24h,
	CAccPrecipPast12h,
	CNoDaysAccPrecip_01,
	CNoDaysAccPrecip_1,
	CNoDaysAccPrecip_10,
	CAccPrecipPast24h,
	CMaxPrecip_30m,
	CNoDaysSnowCover,
	CMeanCloudCover,
	CNoClearDays,
	CNoCloudyDays,
	CMaxSnowDepth,
	CSnowDepth,
	CSnowCover,
	CTempSoil_10,
	CTempSoil_30,
	CLeafMoisture,
	CVapourPressureDeficitMean,
}
