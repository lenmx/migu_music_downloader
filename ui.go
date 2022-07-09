package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/go-resty/resty/v2"
	"strconv"
	"strings"
)

var (
	searchSuccess = "000000"
	datas         = [][]interface{}{
		{TabHeader{Text: "名称"}, TabHeader{Text: "歌手"}, TabHeader{Text: "专辑"}, TabHeader{Text: "无损下载"}, TabHeader{Text: "高品质下载"}},
	}

	myApp           fyne.App
	myWindow        fyne.Window
	loading         *widget.ProgressBarInfinite
	txtKeyword      *widget.Entry
	txtDownloadPath *widget.Entry
	form            *widget.Form
	tab             *widget.Table
	pageContainer   *fyne.Container
	pageButtongs    = []fyne.CanvasObject{}
	pageSize        = 20
	pageIndex       = 1
	pageIndexOffset = 2
)

func MiguMusicInit(setting *MiguMusicSetting) {
	myApp = app.New()
	myApp.Settings().SetTheme(&ThemeDefault{})
	myApp.SetIcon(resourceFavPng)

	myWindow = myApp.NewWindow(setting.Title)
	myWindow.Resize(fyne.NewSize(setting.Width, setting.Height))

	loading = widget.NewProgressBarInfinite()
	loading.Hide()

	txtKeyword = widget.NewEntry()
	txtKeyword.PlaceHolder = "请输入关键字"

	txtDownloadPath = widget.NewEntry()
	txtDownloadPath.PlaceHolder = "请设置保存路径"

	tab = widget.NewTable(
		func() (int, int) { return len(datas), len(datas[0]) },
		func() fyne.CanvasObject { return widget.NewLabel("wide content") },
		func(i widget.TableCellID, o fyne.CanvasObject) {
			// 表头
			if i.Row == 0 {
				data := datas[i.Row][i.Col].(TabHeader)
				o.(*widget.Label).SetText(data.Text)
				return
			}

			// 下载列
			if i.Col == len(datas[0])-2 {
				tmp := o.(*widget.Label)
				tmp.SetText("无损")
				return
			}
			if i.Col == len(datas[0])-1 {
				tmp := o.(*widget.Label)
				tmp.SetText("高品质")
				return
			}

			// 数据
			data := datas[i.Row][i.Col].(TabData)
			o.(*widget.Label).SetText(data.Content.(string))
		},
	)
	tab.SetColumnWidth(0, setting.Width*0.4)
	tab.SetColumnWidth(1, setting.Width*0.2)
	tab.SetColumnWidth(2, setting.Width*0.2)
	tab.SetColumnWidth(3, setting.Width*0.08)
	//tab.SetColumnWidth(4, setting.Width*0.08)
	tab.OnSelected = setting.OnDownload

	form = &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "关键字", Widget: txtKeyword},
			{Text: "保存路径", Widget: txtDownloadPath},
			//{Text: "", Widget: btnBatchDownload},
		},
		OnSubmit: func() {
			pageIndex = 1
			setting.OnSearch(pageIndex)
		},
	}
	form.SubmitText = "提交"

	btnBatchDownloaderSQ := widget.NewButton("下载本页(无损)", func() { setting.OnBatchDownload(SourceType_SQ) })
	btnBatchDownloaderHQ := widget.NewButton("下载本页(高品质)", func() { setting.OnBatchDownload(SourceType_HQ) })

	pageContainer = container.NewHBox()
	downloaderContainer := container.NewHBox(btnBatchDownloaderSQ, btnBatchDownloaderHQ)
	toolContainer := container.NewHBox(pageContainer, downloaderContainer)

	box := container.NewVBox(loading, form, toolContainer)
	container := container.NewVSplit(box, tab)
	container.SetOffset(-100)

	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}

func (m *MiguMusicSetting) GenPagination(total int) {
	for _, b := range pageButtongs {
		pageContainer.Remove(b)
	}

	pageCount := total / pageSize
	if total%pageSize > 0 {
		pageCount += 1
	}

	// 首页
	if pageIndex > pageIndexOffset {
		button := widget.NewButton("首页", func() {
			m.OnSearch(1)
		})
		lab := widget.NewLabel("...")

		pageButtongs = append(pageButtongs, button)
		pageButtongs = append(pageButtongs, lab)
		pageContainer.Add(button)
		pageContainer.Add(lab)
	}

	for i := pageIndex - pageIndexOffset; i < pageIndex+pageIndexOffset+1; i++ {
		if i <= 0 {
			continue
		}

		if i > pageCount-1 {
			continue
		}

		func(_i int) {
			button := widget.NewButton(strconv.Itoa(_i), func() {
				m.OnSearch(_i)
			})
			if _i == pageIndex {
				button.Disable()
			}

			pageButtongs = append(pageButtongs, button)
			pageContainer.Add(button)
		}(i)
	}

	// 末页
	if pageIndex+pageIndexOffset < pageCount-1 {
		lab := widget.NewLabel("...")
		button := widget.NewButton("末页", func() {
			m.OnSearch(pageCount - 1)
		})
		pageButtongs = append(pageButtongs, lab)
		pageButtongs = append(pageButtongs, button)
		pageContainer.Add(lab)
		pageContainer.Add(button)
	}
}

