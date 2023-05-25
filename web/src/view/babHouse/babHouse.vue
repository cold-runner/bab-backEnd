<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
        </el-form-item>
        <el-form-item label="房屋英文名">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="房屋中文名">
          <el-input v-model="searchInfo.chineseName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="房屋城市">
          <el-input v-model="searchInfo.city" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="房屋公司">
          <el-input v-model="searchInfo.company" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="房屋公寓">
          <el-input v-model="searchInfo.apartment" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="房屋类型">
          <el-input v-model="searchInfo.type" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="设施">
          <el-input v-model="searchInfo.facility" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="租期">
          <el-input v-model="searchInfo.leasePeriod" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="介绍">
          <el-input v-model="searchInfo.introduction" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="房屋价格">

          <el-input v-model.number="searchInfo.startPrice" placeholder="搜索条件（起）" />
          —
          <el-input v-model.number="searchInfo.endPrice" placeholder="搜索条件（止）" />

        </el-form-item>
        <el-form-item label="房屋面积">

          <el-input v-model.number="searchInfo.startArea" placeholder="搜索条件（起）" />
          —
          <el-input v-model.number="searchInfo.endArea" placeholder="搜索条件（止）" />

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
        @sort-change="sortChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="100">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="center" show-overflow-tooltip="true" label="房屋英文名" prop="name" width="120" />
        <el-table-column align="center" show-overflow-tooltip="true" label="房屋中文名" prop="chineseName" width="120" />
        <el-table-column align="center" show-overflow-tooltip="true" label="房屋城市" prop="city" width="100" />
        <el-table-column align="center" show-overflow-tooltip="true" label="房屋公司" prop="company" width="120" />
        <el-table-column align="center" show-overflow-tooltip="true" label="房屋公寓" prop="apartment" width="120" />
        <el-table-column align="center" show-overflow-tooltip="true" label="房屋类型" prop="type" width="100" />
        <el-table-column align="center" show-overflow-tooltip="true" label="房屋价格" prop="price" width="120" />
        <el-table-column sortable align="center" show-overflow-tooltip="true" label="房屋面积" prop="area" width="120" />
        <el-table-column align="center" show-overflow-tooltip="true" label="设施" prop="facility" width="120" />
        <el-table-column align="center" show-overflow-tooltip="true" label="租期" prop="leasePeriod" width="120" />
        <el-table-column align="center" show-overflow-tooltip="true" label="介绍" prop="introduction" width="120" />
        <el-table-column align="center" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateBabHouseFunc(scope.row)">变更</el-button>
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
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="150px">
        <el-form-item label="房屋英文名:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房屋中文名:" prop="chineseName">
          <el-input v-model="formData.chineseName" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房屋城市:" prop="city">
          <el-autocomplete
            v-model="formData.city"
            :fetch-suggestions="querySearchCity"
            clearable
            class="inline-input w-50"
            placeholder="请输入房屋所在城市"
          />
        </el-form-item>
        <el-form-item label="房屋公司:" prop="company">
          <el-autocomplete
            v-model="formData.company"
            :fetch-suggestions="querySearchCompany"
            clearable
            class="inline-input w-50"
            placeholder="请输入房屋所属公司"
          />
        </el-form-item>
        <el-form-item label="房屋公寓:" prop="apartment">
          <el-autocomplete
            v-model="formData.apartment"
            :fetch-suggestions="querySearchApartment"
            clearable
            class="inline-input w-80"
            placeholder="请输入房屋所属公寓"
          />
        </el-form-item>
        <el-form-item label="房屋类型:" prop="type">
          <el-input v-model="formData.type" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房屋价格:(单位￡)" prop="price">
          <el-input-number v-model="formData.price" style="width:100%" :precision="2" :clearable="false" />
        </el-form-item>
        <el-form-item label="房屋面积(单位㎡):" prop="area">
          <el-input v-model.number="formData.area" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="设施:" prop="facility">
          <el-input v-model="formData.facility" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="租期:" prop="leasePeriod">
          <el-input v-model="formData.leasePeriod" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="介绍:" prop="introduction">
          <el-input v-model="formData.introduction" type="textarea" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="房型图片:" prop="image">
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
  name: 'BabHouse'
}
</script>

<script setup>
import {
  createBabHouse,
  deleteBabHouse,
  deleteBabHouseByIds,
  updateBabHouse,
  findBabHouse,
  getBabHouseList,

  getCompanyList,
  getCityList,
  getApartmentList
} from '@/api/babHouse'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { useUserStore } from '@/pinia/modules/user'

const userStore = useUserStore()

const loadAll = () => {
  companyList()
  apartmentList()
  citiesList()
}

onMounted(() => {
  loadAll()
})

// 自动化生成的字典（可能为空）以及字段
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
  leasePeriod: '',
  introduction: '',
  image: [],
})
const stored = {
  companyList: [],
  apartmentList: [],
  cityList: []
}

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
const querySearchApartment = (queryString, cb) => {
  const filteredApartmentList = stored.apartmentList.filter(city => city.toLowerCase().indexOf(queryString.toLowerCase()) === 0)
  const cityObjects = filteredApartmentList.map(city => ({ value: city, label: city }))
  cb(cityObjects)
}

