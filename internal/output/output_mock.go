package output

type MockClient struct {
	OutputBuffer []byte
}

func (client MockClient) Write(outputBuffer []byte) error {
	client.OutputBuffer = outputBuffer
	return nil
}