func (m *MiguMusicSetting) OnSearch(_pageIndex int) {
	if loading.Visible() {
		return
	}
	if len(txtKeyword.Text) <= 0 {
		return
	}

	loading.Show()
	form.Disable()
	defer func() {
		loading.Hide()
		form.Enable()
	}()

	searchDataRes, err := m.Search(txtKeyword.Text, _pageIndex, pageSize)
	if err != nil {
		showErr(err, myWindow)
		return
	}
	if searchDataRes.Code != searchSuccess {
		showErr(errors.New(searchDataRes.Info), myWindow)
		return
	}

	datas = datas[:1]
	for _, data := range searchDataRes.SongResultData.Result {
		singers := []string{}
		for _, singer := range data.Singers {
			singers = append(singers, singer.Name)
		}

		albums := []string{}
		for _, album := range data.Albums {
			albums = append(albums, album.Name)
		}

		datas = append(datas, []interface{}{
			TabData{
				Type:    TabDataType_Data,
				ColName: "name",
				Content: data.Name,
			},
			TabData{
				Type:    TabDataType_Data,
				ColName: "singers",
				Content: strings.Join(singers, "，"),
			},
			TabData{
				Type:    TabDataType_Data,
				ColName: "albums",
				Content: strings.Join(albums, "，"),
			},
			TabData{
				Type:    TabDataType_Buttons,
				ColName: string(SourceType_SQ),
				Content: data.ContentId,
			},
			TabData{
				Type:    TabDataType_Buttons,
				ColName: string(SourceType_HQ),
				Content: data.ContentId,
			},
		})
	}
	tab.Refresh()

	pageIndex = _pageIndex
	total, _ := strconv.Atoi(searchDataRes.SongResultData.TotalCount)
	m.GenPagination(total)
}

func (m *MiguMusicSetting) OnDownload(id widget.TableCellID) {
	if loading.Visible() {
		return
	}
	if id.Row == 0 {
		return
	}
	data := datas[id.Row][id.Col].(TabData)
	if data.Type != TabDataType_Buttons {
		return
	}
	if len(txtDownloadPath.Text) <= 0 {
		showErr(errors.New("请设置保存路径"), myWindow)
		return
	}

	loading.Show()
	defer func() {
		loading.Hide()
	}()

	contentId := data.Content.(string)
	name := datas[id.Row][0].(TabData).Content.(string)

	m.Download(SourceType(data.ColName), contentId, name, txtDownloadPath.Text)
	dialog.ShowInformation("信息", fmt.Sprintf("下载完毕"), myWindow)
}

func (m *MiguMusicSetting) OnBatchDownload(sourceType SourceType) {
	if loading.Visible() {
		return
	}
	if len(datas) <= 1 {
		return
	}
	if len(txtDownloadPath.Text) <= 0 {
		showErr(errors.New("请设置保存路径"), myWindow)
		return
	}

	loading.Show()
	defer func() {
		loading.Hide()
	}()

	//var wg sync.WaitGroup
	for i, _datas := range datas {
		if i == 0 {
			continue
		}

		//wg.Add(1)
		//go func() {
		contentId := _datas[len(_datas)-1].(TabData).Content.(string)
		name := _datas[0].(TabData).Content.(string)
		m.Download(sourceType, contentId, name, txtDownloadPath.Text)
		//}()
	}

	//wg.Wait()
	dialog.ShowInformation("信息", fmt.Sprintf("下载完毕"), myWindow)
}

func (m *MiguMusicSetting) Search(keyword string, pageIndex, pageSize int) (*SearchRes, error) {
	// http://pd.musicapp.migu.cn/MIGUM2.0/v1.0/content/search_all.do?ua=Android_migu&version=5.0.1&pageNo=1&pageSize=10&text=周杰伦&searchSwitch=
	url := fmt.Sprintf(m.SearchUrl, pageIndex, pageSize, keyword)

	res, err := resty.New().R().Get(url)
	if err != nil {
		return nil, err
	}

	var resp SearchRes
	err = json.Unmarshal(res.Body(), &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (m *MiguMusicSetting) Download(sourceType SourceType, contentId string, name, path string) error {
	//无损：http://218.205.239.34/MIGUM2.0/v1.0/content/sub/listenSong.do?toneFlag=SQ&formatType=SQ&resourceType=E&netType=00&copyrightId=0&&contentId=600902000006889366&channel=0
	//高品质：http://218.205.239.34/MIGUM2.0/v1.0/content/sub/listenSong.do?toneFlag=HQ&formatType=HQ&resourceType=2&netType=00&copyrightId=0&&contentId=600902000006889366&channel=0

	if path[len(path)-1] != '/' {
		path += "/"
	}

	path += name + SourceType2FileExt[sourceType]
	url := fmt.Sprintf(m.DownloadUrl, string(sourceType), contentId)
	_, err := resty.New().R().SetOutput(path).Get(url)

	return err
}

func showErr(err error, parent fyne.Window) {
	dialog.ShowError(err, parent)
}
