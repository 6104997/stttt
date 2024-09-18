
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="昵称:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="true"  placeholder="请输入昵称" />
       </el-form-item>
        <el-form-item label="头像:" prop="headPortrait">
          <SelectImage v-model="formData.headPortrait" file-type="image"/>
       </el-form-item>
        <el-form-item label="openpid:" prop="openpid">
          <el-input v-model="formData.openpid" :clearable="true"  placeholder="请输入openpid" />
       </el-form-item>
        <el-form-item label="uuid:" prop="uuid">
          <el-input v-model="formData.uuid" :clearable="true"  placeholder="请输入uuid" />
       </el-form-item>
        <el-form-item label="当前积分:" prop="currentPoints">
          <el-input v-model.number="formData.currentPoints" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="签到:" prop="signIn">
          <el-input v-model.number="formData.signIn" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电话号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入电话号" />
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
  createAppUser,
  updateAppUser,
  findAppUser
} from '@/api/st/appUser'

defineOptions({
    name: 'AppUserForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            nickname: '',
            headPortrait: "",
            openpid: '',
            uuid: '',
            currentPoints: undefined,
            signIn: undefined,
            phone: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findAppUser({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
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
               res = await createAppUser(formData.value)
               break
             case 'update':
               res = await updateAppUser(formData.value)
               break
             default:
               res = await createAppUser(formData.value)
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
