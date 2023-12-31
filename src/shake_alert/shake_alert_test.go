package shake_alert

import (
	"testing"
)

func TestOutlierDetection(t *testing.T) {

	data := []float64{1726, 1447, 1875, 3502, 3147, 1686, 2126, 2232, 1672, 1089, 1324, 1280, 1244, 1136, 1414, 1300, 1206, 2253, 2257, 1398, 1530, 1709, 1503, 2488, 1812, 1391, 1568, 1387, 1964, 3101, 1872, 2122, 3883, 7273, 2033, 3870, 2599, 1718, 1725, 1987, 1010, 1264, 1282, 1245, 1053, 1011, 896, 1351, 1574, 1016, 1182, 763, 1053, 1465, 3170, 3102, 1575, 1618, 1057, 1159, 1305, 1213, 1517, 1030, 1213, 1514, 1063, 2096, 1806, 2369, 1731, 2142, 2760, 2583, 1729, 1306, 1419, 4317, 5366, 2456, 1845, 3240, 1884, 1508, 1414, 1133, 1482, 2526, 2690, 1879, 3711, 2825, 1743, 2602, 1713, 1347, 1211, 1840, 1529, 1348, 1230, 3457, 1808, 2711, 4405, 2407, 1835, 1582, 1692, 1495, 1143, 946, 958, 2677, 1629, 1266, 1282, 2518, 3024, 1372, 1148, 1028, 1124, 1099, 1160, 991, 1079, 1484, 1748, 1438, 1355, 3274, 2381, 1293, 1584, 1412, 1396, 1510, 1517, 1731, 1625, 1759, 1892, 1740, 1299, 1514, 1596, 1762, 1070, 1072, 967, 842, 1449, 1447, 1832, 1785, 1008, 1511, 1804, 2583, 1150, 1084, 1359, 1154, 1070, 1595, 1625, 1988, 1821, 1188, 1379, 1072, 1235, 1152, 1461, 1495, 1155, 1220, 1397, 22601, 650958, 1337, 1408, 1162, 1547, 1541, 949, 1424, 1222, 1343, 2050, 1976, 1677, 1214, 2092, 2534, 1737, 2577, 1364, 1556, 1321, 1102, 1664, 1441, 1966, 1443, 1795, 1448, 1004, 1249, 1355, 1362, 1388, 949, 1468, 1362, 2182, 1906, 3288, 1516, 3353, 2668, 1546, 1808, 1832, 1506, 2020, 1491, 1512, 2496, 3729, 3189, 1288, 1249, 3131, 3618, 1402, 1309, 1583, 2066, 1526, 1764, 1528, 1401, 1460, 1384, 1564, 1815, 1631, 2726, 1306, 1645, 2049, 1696, 1533, 2074, 2392, 4392, 3979, 3756, 3254, 2358, 1385, 1631, 2120, 1420, 1669, 1617, 1507, 1657, 1613, 2404, 1710, 1931, 2174, 3191, 2460, 2151, 3726, 4635, 3080, 2812, 1681, 1717, 2186, 2765, 2668, 3544, 3095, 1480, 1193, 1631, 1651, 1985, 2002, 2194, 1953, 1345, 1310, 1625, 1577, 2099, 2006, 1525, 1438, 1136, 1323, 1593, 2113, 1344, 987, 2375, 2552, 1035, 1681, 1722, 1571, 1322, 1652, 2974, 2351, 1099, 1557, 1865, 1233, 2484, 1468, 1253, 1729, 1263, 1493, 2078, 1875, 6717, 3231, 3419, 2584, 2272, 2234, 1294, 1709, 1127, 2095, 1808, 1620, 1807, 3044, 2766, 2660, 2552, 2552, 3644, 1896, 3024, 1724, 1660, 2054, 1788, 1612, 1356, 1454, 1197, 1472, 1355, 1623, 1564, 1379, 1821, 1657, 968, 1637, 1659, 1776, 1365, 911, 1962, 3653, 5018, 3297, 3893, 4307, 2713, 1507, 3164, 2666, 3208, 1837, 1689, 1787, 1346, 4}

	actualOutliers := FindOutliers(data)
	expectedOutlier := []float64{22601, 650958}

	// Check if the arrays are equal
	isEqual := AreArraysEqual(actualOutliers, expectedOutlier)

	if !isEqual {
		t.Errorf("Got %+v want %+v\n", actualOutliers, expectedOutlier)
	}

}

// AreArraysEqual checks if two arrays of integers are equal
func AreArraysEqual(arr1, arr2 []float64) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func TestFindPercentile_5th(t *testing.T) {
	data := []float64{30, 20, 10, 40, 50, 45, 15, 35, 25}
	percentile := 5.0
	expected := 10.0
	result := findPercentile(data, percentile)
	if result != expected {
		t.Errorf("Expected %.2fth percentile to be %.2f, but got %.2f", percentile, expected, result)
	}
}

func TestFindPercentile_50th(t *testing.T) {
	data := []float64{30, 20, 10, 40, 50, 45, 15, 35, 25}
	percentile := 50.0
	expected := 30.0
	result := findPercentile(data, percentile)
	if result != expected {
		t.Errorf("Expected %.2fth percentile to be %.2f, but got %.2f", percentile, expected, result)
	}
}

func TestFindPercentile_95th(t *testing.T) {
	data := []float64{30, 20, 10, 40, 50, 45, 15, 35, 25}
	percentile := 95.0
	expected := 50.0
	result := findPercentile(data, percentile)
	if result != expected {
		t.Errorf("Expected %.2fth percentile to be %.2f, but got %.2f", percentile, expected, result)
	}
}

func TestFindPercentile_EmptyData(t *testing.T) {
	data := []float64{}
	percentile := 10.0
	expected := 0.0
	result := findPercentile(data, percentile)
	if result != expected {
		t.Errorf("Expected %.2fth percentile to be %.2f, but got %.2f", percentile, expected, result)
	}
}

func TestFindPercentile_OutOfBounds(t *testing.T) {
	data := []float64{30, 20, 10, 40, 50, 45, 15, 35, 25}
	percentile := 110.0 // Out of bounds, should return the maximum value in data.
	expected := 50.0
	result := findPercentile(data, percentile)
	if result != expected {
		t.Errorf("Expected %.2fth percentile to be %.2f, but got %.2f", percentile, expected, result)
	}
}
