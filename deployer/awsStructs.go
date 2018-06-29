package deployer

type AWSProvider struct {
	Instances      []AWSInstance      `json:"instances"`
	API            []AWSApi           `json:"api"`
	DomainFront    []AWSDomainFront   `json:"domain_front"`
	SecurityGroups []AWSSecurityGroup `json:"security_group"`
}

type AWSConfigWrapper struct {
	Config    AWSRegionConfig
	RegionMap map[string]int
}

type AWSInstance struct {
	Config  AWSRegionConfig   `json:"config"`
	IPIDMap map[string]string `json:"ip_id"`
}

type AWSRegionConfig struct {
	ModuleName      string
	SecurityGroup   string `json:"hidensneak"`
	SecurityGroupID string `json:"aws_sg_id"`
	Count           string `json:"region_count"`
	CustomAmi       string `json:"custom_ami"`
	InstanceType    string `json:"aws_instance_type"`
	DefaultUser     string `json:"ec2_default_user"`
	Region          string `json:"region"`
	PrivateKeyFile  string `json:"private_key_file"`
	PublicKeyFile   string `json:"public_key_file"`
}

type AWSApi struct {
}

type AWSDomainFront struct{}

type AWSSecurityGroup struct{}

type cloudFrontDeployer struct {
	Origin string
	Region string
}

//Deprecated
type apiGatewayDeployer struct {
	TargetURI string
	StageName string
}