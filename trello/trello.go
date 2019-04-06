package trello

import "time"

/*
Trello
*/

type Member struct {
	ID                       string      `json:"id"`
	AvatarHash               interface{} `json:"avatarHash"`
	AvatarURL                interface{} `json:"avatarUrl"`
	Bio                      string      `json:"bio"`
	BioData                  interface{} `json:"bioData"`
	Confirmed                bool        `json:"confirmed"`
	FullName                 string      `json:"fullName"`
	IDEnterprise             interface{} `json:"idEnterprise"`
	IDEnterprisesDeactivated interface{} `json:"idEnterprisesDeactivated"`
	IDMemberReferrer         interface{} `json:"idMemberReferrer"`
	IDPremOrgsAdmin          interface{} `json:"idPremOrgsAdmin"`
	Initials                 string      `json:"initials"`
	MemberType               string      `json:"memberType"`
	NonPublic                struct {
	} `json:"nonPublic"`
	NonPublicAvailable bool          `json:"nonPublicAvailable"`
	Products           []interface{} `json:"products"`
	Status             string        `json:"status"`
	URL                string        `json:"url"`
	Username           string        `json:"username"`
	AaEmail            interface{}   `json:"aaEmail"`
	AaID               interface{}   `json:"aaId"`
	AvatarSource       interface{}   `json:"avatarSource"`
	Email              interface{}   `json:"email"`
	GravatarHash       interface{}   `json:"gravatarHash"`
	IDBoards           []interface{} `json:"idBoards"`
	IDOrganizations    []interface{} `json:"idOrganizations"`
	IDEnterprisesAdmin []interface{} `json:"idEnterprisesAdmin"`
	Limits             struct {
		Boards struct {
			TotalPerMember struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"totalPerMember"`
		} `json:"boards"`
		Orgs struct {
			TotalPerMember struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"totalPerMember"`
		} `json:"orgs"`
	} `json:"limits"`
	LoginTypes               interface{}   `json:"loginTypes"`
	MarketingOptIn           interface{}   `json:"marketingOptIn"`
	MessagesDismissed        interface{}   `json:"messagesDismissed"`
	OneTimeMessagesDismissed interface{}   `json:"oneTimeMessagesDismissed"`
	Prefs                    interface{}   `json:"prefs"`
	Trophies                 []interface{} `json:"trophies"`
	UploadedAvatarHash       interface{}   `json:"uploadedAvatarHash"`
	UploadedAvatarURL        interface{}   `json:"uploadedAvatarUrl"`
	PremiumFeatures          []interface{} `json:"premiumFeatures"`
	IsAaMastered             interface{}   `json:"isAaMastered"`
	IDBoardsPinned           interface{}   `json:"idBoardsPinned"`
}

