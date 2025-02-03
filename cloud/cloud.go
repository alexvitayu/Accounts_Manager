package cloud

type CloudDb struct {
	url string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func Read() ([]byte, error) {
	return []byte{}, nil
}

func Write([]byte) {

}
