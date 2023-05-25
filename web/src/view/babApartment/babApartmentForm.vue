<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="公寓英文名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="城市英文名称:" prop="city">
          <el-input v-model="formData.city" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="户型:" prop="type">
          <el-input v-model="formData.type" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公寓所属公司:" prop="company">
          <el-input v-model="formData.company" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公寓介绍:" prop="introduction">
          <el-input v-model="formData.introduction" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公寓图片URL:" prop="image">
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
  name: 'BabApartment'
}
</script>

<script setup>
import {
  createBabApartment,
  updateBabApartment,
  findBabApartment
} from '@/api/babApartment'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            name: '',
            city: '',
            type: '',
            company: '',
            introduction: '',
            image: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '请填写公寓英文名称！',
                   trigger: ['input','blur'],
               }],
               city : [{
                   required: true,
                   message: '请选择公寓所属的城市！',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '请填写公寓所属户型！',
                   trigger: ['input','blur'],
               }],
               company : [{
                   required: true,
                   message: '请填写公寓所属的公司！',
                   trigger: ['input','blur'],
               }],
               introduction : [{
                   required: true,
                   message: '请填写公寓介绍信息！',
                   trigger: ['input','blur'],
               }],
               image : [{
                   required: true,
                   message: '请上传公寓图片！',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBabApartment({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rebabInfo
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
               res = await createBabApartment(formData.value)
               break
             case 'update':
               res = await updateBabApartment(formData.value)
               break
             default:
               res = await createBabApartment(formData.value)
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
