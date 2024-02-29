package schemas

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"testing/fstest"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

var cueImportsPath = filepath.Join("packages", "grafana-schema", "src", "common")
var importPath = "github.com/grafana/grafana/packages/grafana-schema/src/common"

type ComposableKind struct {
	Name     string
	Filename string
	CueFile  cue.Value
}

func GetComposableKinds() ([]ComposableKind, error) {
	kinds := make([]ComposableKind, 0)

	_, caller, _, _ := runtime.Caller(0)
	root := filepath.Join(caller, "../../../..")

	azuremonitorCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/datasource/azuremonitor/dataquery.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "azuremonitor",
		Filename: "dataquery.cue",
		CueFile:  azuremonitorCue,
	})

	googlecloudmonitoringCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/datasource/cloud-monitoring/dataquery.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "googlecloudmonitoring",
		Filename: "dataquery.cue",
		CueFile:  googlecloudmonitoringCue,
	})

	cloudwatchCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/datasource/cloudwatch/dataquery.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "cloudwatch",
		Filename: "dataquery.cue",
		CueFile:  cloudwatchCue,
	})

	elasticsearchCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/datasource/elasticsearch/dataquery.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "elasticsearch",
		Filename: "dataquery.cue",
		CueFile:  elasticsearchCue,
	})

	grafanapyroscopeCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/datasource/grafana-pyroscope-datasource/dataquery.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "grafanapyroscope",
		Filename: "dataquery.cue",
		CueFile:  grafanapyroscopeCue,
	})

	grafanatestdatadatasourceCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/datasource/grafana-testdata-datasource/dataquery.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "grafanatestdatadatasource",
		Filename: "dataquery.cue",
		CueFile:  grafanatestdatadatasourceCue,
	})

	lokiCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/datasource/loki/dataquery.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "loki",
		Filename: "dataquery.cue",
		CueFile:  lokiCue,
	})

	parcaCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/datasource/parca/dataquery.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "parca",
		Filename: "dataquery.cue",
		CueFile:  parcaCue,
	})

	prometheusCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/datasource/prometheus/dataquery.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "prometheus",
		Filename: "dataquery.cue",
		CueFile:  prometheusCue,
	})

	tempoCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/datasource/tempo/dataquery.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "tempo",
		Filename: "dataquery.cue",
		CueFile:  tempoCue,
	})

	alertgroupsCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/alertGroups/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "alertgroups",
		Filename: "panelcfg.cue",
		CueFile:  alertgroupsCue,
	})

	annotationslistCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/annolist/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "annotationslist",
		Filename: "panelcfg.cue",
		CueFile:  annotationslistCue,
	})

	barchartCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/barchart/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "barchart",
		Filename: "panelcfg.cue",
		CueFile:  barchartCue,
	})

	bargaugeCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/bargauge/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "bargauge",
		Filename: "panelcfg.cue",
		CueFile:  bargaugeCue,
	})

	candlestickCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/candlestick/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "candlestick",
		Filename: "panelcfg.cue",
		CueFile:  candlestickCue,
	})

	canvasCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/canvas/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "canvas",
		Filename: "panelcfg.cue",
		CueFile:  canvasCue,
	})

	dashboardlistCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/dashlist/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "dashboardlist",
		Filename: "panelcfg.cue",
		CueFile:  dashboardlistCue,
	})

	datagridCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/datagrid/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "datagrid",
		Filename: "panelcfg.cue",
		CueFile:  datagridCue,
	})

	debugCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/debug/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "debug",
		Filename: "panelcfg.cue",
		CueFile:  debugCue,
	})

	gaugeCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/gauge/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "gauge",
		Filename: "panelcfg.cue",
		CueFile:  gaugeCue,
	})

	geomapCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/geomap/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "geomap",
		Filename: "panelcfg.cue",
		CueFile:  geomapCue,
	})

	heatmapCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/heatmap/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "heatmap",
		Filename: "panelcfg.cue",
		CueFile:  heatmapCue,
	})

	histogramCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/histogram/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "histogram",
		Filename: "panelcfg.cue",
		CueFile:  histogramCue,
	})

	logsCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/logs/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "logs",
		Filename: "panelcfg.cue",
		CueFile:  logsCue,
	})

	newsCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/news/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "news",
		Filename: "panelcfg.cue",
		CueFile:  newsCue,
	})

	nodegraphCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/nodeGraph/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "nodegraph",
		Filename: "panelcfg.cue",
		CueFile:  nodegraphCue,
	})

	piechartCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/piechart/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "piechart",
		Filename: "panelcfg.cue",
		CueFile:  piechartCue,
	})

	statCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/stat/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "stat",
		Filename: "panelcfg.cue",
		CueFile:  statCue,
	})

	statetimelineCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/state-timeline/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "statetimeline",
		Filename: "panelcfg.cue",
		CueFile:  statetimelineCue,
	})

	statushistoryCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/status-history/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "statushistory",
		Filename: "panelcfg.cue",
		CueFile:  statushistoryCue,
	})

	tableCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/table/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "table",
		Filename: "panelcfg.cue",
		CueFile:  tableCue,
	})

	textCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/text/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "text",
		Filename: "panelcfg.cue",
		CueFile:  textCue,
	})

	timeseriesCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/timeseries/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "timeseries",
		Filename: "panelcfg.cue",
		CueFile:  timeseriesCue,
	})

	trendCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/trend/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "trend",
		Filename: "panelcfg.cue",
		CueFile:  trendCue,
	})

	xychartCue, err := loadCueFileWithCommon(root, filepath.Join(root, "./public/app/plugins/panel/xychart/panelcfg.cue"))
	if err != nil {
		return nil, err
	}
	kinds = append(kinds, ComposableKind{
		Name:     "xychart",
		Filename: "panelcfg.cue",
		CueFile:  xychartCue,
	})

	return kinds, nil
}

