<template>
  <q-page>
    <q-form class="form-container" dense>
      <!-- 顶部 sticky 工具栏 -->
      <div class="q-px-sm q-py-xs box">
        <q-btn class="publish-btn q-px-sm q-py-xs" dense label="发布" @click="submitForm" />

        <!-- 主体表单 -->
        <div class="q-pa-md">
          <!-- 任务类型 -->
          <q-item-label header class="q-px-none" dense>任务类型</q-item-label>
          <q-option-group v-model="typ" :options="ts" inline dense />

          <!-- 文章标题 -->
          <q-input
            v-model="postForm.title"
            label="文章标题"
            maxlength="100"
            clearable
            dense
            class="q-mt-md"
            style="width: 80%"
          />

          <div class="row q-col-gutter-md q-mt-xs">
            <!-- 项目名称 -->
            <q-select
              v-model="postForm.project_id"
              label="项目名称"
              :options="projects"
              option-value="id"
              option-label="name"
              emit-value
              map-options
              dense
              @update:model-value="changeProject"
              style="width: 400px"
            />

            <!-- 以下四项仅在 typ == 1 时显示 -->
            <template v-if="typ == 1">
              <!-- 运行环境 -->
              <q-select
                v-model="postForm.env_id"
                label="运行环境"
                :options="envnames"
                option-value="id"
                option-label="name"
                emit-value
                dense
                map-options
                style="width: 400px"
              />

              <!-- 应用版本 -->
              <!-- <q-select
            v-model="postForm.vid"
            label="应用版本"
            :options="versions"
            option-value="id"
            option-label="name"
            dense
            emit-value
            map-options
            style="width: 400px"
          /> -->

              <!-- 优先级别 -->
              <q-select
                v-model="postForm.level_id"
                label="优先级别"
                :options="levels"
                option-value="id"
                option-label="name"
                emit-value
                dense
                map-options
                style="width: 400px"
              />

              <!-- 重要性 -->
              <q-select
                v-model="postForm.important_id"
                label="重要性"
                :options="importants"
                option-value="id"
                dense
                option-label="name"
                emit-value
                map-options
                style="width: 400px"
              />

              <!-- 分配任务（多选） -->
              <q-select
                v-model="postForm.handle_uid"
                label="分配任务"
                :options="users"
                option-value="id"
                dense
                option-label="name"
                emit-value
                map-options
                use-chips
                style="width: 400px"
              />
            </template>
          </div>

          <!-- 富文本编辑器 -->
          <div id="main" class="q-mt-lg">
            <!-- 如果你仍在用 mavon-editor，直接保留即可 -->
            <v-md-editor ref="md" v-model="postForm.content" @imgAdd="imgAdd" style="width: 100%; min-height: 400px"/>
            <!-- 想换成 Quasar 官方 Markdown 可换 q-editor -->
          </div>
        </div>
      </div>
    </q-form>
  </q-page>
</template>

<script setup lang="ts">
// import Sticky from 'src/components/Sticky'; // 粘性header组件
import { createBug, updateBug } from 'src/api/bugs';
import {
  getProjectKeyName,
  getLevels,
  getImportants,
  getTyp,
  getUserKeyNameByProject,
  // getVersionKeyNameByProject,
} from 'src/api/get';
import { getEnvName } from 'src/api/env';
import type { Option, KeyName } from 'src/types/response';
import { reactive, ref, useTemplateRef } from 'vue';
import type { Bug } from 'src/api/types/type';
import { uploadImg } from 'src/api/uploadimg';
import { useQuasar } from 'quasar';



// import { laAws } from '@quasar/extras/line-awesome';

const $q = useQuasar();
let postForm = reactive<Bug>({
  // status: 'draft',
  title: '', // 文章题目
  content: '', // 文章内容
  id: 0,
  handle_uid: null,
  project_id: null,
  level_id: null,
  env_id: null,
  important_id: null,
  create_id: 0,
  status_id: 0,
  type_id: 0,
  deadline: 0,
  dustbin: false,
});
const isEdit = ref(false);
// let versions: KeyName[] = reactive([]); // 随项目变化的版本号
let importants: KeyName[] = reactive([]);
let levels: KeyName[] = reactive([]);
// const oses = reactive([]);
let users: KeyName[] = reactive([]); // 随项目变化的用户
let projects: KeyName[] = reactive([]);
let envnames: KeyName[] = reactive([]);
let ts: Option[] = reactive([]);
const typ = ref(0);
const md = useTemplateRef('md');

