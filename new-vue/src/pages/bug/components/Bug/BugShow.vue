<template>
  <div>
    <q-table
      :rows="list"
      :columns="columns"
      row-key="id"
      :loading="listLoading"
      flat
      bordered
      binary-state-sort
      style="width: 100%"
    >
      <!-- 自定义列渲染 -->
      <!-- <template v-slot:body-cell-createtime="props">
        <q-td :props="props" class="text-center">
          {{ formatDate(props.row.createtime) }}
        </q-td>
      </template>

      <template v-slot:body-cell-deadline="props">
        <q-td :props="props" class="text-center">
          {{ formatDate(props.row.deadline) }}
        </q-td>
      </template>

      <template v-slot:body-cell-project="props">
        <q-td :props="props" class="text-center">
          {{ props.row.pid || '-' }}
        </q-td>
      </template>

      <template v-slot:body-cell-level="props">
        <q-td :props="props" class="text-center">
          {{ props.row.lid || '-' }}
        </q-td>
      </template>

      <template v-slot:body-cell-important="props">
        <q-td :props="props" class="text-center">
          {{ importantMap.get(props.row.iid) || '-' }}
        </q-td>
      </template>

      <template v-slot:body-cell-handler="props">
        <q-td :props="props" class="text-center">
          {{ formatUsers(props.row.spusers) }}
        </q-td>
      </template>

      <template v-slot:body-cell-status="props">
        <q-td :props="props" class="text-center">
          <q-chip dense color="primary" text-color="white" size="sm">
            {{ props.row.sid || '-' }}
          </q-chip>
        </q-td>
      </template> -->

      <template v-slot:body-cell-title="props">
        <q-td :props="props" class="text-center">
          <router-link
            :to="'/showbug/' + props.row.id"
            class="text-primary text-decoration-none"
            style="cursor: pointer"
            target="_blank"
          >
            {{ props.row.title }}
          </router-link>
        </q-td>
      </template>

      <template v-slot:body-cell-actions="props">
        <q-td :props="props" class="text-center" v-if="pageType != 3">
          <div class="q-gutter-xs">
            <q-btn
              v-if="pageType == 2"
              color="primary"
              size="sm"
              label="编辑"
              :to="'/bug/edit/' + props.row.id"
              unelevated
            />

            <q-btn
              v-if="pageType == 3"
              color="positive"
              size="sm"
              label="恢复"
              @click="resume(props.row.id)"
              unelevated
            />

            <q-btn
              v-if="pageType == 4"
              color="positive"
              size="sm"
              label="领取"
              @click="Receive(props.row)"
              unelevated
            />

            <q-btn
              v-if="pageType == 2"
              color="negative"
              size="sm"
              label="删除"
              @click="handleRemove(props.row.id)"
              unelevated
            />

            <q-btn
              v-if="pageType == 4"
              color="primary"
              size="sm"
              label="转交"
              @click="handlePass(props.row)"
              unelevated
            />

            <q-btn
              v-if="pageType == 4 && props.row.uid == currentUserId"
              color="secondary"
              size="sm"
              label="完成"
              @click="handleComplete(props.row)"
              unelevated
            />
          </div>
        </q-td>
      </template>
    </q-table>

    <!-- 转交任务对话框 -->
    <q-dialog v-model="dialogFormVisible" persistent v-if="pageType == 4">
      <q-card style="min-width: 400px">
        <q-card-section>
          <div class="text-h6">任务完成转交</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-form ref="dataForm" class="q-gutter-md">
            <q-select
              v-model="temp.spusers"
              :options="userOptions"
              option-value="id"
              option-label="name"
              emit-value
              map-options
              multiple
              use-input
              use-chips
              new-value-mode="add-unique"
              label="任务给"
              outlined
              dense
              hint="请选择或输入指定的用户"
            />

            <q-input
              v-model="temp.remark"
              type="textarea"
              label="说明"
              outlined
              autogrow
              placeholder="请输入说明"
            />
          </q-form>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="取消" color="grey" v-close-popup />
          <q-btn flat label="确认" color="primary" @click="updateData" />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <!-- 领取任务对话框 -->
    <q-dialog v-model="openReceive" persistent v-if="pageType == 4">
      <q-card style="min-width: 350px">
        <q-card-section>
          <div class="text-h6">领取任务</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-form class="q-gutter-md">
            <q-input
              v-model="receive.deadline"
              type="datetime-local"
              label="完成时间"
              outlined
              dense
            />
          </q-form>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="取消" color="grey" v-close-popup />
          <q-btn flat label="确认" color="primary" @click="receiveHandle" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted,watch } from 'vue';
