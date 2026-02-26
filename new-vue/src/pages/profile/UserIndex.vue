<template>
  <div class="app-container">
    <div v-if="user">
      <el-row :gutter="20">
        <el-col :span="6" :xs="24">
          <user-card :user="user" />
        </el-col>

        <el-col :span="18" :xs="24">
          <el-card>
            <el-tabs v-model="activeTab">
              <el-tab-pane label="Activity" name="activity">
                <profile-activity />
              </el-tab-pane>
              <el-tab-pane label="Timeline" name="timeline">
                <timeline />
              </el-tab-pane>
              <el-tab-pane label="Account" name="account">
                <account :user="user" />
              </el-tab-pane>
            </el-tabs>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script setup lang="ts">
// import { mapGetters } from 'vuex'
import UserCard from './components/UserCard.vue';
import ProfileActivity from './components/ProfileActivity.vue';
import Timeline from './components/TimeLine.vue';
import Account from './components/ProfileAccount.vue';
import { reactive, ref } from 'vue';
let user = reactive({});
const activeTab = ref('activity');
// computed: {
//   ...mapGetters([
//     'name',
//     'avatar',
//     'roles'
//   ])
// },
const getUser = () => {
  user = {
    name: '',
    role: '',
    email: 'admin@test.com',
    avatar: '',
  };
};
getUser();
</script>
