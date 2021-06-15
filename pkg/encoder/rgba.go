package encoder

type ChannelSelect byte

const (
	Red ChannelSelect = iota
	Green
	Blue
	Alpha
)

type RGBAMode struct {
	Encoder
	ChannelEncode ChannelSelect
}

func Encode() {

}

func LoadKey() {

}

func Mode() string {
	return ""
}
