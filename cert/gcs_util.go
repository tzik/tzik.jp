package cert

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/storage"
)

type GCSUtil struct {
	client *storage.Client
}

func (g *GCSUtil) Close() {
	g.client.Close()
}

func getObject(c *storage.Client, key string) *storage.ObjectHandle {
	return c.Bucket("storage.tzik.jp").Object("autocert/" + key)
}

func (g *GCSUtil) Load(ctx context.Context, key string) ([]byte, error) {
	r, err := getObject(g.client, key).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return ioutil.ReadAll(r)
}

func (g *GCSUtil) Store(ctx context.Context, key string, data []byte) error {
	w := getObject(g.client, key).NewWriter(ctx)

	var err error = nil
	defer func() {
		if err == nil {
			w.Close()
		} else {
			w.CloseWithError(err)
		}
	}()

	for len(data) != 0 {
		n, err := w.Write(data)
		if err != nil {
			return err
		}
		data = data[n:]
	}
	return nil
}

func (g *GCSUtil) Delete(ctx context.Context, key string) error {
	return getObject(g.client, key).Delete(ctx)
}
