<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="房屋英文名:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房屋中文名:" prop="chineseName">
          <el-input v-model="formData.chineseName" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房屋所在的城市:" prop="city">
          <el-input v-model="formData.city" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房屋所属的公司:" prop="company">
          <el-input v-model="formData.company" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房屋所属的公寓:" prop="apartment">
          <el-input v-model="formData.apartment" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房屋类型:" prop="type">
          <el-input v-model="formData.type" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房屋价格:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="false"></el-input-number>
        </el-form-item>
        <el-form-item label="房屋面积:" prop="area">
          <el-input v-model.number="formData.area" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="设施:" prop="facility">
          <el-input v-model="formData.facility" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="介绍:" prop="introduction">
          <el-input v-model="formData.introduction" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房型图片:" prop="image">
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
  name: 'BabHouse'
}
</script>

<script setup>
import {
  createBabHouse,
  updateBabHouse,
  findBabHouse
} from '@/api/babHouse'

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
            chineseName: '',
            city: '',
            company: '',
            apartment: '',
            type: '',
            price: 0,
            area: 0,
            facility: '',
            introduction: '',
            image: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '请填写房屋英文名！',
                   trigger: ['input','blur'],
               }],
               chineseName : [{
                   required: true,
                   message: '请填写房屋中文名！',
                   trigger: ['input','blur'],
               }],
               city : [{
                   required: true,
                   message: '请选择房屋所在的城市！',
                   trigger: ['input','blur'],
               }],
               company : [{
                   required: true,
                   message: '请选择房屋所属的公司！',
                   trigger: ['input','blur'],
               }],
               apartment : [{
                   required: true,
                   message: '请填写房屋所属的公寓！',
                   trigger: ['input','blur'],
               }],
               type : [{
                   required: true,
                   message: '请填写房屋类型！',
                   trigger: ['input','blur'],
               }],
               price : [{
                   required: true,
                   message: '请填写房屋价格！',
                   trigger: ['input','blur'],
               }],
               area : [{
                   required: true,
                   message: '请填写房屋面积！',
                   trigger: ['input','blur'],
               }],
               introduction : [{
                   required: true,
                   message: '请填写房型介绍！',
                   trigger: ['input','blur'],
               }],
               image : [{
                   required: true,
                   message: '请上传房型图片！',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBabHouse({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rehouseInfo
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
               res = await createBabHouse(formData.value)
               break
             case 'update':
               res = await updateBabHouse(formData.value)
               break
             default:
               res = await createBabHouse(formData.value)
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
