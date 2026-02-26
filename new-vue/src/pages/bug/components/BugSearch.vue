<template>
  <div class="q-pa-md">
    <!-- 筛选栏 -->
    <div class="row q-col-gutter-sm items-center q-mb-md">
      <div class="col-auto">
        <q-input
          v-model="listQuery.title"
          placeholder="标题"
          outlined
          dense
          clearable
          style="width: 200px"
          @keyup.enter="searchHandle"
        />
      </div>

      <div class="col-auto">
        <q-select
          v-model="listQuery.level_id"
          :options="levels"
          option-value="id"
          option-label="name"
          emit-value
          map-options
          placeholder="级别"
          clearable
          outlined
          dense
          style="width: 120px"
        />
      </div>

      <div class="col-auto">
        <q-select
          v-model="listQuery.project_id"
          :options="projectnames"
          option-value="id"
          option-label="name"
          emit-value
          map-options
          placeholder="项目名"
          clearable
          outlined
          dense
          style="width: 150px"
        />
      </div>


       <div class="col-auto">
        <q-select
          v-model="showstatus"
          :options="allStatus"
          option-value="id"
          option-label="name"
          emit-value
          map-options
          placeholder="项目名"
          clearable
          outlined
          dense
          style="width: 150px"
        />
      </div>
    
      <!-- <div class="col-auto">
        <q-btn-dropdown :label="`状态(${statuslength})`" flat no-caps persistent style="width: 150px">
          <q-list dense padding style="min-width: 150px">
            <q-item tag="label" v-for="item in allStatus" :key="item.id" v-ripple>
              <q-item-section avatar>
                <q-checkbox
                  v-model="showstatus"
                  :val="item.name"
                  @update:model-value="HandleChange()"
                  dense
                />
              </q-item-section>
              <q-item-section>
                <q-item-label>{{ item.name }}</q-item-label>
              </q-item-section>
            </q-item>
          </q-list>
        </q-btn-dropdown>
      </div> -->
       <div class="col-auto">
        <q-btn color="primary" icon="search" label="搜索" @click="searchHandle" unelevated />
      </div>
    </div>
   

    <!-- 列表组件 -->
    <Show
      :list="list"
    
      :statusMap="statusidMap"
      :projectMap="projectMap"
      :levelMap="levelMap"
    />

    <!-- 分页 -->
    <div class="row justify-center q-mt-md">
      <q-pagination
        v-model="listQuery.page"
        :max="Math.ceil(total / listQuery.limit)"
        :max-pages="6"
        boundary-numbers
        direction-links
        @update:model-value="handleCurrentChange"
      />
      <div class="q-ml-md flex items-center">
        <q-select
          v-model="listQuery.limit"
          :options="[10, 15, 20, 30]"
          label="每页"
          dense
          outlined
          style="width: 80px"
          emit-value
          @update:model-value="handleSizeChange"
        />
        <span class="q-ml-sm text-caption">共 {{ total }} 条</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
// import { useQuasar } from 'quasar';
import { searchAllBugs } from 'src/api/search';
// import { resumeBug } from 'src/api/bugs';
// import { statusFilter } from 'src/api/status';
import type { ReqMyBugFilter, Status } from 'src/api/types/type';
import Show from './Bug/BugShow.vue';
import { getProjectKeyName, getLevels, getStatus } from 'src/api/get';

// const $q = useQuasar();

interface Props {
  pageType?: number; // 1: 垃圾箱, 2: 我创建的, 3: 所有的, 4: 我的任务
}

const props = withDefaults(defineProps<Props>(), {
  pageType: 1,
});

// 数据定义
const list = ref([]);
const total = ref(0);
const projectnames = ref<Status[]>([]);
const projectMap = ref<Map<number, string>>(new Map());
const levels = ref<Status[]>([]);
const levelMap = ref<Map<number, string>>(new Map());
const allStatus = ref<Status[]>([]);
const showstatus = ref<number| null>(null);
const statusMap = ref<Map<string, number>>(new Map());
const statusidMap = ref<Map<number, string>>(new Map());
// const statuslength = ref(0);

const listQuery: ReqMyBugFilter = reactive({
  limit: 15,
  page: 1,
  level_id: null as number | null,
  project_id: null as number | null,
  showstatus: null as number | null,
  title: '',
  page_type: 0,
});

// 获取状态、级别和显示状态
const get_status = async () => {
  await getStatus().then((resp) => {
    allStatus.value = resp.data.data;
    statusMap.value.clear();
    statusidMap.value.clear();

    for (const v of allStatus.value) {
      statusMap.value.set(v.name, v.id);
      statusidMap.value.set(v.id, v.name);
    }
  });
  await getLevels().then((resp) => {
    levels.value = resp.data.data;
    levelMap.value.clear();
    for (const v of levels.value) {
      levelMap.value.set(v.id, v.name);
    }
  });
  // await getShowStatus().then((resp) => {
  //   showstatus.value = [];
  //   for (const v of resp.data.data) {
  //     const name = statusidMap.value.get(v);
  //     if (name) showstatus.value.push(name);
  //   }
  //   statuslength.value = showstatus.value.length;
  // });
  // 处理显示状态

  await handleFilter();
};

const getPname = async () => {
  try {
    const resp = await getProjectKeyName();
    projectnames.value = resp.data.data;
    projectMap.value.clear();
    for (const v of projectnames.value) {
      projectMap.value.set(v.id, v.name);
    }
  } catch (error) {
    console.error('获取项目失败:', error);
  }
};

// const getSidByShowStatus = () => {
//   const sids: number[] = [];
//   for (const v of showstatus.value) {
//     const id = statusMap.value.get(v);
//     if (id !== undefined) sids.push(id);
//   }
//   return sids;
// };

// const HandleChange = () => {
//   console.log(showstatus)
  
//   // const sids = getSidByShowStatus();
//   // statuslength.value = showstatus.value.length;
//   // statusFilter({ showstatus: sids });
// };

const handleFilter = async () => {
  // listQuery.showstatus = getSidByShowStatus();
  listQuery.page_type = props.pageType
  listQuery.showstatus = showstatus.value
  // try {
  //   let resp;
  //   switch (props.pageType) {
  //     case 1:
  //       resp = await bugFilter(listQuery);
  //       break;
  //     case 2:
  //       resp = await searchMyBugs(listQuery);
  //       break;
  //     case 3:
     const   resp = await searchAllBugs(listQuery);
    //     break;
    //   case 4:
    //     resp = await searchMyTasks(listQuery);
    //     break;
    //   default:
    //     resp = await bugFilter(listQuery);
    // }

    list.value = resp.data.data;
    console.log(list);
    total.value = resp.data.total;
    listQuery.page = resp.data.page;
  // } catch (error) {
  //   console.error('搜索失败:', error);
  //   $q.notify({
  //     color: 'negative',
  //     message: '获取数据失败',
  //     position: 'top',
  //   });
  // }
};

const searchHandle = async () => {
  listQuery.page = 1;
  await handleFilter();
};

const handleSizeChange = async (val: number) => {
  listQuery.limit = val;
  listQuery.page = 1;
  await handleFilter();
};

const handleCurrentChange = async (val: number) => {
  listQuery.page = val;
  await handleFilter();
};

// const resume = async (id: number) => {
//   await resumeBug(id);
//   const index = list.value.findIndex((item) => item.id === id);
//   if (index > -1) {
//     list.value.splice(index, 1);
//     $q.notify({
//       color: 'positive',
//       message: '恢复成功',
//       position: 'top',
//       timeout: 2000,
//     });
//   }
// };

await getPname();

await get_status();
</script>
