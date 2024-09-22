
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="房间号:" prop="roomNumber">
          <el-input v-model="formData.roomNumber" :clearable="true"  placeholder="请输入房间号" />
       </el-form-item>
        <el-form-item label="用户id1:" prop="userID1">
          <el-input v-model="formData.userID1" :clearable="true"  placeholder="请输入用户id1" />
       </el-form-item>
        <el-form-item label="用户id2:" prop="userID2">
          <el-input v-model="formData.userID2" :clearable="true"  placeholder="请输入用户id2" />
       </el-form-item>
        <el-form-item label="用户id3:" prop="userID3">
          <el-input v-model="formData.userID3" :clearable="true"  placeholder="请输入用户id3" />
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
  createUserGoldData,
  updateUserGoldData,
  findUserGoldData
} from '@/api/goldenHouseManagement/userGoldData'

defineOptions({
    name: 'UserGoldDataForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            roomNumber: '',
            userID1: '',
            userID2: '',
            userID3: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUserGoldData({ ID: route.query.id })
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
               res = await createUserGoldData(formData.value)
               break
             case 'update':
               res = await updateUserGoldData(formData.value)
               break
             default:
               res = await createUserGoldData(formData.value)
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
