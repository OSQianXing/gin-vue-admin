package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
    "gorm.io/gorm"
)

type FragmentAuditService struct {
}

// CreateFragmentAudit 创建FragmentAudit记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentAuditService *FragmentAuditService) CreateFragmentAudit(fragmentAudit *video.FragmentAudit) (err error) {
	err = global.GVA_DB.Create(fragmentAudit).Error
	return err
}

// DeleteFragmentAudit 删除FragmentAudit记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentAuditService *FragmentAuditService)DeleteFragmentAudit(fragmentAudit video.FragmentAudit) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.FragmentAudit{}).Where("id = ?", fragmentAudit.ID).Update("deleted_by", fragmentAudit.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&fragmentAudit).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteFragmentAuditByIds 批量删除FragmentAudit记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentAuditService *FragmentAuditService)DeleteFragmentAuditByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.FragmentAudit{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&video.FragmentAudit{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateFragmentAudit 更新FragmentAudit记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentAuditService *FragmentAuditService)UpdateFragmentAudit(fragmentAudit video.FragmentAudit) (err error) {
	err = global.GVA_DB.Save(&fragmentAudit).Error
	return err
}

// GetFragmentAudit 根据id获取FragmentAudit记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentAuditService *FragmentAuditService)GetFragmentAudit(id uint) (fragmentAudit video.FragmentAudit, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&fragmentAudit).Error
	return
}

// GetFragmentAuditInfoList 分页获取FragmentAudit记录
// Author [piexlmax](https://github.com/piexlmax)
func (fragmentAuditService *FragmentAuditService)GetFragmentAuditInfoList(info videoReq.FragmentAuditSearch) (list []video.FragmentAudit, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&video.FragmentAudit{})
    var fragmentAudits []video.FragmentAudit
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&fragmentAudits).Error
	return  fragmentAudits, total, err
}
