<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="首页展示图:" prop="homePage">
          <el-input v-model="formData.homePage" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="询问页展示图:" prop="enquiry">
          <el-input v-model="formData.enquiry" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="咨询页展示图:" prop="news">
          <el-input v-model="formData.news" :clearable="false" placeholder="请输入" />
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
  name: 'BabIdxImage'
}
</script>

<script setup>
import {
  createBabIdxImage,
  updateBabIdxImage,
  findBabIdxImage
} from '@/api/babIdxImage'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            homePage: '',
            enquiry: '',
            news: '',
        })
// 验证规则
const rule = reactive({
               homePage : [{
                   required: true,
                   message: '请上传首页展示图！',
                   trigger: ['input','blur'],
               }],
               enquiry : [{
                   required: true,
                   message: '请上传询问页展示图！',
                   trigger: ['input','blur'],
               }],
               news : [{
                   required: true,
                   message: '请上传咨询页展示图！',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBabIdxImage({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reimg
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
               res = await createBabIdxImage(formData.value)
               break
             case 'update':
               res = await updateBabIdxImage(formData.value)
               break
             default:
               res = await createBabIdxImage(formData.value)
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