import { useRouter } from 'vue-router';
import { useQuasar,date } from 'quasar';
import { getImportants, getUserKeyName, getUserKeyNameByProject } from 'src/api/get';
import { defaultValue } from 'src/api/defaultvalue';
import { passBug, completeBug, delBug } from 'src/api/bugs';
// import { passBug, completeBug, delBug, receiveBug } from 'src/api/bugs';
import type { RequestPass, DefaultValue, Bug } from 'src/api/types/type';

const $q = useQuasar();
const router = useRouter();




const props = defineProps(['list', 'pageType', 'projectMap', 'levelMap']);

watch(
  () => props.list, // 用函数返回 props.list，确保响应式追踪
  (newVal, oldVal) => {
    // 注意：数组是引用类型，newVal 和 oldVal 指向同一引用，需对比内容
    // watchMsg.value = `列表从 ${oldVal.length} 条变为 ${newVal.length} 条`
    console.log('list 变化了：', newVal, oldVal)
  },
  {
    deep: true, // 监听数组内部元素/对象属性变化（必加，否则仅监听数组引用变化）
    immediate: true // 可选：页面加载时立即执行一次
  }
)
// Props 定义
// interface Props {
//   list: any[];
//   projectMap: Map<number, string>;
//   levelMap: Map<number, string>;
//   statusMap: Map<number, string>;
//   pageType: number; // 1: 垃圾箱, 2: 我创建的bug, 3: 所有bug, 4: 我的任务
//   currentUserId?: number; // 当前用户ID，用于判断是否是任务所有者
// }

// const props = withDefaults(defineProps<Props>(), {
//   list: () => [],
//   currentUserId: 0,
// });

// 事件 emits
const emit = defineEmits(['refresh', 'resume']);

// 列定义
const columns = [
  { name: 'id', label: 'ID', field: 'id', align: 'center' as const, style: 'width: 50px' },
  {
    name: 'create_time',
    label: '日期',
    field: 'create_time',
    align: 'center' as const,
    format: (val: string) => {
      return date.formatDate(val, 'YYYY-MM-DD HH:mm:ss')
    },
    style: 'width: 150px',
  },
  {
    name: 'deadline',
    label: '完成时间',
    field: 'deadline',
     format: (val: string) => {
      if (val == '0001-01-01T00:00:00Z') {
        return ''
      }
      return date.formatDate(val, 'YYYY-MM-DD HH:mm:ss')
    },
    align: 'center' as const,
    style: 'width: 150px',
  },
  { name: 'project_name', label: '项目', field: 'project_name', align: 'center' as const, style: 'width: 100px' },
  { name: 'level_name', label: '优先级', field: 'level_name', align: 'center' as const, style: 'width: 80px' },
  {
    name: 'important_name',
    label: '重要性',
    field: 'important_name',
    align: 'center' as const,
    style: 'width: 100px',
  },
  {
    name: 'handle_name',
    label: '处理者',
    field: 'handle_name',
    align: 'center' as const,
    style: 'width: 100px',
  },
  { name: 'status_name', label: '状态', field: 'status_name', align: 'center' as const, style: 'width: 110px' },
  {
    name: 'title',
    label: '标题',
    field: 'title',
    align: 'center' as const,
    style: 'min-width: 300px',
  },
  {
    name: 'actions',
    label: '操作',
    field: 'actions',
    align: 'center' as const,
    style: 'width: 230px',
  },
];

const listLoading = ref(false);
const dialogFormVisible = ref(false);
const openReceive = ref(false);
const importantMap = ref<Map<number, string>>(new Map());
const userMap = ref<Map<number, string>>(new Map());
const defaultVal = reactive<DefaultValue>({
  created: 0,
  pass: 0,
  completed: 0,
  receive: 0,
});
const currentUserId = ref(0);
const temp = reactive<RequestPass>({
  bid: 0,
  remark: '',
  spusers: [],
});

