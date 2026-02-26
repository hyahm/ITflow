import { defineStore } from 'pinia';
import { reactive, ref } from 'vue';
import type { KeyName } from '../types/response';

export const useUserStore = defineStore('user', () => {
  // 存储 所有用户信息
  const userMap = reactive(new Map<number, string>());
  const headImg = ref('');
  // 填充用户
  function GetAll(users: KeyName[]) {
    userMap.clear();
    for (const user of users) {
      userMap.set(user.id, user.name);
    }
  }

  function UpdateUser(user: KeyName) {
    userMap.set(user.id, user.name);
  }

  function DeleteUser(id: number) {
    userMap.delete(id);
  }


  function SetHead(url: string) {
    headImg.value = url
  }
  return {
    userMap,
    GetAll,
    UpdateUser,
    DeleteUser,
    SetHead,
    headImg
  };
});
