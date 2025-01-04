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

	Files *SchemaModFiles
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

type SchemaModFile struct {
	Id         []int `json:"id"`
	Uid        int   `json:"uid"`
	FileId     int   `json:"file_id"`
	CategoryId int   `json:"category_id"`

	Name         string `json:"name"`
	CategoryName string `json:"category_name"`
	FileName     string `json:"file_name"`
	Description  string `json:"description"`

	Version       string `json:"version"`
	ModVersion    string `json:"mod_version"`
	IsPrimary     bool   `json:"is_primary"`
	Size          int    `json:"size"`
	SizeKb        int    `json:"size_kb"`
	SizeKbInBytes int    `json:"size_kb_in_bytes"`

	UploadedTimestamp    int    `json:"uploaded_timestamp"`
	UploadedTime         string `json:"uploaded_time"`
	ExternalVirusScanUrl string `json:"external_virus_scan_url"`

	ChangelogHtml      string `json:"changelog_html"`
	ContentPreviewLink string `json:"content_preview_link"`
}

type SchemaModFileUpdate struct {
	OldFileId         int    `json:"old_file_id"`
	NewFileId         int    `json:"new_file_id"`
	OldFileName       string `json:"old_file_name"`
	NewFileName       string `json:"new_file_name"`
	UploadedTimestamp int    `json:"uploaded_timestamp"`
	UploadedTime      string `json:"uploaded_time"`
}

type SchemaModFiles struct {
	Files       []SchemaModFile       `json:"files"`
	FileUpdates []SchemaModFileUpdate `json:"file_updates"`
}
