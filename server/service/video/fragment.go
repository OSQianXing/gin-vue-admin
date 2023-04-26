package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
)

type FragmentService struct {
}

// CreateFragment 创建Fragment记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentService *FragmentService) CreateFragment(fragment *video.Fragment) (err error) {
	err = global.GVA_DB.Create(fragment).Error
	return err
}

// DeleteFragment 删除Fragment记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentService *FragmentService)DeleteFragment(fragment video.Fragment) (err error) {
	err = global.GVA_DB.Delete(&fragment).Error
	return err
}

// DeleteFragmentByIds 批量删除Fragment记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentService *FragmentService)DeleteFragmentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]video.Fragment{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFragment 更新Fragment记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentService *FragmentService)UpdateFragment(fragment video.Fragment) (err error) {
	err = global.GVA_DB.Save(&fragment).Error
	return err
}

// GetFragment 根据id获取Fragment记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentService *FragmentService)GetFragment(id uint) (fragment video.Fragment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fragment).Error
	return
}

// GetFragmentInfoList 分页获取Fragment记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentService *FragmentService)GetFragmentInfoList(info videoReq.FragmentSearch) (list []video.Fragment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&video.Fragment{})
    var fragments []video.Fragment
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&fragments).Error
	return  fragments, total, err
}
