package service

import (
	"context"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
	"helloworld/pkg"
)

// 查询地址
func (c *GreeterService) AddressList(ctx context.Context, in *v1.AddressListRequest) (*v1.AddressListReply, error) {
	var maps []biz.MapPoi
	var maplist []*v1.AddressList
	page := in.Page
	if page <= 0 {
		page = 1
	}

	pageSize := in.Size
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := c.db.GetDB().Offset(int(offset)).Limit(int(pageSize))
	if in.City != "" {
		query = query.Where("city LIKE ?", "%"+in.City+"%")
	}
	if in.District != "" {
		query = query.Where("district LIKE ?", "%"+in.District+"%")
	}
	query.Find(&maps)

	for _, i := range maps {
		maplist = append(maplist, &v1.AddressList{
			Address:  i.Address,
			Province: i.Province,
			City:     i.City,
			District: i.District,
			Tel:      i.Tel,
		})

	}

	return &v1.AddressListReply{List: maplist}, nil

}

// 获取当前定位
func (c *GreeterService) UserAddress(_ context.Context, in *v1.UserAddressRequest) (*v1.UserAddressReply, error) {
	address, province, city, district, err := pkg.GetLocationByLatLng(in.Latitude, in.Longitude)
	if err != nil {
		return &v1.UserAddressReply{
			Code:    500,
			Message: err.Error(),
		}, nil
	}
	return &v1.UserAddressReply{
		Code:     200,
		Message:  "获取成功",
		Address:  address,
		Province: province,
		City:     city,
		District: district,
	}, nil
}
