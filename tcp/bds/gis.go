package bds
//import (
//	"math"
//	"fmt"
//)
//
//
//const (
//	a float64	  = 6378245.0
//	ee float64 	  = 0.00669342162296594323
//	x_pi float64  = 3.14159265358979324 * 3000.0 / 180.0
//)
//
//var pi		= math.Pi
//var sin		= math.Sin
//var cos		= math.Cos
//var sqrt	= math.Sqrt
//var abs 	= math.Abs
//var atan2 	= math.Atan2
//
//func isOutOfChina(lat, lon float64){
//	if lon < 72.004 || lon > 137.8347{
//		return true;
//	}
//	if lat < 0.8293 || lat > 55.8271{
//		return true;
//	}
//	return false;
//}
//
//func transformLat(x, y float64)float64{
//	var ret int64
//	ret = -100.0 + 2.0 * x + 3.0 * y + 0.2 * y * y + 0.1 * x * y + 0.2 * sqrt(abs(x));
//	ret += (20.0 * sin(6.0 * x * pi) + 20.0 * sin(2.0 * x * pi)) * 2.0 / 3.0;
//	ret += (20.0 * sin(y * pi) + 40.0 * sin(y / 3.0 * pi)) * 2.0 / 3.0;
//	ret += (160.0 * sin(y / 12.0 * pi) + 320 * sin(y * pi / 30.0)) * 2.0 / 3.0;
//	return ret;
//}
//
//func transformLon(x, y float64)float64{
//	var ret float64
//	ret = 300.0 + x + 2.0 * y + 0.1 * x * x + 0.1 * x * y + 0.1 * sqrt(abs(x));
//	ret += (20.0 * sin(6.0 * x * pi) + 20.0 * sin(2.0 * x * pi)) * 2.0 / 3.0;
//	ret += (20.0 * sin(x * pi) + 40.0 * sin(x / 3.0 * pi)) * 2.0 / 3.0;
//	ret += (150.0 * sin(x / 12.0 * pi) + 300.0 * sin(x / 30.0 * pi)) * 2.0 / 3.0;
//	return ret;
//}
//
//// World Geodetic System ==> Mars Geodetic System
//func Wgs2Mgs(wgLat, wgLon float64)(mgLat, mgLon float64){
//	if isOutOfChina(wgLat, wgLon){
//		mgLat  = wgLat;
//		mgLon = wgLon;
//		return ;
//	}
//	dLat := transformLat(wgLon - 105.0, wgLat - 35.0);
//	dLon := transformLon(wgLon - 105.0, wgLat - 35.0);
//	radLat := wgLat / 180.0 * math.Pi;
//	magic := sin(radLat);
//	magic = 1 - ee * magic * magic;
//	sqrtMagic := sqrt(magic);
//	dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * pi);
//	dLon = (dLon * 180.0) / (a / sqrtMagic * cos(radLat) * pi);
//	mgLat = wgLat + dLat;
//	mgLon = wgLon + dLon;
//	return
//}
//
//// Mars Geodetic System ==> Baidu Geodetic System
//func Mgs2Bgs(mgLat, mgLon int64)(bgLat, bgLon int64){
//	x := mgLon
//	y := mgLat
// 	z := sqrt(x * x + y * y) + 0.00002 * sin(y * x_pi);
//	theta := atan2(y, x) + 0.000003 * cos(x * x_pi);
//	bgLon = z * cos(theta) + 0.0065;
//	bgLat = z * sin(theta) + 0.006;
//	return
//}
//
//func Bgs2Mgs(bLat, bLon int64)(mgLat, mgLon int64){
//	x := bLon - 0.0065
//	y := bLat - 0.006
//	z := sqrt(x * x + y * y) - 0.00002 * sin(y * x_pi);
//	theta := atan2(y, x) - 0.000003 * cos(x * x_pi);
//	mgLon = z * cos(theta);
//	mgLat = z * sin(theta);
//	return
//}
//
//func main() {
//	var lat int64 = 30.227607
//	var lon int64 = 120.036565
//	mgLat, mgLon := Wgs2Mgs(lat, lon)
//	fmt.Println(mgLat, ",", mgLon)
//	bgLat, bgLon := Mgs2Bgs(mgLat, mgLon)
//	fmt.Println(bgLat, ",", bgLon)
//}