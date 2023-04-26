package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
    "gorm.io/gorm"
)

type PlayListVideoService struct {
}

// CreatePlayListVideo 创建PlayListVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (playListVideoService *PlayListVideoService) CreatePlayListVideo(playListVideo *video.PlayListVideo) (err error) {
	err = global.GVA_DB.Create(playListVideo).Error
	return err
}

// DeletePlayListVideo 删除PlayListVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (playListVideoService *PlayListVideoService)DeletePlayListVideo(playListVideo video.PlayListVideo) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.PlayListVideo{}).Where("id = ?", playListVideo.ID).Update("deleted_by", playListVideo.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&playListVideo).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeletePlayListVideoByIds 批量删除PlayListVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (playListVideoService *PlayListVideoService)DeletePlayListVideoByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.PlayListVideo{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&video.PlayListVideo{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdatePlayListVideo 更新PlayListVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (playListVideoService *PlayListVideoService)UpdatePlayListVideo(playListVideo video.PlayListVideo) (err error) {
	err = global.GVA_DB.Save(&playListVideo).Error
	return err
}

// GetPlayListVideo 根据id获取PlayListVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (playListVideoService *PlayListVideoService)GetPlayListVideo(id uint) (playListVideo video.PlayListVideo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&playListVideo).Error
	return
}

// GetPlayListVideoInfoList 分页获取PlayListVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (playListVideoService *PlayListVideoService)GetPlayListVideoInfoList(info videoReq.PlayListVideoSearch) (list []video.PlayListVideo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&video.PlayListVideo{})
    var playListVideos []video.PlayListVideo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&playListVideos).Error
	return  playListVideos, total, err
}
