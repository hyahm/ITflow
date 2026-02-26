<template>
  <q-page class="row q-pa-md">
    <!-- 1. 图片选择 + 操作按钮区域 -->
    <div class="col-12 mb-6">
      <!-- Quasar 文件选择组件：仅允许图片，限制格式 -->
      <q-file
      v-model="selectedFile"
        label="选择头像图片"
        accept="image/jpeg,image/png,image/webp"
        class="mb-4"
        square
        outlined
        color="primary"
      />

      <!-- 操作按钮组 -->
      <q-btn
        label="设置头像"
        color="primary"
        icon="photo_camera"
        @click="toggleCropShow"
        :disabled="!selectedFile"
        class="mr-2"
      />
      <q-btn
        label="重置"
        color="grey"
        icon="refresh"
        @click="resetAll"
        flat
      />
    </div>

    <!-- 2. 裁剪弹窗：Quasar 对话框封装，更符合生态 -->
    <q-dialog v-model="cropShow"  maximized>
      <q-card style="max-width: 800px; margin: 0 auto;">
        <q-card-header>
          <!-- <q-card-title>裁剪头像</q-card-title> -->
          <q-card-actions>
            <q-btn icon="close" flat @click="toggleCropShow" />
          </q-card-actions>
        </q-card-header>

        <q-card-section>
          <!-- 核心裁剪组件 -->
          <cropper
            class="cropper-container"
            :src="cropImageSrc"
            :stencil-props="{
              aspectRatio: 1, // 1:1 正方形裁剪（头像常用）
              movable: true,
              resizable: true
            }"
            @change="handleCropChange"
            :image-restriction="'stencil'"
          />
        </q-card-section>

        <q-card-actions align="right" class="q-pa-sm">
          <q-btn label="取消" flat @click="toggleCropShow" />
          <q-btn
            label="确认上传"
            color="primary"
            @click="handleConfirmUpload"
            :loading="uploading"
          />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <!-- 3. 头像预览区域 -->
    <!-- <div class="col-12 mt-4">
      <q-typography text="当前头像预览" class="text-subtitle2 mb-2" />
      <q-img
        :src="avatarUrl || 'https://cdn.quasar.dev/img/avatar.png'"
        alt="头像预览"
        style="width: 150px; height: 150px; border-radius: 50%;"
        spinner-color="primary"
        fit="cover"
      />
    </div>
  -->
  </q-page>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue';
import { Cropper } from 'vue-advanced-cropper';
import 'vue-advanced-cropper/dist/style.css';
import { useQuasar } from 'quasar';
import { api } from "boot/axios";
import { useUserStore } from 'src/stores/user'

// Quasar 工具：通知、加载等
const $q = useQuasar();

// 核心状态变量
const selectedFile = ref<File | null>(null); // 选中的本地图片文件
const cropImageSrc = ref(''); // 裁剪组件的图片源（base64）
const cropShow = ref(false); // 裁剪弹窗显隐
const croppedCanvas = ref<HTMLCanvasElement | null>(null); // 裁剪后的Canvas
const avatarUrl = ref(''); // 最终头像地址（后端返回/本地预览）
const uploading = ref(false); // 上传加载状态
const user = useUserStore()

// 核心修复：用 watch 监听 selectedFile 变化，替代 @input 事件
watch(selectedFile, (newFile) => {
  if (!newFile) return;
  handleFileSelect(newFile); // 传入纯 File 对象
}, { immediate: false });

// 1. 选择本地图片：转为base64传给裁剪组件
const handleFileSelect = (file: File | null) => {
  if (!file) return;
  // 校验图片格式/大小（可选）
  const validTypes = ['image/jpeg', 'image/jpg','image/png', 'image/webp'];
  if (!validTypes.includes(file.type)) {
    $q.notify({ type: 'warning', message: '仅支持 JPG/PNG/WEBP 格式！' });
    return;
  }
  if (file.size > 5 * 1024 * 1024) { // 5MB 限制
    $q.notify({ type: 'warning', message: '图片大小不能超过 5MB！' });
    return;
  }

  selectedFile.value = file;
  // 将 File 转为 base64，供裁剪组件使用
  const reader = new FileReader();
  reader.onload = (e) => {
    cropImageSrc.value = e.target?.result as string;
  };
  reader.readAsDataURL(file);
};

// 2. 裁剪变化：保存裁剪后的Canvas
const handleCropChange = ({ canvas }: { canvas: HTMLCanvasElement }) => {
  croppedCanvas.value = canvas;
};

// 3. 切换裁剪弹窗显隐
const toggleCropShow = () => {
  cropShow.value = !cropShow.value;
};

// 4. 确认裁剪并上传
const handleConfirmUpload = async () => {
  if (!croppedCanvas.value) {
    $q.notify({ type: 'warning', message: '请先裁剪头像！' });
    return;
  }

  uploading.value = true;
  try {
    // 核心修复：处理 toBlob 回调的 null 类型，显式声明 Promise 类型
    const blob = await new Promise<Blob | null>((resolve) => {
      // 1. 声明 Promise 接收 Blob | null 类型
      croppedCanvas.value?.toBlob((blob) => {
        // 2. 回调参数是 Blob | null，直接 resolve
        resolve(blob);
      }, 'image/jpeg', 0.8);
    });

    // 3. 新增：校验 blob 是否为 null，避免后续上传错误
    if (!blob) {
      throw new Error('裁剪后的图片转换失败，请重新裁剪！');
    }

    // 后续上传逻辑不变
    const formData = new FormData();
    formData.append('image', blob, `avatar_${Date.now()}.jpg`);
    const response = await api({
        url: '/uploadimg',
        method: 'post',
      data: formData
    });
    if (response.status == 200) {
      cropShow.value = false
      avatarUrl.value = response.data.data.url
      user.SetHead(avatarUrl.value)
      $q.notify({
      type: 'negative',
      message: `上传成功`,
      icon: 'success'
    });
    }
  } catch (error) {
    $q.notify({
      type: 'negative',
      message: `上传失败：${(error as Error).message}`,
      icon: 'error'
    });
  } finally {
    uploading.value = false;
  }
};

// 5. 重置所有状态
const resetAll = () => {
  selectedFile.value = null;
  cropImageSrc.value = '';
  cropShow.value = false;
  croppedCanvas.value = null;
  // 可选：不重置已上传的头像
  // avatarUrl.value = '';
};

// 可选：监听裁剪弹窗关闭，清空裁剪缓存
watch(cropShow, (val) => {
  if (!val) {
    croppedCanvas.value = null;
  }
});
</script>

<style scoped>
/* 裁剪组件容器样式 */
.cropper-container {
  width: 100%;
  height: 400px;
  border: 1px solid #eee;
  border-radius: 4px;
}

/* 适配 Quasar 对话框间距 */
.q-dialog .q-card {
  padding: 0;
}
</style>