type Board struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Desc           string      `json:"desc"`
	DescData       interface{} `json:"descData"`
	Closed         bool        `json:"closed"`
	IDOrganization string      `json:"idOrganization"`
	Limits         struct {
		Attachments struct {
			PerBoard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perBoard"`
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"attachments"`
		Boards struct {
			TotalMembersPerBoard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"totalMembersPerBoard"`
		} `json:"boards"`
		Cards struct {
			OpenPerBoard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"openPerBoard"`
			OpenPerList struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"openPerList"`
			TotalPerBoard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"totalPerBoard"`
			TotalPerList struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"totalPerList"`
		} `json:"cards"`
		Checklists struct {
			PerBoard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perBoard"`
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"checklists"`
		CheckItems struct {
			PerChecklist struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perChecklist"`
		} `json:"checkItems"`
		CustomFields struct {
			PerBoard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perBoard"`
		} `json:"customFields"`
		Labels struct {
			PerBoard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perBoard"`
		} `json:"labels"`
		Lists struct {
			OpenPerBoard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"openPerBoard"`
			TotalPerBoard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"totalPerBoard"`
		} `json:"lists"`
		Stickers struct {
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"stickers"`
		Reactions struct {
			PerAction struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perAction"`
			UniquePerAction struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"uniquePerAction"`
		} `json:"reactions"`
	} `json:"limits"`
	Memberships []struct {
		ID          string `json:"id"`
		IDMember    string `json:"idMember"`
		MemberType  string `json:"memberType"`
		Unconfirmed bool   `json:"unconfirmed"`
	} `json:"memberships"`
	Pinned  bool        `json:"pinned"`
	Starred interface{} `json:"starred"`
	URL     string      `json:"url"`
	Prefs   struct {
		PermissionLevel       string `json:"permissionLevel"`
		Voting                string `json:"voting"`
		Comments              string `json:"comments"`
		Invitations           string `json:"invitations"`
		SelfJoin              bool   `json:"selfJoin"`
		CardCovers            bool   `json:"cardCovers"`
		CardAging             string `json:"cardAging"`
		CalendarFeedEnabled   bool   `json:"calendarFeedEnabled"`
		Background            string `json:"background"`
		BackgroundImage       string `json:"backgroundImage"`
		BackgroundImageScaled []struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
		} `json:"backgroundImageScaled"`
		BackgroundTile        bool   `json:"backgroundTile"`
		BackgroundBrightness  string `json:"backgroundBrightness"`
		BackgroundBottomColor string `json:"backgroundBottomColor"`
		BackgroundTopColor    string `json:"backgroundTopColor"`
		CanBePublic           bool   `json:"canBePublic"`
		CanBeEnterprise       bool   `json:"canBeEnterprise"`
		CanBeOrg              bool   `json:"canBeOrg"`
		CanBePrivate          bool   `json:"canBePrivate"`
		CanInvite             bool   `json:"canInvite"`
	} `json:"prefs"`
	ShortLink  string      `json:"shortLink"`
	Subscribed interface{} `json:"subscribed"`
	LabelNames struct {
		Green  string `json:"green"`
		Yellow string `json:"yellow"`
		Orange string `json:"orange"`
		Red    string `json:"red"`
		Purple string `json:"purple"`
		Blue   string `json:"blue"`
		Sky    string `json:"sky"`
		Lime   string `json:"lime"`
		Pink   string `json:"pink"`
		Black  string `json:"black"`
	} `json:"labelNames"`
	PowerUps          []interface{} `json:"powerUps"`
	DateLastActivity  time.Time     `json:"dateLastActivity"`
	DateLastView      interface{}   `json:"dateLastView"`
	ShortURL          string        `json:"shortUrl"`
	IDTags            []interface{} `json:"idTags"`
	DatePluginDisable interface{}   `json:"datePluginDisable"`
	CreationMethod    interface{}   `json:"creationMethod"`
}

type Tasks []struct {
	ID                string        `json:"id"`
	Address           interface{}   `json:"address"`
	CheckItemStates   interface{}   `json:"checkItemStates"`
	Closed            bool          `json:"closed"`
	Coordinates       interface{}   `json:"coordinates"`
	CreationMethod    interface{}   `json:"creationMethod"`
	DateLastActivity  time.Time     `json:"dateLastActivity"`
	Desc              string        `json:"desc"`
	DescData          interface{}   `json:"descData"`
	DueReminder       interface{}   `json:"dueReminder"`
	IDBoard           string        `json:"idBoard"`
	IDLabels          []interface{} `json:"idLabels"`
	IDList            string        `json:"idList"`
	IDMembersVoted    []interface{} `json:"idMembersVoted"`
	IDShort           int           `json:"idShort"`
	IDAttachmentCover interface{}   `json:"idAttachmentCover"`
	Limits            struct {
		Attachments struct {
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"attachments"`
		Checklists struct {
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"checklists"`
		Stickers struct {
			PerCard struct {
				Status    string `json:"status"`
				DisableAt int    `json:"disableAt"`
				WarnAt    int    `json:"warnAt"`
			} `json:"perCard"`
		} `json:"stickers"`
	} `json:"limits"`
	LocationName          interface{} `json:"locationName"`
	ManualCoverAttachment bool        `json:"manualCoverAttachment"`
	Name                  string      `json:"name"`
	Pos                   int         `json:"pos"`
	ShortLink             string      `json:"shortLink"`
	Badges                struct {
		AttachmentsByType struct {
			Trello struct {
				Board int `json:"board"`
				Card  int `json:"card"`
			} `json:"trello"`
		} `json:"attachmentsByType"`
		Location           bool        `json:"location"`
		Votes              int         `json:"votes"`
		ViewingMemberVoted bool        `json:"viewingMemberVoted"`
		Subscribed         bool        `json:"subscribed"`
		Fogbugz            string      `json:"fogbugz"`
		CheckItems         int         `json:"checkItems"`
		CheckItemsChecked  int         `json:"checkItemsChecked"`
		Comments           int         `json:"comments"`
		Attachments        int         `json:"attachments"`
		Description        bool        `json:"description"`
		Due                interface{} `json:"due"`
		DueComplete        bool        `json:"dueComplete"`
	} `json:"badges"`
	DueComplete  bool          `json:"dueComplete"`
	Due          interface{}   `json:"due"`
	IDChecklists []interface{} `json:"idChecklists"`
	IDMembers    []interface{} `json:"idMembers"`
	Labels       []interface{} `json:"labels"`
	ShortURL     string        `json:"shortUrl"`
	Subscribed   bool          `json:"subscribed"`
	URL          string        `json:"url"`
}