// if (isEdit) {
//   const id = this.$route.params && this.$route.params.id;
//   this.postForm.id = parseInt(id);
//   this.fetchData(id);
// } else {
//   this.postForm = Object.assign({}, defaultForm);
// }
// const getversionMap = async (projectid: number) => {
//   await getVersionKeyNameByProject({ project_id: projectid }).then((resp) => {
//     versions = resp.data.data;
//   });
// };

const gettyp = async () => {
  await getTyp().then((resp) => {
    ts = resp.data.data;
    if (ts.length > 0) {
      typ.value = ts[0]?.value as number;
    }
  });
};
const changeProject = async () => {
  // 选择不用项目会显示不同的用户
  if (postForm.project_id === null || postForm.project_id <= 0) {
    users = reactive([]);
    Object.assign(users, []);
    return;
  }
  await getUserKeyNameByProject({ project_id: postForm.project_id }).then((resp) => {
    console.log(resp.data.data);
    Object.assign(users, resp.data.data);
  });
  // await getversionMap(postForm.pid as number);
};

const getimportants = async () => {
  await getImportants().then((resp) => {
    importants = resp.data.data;
  });
};

const getlevels = async () => {
  await getLevels().then((resp) => {
    levels = resp.data.data;
  });
};
const getenv = async () => {
  await getEnvName().then((resp) => {
    console.log(resp.data.data);
    envnames = resp.data.data;
  });
};
const getproject = async () => {
  await getProjectKeyName().then((resp) => {
    Object.assign(projects, resp.data.data);
  });
};

// const fetchData = (id: number) => {
//   fetchBug(id).then((resp) => {
//     const dd = resp.data.data;
//     postForm = dd;
//     getuserMap(postForm.pid);
//     getversionMap(postForm.pid);
//   });
// };
const submitForm = async () => {
  // this.postForm.display_time = parseInt(this.display_time / 1000)
  if (postForm.title.length > 40) {
    // this.$message({
    //   message: '标题长度必须小于40位',
    //   type: 'error',
    // });
    return;
  }
  if (typ.value == 1 && !postForm.handle_uid) {
    // this.$message({
    //   message: '请选择指定给谁',
    //   type: 'error',
    // });
    return;
  }

  if (typ.value == 1 && !postForm.level_id) {
    // this.$message({
    //   message: '请选择优先级别',
    //   type: 'error',
    // });
    return;
  }
  if (typ.value == 1 && !postForm.important_id) {
    // this.$message({
    //   message: '请选择项目严重程度',
    //   type: 'error',
    // });
    return;
  }
  if (postForm.content.length < 1) {
    // this.$message({
    //   message: '请填写内容',
    //   type: 'error',
    // });
    return;
  }
  if (typ.value == 1 && !postForm.env_id) {
    // this.$message({
    //   message: '请选择运行环境',
    //   type: 'error',
    // });
    return;
  }

  if (isEdit.value) {
    // postForm.tid = parseInt(tid);
    await updateBug(postForm).then(() => {
      $q.notify({ type: 'positive', message: '保存成功', position: 'top' });
    });
  } else {
    await createBug(postForm).then(() => {
      $q.notify({ type: 'positive', message: '创建成功', position: 'top' });
    });
  }
};
const imgAdd = async (pos: number, file: File) => {
  // 第一步.将图片上传到服务器.
  const formdata = new FormData();
  formdata.append('image', file);
  await uploadImg(formdata).then((resp) => {
    md.value.$img2Url(pos, resp.data.data.url);
    // this.$refs.md.$img2Url(pos, resp.data.url);
  });
};
// const draftForm = () => {
//   if (postForm.content.length === 0 || postForm.title.length === 0) {
//     // this.$message({
//     //   message: '请填写必要的标题和内容',
//     //   type: 'warning',
//     // });
//     return;
//   }
//   // this.$message({
//   //   message: '保存成功',
//   //   type: 'success',
//   //   duration: 1000,
//   // });
// };

await getimportants();
await getproject();
await getlevels();
await getenv();
await gettyp();
</script>

<style lang="css" scoped>
.box > .publish-btn {
  right: 0;
  top: 50%;
  color: red;
  z-index: 1;
}
</style>
