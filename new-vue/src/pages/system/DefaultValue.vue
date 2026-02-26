<template>
  <div class="q-pl-md">
    <p class="text-warning">某些选项的默认值</p>
    
    <div class="q-mt-md">
      bug创建时的状态:
      <q-select
        v-model="form.created"
        :options="statusOptions"
        option-value="id"
        option-label="name"
        label="Select"
        outlined
        dense
        emit-value
        map-options
      />
    </div>
    
    <div class="q-mt-md">
      bug完成时的状态:
      <q-select
        v-model="form.completed"
        :options="statusOptions"
        option-value="id"
        option-label="name"
        label="Select"
        outlined
        dense
        emit-value
        map-options
      />
    </div>
    
    <div class="q-mt-md">
      bug转交时的状态:
      <q-select
        v-model="form.pass"
        :options="statusOptions"
        option-value="id"
        option-label="name"
        label="Select"
        outlined
        dense
        emit-value
        map-options
      />
    </div>
    
    <div class="q-mt-md">
      bug领取后的状态:
      <q-select
        v-model="form.receive"
        :options="statusOptions"
        option-value="id"
        option-label="name"
        label="Select"
        outlined
        dense
        emit-value
        map-options
      />
    </div>
    
    <q-btn 
      class="q-mt-md" 
      color="primary" 
      outline 
      label="保存"
      @click="handleSave"
    />
  </div>
</template>

<script setup lang="ts">
import { defaultValue, save } from 'src/api/defaultvalue'
import type { DefaultValue } from 'src/api/types/type'
import { getStatus } from 'src/api/get'
import { reactive, ref, onMounted } from 'vue'
import type { KeyName } from 'src/types/response'
import { useQuasar } from 'quasar'

const $q = useQuasar()

const form: DefaultValue = reactive({
  created: 0,
  completed: 0,
  pass: 0,
  receive: 0,
})

const statusOptions = ref<KeyName[]>([])

async function getDefaultValue() {
  try {
    const resp = await defaultValue()
    const data = resp.data.data
    form.created = data.created
    form.completed = data.completed
    form.pass = data.pass
    form.receive = data.receive
  } catch (error) {
    console.error('获取默认值失败:', error)
  }
}

async function getStatusList() {
  try {
    const resp = await getStatus()
    statusOptions.value = resp.data.data
  } catch (error) {
    console.error('获取状态列表失败:', error)
  }
}

async function handleSave() {
  try {
    await save(form)
    $q.notify({
      type: 'positive',
      message: '保存成功',
      position: 'top',
      timeout: 2000
    })
  } catch (error) {
    $q.notify({
      type: 'negative',
      message: '保存失败',
      position: 'top',
      timeout: 2000
    })
    console.error('保存失败:', error)
  }
}

onMounted(async () => {
  await getDefaultValue()
  await getStatusList()
})
</script>

<style scoped>
.text-warning {
  color: #f2c037;
}
</style>