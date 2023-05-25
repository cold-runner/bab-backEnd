<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
        </el-form-item>
        <el-form-item label="公寓英文名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="城市英文名称">
          <el-input v-model="searchInfo.city" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="户型">
          <el-input v-model="searchInfo.type" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="公寓所属公司">
          <el-input v-model="searchInfo.company" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="center" label="日期" width="100">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="center" show-overflow-tooltip="true" label="公寓英文名称" prop="name" width="300" />
        <el-table-column align="center" label="城市英文名称" prop="city" width="120" />
        <el-table-column align="center" show-overflow-tooltip="true" label="户型" prop="type" width="120" />
        <el-table-column align="center" show-overflow-tooltip="true" label="公寓所属公司" prop="company" width="120" />
        <el-table-column align="center" show-overflow-tooltip="true" label="公寓介绍" prop="introduction" width="500" />
        <el-table-column align="center" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateBabApartmentFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
        <el-form-item label="公寓英文名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="城市英文名称:" prop="city">
          <el-autocomplete
              v-model="formData.city"
              :fetch-suggestions="querySearchCity"
              clearable
              class="inline-input w-50"
              placeholder="请输入房屋所在城市"
          />
        </el-form-item>
        <el-form-item label="公寓所属公司:" prop="company">
            <el-autocomplete
                v-model="formData.company"
                :fetch-suggestions="querySearchCompany"
                clearable
                class="inline-input w-50"
                placeholder="请输入房屋所属公司"
            />
          </el-form-item>
        <el-form-item label="户型:" prop="type">
          <el-input v-model="formData.type" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公寓介绍:" prop="introduction">
          <el-input type="textarea" v-model="formData.introduction" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公寓图片:" prop="image">
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
  name: 'BabApartment'
}
</script>

<script setup>
import {
  createBabApartment,
  deleteBabApartment,
  deleteBabApartmentByIds,
  updateBabApartment,
  findBabApartment,
  getBabApartmentList
} from '@/api/babApartment'
import { useUserStore } from '@/pinia/modules/user'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import {ref, reactive, onMounted} from 'vue'
import { Plus } from '@element-plus/icons-vue'
import {getCityList, getCompanyList} from "@/api/babHouse";

const stored = {
  companyList: [],
  cityList: []
}
const loadAll = () => {
  companyList()
  citiesList()
}

onMounted(() => {
  loadAll()
})

const querySearchCity = (queryString, cb) => {
  const filteredCityList = stored.cityList.filter(city => city.toLowerCase().indexOf(queryString.toLowerCase()) === 0)
  const cityObjects = filteredCityList.map(city => ({ value: city, label: city }))
  cb(cityObjects)
}

const querySearchCompany = (queryString, cb) => {
  const filteredCompanyList = stored.companyList.filter(city => city.toLowerCase().indexOf(queryString.toLowerCase()) === 0)
  const cityObjects = filteredCompanyList.map(city => ({ value: city, label: city }))
  cb(cityObjects)
}
const companyList = async() => {
  const resp = await getCompanyList()
  for (let i = 0; i < resp.data.length; i++) {
    stored.companyList.push(resp.data[i])
  }
}
const citiesList = async() => {
  const resp = await getCityList()
  for (let i = 0; i < resp.data.length; i++) {
    stored.cityList.push(resp.data[i])
  }
}

const userStore = useUserStore()
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  city: '',
  type: '',
  company: '',
  introduction: '',
  image: [],
})

// 验证规则
const rule = reactive({
  name: [{
    required: true,
    message: '请填写公寓英文名称！',
    trigger: ['input', 'blur'],
  }],
  city: [{
    required: true,
    message: '请选择公寓所属的城市！',
    trigger: ['input', 'blur'],
  }],
  type: [{
    required: true,
    message: '请填写公寓所属户型！',
    trigger: ['input', 'blur'],
  }],
  company: [{
    required: true,
    message: '请填写公寓所属的公司！',
    trigger: ['input', 'blur'],
  }],
  introduction: [{
    required: true,
    message: '请填写公寓介绍信息！',
    trigger: ['input', 'blur'],
  }],
  image: [{
    required: true,
    message: '请上传公寓图片！',
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
  const table = await getBabApartmentList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteBabApartmentFunc(row)
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
  const res = await deleteBabApartmentByIds({ ids })
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
const updateBabApartmentFunc = async(row) => {
  const res = await findBabApartment({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    const resp = res.data.rebabInfo
    formData.value.CreatedAt = resp.CreatedAt
    formData.value.CreatedBy = resp.CreatedBy
    formData.value.DeletedBy = resp.DeletedBy
    formData.value.ID = resp.ID
    formData.value.UpdatedAt = resp.UpdatedAt
    for (let i = 0; i < resp.image.length; i++) {
      formData.value.image.push(resp.image[i])
    }
    formData.value.name = resp.name
    formData.value.city = resp.city
    formData.value.type = resp.type
    formData.value.company = resp.company
    formData.value.introduction = resp.introduction
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteBabApartmentFunc = async(row) => {
  const res = await deleteBabApartment({ ID: row.ID })
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
    name: '',
    city: '',
    type: '',
    company: '',
    introduction: '',
    image: [],
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       const multiForm = new FormData()
       multiForm.append('name', formData.value.name)
       multiForm.append('city', formData.value.city)
       multiForm.append('type', formData.value.type)
       multiForm.append('company', formData.value.company)
       multiForm.append('introduction', formData.value.introduction)

       const blob = new Blob(['nil blob'], { type: 'text/plain' })
       for (let i = 0; i < formData.value.image.length; i++) {
         const file = formData.value.image[i]
         if (file.raw !== undefined) { multiForm.append('images[]', file.raw, file.name) } else {
           multiForm.append('images[]', blob, file.name)
         }
       }
       switch (type.value) {
         case 'create':
           res = await createBabApartment(multiForm)
           break
         case 'update':
           multiForm.append('CreatedAt', formData.value.CreatedAt)
           multiForm.append('UpdatedAt', formData.value.UpdatedAt)
           multiForm.append('CreatedBy', formData.value.CreatedBy)

           multiForm.append('ID', formData.value.ID)
           res = await updateBabApartment(multiForm)
           break
         default:
           res = await createBabApartment(multiForm)
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
