<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="审核不通过原因:" prop="auditReason">
          <el-input v-model="formData.auditReason" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="auditedAt字段:" prop="auditedAt">
          <el-date-picker v-model="formData.auditedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="播单ID:" prop="playListId">
          <el-input v-model.number="formData.playListId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="成品审核不通过原因:" prop="prodAuditReason">
          <el-input v-model="formData.prodAuditReason" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="prodAuditedAt字段:" prop="prodAuditedAt">
          <el-date-picker v-model="formData.prodAuditedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="成品ID:" prop="prodId">
          <el-input v-model.number="formData.prodId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="成品状态 0-待审核 1-审核通过 2-审核不通过:" prop="prodStatus">
          <el-switch v-model="formData.prodStatus" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="prodUserId字段:" prop="prodUserId">
          <el-input v-model.number="formData.prodUserId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="短视频状态 0-待审核 1-审核通过 2-审核不通过:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="合成视频地址:" prop="url">
          <el-input v-model="formData.url" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核人ID:" prop="userId">
          <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="版本:" prop="version">
          <el-input v-model="formData.version" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PlayListVideo'
}
</script>

<script setup>
import {
  createPlayListVideo,
  updatePlayListVideo,
  findPlayListVideo
} from '@/api/playListVideo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            auditReason: '',
            auditedAt: new Date(),
            playListId: 0,
            prodAuditReason: '',
            prodAuditedAt: new Date(),
            prodId: 0,
            prodStatus: false,
            prodUserId: 0,
            status: false,
            url: '',
            userId: 0,
            version: '',
        })
// 验证规则
const rule = reactive({
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPlayListVideo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.replayListVideo
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createPlayListVideo(formData.value)
               break
             case 'update':
               res = await updatePlayListVideo(formData.value)
               break
             default:
               res = await createPlayListVideo(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
