package globals

// Promo 生成推广码
func Promo() string {
	return Bytes2str(Krand(8, KC_RAND_KIND_ALL))
}