const receive: Bug = reactive({
  id: 0,
  title: '',
  status_id: 0,
  create_id: 0,
  handle_uid: null,
  content: '',
  important_id: 0,
  level_id: 0,
  env_id: 0,
  type_id: 0,
  project_id: 0,
  deadline: new Date().getTime(),
  dustbin: false,
});

// 计算属性：用户选项列表
const userOptions = computed(() => {
  const options = [];
  for (const [id, name] of userMap.value.entries()) {
    options.push({ id, name });
  }
  return options;
});

// 方法
// const formatDate = (timestamp: number | string) => {
//   if (!timestamp) return '-';
//   return quasarDate.formatDate(Number(timestamp), 'YYYY-MM-DD HH:mm');
// };

// const formatUsers = (ids: number[]) => {
//   if (!ids || !Array.isArray(ids)) return '-';
//   const names = ids.map((id) => userMap.value.get(id)).filter(Boolean);
//   return names.join(', ') || '-';
// };

const handleComplete = async (row: Bug) => {
  receive.id = row.id;
  await completeBug(receive);
  $q.notify({
    color: 'positive',
    message: '操作成功',
    position: 'top',
    timeout: 2000,
  });
  emit('refresh');
};

const handleRemove = async (id: number) => {
  await delBug(id);
  $q.notify({
    color: 'positive',
    message: '删除成功',
    position: 'top',
    timeout: 2000,
  });
  emit('refresh');
};

const updateData = async () => {
  if (!temp.spusers || temp.spusers.length === 0) {
    $q.notify({
      color: 'warning',
      message: '至少选择一个处理人',
      position: 'top',
      timeout: 2000,
    });
    return;
  }

  await passBug(temp);
  $q.notify({
    color: 'positive',
    message: '操作成功',
    position: 'top',
    timeout: 2000,
  });
  dialogFormVisible.value = false;
  setTimeout(() => {
    router.go(0);
  }, 2000);
};

const getImportant = async () => {
  try {
    const resp = await getImportants();
    for (const v of resp.data.data) {
      importantMap.value.set(v.id, v.name);
    }
    const resp2 = await defaultValue();
    Object.assign(defaultVal, resp2.data.data);
  } catch (error) {
    console.error('获取重要级别失败:', error);
  }
};

const getUsers = async () => {
  try {
    const resp = await getUserKeyName();
    for (const v of resp.data.data) {
      userMap.value.set(v.id, v.name);
    }
  } catch (error) {
    console.error('获取用户列表失败:', error);
  }
};

const handlePass = async (row: Bug) => {
  temp.bid = row.id;
  temp.spusers = [];
  temp.remark = '';

  await getUserKeyNameByProject({ project_id: row.project_id as number }).then((resp) => {
    const users = resp.data.data;
    console.log(users);
    // 这里逻辑有点奇怪，原代码是筛选，但可能应该直接赋值可用用户
    // 保留原逻辑：
    // for (let i = 0; i < users.length; i++) {
    //   for (let j = 0; j < row.spusers.length; j++) {
    //     if (row.spusers[j] === users[i]) {
    //       temp.spusers.push(users[i]);
    //     }
    //   }
    // }
  });

  dialogFormVisible.value = true;
};

const Receive = (row: Bug) => {
  // if (row.sid != defaultVal.pass && row.sid != defaultVal.created) {
  //   $q.notify({
  //     color: 'warning',
  //     message: '此状态无法领取',
  //     position: 'top',
  //     timeout: 2000,
  //   });
  //   return;
  // }
  openReceive.value = true;
  receive.id = row.id;
  receive.deadline = new Date().getTime();
};

const receiveHandle = async () => {
  // await receiveBug();
  // $q.notify({
  //   color: 'positive',
  //   message: '领取成功',
  //   position: 'top',
  //   timeout: 2000,
  // });
  // openReceive.value = false;
  // setTimeout(() => {
  //   router.go(0);
  // }, 2000);
};

const resume = (id: number) => {
  emit('resume', id);
};

// 初始化
onMounted(async () => {
  await getImportant();
  await getUsers();
});
</script>

<style scoped>
.text-decoration-none {
  text-decoration: none;
}
</style>
