package sdp

import (
	"github.com/pion/sdp/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUndefinedValue(t *testing.T) {
	sdpStr := `v=0
o=- 0 0 IN IP4 127.0.0.1
s=-
c=IN IP4 127.0.0.1
t=0 0
m=video 0 RTP/AVP 96
a=rtpmap:96 H264/90000
a=recvonly
a=somefield:undefined
`
	_, err := sdp.Unmarshal([]byte(sdpStr))
	require.NoError(t, err)
}
