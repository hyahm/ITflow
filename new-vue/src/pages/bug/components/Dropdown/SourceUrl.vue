<template>
  <div></div>
  <!-- <el-dropdown :show-timeout="100" trigger="click">
    <el-button plain>
      外链
      <i class="el-icon-caret-bottom el-icon--right" />
    </el-button>
    <template v-slot:dropdown>
      <el-dropdown-menu class="no-padding no-border" style="width: 400px">
        <el-form-item label-width="0px" style="margin-bottom: 0px" prop="source_uri">
          <el-input v-model="value" placeholder="请输入内容">
            <template v-slot:prepend>填写url</template>
          </el-input>
        </el-form-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown> -->
</template>

<script setup lang="ts">
// const value = defineProps<{
//   value: string;
// }>();

// // 定义 update:modelValue 事件
// const emit = defineEmits<{
//   (e: 'update:modelValue', value: string): void;
// }>();

// // 处理输入
// const handleInput = (event: Event) => {
//   const target = event.target as HTMLInputElement;
//   emit('update:modelValue', target.value);
// };
defineProps(['value']);
// export default {
//   props: {
//     value: {
//       type: String,
//       default: '',
//     },
//   },
//   computed: {
//     source_uri: {
//       get() {
//         return this.value;
//       },
//       set(val) {
//         this.$emit('input', val);
//       },
//     },
//   },
// };
</script>
