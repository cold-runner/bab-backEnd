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

        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateBabIdxImageFunc(scope.row)">变更</el-button>
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
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="200px">
        <el-form-item label="首页展示图:" prop="homePage">
          <el-upload
            v-model:file-list="formData.homePage"
            list-type="picture-card"
            :on-preview="handlePictureCardPreview"
            :auto-upload="false"
          >
            <el-icon><Plus /></el-icon>
          </el-upload>


        </el-form-item>
        <el-form-item label="询问页展示图:" prop="enquiry">
          <el-upload
            v-model:file-list="formData.enquiry"
            list-type="picture-card"
            :on-preview="handlePictureCardPreview"
            :auto-upload="false"
            :limit="1"
          >
            <el-icon><Plus /></el-icon>
          </el-upload>


        </el-form-item>
        <el-form-item label="咨询页展示图:" prop="news">
          <el-upload
            v-model:file-list="formData.news"
            list-type="picture-card"
            :on-preview="handlePictureCardPreview"
            :auto-upload="false"
            :limit="1"
          >
            <el-icon><Plus /></el-icon>
          </el-upload>


        </el-form-item>
        <el-dialog v-model="dialogVisible" :align-center="true" style="text-align: center">
          <img width="800" height="800" :src="dialogImageUrl" alt="Preview Image">
        </el-dialog>
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
  name: 'BabIdxImage'
}
</script>

<script setup>
import {
  createBabIdxImage,
  deleteBabIdxImage,
  deleteBabIdxImageByIds,
  updateBabIdxImage,
  findBabIdxImage,
  getBabIdxImageList
} from '@/api/babIdxImage'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { useUserStore } from '@/pinia/modules/user'
const userStore = useUserStore()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  homePage: [],
  enquiry: [],
  news: [],
})

// 验证规则
const rule = reactive({
  homePage: [{
    required: true,
    message: '请上传首页展示图！',
    trigger: ['input', 'blur'],
  }],
  enquiry: [{
    required: true,
    message: '请上传询问页展示图！',
    trigger: ['input', 'blur'],
  }],
  news: [{
    required: true,
    message: '请上传咨询页展示图！',
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
  const table = await getBabIdxImageList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteBabIdxImageFunc(row)
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
  const res = await deleteBabIdxImageByIds({ ids })
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
const updateBabIdxImageFunc = async(row) => {
  const res = await findBabIdxImage({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    const resp = res.data.reimg
    formData.value.CreatedAt = resp.CreatedAt
    formData.value.CreatedBy = resp.CreatedBy
    formData.value.DeletedBy = resp.DeletedBy
    formData.value.ID = resp.ID
    formData.value.UpdatedAt = resp.UpdatedAt
    for (let i = 0; i < resp.Hp.length; i++) {
      formData.value.homePage.push(resp.Hp[i])
    }
    for (let i = 0; i < resp.Enq.length; i++) {
      formData.value.enquiry.push(resp.Enq[i])
    }
    for (let i = 0; i < resp.Ne.length; i++) {
      formData.value.news.push(resp.Ne[i])
    }

    dialogFormVisible.value = true
  }
}

// 删除行
const deleteBabIdxImageFunc = async(row) => {
  const res = await deleteBabIdxImage({ ID: row.ID })
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
    homePage: [],
    enquiry: [],
    news: [],
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       const multiForm = new FormData()
       const blob = new Blob(['nil blob'], { type: 'text/plain' })

       for (let i = 0; i < formData.value.homePage.length; i++) {
         const file = formData.value.homePage[i]
         console.log(file)
         if (file.raw !== undefined) {
           multiForm.append('homePage[]', file.raw, file.name)
           console.log(multiForm)
         } else {
           multiForm.append('homePage[]', blob, file.name)
         }
       }

       for (let i = 0; i < formData.value.enquiry.length; i++) {
         const file = formData.value.enquiry[i]
         if (file.raw !== undefined) {
           multiForm.append('enquiry[]', file.raw, file.name) }
         else {
           multiForm.append('enquiry[]', blob, file.name)
         }
       }

       for (let i = 0; i < formData.value.news.length; i++) {
         const file = formData.value.news[i]
         if (file.raw !== undefined) {
           multiForm.append('news[]', file.raw, file.name) }
         else {
           multiForm.append('news[]', blob, file.name)
         }
       }

       switch (type.value) {
         case 'create':
           res = await createBabIdxImage(multiForm)
           break
         case 'update':
           multiForm.append('CreatedAt', formData.value.CreatedAt)
           multiForm.append('UpdatedAt', formData.value.UpdatedAt)
           multiForm.append('CreatedBy', formData.value.CreatedBy)
           multiForm.append('ID', formData.value.ID)
           res = await updateBabIdxImage(multiForm)
           break
         default:
           res = await createBabIdxImage(multiForm)
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


</script>

<style>
</style>
