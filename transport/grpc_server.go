package transport

ssov1 "github.com/antoha2/auth/proto/gen/go/sso"

type serverAPI struct {
	ssov1.UnimplementedAuthServer // Хитрая штука, о ней ниже
	auth                          Auth
}
