<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户名字:" prop="firstName">
          <el-input v-model="formData.firstName" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户姓氏:" prop="lastName">
          <el-input v-model="formData.lastName" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户要咨询的城市:" prop="city">
          <el-input v-model="formData.city" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户手机号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户微信号:" prop="wechatNo">
          <el-input v-model="formData.wechatNo" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户邮箱地址:" prop="email">
          <el-input v-model="formData.email" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="用户留言:" prop="message">
          <el-input v-model="formData.message" :clearable="false" placeholder="请输入" />
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
  name: 'BabEnquiry'
}
</script>

<script setup>
import {
  createBabEnquiry,
  updateBabEnquiry,
  findBabEnquiry
} from '@/api/babEnquiry'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            firstName: '',
            lastName: '',
            city: '',
            phone: '',
            wechatNo: '',
            email: '',
            message: '',
        })
// 验证规则
const rule = reactive({
               firstName : [{
                   required: true,
                   message: '请填写用户名字！',
                   trigger: ['input','blur'],
               }],
               lastName : [{
                   required: true,
                   message: '请填写用户姓氏！',
                   trigger: ['input','blur'],
               }],
               city : [{
                   required: true,
                   message: '请输入你要咨询的城市！',
                   trigger: ['input','blur'],
               }],
               phone : [{
                   required: true,
                   message: '请填写用户的手机号码！',
                   trigger: ['input','blur'],
               }],
               email : [{
                   required: true,
                   message: '请填写用户邮箱地址！',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBabEnquiry({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reenquiryInfo
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
               res = await createBabEnquiry(formData.value)
               break
             case 'update':
               res = await updateBabEnquiry(formData.value)
               break
             default:
               res = await createBabEnquiry(formData.value)
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
