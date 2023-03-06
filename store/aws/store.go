package aws

func NewAwsOssStore() *AwsOssStore {
	return &AwsOssStore{}
}

//定义对象
type AwsOssStore struct {
}

//f方法
func (s *AwsOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	return nil
}
