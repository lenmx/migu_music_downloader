package main

func main() {
	MiguMusicInit(&MiguMusicSetting{
		UiBase:      UiBase{
			Width:  1000,
			Height: 600,
			Title:  "migu_music_downloader",
		},
		SearchUrl:   "http://pd.musicapp.migu.cn/MIGUM2.0/v1.0/content/search_all.do?ua=Android_migu&version=5.0.1&pageNo=%d&pageSize=%d&text=%s&searchSwitch=",
		DownloadUrl: "http://218.205.239.34/MIGUM2.0/v1.0/content/sub/listenSong.do?toneFlag=HQ&formatType=HQ&resourceType=2&netType=00&copyrightId=0&&contentId=%s&channel=0",
	})
}