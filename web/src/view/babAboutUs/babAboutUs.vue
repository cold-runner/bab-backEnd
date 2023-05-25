<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
<!--      <div class="gva-btn-list">-->
<!--        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>-->
<!--        <el-popover v-model:visible="deleteVisible" placement="top" width="160">-->
<!--          <p>确定要删除吗？</p>-->
<!--          <div style="text-align: right; margin-top: 8px;">-->
<!--            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>-->
<!--            <el-button type="primary" @click="onDelete">确定</el-button>-->
<!--          </div>-->
<!--          <template #reference>-->
<!--            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>-->
<!--          </template>-->
<!--        </el-popover>-->
<!--      </div>-->
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="100">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="center" label="关于我们介绍" prop="message" width="800" />
        <el-table-column align="center" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateBabAboutUsFunc(scope.row)">变更</el-button>
<!--            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>-->
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="180px">
        <el-form-item label="关于我们介绍:" prop="message">
          <el-input type="textarea" v-model="formData.message" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图片:" prop="image">
          <el-upload
            v-model:file-list="formData.image"
            list-type="picture-card"
            :headers="{ 'x-token': userStore.token }"
            :on-preview="handlePictureCardPreview"
            :on-success="handleSuccess"
            :auto-upload="false"
          >
            <el-icon><Plus /></el-icon>
          </el-upload>

          <el-dialog v-model="dialogVisible" :align-center="true" style="text-align: center">
            <img :src="dialogImageUrl" alt="Preview Image">
          </el-dialog>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
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
  deleteBabAboutUs,
  deleteBabAboutUsByIds,
  updateBabAboutUs,
  findBabAboutUs,
  getBabAboutUsList
} from '@/api/babAboutUs'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  message: '',
  image: [],
})

// 验证规则
const rule = reactive({
  message: [{
    required: true,
    message: '请填写关于我们的介绍信息！',
    trigger: ['input', 'blur'],
  }],
  image: [{
    required: true,
    message: '请上传图片！',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getBabAboutUsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteBabAboutUsFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteBabAboutUsByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateBabAboutUsFunc = async(row) => {
  const res = await findBabAboutUs({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    const resp = res.data.reabus
    formData.value.CreatedAt = resp.CreatedAt
    formData.value.CreatedBy = resp.CreatedBy
    formData.value.DeletedBy = resp.DeletedBy
    formData.value.ID = resp.ID
    formData.value.UpdatedAt = resp.UpdatedAt
    formData.value.message = resp.message
    for (let i = 0; i < resp.image.length; i++) {
      formData.value.image.push(resp.image[i])
    }
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteBabAboutUsFunc = async(row) => {
  const res = await deleteBabAboutUs({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    message: '',
    image: [],
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       const multiForm = new FormData()
       multiForm.append('message', formData.value.message)
       const blob = new Blob(['nil blob'], { type: 'text/plain' })
       for (let i = 0; i < formData.value.image.length; i++) {
         const file = formData.value.image[i]
         if (file.raw !== undefined) { multiForm.append('images[]', file.raw, file.name) } else {
           multiForm.append('images[]', blob, file.name)
         }
       }
       switch (type.value) {
         case 'create':
           res = await createBabAboutUs(multiForm)
           break
         case 'update':
           multiForm.append('CreatedAt', formData.value.CreatedAt)
           multiForm.append('UpdatedAt', formData.value.UpdatedAt)
           multiForm.append('CreatedBy', formData.value.CreatedBy)
           multiForm.append('ID', formData.value.ID)
           res = await updateBabAboutUs(multiForm)
           break
         default:
           res = await createBabAboutUs(multiForm)
           break
       }
       if (res.code === 0) {
         ElMessage({
           type: 'success',
           message: '创建/更改成功'
         })
         closeDialog()
         getTableData()
       }
     })
}
const dialogImageUrl = ref('')
const dialogVisible = ref(false)

const handlePictureCardPreview = (uploadFile) => {
  dialogImageUrl.value = uploadFile.url
  dialogVisible.value = true
}

const handleSuccess = (response, file, fileList) => {
  const { name, url } = response.data
  for (let i = 0; i < fileList.length; i++) {
    if (fileList[i].uid === file.uid) {
      // 更新该文件对象的name和url属性
      fileList[i].name = name
      fileList[i].url = url
      break
    }
  }
}
</script>

<style>
</style>
