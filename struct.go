package main

type mainConfig struct {
	token      string
	parallel   int
	interval   int
	prefix     string
	debugMode  bool
	singleMode bool
	version    bool
	keepMode   bool
	blockSize  int
	hashCheck  bool
	passCode   string
	silentMode bool
}

type prepareSendResp struct {
	UploadToken  string `json:"uptoken"`
	TransferGUID string `json:"transferguid"`
	UniqueURL    string `json:"uniqueurl"`
	Prefix       string `json:"prefix"`
	QRCode       string `json:"qrcode"`
	Error        bool   `json:"error"`
	ErrorMessage string `json:"error_message"`
}

type uploadResponse struct {
	Ticket string `json:"ctx"`
	Hash   int    `json:"crc32"`
}

type finishResponse struct {
	TempDownloadCode string `json:"tempDownloadCode"`
	Status           bool   `json:"complete"`
}