const companyList = async() => {
  const resp = await getCompanyList()
  for (let i = 0; i < resp.data.length; i++) {
    stored.companyList.push(resp.data[i])
  }
}

const apartmentList = async() => {
  const resp = await getApartmentList()
  for (let i = 0; i < resp.data.length; i++) {
    stored.apartmentList.push(resp.data[i])
  }
}

const citiesList = async() => {
  const resp = await getCityList()
  for (let i = 0; i < resp.data.length; i++) {
    stored.cityList.push(resp.data[i])
  }
}

// 验证规则
const rule = reactive({
  name: [{
    required: true,
    message: '请填写房屋英文名！',
    trigger: ['input', 'blur'],
  }],
  chineseName: [{
    required: true,
    message: '请填写房屋中文名！',
    trigger: ['input', 'blur'],
  }],
  city: [{
    required: true,
    message: '请选择房屋所在的城市！',
    trigger: ['input', 'blur'],
  }],
  company: [{
    required: true,
    message: '请选择房屋所属的公司！',
    trigger: ['input', 'blur'],
  }],
  apartment: [{
    required: true,
    message: '请填写房屋所属的公寓！',
    trigger: ['input', 'blur'],
  }],
  type: [{
    required: true,
    message: '请填写房屋类型！',
    trigger: ['input', 'blur'],
  }],
  price: [{
    required: true,
    message: '请填写房屋价格！',
    trigger: ['input', 'blur'],
  }],
  area: [{
    required: true,
    message: '请填写房屋面积！',
    trigger: ['input', 'blur'],
  }],
  introduction: [{
    required: true,
    message: '请填写房型介绍！',
    trigger: ['input', 'blur'],
  }],
  image: [{
    required: true,
    message: '请上传房型图片！',
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
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}

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
  const table = await getBabHouseList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteBabHouseFunc(row)
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
  const res = await deleteBabHouseByIds({ ids })
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
const updateBabHouseFunc = async(row) => {
  const res = await findBabHouse({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    const resp = res.data.rehouseInfo
    formData.value.CreatedAt = resp.CreatedAt
    formData.value.CreatedBy = resp.CreatedBy
    formData.value.DeletedBy = resp.DeletedBy
    formData.value.ID = resp.ID
    formData.value.UpdatedAt = resp.UpdatedAt
    formData.value.name = resp.name
    formData.value.chineseName = resp.chineseName
    formData.value.city = resp.city
    formData.value.company = resp.company
    formData.value.apartment = resp.apartment
    formData.value.type = resp.type
    formData.value.price = resp.price
    formData.value.introduction = resp.introduction
    formData.value.area = resp.area

    // if (resp.facility !== null) {
    //   for (let i = 0; i < resp.facility.length; i++) {
    //     formData.value.facility.push(resp.facility[i])
    //   }
    // }
    formData.value.facility = resp.facility
    // if (resp.leasePeriod !== null) {
    //   for (let i = 0; i < resp.leasePeriod; i++) {
    //     formData.value.leasePeriod.push(resp.leasePeriod[i])
    //   }
    // }
    formData.value.leasePeriod = resp.leasePeriod
    for (let i = 0; i < resp.image.length; i++) {
      formData.value.image.push(resp.image[i])
    }

    dialogFormVisible.value = true
  }
}

// 删除行
const deleteBabHouseFunc = async(row) => {
  const res = await deleteBabHouse({ ID: row.ID })
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
    chineseName: '',
    city: '',
    company: '',
    apartment: '',
    type: '',
    price: 0,
    area: 0,
    introduction: '',
    facility: '',
    leasePeriod: '',
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
       multiForm.append('chineseName', formData.value.chineseName)
       multiForm.append('city', formData.value.city)
       multiForm.append('company', formData.value.company)
       multiForm.append('apartment', formData.value.apartment)
       multiForm.append('type', formData.value.type)
       multiForm.append('price', formData.value.price)
       multiForm.append('introduction', formData.value.introduction)
       multiForm.append('area', formData.value.area)

       // for (let i = 0; i < formData.value.leasePeriod.length; i++) {
       //   multiForm.append('leasePeriod[]', formData.value.leasePeriod[i])
       // }
       multiForm.append('leasePeriod', formData.value.leasePeriod)

       // for (let i = 0; i < formData.value.facility.length; i++) {
       //   multiForm.append('facility[]', formData.value.facility[i])
       // }
       multiForm.append('facility', formData.value.facility)

       const blob = new Blob(['nil blob'], { type: 'text/plain' })
       for (let i = 0; i < formData.value.image.length; i++) {
         const file = formData.value.image[i]
         if (file.raw !== undefined) { multiForm.append('images[]', file.raw, file.name) } else {
           multiForm.append('images[]', blob, file.name)
         }
       }
       switch (type.value) {
         case 'create':
           res = await createBabHouse(multiForm)
           break
         case 'update':
           multiForm.append('CreatedAt', formData.value.CreatedAt)
           multiForm.append('UpdatedAt', formData.value.UpdatedAt)
           multiForm.append('CreatedBy', formData.value.CreatedBy)
           multiForm.append('ID', formData.value.ID)
           res = await updateBabHouse(multiForm)
           break
         default:
           res = await createBabHouse(multiForm)
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
