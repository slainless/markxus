package nexus

type SchemaMod struct {
	Uid        int    `json:"uid"`
	ModId      int    `json:"mod_id"`
	GameId     int    `json:"game_id"`
	DomainName string `json:"domain_name"`

	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Version     string `json:"version"`
	CategoryId  int    `json:"category_id"`
	Status      string `json:"status"`

	MetadataPageUrl      string
	MetadataCategoryIcon string

	Author                  string `json:"author"`
	UploadedBy              string `json:"uploaded_by"`
	UploadedUsersProfileUrl string `json:"uploaded_users_profile_url"`

	PictureUrl string `json:"picture_url"`

	AllowRating          bool `json:"allow_rating"`
	ContainsAdultContent bool `json:"contains_adult_content"`
	Available            bool `json:"available"`

	EndorsementCount   int `json:"endorsement_count"`
	ModDownloads       int `json:"mod_downloads"`
	ModUniqueDownloads int `json:"mod_unique_downloads"`

	CreatedTimestamp int `json:"created_timestamp"`
	UpdatedTimestamp int `json:"updated_timestamp"`

	User        *SchemaModUser        `json:"user"`
	Endorsement *SchemaModEndorsement `json:"endorsement"`
}

type SchemaModUser struct {
	MemberId      int    `json:"member_id"`
	MemberGroupId int    `json:"member_group_id"`
	Name          string `json:"name"`
}

type SchemaModEndorsement struct {
	EndorseStatus string `json:"endorse_status"`
	Timestamp     int    `json:"timestamp"`
	Version       string `json:"version"`
}
