<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="城市英文名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="城市中文名称:" prop="chineseName">
          <el-input v-model="formData.chineseName" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="城市介绍:" prop="introdcution">
          <el-input v-model="formData.introduction" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="城市展示图片URL:" prop="image">
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
  name: 'BabCity'
}
</script>

<script setup>
import {
  createBabCity,
  updateBabCity,
  findBabCity
} from '@/api/babCity'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  name: '',
  chineseName: '',
  introdcution: '',
  image: '',
})
// 验证规则
const rule = reactive({
  name: [{
    required: true,
    message: '请填写城市英文名称',
    trigger: ['input', 'blur'],
  }],
  image: [{
    required: true,
    message: '请上传城市展示图片！',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findBabCity({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.recityInfo
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate(async(valid) => {
        if (!valid) return
        let res
        switch (type.value) {
          case 'create':
            res = await createBabCity(formData.value)
            break
          case 'update':
            res = await updateBabCity(formData.value)
            break
          default:
            res = await createBabCity(formData.value)
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
