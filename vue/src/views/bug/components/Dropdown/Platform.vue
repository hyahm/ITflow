<template>
  <el-dropdown :hide-on-click="false" :show-timeout="100" trigger="click">
    <el-button plain>
      平台({{ platforms.length }})
      <i class="el-icon-caret-bottom el-icon--right" />
    </el-button>
    <el-dropdown-menu slot="dropdown" class="no-border">
      <el-checkbox-group v-model="platforms" style="padding-left: 15px;">
        <el-checkbox v-for="(item, index) in platformsOptions" :key="index" :label="item">
          {{ item }}
        </el-checkbox>
      </el-checkbox-group>
    </el-dropdown-menu>
  </el-dropdown>
</template>

<script>
import { getStatus } from '@/api/get'
export default {
  // props: ['value'],
  data() {
    return {
      platformsOptions: []
    }
  },
  computed: {
    platforms: {
      get() {
        return this.value
      },
      set(val) {
        this.$emit('input', val)
      }
    }
  },
  created() {
    this.getstatus()
  },
  methods: {
    getstatus() {
      getStatus().then(resp => {
        if (resp.data.code === 0) {
          this.platformsOptions = resp.data.statuslist
        }
      })
    }
  }
}
</script>
