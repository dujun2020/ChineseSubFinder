package common

import "time"

const HTMLTimeOut = 60 * time.Second	// HttpClient 超时时间

const DownloadSubsPerSite = 1 // 默认，每个网站下载一个字幕，允许额外传参调整

const DebugFolder = "debugThings"