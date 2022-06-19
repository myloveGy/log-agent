<template>
  <div class="body">
    <div class="container">
      <div class="div-description">
        <img src="./1.jpg" alt="">
        <span>日志收集系统</span>
      </div>
      <div class="div-form">
        <form class="form-login" @submit.stop.prevent="submit">
          <h1>日志收集系统</h1>
          <input type="text" name="username" v-model="form.username" placeholder="帐号">
          <input type="password" v-model="form.password" placeholder="密码">
          <button type="submit">登录</button>
        </form>
      </div>
      <div class="div-description">
        <img src="./1.jpg" alt="">
        <span>日志收集系统</span>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import {ElMessage} from 'element-plus'
import {userLoginApi, UserRequest} from '../api'
import {ref, onMounted} from 'vue'
import {userStore} from '../utils'
import {useRouter} from 'vue-router'
const router = useRouter()
const form = ref<UserRequest>({username: '', password: ''})
const submit = () => {
  userLoginApi(form.value).then(data => {
    userStore.save(data)
    router.push('/database')
  }).catch(e => {
    ElMessage.error('登录失败: ' + e?.message)
  })
}

onMounted(() => {
  if (userStore.fetch('token')) {
    router.push('/database')
  }
})
</script>
<style scoped>
.body {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  margin: 0;
  background-image: linear-gradient(120deg, #a1c4fd, #c2e9fb);
}

.container {
  position: relative;
  display: flex;
  /* 视域 */
  perspective: 1200px;
  transform-style: preserve-3d;
}

.div-description {
  width: 300px;
  height: 400px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  background-color: #fff;
  border-radius: 5px 0 0 5px;
}

.div-description:nth-last-child(1) {
  border-radius: 0 5px 5px 0;
}

.div-description img {
  width: 75%;
  margin-bottom: 15px;
}

.div-description span {
  font-size: 12px;
  color: #555;
}

.div-form {
  width: 300px;
  height: 400px;
  position: absolute;
  left: 0;
  transition: .5s;
  transform-origin: right;
}

.div-form .form-login {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  flex-direction: column;
  background-image: linear-gradient(to right, #48c6ef 0%, #6f86d6 200%);
  border-radius: 5px 0 0 5px;
}

.div-form h1 {
  margin: 50px 0 25px 0;
  /* 大写 */
  text-transform: uppercase;
  color: #fff;
  letter-spacing: 5px;
}

.div-form input {
  background-color: transparent;
  width: 70%;
  color: #fff;
  outline: none;
  border: none;
  border-bottom: 2px solid rgba(255, 255, 255, 0.7);
  padding: 5px 0;
  margin: 8px 0;
  text-indent: 5px;
  font-size: 14px;
  letter-spacing: 1px;
}

.div-form input::placeholder {
  color: #fff;
}

.div-form button[type="submit"] {
  width: 70%;
  height: 32px;
  margin: 30px auto 10px;
  font-size: 14px;
  color: #fff;
  border: none;
  border-radius: 16px;
  background-color: rgba(255, 255, 255, 0.2);
  transition: .4s;
}

.div-form button[type="submit"]:hover {
  letter-spacing: 3px;
  background-color: #fff;
  color: #000;
}

.div-form .control a {
  color: #fff;
  margin: 0 5px;
  letter-spacing: 1px;
}

.div-form .control a:hover {
  color: #000;
}
</style>

