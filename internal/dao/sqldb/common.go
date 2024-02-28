package sqldb

import "github.com/Forget-C/http-structer/internal/model"

type Common[t model.Interface] struct {
}

func (c *Common[t]) Get(obj model.Interface) (*t, error) {
	expr, err := obj.GetGettingQuery()
	if err != nil {
		return nil, err
	}
	var data = new(t)
	err = Client.Where(expr).First(data).Error
	return data, err
}

func (c *Common[t]) Create(obj t) error {
	if err := obj.WriteAvailable(); err != nil {
		return err
	}
	return Client.Create(obj).Error
}

func (c *Common[t]) Update(obj model.Interface) error {
	expr, err := obj.GetGettingQuery()
	if err != nil {
		return err
	}
	return Client.Where(expr).Updates(obj).Error
}

func (c *Common[t]) Delete(obj model.Interface) error {
	expr, err := obj.GetGettingQuery()
	if err != nil {
		return err
	}
	return Client.Where(expr).Delete(obj).Error
}
