<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="碎片视频作者:" prop="author">
          <el-input v-model="formData.author" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="碎片视频名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="碎片视频审核状态（0 未审核 1 初审中 3 初审完成 4 复审审核中 5 复审完成）:" prop="status">
          <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="碎片视频字幕:" prop="subtitles">
          <el-input v-model="formData.subtitles" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="碎片视频上传时间:" prop="uploadedAt">
          <el-date-picker v-model="formData.uploadedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="碎片视频存储地址:" prop="url">
          <el-input v-model="formData.url" :clearable="true" placeholder="请输入" />
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
  name: 'Fragment'
}
</script>

<script setup>
import {
  createFragment,
  updateFragment,
  findFragment
} from '@/api/fragment'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            author: '',
            name: '',
            status: false,
            subtitles: '',
            uploadedAt: new Date(),
            url: '',
        })
// 验证规则
const rule = reactive({
               author : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               subtitles : [{
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
      const res = await findFragment({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.refragment
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
               res = await createFragment(formData.value)
               break
             case 'update':
               res = await updateFragment(formData.value)
               break
             default:
               res = await createFragment(formData.value)
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
