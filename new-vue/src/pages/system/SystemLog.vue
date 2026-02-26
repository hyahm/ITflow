<template>
  <div id="log" class="q-pl-md">
    <q-banner class="bg-orange-1 text-brown q-mb-md rounded-borders">
      <template v-slot:avatar>
        <q-icon name="warning" color="brown" />
      </template>
      操作日志
    </q-banner>

    <div class="row q-col-gutter-md q-mb-md items-center q-ml-sm">
      <div class="col-auto">
        <q-select
          v-model="listQuery.classify"
          :options="classifys"
          label="分类"
          clearable
          outlined
          dense
          style="min-width: 150px"
          emit-value
          map-options
        />
      </div>

      <div class="col-auto">
           <q-popup-proxy
    v-model="dateProxyenable" 
     :persistent="false"
    cover            
    transition-show="scale" 
    transition-hide="scale"
  >
    <q-date
      v-model="dateRange"
      range
      mask="YYYY-MM-DD"
      bordered  
      
    />
  </q-popup-proxy>

  <!-- 输入框：点击触发弹窗，显示格式化后的日期范围 -->
  <q-input
    v-model="displayDateRange"
    label="日期范围"
    readonly  
    dense     
    :disable="false"
  
    @click="!dateProxyenable"
    append-icon="clear"
    :append-icon-clickable="displayDateRange !== ''"
    @append-icon-click="clearDate"
    :append-icon-color="displayDateRange !== '' ? 'grey-5' : 'transparent'"
    style="cursor: pointer;width: 200px;"
  />
         <!-- <q-input filled v-model="date" mask="date" :rules="['date']">
      <template v-slot:append>
        <q-icon name="event" class="cursor-pointer">
          <q-popup-proxy cover transition-show="scale" transition-hide="scale">
            <q-date v-model="date">
              <div class="row items-center justify-end">
                <q-btn v-close-popup label="Close" color="primary" flat />
              </div>
            </q-date>
          </q-popup-proxy>
        </q-icon>
      </template>
    </q-input> -->
        <!-- <q-date v-model="dateRange" range mask="YYYY-MM-DD" @update:model-value="changeDate">
          <template v-slot:default>
            <q-input
              v-model="displayDateRange"
              label="日期范围"
              outlined
              dense
              readonly
              clearable
              @clear="clearDateRange"
              style="min-width: 220px"
            >
              <template v-slot:append>
                <q-icon name="event" class="cursor-pointer">
                  <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                    <q-date v-model="dateRange" range mask="YYYY-MM-DD">
                      <div class="row items-center justify-end">
                        <q-btn v-close-popup label="关闭" color="primary" flat />
                      </div>
                    </q-date>
                  </q-popup-proxy>
                </q-icon>
              </template>
            </q-input>
          </template>
        </q-date> -->
      </div>

      <div class="col-auto">
        <q-btn color="primary" icon="search" label="搜索" @click="handleFilter" unelevated />
      </div>
    </div>

    <q-table
      :rows="list"
      :columns="columns"
      row-key="id"
      v-model:pagination="pagination"
      @request="onRequest"
      binary-state-sort
      flat
      bordered
      style="min-height: 350px"
    >
      <template v-slot:loading>
        <q-inner-loading showing color="primary" />
      </template>

 
    </q-table>

    <div class="row justify-center q-mt-md">
      <q-pagination
        v-model="pagination.page"
        :max="Math.ceil(count / pagination.rowsPerPage)"
        :max-pages="6"
        boundary-numbers
        direction-links
        @update:model-value="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted,watch, computed } from 'vue';
import { date,QPopupProxy } from 'quasar';
import { searchLog, logClassify } from 'src/api/log';
import type { SearchLog } from 'src/api/types/type';
import type { Option } from 'src/types/response';

const list = ref<SearchLog[]>([]);
const count = ref(0);
const classifys = ref<Option[]>([]);
const dateRange = ref<{ from: string; to: string } | string | null>(null);
// const date = ref(new Date())
// 日期选择器的原始值（数组格式：[开始日期, 结束日期]）
// const dateRange = ref<[string, string] | null>(null);
// 输入框显示的格式化值（字符串："2026-02-01 - 2026-02-23"）
// const displayDateRange = ref<string>('');
// 弹窗代理的 ref，用于手动控制显示/隐藏
const dateProxyenable = ref(false);


// 监听日期范围变化（防止手动修改 dateRange 时输入框不同步）
watch(dateRange, (newVal) => {
    console.log(newVal)
    console.log(dateRange.value)
  // handleDateChange(newVal);

  dateProxyenable.value = false
  console.log(dateProxyenable.value)
});
const listQuery: SearchLog = reactive({
  page: 1,
  limit: 20,
  classify: '',
  create_time: '',
  endtime: 0,
  count: 0,
  ip: '',
});

const pagination = ref({
  page: 1,
  rowsPerPage: 20,
  rowsNumber: 0,
});

const columns = [
  { name: 'id', label: 'ID', field: 'id', align: 'center' as const, width: '65px' },
  { name: 'create_time', label: '日期',  field: (row: SearchLog) =>{
    return date.formatDate(row.create_time, "YYYY-MM-DD HH:mm:ss")
  }, align: 'center' as const, width: '150px' },
  { name: 'classify', label: '分类', field: 'classify', align: 'center' as const, width: '150px' },
  { name: 'ip', label: 'IP', field: 'ip', align: 'center' as const, width: '150px' },
  {
    name: 'username',
    label: '操作者',
    field: 'username',
    align: 'center' as const,
    width: '150px',
  },
  { name: 'action', label: '操作', field: 'action', align: 'center' as const, width: '400px' },
];

const displayDateRange = computed(() => {
  if (!dateRange.value) return '';
  if (typeof  dateRange.value == "string") {
    return  `${dateRange.value} 至 ${dateRange.value}`;
  }
  return `${dateRange.value.from} 至 ${dateRange.value.to}`;
  // return ''
});

// function formatDate(timestamp: number) {
//   return date.formatDate(timestamp, 'YYYY-MM-DD HH:mm:ss');
// }

function clearDate() {
  dateRange.value = {from: '', to: ''};
}

async function classifyList() {
  try {
    const resp = await logClassify();
    classifys.value = resp.data.data;
  } catch (error) {
    console.error('获取分类失败:', error);
  }
}

// function changeDate(val: { from: string; to: string }) {
//   if (!val) {
//     listQuery.starttime = 0;
//     listQuery.endtime = 0;
//     return;
//   }
//   const fromDate = new Date(val.from);
//   const toDate = new Date(val.to);
//   listQuery.starttime = Math.floor(fromDate.getTime() / 1000);
//   listQuery.endtime = Math.floor(toDate.getTime() / 1000);
// }

async function handleCurrentChange(val: number) {
  listQuery.page = val;
  await handleFilter();
}

async function onRequest(props: { pagination: { page: number; rowsPerPage: number } }) {
  const { page } = props.pagination;
  listQuery.page = page;
  await handleFilter();
}

async function handleFilter() {
  await searchLog(listQuery).then((resp) => {
    console.log(resp.data.data)
    list.value = resp.data.data.loglist;
    console.log(list.value)
    count.value = resp.data.data.count;
    pagination.value.rowsNumber = count.value;
    pagination.value.page = listQuery.page;
  });
}

onMounted(async () => {
  await classifyList();
  await handleFilter();
});
</script>

<style scoped>
#log :deep(.q-table th) {
  font-weight: 500;
}

:deep(.q-input__append-icon) {
  display: inline-block !important;
  color: grey !important;
}
</style>
