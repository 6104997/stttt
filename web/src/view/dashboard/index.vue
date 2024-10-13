<template>
  <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-7 py-2 gap-4 md:gap-2 gva-container2">
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32 ">
      <gva-chart :type="2" title="注册用户总数" :number="total" />
    </gva-card>
    <gva-card custom-class="col-span-1 lg:col-span-2 h-32 ">
      <gva-chart :type="2" title="今日新增客户" :number="number" />
    </gva-card>
  </div>
</template>

<script setup>
import { GvaPluginTable,GvaTable, GvaChart, GvaWiki , GvaNotice , GvaQuickLink , GvaCard , GvaBanner } from "./componenst"
import  {getTheNumberOfRegisteredUsersTodays,totalNumberOfRegisteredUsers}  from '@/api/st/appUser.js'
import { ref } from 'vue'
const number = ref()
const total = ref()
const getTheNumberOfRegisteredUsers = async () => {
  const res = await getTheNumberOfRegisteredUsersTodays()
  number.value = res.data.number
  console.log(res.data.number)
  const totaltes = await totalNumberOfRegisteredUsers()
  total.value = totaltes.data.total
  
  console.log(total.value)
  console.log(totaltes)
}

getTheNumberOfRegisteredUsers()


</script>

<style lang="scss" scoped>
</style>
