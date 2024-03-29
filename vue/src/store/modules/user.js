import { login, logout, getInfo } from "@/api/user";
import { getToken, setToken, removeToken, setTimeout } from "@/utils/auth";
// import router, { resetRouter } from '@/router'
// import router from '@/router'

const state = {
  token: getToken(),
  name: "",
  avatar: "",
  introduction: "",
  roles: [],
  uid: 0
};

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token;
  },
  SET_INTRODUCTION: (state, introduction) => {
    state.introduction = introduction;
  },
  SET_NAME: (state, name) => {
    state.name = name;
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar;
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles;
  },
  SET_UID: (state, uid) => {
    state.uid = uid;
  }
};

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { username, password } = userInfo;
    return new Promise((resolve, reject) => {
      login({ username: username.trim(), password: password })
        .then(response => {
          const { data } = response;
          console.log(data)
          if (data.code === 0) {
            commit("SET_TOKEN", data.data);
            setToken(data.data);
            setTimeout();
            resolve();
          } else {
            reject();
          }
        })
        .catch(error => {
          reject(error);
        });
    });
  },

  setHeadImage({ commit, state }) {
    commit("SET_AVATAR", state.avatar);
  },

  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      getInfo().then(response => {
          const { data } = response;
          if (data.code !== 0) {
            reject("Verification failed, please Login again.");
          }

          const { roles, name, avatar, uid } = data.data;
          // roles must be a non-empty array
          if (!roles || roles.length <= 0) {
            reject("getInfo: roles must be a non-null array!");
          }

          commit("SET_ROLES", roles);
          commit("SET_NAME", name);
          commit("SET_AVATAR", avatar);
          commit("SET_UID", uid);
          // commit('SET_INTRODUCTION', introduction)
          resolve(roles);
        })
        .catch(error => {
          reject(error);
        });
    });
  },

  // user logout
  logout({ commit, state, dispatch }) {
    return new Promise((resolve, reject) => {
      commit("SET_TOKEN", "");
      commit("SET_ROLES", []);
      removeToken();
      // resetRouter()

      // reset visited views and cached views
      // to fixed https://github.com/PanJiaChen/vue-element-admin/issues/2485
      dispatch("tagsView/delAllViews", null, { root: true });
      resolve();
    });
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      commit("SET_TOKEN", "");
      commit("SET_ROLES", []);
      removeToken();
      resolve();
    });
  }

  // dynamically modify permissions
  // changeRoles({ commit, dispatch }, role) {
  //   return new Promise(async resolve => {
  //     const token = role + '-token'

  //     commit('SET_TOKEN', token)
  //     setToken(token)

  //     const { roles } = await dispatch('getInfo')

  //     // resetRouter()

  //     // generate accessible routes map based on roles
  //     const accessRoutes = await dispatch('permission/generateRoutes', roles, { root: true })

  //     // dynamically add accessible routes
  //     router.addRoutes(accessRoutes)

  //     // reset visited views and cached views
  //     dispatch('tagsView/delAllViews', null, { root: true })

  //     resolve()
  //   })
  // }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
};
