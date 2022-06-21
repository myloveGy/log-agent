<template>
  <div class="body">
    <div class="container">
      <div class="div-description">
      </div>
      <div class="div-form" :style="{transform: action === 'register' ? 'rotateY(180deg)' : 'none'}">
        <form action="" class="form-login" :class="{disappear: action === 'register'}">
          <h1>日志收集系统</h1>
          <p>请输入账号密码登录</p>
          <input type="text" placeholder="帐号">
          <input type="password" placeholder="密码">
          <button type="submit">登录</button>
          <div class="control">
            <span>没有帐号？<a @click="action='register'">去注册</a></span>
          </div>
        </form>
        <form action="" class="form-register" :class="{disappear: action === 'login'}">
          <h1>日志收集系统</h1>
          <p>请输入账号密码注册</p>
          <input type="text" placeholder="帐号">
          <input type="password" placeholder="密码">
          <button type="submit">注册</button>
          <div class="control">
            <span>已有帐号？<a @click="action='login'">去登录</a></span>
          </div>
        </form>
      </div>
      <div class="div-description"></div>
    </div>
  </div>
</template>
<script setup lang="ts">
import {ElMessage} from 'element-plus'
import {loginApi, UserRequest} from '@/api'
import {ref, onMounted} from 'vue'
import {userStore} from '@/utils'
import {useRouter} from 'vue-router'
import Form from '@/views/user/Form.vue'

const action = ref('login')
const router = useRouter()
const form = ref<UserRequest>({username: '', password: ''})
const submit = () => {
  loginApi(form.value).then(data => {
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
  background: url("@/assets/background.jpg") no-repeat fixed center;
}

.container {
  position: relative;
  display: flex;
  perspective: 1200px;
  transform-style: preserve-3d;
  background: url("@/assets/background.jpg") no-repeat fixed center;
  background-size: cover;
  transform: translateX(0);
  transition: transform 0.6s ease-in-out;
  border-radius: 5px;
}

.div-description {
  width: 300px;
  height: 400px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  border-radius: 5px 0 0 5px;

}

.div-description:nth-last-child(1) {
  border-radius: 0 5px 5px 0;
}

.div-description img {
  width: 75%;
  margin-bottom: 15px;
}

.div-description span, form p {
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

.div-form .form-register {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  flex-direction: column;
  background-image: linear-gradient(to right, #48c6ef -100%, #6f86d6 100%);
  border-radius: 0 5px 5px 0;
  transform: rotateY(180deg);
}

.div-form form.disappear {
  display: none;
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

.div-form .control {
  color: #fff;
  margin: 5px;
  font-size: 13px;
}

.div-form .control a {
  color: #fff;
  margin: 0 5px;
  letter-spacing: 1px;
  cursor: pointer;
}

.div-form .control a:hover {
  color: #000;
}
</style>

