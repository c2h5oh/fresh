package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/BurntSushi/toml"
)

type config struct {
	Root            string   `toml:"root"`
	WatchPaths      []string `toml:"watch_paths"`
	ExcludePaths    []string `toml:"exclude_paths"`
	ConfigPath      string   `toml:"config_path"`
	TmpPath         string   `toml:"tmp_path"`
	BuildName       string   `toml:"build_name"`
	BuildArgs       string   `toml:"build_args"`
	RunArgs         string   `toml:"run_args"`
	BuildLog        string   `toml:"build_log"`
	ValidExtensions []string `toml:"valid_ext"`
	BuildDelay      int32    `toml:"build_delay"`
	Colors          bool     `toml:"colors"`
	LogColorMain    string   `toml:"log_color_main"`
	LogColorBuild   string   `toml:"log_color_build"`
	LogColorRunner  string   `toml:"log_color_runner"`
	LogColorWatcher string   `toml:"log_color_watcher"`
	LogColorApp     string   `toml:"log_color_app"`

	BuildErrorPath string
	BinaryPath     string
}

var (
	settings = config{
		Root:            ".",
		WatchPaths:      []string{},
		ExcludePaths:    []string{},
		ConfigPath:      "./runner.conf",
		TmpPath:         "./tmp",
		BuildName:       "runner-build",
		BuildArgs:       "",
		RunArgs:         "",
		BuildLog:        "runner-build-errors.log",
		ValidExtensions: []string{".go", ".tpl", ".tmpl", ".html"},
		BuildDelay:      600,
		Colors:          true,
		LogColorMain:    "cyan",
		LogColorBuild:   "yellow",
		LogColorRunner:  "green",
		LogColorWatcher: "magenta",
	}
	colors = map[string]string{
		"reset":          "0",
		"black":          "30",
		"red":            "31",
		"green":          "32",
		"yellow":         "33",
		"blue":           "34",
		"magenta":        "35",
		"cyan":           "36",
		"white":          "37",
		"bold_black":     "30;1",
		"bold_red":       "31;1",
		"bold_green":     "32;1",
		"bold_yellow":    "33;1",
		"bold_blue":      "34;1",
		"bold_magenta":   "35;1",
		"bold_cyan":      "36;1",
		"bold_white":     "37;1",
		"bright_black":   "30;2",
		"bright_red":     "31;2",
		"bright_green":   "32;2",
		"bright_yellow":  "33;2",
		"bright_blue":    "34;2",
		"bright_magenta": "35;2",
		"bright_cyan":    "36;2",
		"bright_white":   "37;2",
	}
)

func initSettings(confFile, buildArgs, runArgs, buildPath, outputBuildPath *string, watchList, excludeList Multiflag) error {
	defer buildPaths()

	if *confFile != "" {
		if _, err := os.Stat(*confFile); os.IsNotExist(err) {
			return fmt.Errorf("Config file %s does not exist", *confFile)
		}
		settings.ConfigPath = *confFile

		if _, err := toml.DecodeFile(settings.ConfigPath, &settings); err != nil {
			return fmt.Errorf("Reading config file failed: %v", err)
		}
	}

	if *buildArgs != "" {
		settings.BuildArgs = *buildArgs
	}
	if *runArgs != "" {
		settings.RunArgs = *runArgs
	}
	if *buildPath != "" {
		settings.Root = *buildPath
	}
	if *outputBuildPath != "" {
		settings.BuildName = *outputBuildPath
	}
	if len(watchList) > 0 {
		settings.WatchPaths = watchList
	}
	if len(excludeList) > 0 {
		settings.ExcludePaths = excludeList
	}

	return nil
}

func logColor(logName string) string {
	switch strings.ToLower(logName) {
	case "main":
		return colors[settings.LogColorMain]
	case "build":
		return colors[settings.LogColorBuild]
	case "runner":
		return colors[settings.LogColorRunner]
	case "watcher":
		return colors[settings.LogColorWatcher]
	default:
		return colors[settings.LogColorApp]
	}
}

func buildPaths() {
	settings.BinaryPath = filepath.Join(settings.TmpPath, settings.BuildName)
	if runtime.GOOS == "windows" && filepath.Ext(settings.BinaryPath) != ".exe" {
		settings.BinaryPath += ".exe"
	}
	settings.BuildErrorPath = filepath.Join(settings.TmpPath, settings.BuildLog)

	settings.WatchPaths = append(settings.WatchPaths, settings.Root)

	settings.ExcludePaths = append(settings.ExcludePaths, settings.TmpPath)
}
