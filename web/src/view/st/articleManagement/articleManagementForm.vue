
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        
        <el-form-item label="文章标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入文章标题" />
       </el-form-item>
        <el-form-item label="文章状态:" prop="article_status">
           <el-select v-model="formData.article_status" placeholder="请选择文章状态" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in stateOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
       <el-form-item label="文章状态:" prop="articleClassificationId">
           <el-select v-model="formData.articleClassificationId" placeholder="请选择文章类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in ArticleClassificationIdOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>


        <el-form-item label="作者:" prop="author">
          <el-input v-model="formData.author" :clearable="true"  placeholder="请输入作者" />
       </el-form-item>
        <el-form-item label="阅读次数:" prop="view_count">
          <el-input v-model.number="formData.view_count" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章内容:" prop="content">
          <RichEdit v-model="formData.content"/>
       </el-form-item>
        <el-form-item label="置顶:" prop="select">
          <el-switch v-model="formData.select" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="账号编号:" prop="accountNumber">
          <el-input v-model.number="formData.accountNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="联系方式:" prop="contact">
          <el-input v-model="formData.contact" :clearable="true"  placeholder="请输入联系方式" />
       </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input v-model.number="formData.price" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预览图:" prop="previewTheImage">
          <SelectImage v-model="formData.previewTheImage" file-type="image"/>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createArticleManagement,
  updateArticleManagement,
  findArticleManagement
} from '@/api/st/articleManagement'

defineOptions({
    name: 'ArticleManagementForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const stateOptions = ref([])
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
            articleClassificationId:'',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findArticleManagement({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    stateOptions.value = await getDictFunc('state')
    ArticleClassificationIdOptions.value = await getDictFunc('ArticleClassification')
}

init()
// 保存按钮
const save = async() => {
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