func loadCueFileWithCommon(root string, entrypoint string) (cue.Value, error) {
	commonFS, err := mockCommonFS(root)
	if err != nil {
		fmt.Printf("cannot load common cue files: %s\n", err)
		return cue.Value{}, err
	}

	overlay, err := buildOverlay(commonFS)
	if err != nil {
		fmt.Printf("Cannot build overlay: %s\n", err)
		return cue.Value{}, err
	}

	bis := load.Instances([]string{entrypoint}, &load.Config{
		ModuleRoot: "/",
		Overlay:    overlay,
	})

	values, err := cuecontext.New().BuildInstances(bis)
	if err != nil {
		fmt.Printf("Cannot build instance: %s\n", err)
		return cue.Value{}, err
	}

	return values[0], nil
}

func mockCommonFS(root string) (fs.FS, error) {
	path := filepath.Join(root, cueImportsPath)
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open common cue files directory: %s", err)
	}

	prefix := "cue.mod/pkg/" + importPath

	commonFS := fstest.MapFS{}
	for _, d := range dir {
		if d.IsDir() {
			continue
		}

		b, err := os.ReadFile(filepath.Join(path, d.Name()))
		if err != nil {
			return nil, err
		}

		commonFS[filepath.Join(prefix, d.Name())] = &fstest.MapFile{Data: b}
	}

	return commonFS, nil
}

// It loads common cue files into the schema to be able to make import works
func buildOverlay(commonFS fs.FS) (map[string]load.Source, error) {
	overlay := make(map[string]load.Source)

	err := fs.WalkDir(commonFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		f, err := commonFS.Open(path)
		if err != nil {
			return err
		}
		defer func() { _ = f.Close() }()

		b, err := io.ReadAll(f)
		if err != nil {
			return err
		}

		overlay[filepath.Join("/", path)] = load.FromBytes(b)

		return nil
	})

	return overlay, err
}
