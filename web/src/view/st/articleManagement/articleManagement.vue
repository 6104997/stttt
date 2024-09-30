
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="文章标题" prop="title">
         <el-input v-model="searchInfo.title" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="文章状态" prop="article_status">
            <el-select v-model="searchInfo.article_status" clearable placeholder="请选择" @clear="()=>{searchInfo.article_status=undefined}">
              <el-option v-for="(item,key) in stateOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
        </el-form-item>
        <el-form-item label="文章类型" prop="article_classification_id">
            <el-select v-model="searchInfo.article_classification_id" clearable placeholder="请选择" @clear="()=>{searchInfo.article_classification_id=undefined}">
              <el-option v-for="(item,key) in ArticleClassificationIdOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
        </el-form-item>
        <!-- accountTypeOptions -->
        <el-form-item label="账号类型" prop="account_type">
            <el-select v-model="searchInfo.account_type" clearable placeholder="请选择" @clear="()=>{searchInfo.account_type=undefined}">
              <el-option v-for="(item,key) in accountTypeOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
        </el-form-item>

        
        <el-form-item label="作者" prop="author">
         <el-input v-model="searchInfo.author" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="阅读次数" prop="view_count">
            
             <el-input v-model.number="searchInfo.view_count" placeholder="搜索条件" />

        </el-form-item>
            <el-form-item label="置顶" prop="select">
            <el-select v-model="searchInfo.select" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
        <el-form-item label="账号编号" prop="accountNumber">
            
             <el-input v-model.number="searchInfo.accountNumber" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="联系方式" prop="contact">
         <el-input v-model="searchInfo.contact" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="价格" prop="price">
            
             <el-input v-model.number="searchInfo.price" placeholder="搜索条件" />

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
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
        
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
          <el-table-column align="left" label="文章标题" prop="title" width="120" />
        <el-table-column align="left" label="文章状态" prop="article_status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.article_status,stateOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="文章类型" prop="article_classification_id" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.article_classification_id,ArticleClassificationIdOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="账号类型" prop="account_type" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.account_type,accountTypeOptions) }}
            </template>
        </el-table-column>

          <el-table-column align="left" label="作者" prop="author" width="120" />
          <el-table-column align="left" label="阅读次数" prop="view_count" width="120" />
        <el-table-column align="left" label="置顶" prop="select" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.select) }}</template>
        </el-table-column>
          <el-table-column align="left" label="账号编号" prop="accountNumber" width="120" />
          <el-table-column align="left" label="联系方式" prop="contact" width="120" />
          <el-table-column align="left" label="价格" prop="price" width="120" />
          <el-table-column label="预览图" prop="previewTheImage" width="200">
              <template #default="scope">
                <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.previewTheImage)" fit="cover"/>
              </template>
          </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateArticleManagementFunc(scope.row)">变更</el-button>
            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            

            <el-form-item label="文章标题:"  prop="title" >
              <el-input v-model="formData.title" :clearable="true"  placeholder="请输入资讯标题" />
            </el-form-item>
            <el-form-item label="文章内容:"  prop="content" >
              <RichEdit v-model="formData.content"/>
            </el-form-item>
            
            <el-form-item label="作者:"  prop="author" >
              <el-input v-model="formData.author" :clearable="true"  placeholder="请输入作者" />
            </el-form-item>
            <el-form-item label="文章状态:"  prop="article_status" >
              <el-select v-model="formData.article_status" placeholder="请选择文章状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in stateOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <!--  -->
            <el-form-item label="文章类型:"  prop="article_classification_id" >
              <el-select v-model="formData.article_classification_id" placeholder="请选择文章状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in ArticleClassificationIdOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>



            <el-form-item label="阅读次数:"  prop="view_count" >
              <el-input v-model.number="formData.view_count" :clearable="true" placeholder="请输入阅读次数" />
            </el-form-item>
            <el-form-item label="置顶:"  prop="select" >
              <el-switch v-model="formData.select" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="账号编号:"  prop="accountNumber" >
              <el-input v-model.number="formData.accountNumber" :clearable="true" placeholder="请输入账号编号" />
            </el-form-item>
            <el-form-item label="账号类型:"  prop="account_type" >
              <el-select v-model="formData.account_type" placeholder="请选择账号类型" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in accountTypeOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="联系方式:"  prop="contact" >
              <el-input v-model="formData.contact" :clearable="true"  placeholder="请输入联系方式" />
            </el-form-item>
            <el-form-item label="价格:"  prop="price" >
              <el-input v-model.number="formData.price" :clearable="true" placeholder="请输入价格" />
            </el-form-item>
            <el-form-item label="预览图:"  prop="previewTheImage" >
                <SelectImage
                 v-model="formData.previewTheImage"
                 file-type="image"
                />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="uuid">
                        {{ detailFrom.uuid }}
                    </el-descriptions-item>
                    <el-descriptions-item label="文章标题">
                        {{ detailFrom.title }}
                    </el-descriptions-item>
                    <el-descriptions-item label="文章状态">
                        {{ detailFrom.article_status }}
                    </el-descriptions-item>
                    <el-descriptions-item label="文章类型">
                        {{ detailFrom.article_classification_id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="作者">
                        {{ detailFrom.author }}
                    </el-descriptions-item>
                    <el-descriptions-item label="阅读次数">
                        {{ detailFrom.view_count }}
                    </el-descriptions-item>
                    <el-descriptions-item label="文章内容">
                        {{ detailFrom.content }}
                    </el-descriptions-item>
                    <el-descriptions-item label="置顶">
                        {{ detailFrom.select }}
                    </el-descriptions-item>
                    <el-descriptions-item label="账号编号">
                        {{ detailFrom.accountNumber }}
                    </el-descriptions-item>
                    <el-descriptions-item label="账号类型">
                        {{ detailFrom.account_type }}
                    </el-descriptions-item>
                    <!--  -->
                    <el-descriptions-item label="联系方式">
                        {{ detailFrom.contact }}
                    </el-descriptions-item>
                    <el-descriptions-item label="价格">
                        {{ detailFrom.price }}
                    </el-descriptions-item>
                    <el-descriptions-item label="预览图">
                            <el-image style="width: 50px; height: 50px" :preview-src-list="returnArrImg(detailFrom.previewTheImage)" :src="getUrl(detailFrom.previewTheImage)" fit="cover" />
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createArticleManagement,
  deleteArticleManagement,
  deleteArticleManagementByIds,
  updateArticleManagement,
  findArticleManagement,
  getArticleManagementList
} from '@/api/st/articleManagement'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
    name: 'ArticleManagement'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const stateOptions = ref([])
const accountTypeOptions = ref([])
const ArticleClassificationIdOptions = ref([])
const formData = ref({
            
            title: '',
            article_status: '',
            author: '',
            view_count: undefined,
            content: '',
            select: false,
            accountNumber: undefined,
            contact: '',
            price: undefined,
            previewTheImage: "",
            article_classification_id:'',
            account_type:''
        })



// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

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
  console.log(searchInfo.value);
  
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    if (searchInfo.value.select === ""){
        searchInfo.value.select=null
    }
    getTableData()
  })
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
  const table = await getArticleManagementList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const setOptions = async () =>{
    stateOptions.value = await getDictFunc('state')
    ArticleClassificationIdOptions.value = await getDictFunc('ArticleClassification')
    accountTypeOptions.value = await getDictFunc('AccountType')

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
            deleteArticleManagementFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteArticleManagementByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateArticleManagementFunc = async(row) => {
    const res = await findArticleManagement({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteArticleManagementFunc = async (row) => {
    const res = await deleteArticleManagement({ ID: row.ID })
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
        
        title: '',
        article_status: '',
        author: '',
        view_count: undefined,
        content: '',
        select: false,
        accountNumber: undefined,
        contact: '',
        price: undefined,
        previewTheImage: "",
        article_classification_id:'',
        account_type:''
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createArticleManagement(formData.value)
                  break
                case 'update':
                  res = await updateArticleManagement(formData.value)
                  break
                default:
                  res = await createArticleManagement(formData.value)
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


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findArticleManagement({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>