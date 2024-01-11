package geohash

import (
	"fmt"
	"strconv"
	"strings"
)

// https://zhuanlan.zhihu.com/p/645078866
type GeoService struct {
	root *TrieNode
}
type TrieNode struct {
	childs [32]*TrieNode
	cnt    int
	Entry
}
type Entry struct {
	Points map[string]interface{}
	Hash   string
}
type Point struct {
	Lng, Lat float64
	Val      interface{}
}

func (e *Entry) GetPoints() []*Point {
	points := make([]*Point, 0, len(e.Points))
	for key, val := range e.Points {
		lng, lat := e.LngLat(key)
		points = append(points, &Point{Lng: lng, Lat: lat, Val: val})
	}
	return points
}
func (e *Entry) add(lng, lat float64, val interface{}) {
	if e.Points == nil {
		e.Points = map[string]interface{}{}
	}
	e.Points[e.pointKey(lng, lat)] = val
}
func (e *Entry) LngLat(key string) (lng float64, lat float64) {
	info := strings.Split(key, "_")
	lng, _ = strconv.ParseFloat(info[0], 64)
	lat, _ = strconv.ParseFloat(info[1], 64)
	return
}
func (e *Entry) pointKey(lng, lat float64) string {
	return fmt.Sprintf("%v_%v", lng, lat)
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}
func (g *GeoService) Add(lng, lat float64, val string, precision int) string {
	hashValue, _ := Encode(lng, lat, precision)
	target := g.get(hashValue)
	if target != nil {
		target.add(lng, lat, val)
		return hashValue
	}
	move := g.root
	for _, v := range hashValue {
		index := Base32ToIndex(byte(v))
		if move.childs[index] == nil {
			move.childs[index] = &TrieNode{}
		}
		move.cnt++
	}
	move.add(lng, lat, val)
	move.Hash = hashValue
	return hashValue
}

func (g *GeoService) get(prefixHash string) *TrieNode {
	move := g.root
	for i := 0; i < len(prefixHash); i++ {
		index := Base32ToIndex(prefixHash[i])
		if index == -1 || move.childs[index] == nil {
			return nil
		}
		move = move.childs[index]
	}
	return move
}

func (g *GeoService) GetEntry(hashValue string) (*Entry, bool) {
	target := g.get(hashValue)
	if target != nil {
		return nil, false
	}
	return &target.Entry, true
}
func (g *GeoService) Del(hashValue string) bool {
	target := g.get(hashValue)
	if target == nil || target.cnt == 0 {
		return false
	}
	move := g.root
	for i := 0; i < len(move.childs); i++ {
		index := Base32ToIndex(hashValue[i])
		move.childs[index].cnt--
		if move.childs[index].cnt == 0 {
			move.childs[index] = nil
			return true
		}
		move = move.childs[index]
	}
	return true
}

func (g *GeoService) ListByPrefix(prefix string) []*Entry {
	target := g.get(prefix)
	if target == nil {
		return nil
	}
	return Dfs(target)
}
func Dfs(root *TrieNode) []*Entry {
	var entries []*Entry
	for i := 0; i < len(root.childs); i++ {
		if root.childs[i] == nil {
			continue
		}
		entries = append(entries, Dfs(root.childs[i])...)
	}
	return entries
}

func Base32ToIndex(base32 byte) int {
	if base32 >= '0' && base32 <= '9' {
		return int(base32 - '0')
	}
	if base32 >= 'b' && base32 <= 'h' {
		return int(base32 - 'b' + 10)
	}
	if base32 >= 'j' && base32 <= 'k' {
		return int(base32 - 'j' + 17)
	}
	if base32 >= 'm' && base32 <= 'n' {
		return int(base32 - 'm' + 19)
	}
	if base32 >= 'p' && base32 <= 'z' {
		return int(base32 - 'p' + 21)
	}
	return -1
}
