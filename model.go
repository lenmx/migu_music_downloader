package main

type UiBase struct {
	Width  float32
	Height float32
	Title  string
}

type MiguMusicSetting struct {
	UiBase

	SearchUrl   string // http://pd.musicapp.migu.cn/MIGUM2.0/v1.0/content/search_all.do?ua=Android_migu&version=5.0.1&pageNo=1&pageSize=10&text=周杰伦&searchSwitch=
	DownloadUrl string // http://218.205.239.34/MIGUM2.0/v1.0/content/sub/listenSong.do?toneFlag=HQ&formatType=HQ&resourceType=2&netType=00&copyrightId=0&&contentId=600902000006889366&channel=0
}

type TabHeader struct {
	Text    string
	ColName string
}

type TabDataType int

const (
	TabDataType_Placeholder TabDataType = 0
	TabDataType_Data        TabDataType = 1
	TabDataType_Buttons     TabDataType = 2
)

type TabData struct {
	Type    TabDataType
	ColName string
	Content interface{}
}

type SearchRes struct {
	Code           string `json:"code"`
	Info           string `json:"info"`
	SongResultData struct {
		TotalCount  string        `json:"totalCount"`
		Correct     []interface{} `json:"correct"`
		ResultType  string        `json:"resultType"`
		IsFromCache string        `json:"isFromCache"`
		Result      []struct {
			Id           string   `json:"id"`
			ResourceType string   `json:"resourceType"`
			ContentId    string   `json:"contentId"`
			CopyrightId  string   `json:"copyrightId"`
			Name         string   `json:"name"`
			HighlightStr []string `json:"highlightStr"`
			Singers      []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"singers"`
			Albums []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"albums"`
			Tags     []string `json:"tags"`
			LyricUrl string   `json:"lyricUrl"`
			TrcUrl   string   `json:"trcUrl"`
			ImgItems []struct {
				ImgSizeType string `json:"imgSizeType"`
				Img         string `json:"img"`
			} `json:"imgItems"`
			TelevisionNames []string `json:"televisionNames"`
			Tones           []struct {
				Id          string `json:"id"`
				CopyrightId string `json:"copyrightId"`
				Price       string `json:"price"`
				ExpireDate  string `json:"expireDate"`
			} `json:"tones"`
			MvList []struct {
				Id           string `json:"id"`
				CopyrightId  string `json:"copyrightId"`
				ResourceType string `json:"resourceType"`
				Price        string `json:"price"`
				ExpireDate   string `json:"expireDate"`
				MvPicUrl     []struct {
					ImgSizeType string `json:"imgSizeType"`
					Img         string `json:"img"`
					FileId      string `json:"fileId"`
				} `json:"mvPicUrl"`
				IsInDAlbum string `json:"isInDAlbum"`
				PlayNum    string `json:"playNum"`
				MvType     string `json:"mvType"`
			} `json:"mvList,omitempty"`
			RelatedSongs []struct {
				ResourceType     string `json:"resourceType"`
				ResourceTypeName string `json:"resourceTypeName"`
				CopyrightId      string `json:"copyrightId"`
				ProductId        string `json:"productId"`
			} `json:"relatedSongs"`
			ToneControl string `json:"toneControl"`
			RateFormats []struct {
				ResourceType         string   `json:"resourceType"`
				FormatType           string   `json:"formatType"`
				Url                  string   `json:"url,omitempty"`
				Format               string   `json:"format"`
				Size                 string   `json:"size"`
				FileType             string   `json:"fileType,omitempty"`
				Price                string   `json:"price"`
				ShowTag              []string `json:"showTag,omitempty"`
				IosUrl               string   `json:"iosUrl,omitempty"`
				AndroidUrl           string   `json:"androidUrl,omitempty"`
				AndroidFileType      string   `json:"androidFileType,omitempty"`
				IosFileType          string   `json:"iosFileType,omitempty"`
				IosSize              string   `json:"iosSize,omitempty"`
				AndroidSize          string   `json:"androidSize,omitempty"`
				IosFormat            string   `json:"iosFormat,omitempty"`
				AndroidFormat        string   `json:"androidFormat,omitempty"`
				IosAccuracyLevel     string   `json:"iosAccuracyLevel,omitempty"`
				AndroidAccuracyLevel string   `json:"androidAccuracyLevel,omitempty"`
			} `json:"rateFormats"`
			NewRateFormats []struct {
				ResourceType         string   `json:"resourceType"`
				FormatType           string   `json:"formatType"`
				Url                  string   `json:"url,omitempty"`
				Format               string   `json:"format,omitempty"`
				Size                 string   `json:"size,omitempty"`
				FileType             string   `json:"fileType,omitempty"`
				Price                string   `json:"price"`
				ShowTag              []string `json:"showTag,omitempty"`
				IosUrl               string   `json:"iosUrl,omitempty"`
				AndroidUrl           string   `json:"androidUrl,omitempty"`
				AndroidFileType      string   `json:"androidFileType,omitempty"`
				IosFileType          string   `json:"iosFileType,omitempty"`
				IosSize              string   `json:"iosSize,omitempty"`
				AndroidSize          string   `json:"androidSize,omitempty"`
				IosFormat            string   `json:"iosFormat,omitempty"`
				AndroidFormat        string   `json:"androidFormat,omitempty"`
				IosAccuracyLevel     string   `json:"iosAccuracyLevel,omitempty"`
				AndroidAccuracyLevel string   `json:"androidAccuracyLevel,omitempty"`
				AndroidNewFormat     string   `json:"androidNewFormat,omitempty"`
				IosBit               int      `json:"iosBit,omitempty"`
				AndroidBit           int      `json:"androidBit,omitempty"`
			} `json:"newRateFormats"`
			Z3DCode struct {
				ResourceType    string `json:"resourceType"`
				FormatType      string `json:"formatType"`
				Price           string `json:"price"`
				IosUrl          string `json:"iosUrl"`
				AndroidUrl      string `json:"androidUrl"`
				AndroidFileType string `json:"androidFileType"`
				IosFileType     string `json:"iosFileType"`
				IosSize         string `json:"iosSize"`
				AndroidSize     string `json:"androidSize"`
				IosFormat       string `json:"iosFormat"`
				AndroidFormat   string `json:"androidFormat"`
				AndroidFileKey  string `json:"androidFileKey"`
				IosFileKey      string `json:"iosFileKey"`
				H5Url           string `json:"h5Url"`
				H5Size          string `json:"h5Size"`
				H5Format        string `json:"h5Format"`
			} `json:"z3dCode,omitempty"`
			SongType         string   `json:"songType"`
			IsInDAlbum       string   `json:"isInDAlbum"`
			Copyright        string   `json:"copyright"`
			DigitalColumnId  string   `json:"digitalColumnId"`
			Mrcurl           string   `json:"mrcurl"`
			SongDescs        string   `json:"songDescs"`
			SongAliasName    string   `json:"songAliasName"`
			TranslateName    string   `json:"translateName"`
			InvalidateDate   string   `json:"invalidateDate"`
			IsInSalesPeriod  string   `json:"isInSalesPeriod"`
			DalbumId         string   `json:"dalbumId"`
			IsInSideDalbum   string   `json:"isInSideDalbum"`
			VipType          string   `json:"vipType"`
			ChargeAuditions  string   `json:"chargeAuditions"`
			ScopeOfcopyright string   `json:"scopeOfcopyright"`
			MvCopyright      string   `json:"mvCopyright,omitempty"`
			MovieNames       []string `json:"movieNames,omitempty"`
		} `json:"result"`
		TipStatus string `json:"tipStatus"`
	} `json:"songResultData"`
	TagSongResultData struct {
		TotalCount string        `json:"totalCount"`
		Correct    []interface{} `json:"correct"`
		Result     []interface{} `json:"result"`
	} `json:"tagSongResultData"`
	BestShowResultToneData struct {
	} `json:"bestShowResultToneData"`
}
