package geohash

const (
	BASE32   = "0123456789bcdefghjkmnpqrstuvwxyz"
	MAX_LAT  = 90
	MIN_LAT  = -90
	MAX_LONG = 180
	MIN_LONG = -180
)

type Box struct {
	MinLat, MaxLat   float64
	MinLong, MaxLong float64
}

func Encode(longitude, latitude float64, precision int) (string, Box) {
	length := 0
	var minLat, maxLat, minLong, maxLong float64
	minLat, maxLat, minLong, maxLong = MIN_LAT, MAX_LAT, MIN_LONG, MAX_LONG
	result := make([]byte, 0)
	var bit byte
	for len(result) < precision {
		midLong := (minLong + maxLong) / 2.0
		bit <<= 1
		if longitude > midLong {
			bit += 1
			minLong = midLong
		} else {
			maxLong = midLong
		}
		length++
		if length%5 == 0 {
			result = append(result, BASE32[bit])
			bit = 0
		}
		if len(result) >= precision {
			break
		}
		midLat := (minLat + maxLat) / 2.0
		bit <<= 1
		if latitude > midLat {
			minLat = midLat
			bit += 1
		} else {
			maxLat = midLat
		}
		length++
		if length%5 == 0 {
			result = append(result, BASE32[bit])
			bit = 0
		}
		if len(result) >= precision {
			break
		}
	}
	box := Box{
		MinLat:  minLat,
		MaxLat:  maxLat,
		MinLong: minLong,
		MaxLong: maxLong,
	}
	return string(result), box
}
