<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="关于我们介绍:" prop="message">
          <el-input v-model="formData.message" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="关于我们图片:" prop="image">
          <el-input v-model="formData.image" :clearable="false" placeholder="请输入" />
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
  name: 'BabAboutUs'
}
</script>

<script setup>
import {
  createBabAboutUs,
  updateBabAboutUs,
  findBabAboutUs
} from '@/api/babAboutUs'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            message: '',
            image: '',
        })
// 验证规则
const rule = reactive({
               message : [{
                   required: true,
                   message: '请填写关于我们的介绍信息！',
                   trigger: ['input','blur'],
               }],
               image : [{
                   required: true,
                   message: '请上传图片！',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBabAboutUs({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reabus
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
               res = await createBabAboutUs(formData.value)
               break
             case 'update':
               res = await updateBabAboutUs(formData.value)
               break
             default:
               res = await createBabAboutUs(formData.value)
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
