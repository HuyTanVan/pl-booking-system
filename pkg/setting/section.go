package setting

type Config struct {
	HTTPServer        HTTPServerSetting        `mapstructure:"httpserver"`
	GRPCServer        GRPCServerSetting        `mapstructure:"grpcserver"`
	GRPCGatewayServer GRPCGatewayServerSetting `mapstructure:"grpcgatewayserver"`
	PostgreSQL        PostgreSQLSetting        `mapstructure:"postgresql"`
	Redis             RedisSeting              `mapstructure:"redis"`
	EmailSender       EmailSenderSetting       `mapstructure:"emailsender"`
	JWTToken          JWTTokenSetting          `mapstructure:"jwttoken"`
	Stripe            StripeSetting            `mapstructure:"stripe"`
}
type EmailSenderSetting struct {
	EmailSenderName     string `mapstructure:"emailsendername"`
	EmailSenderAddress  string `mapstructure:"emailsenderaddress"`
	EmailSenderPassword string `mapstructure:"emailsenderpassword"`
}
type HTTPServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type GRPCServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}
type GRPCGatewayServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}
type PostgreSQLSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type RedisSeting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type JWTTokenSetting struct {
	TokenSymmetricKey    string `mapstructure:"tokensymmetrickey"`
	AccessTokenDuration  string `mapstructure:"accesstokenduration"`
	RefreshTokenDuration string `mapstructure:"refreshtokenduration"`
}

type StripeSetting struct {
	StripePublishableKey string `mapstructure:"stripepublishablekey"`
	StripeSecretKey      string `mapstructure:"stripesecretkey"`
	Visa                 string `mapstructure:"visa"`
}
