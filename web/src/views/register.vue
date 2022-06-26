<template>
  <div class="body">
    <div class="container">
      <div class="div-description"></div>
      <div
        class="div-form"
        :style="{transform: action === 'register' ? 'rotateY(180deg)' : 'none'}"
      >
        <user-register
          title="日志收集系统"
          description="请输入账号密码登录"
          submit-text="登录"
          :api="loginApi"
          :class="{disappear: action === 'register', login: true}"
          @submit="handler"
        >
          <template #footer v-if="allow === 'Y'">
            <span>没有帐号？<a @click="action = 'register'">去注册</a></span>
          </template>
        </user-register>
        <user-register
          title="日志收集系统"
          description="请输入账号密码注册"
          submit-text="注册"
          :api="registerApi"
          :class="{disappear: action === 'login', register: true}"
          @submit="handler"
        >
          <template #footer>
            <span>已有帐号？<a @click="action = 'login'">去登录</a></span>
          </template>
        </user-register>
      </div>
      <div class="div-description"></div>
    </div>
  </div>
</template>
<script setup lang="ts">
import {useLogin} from '@/hooks'
import UserRegister from '@/components/UserRegister.vue'
import {loginApi, registerApi} from '@/api'

const {action, handler, allow} = useLogin()
</script>
<style scoped>
.body {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  margin: 0;
  background: url('@/assets/background.jpg') no-repeat fixed center;
}

.container {
  position: relative;
  display: flex;
  perspective: 1200px;
  transform-style: preserve-3d;
  background: url('@/assets/background.jpg') no-repeat fixed center;
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

.div-description span {
  font-size: 12px;
  color: #555;
}

.div-form {
  width: 300px;
  height: 400px;
  position: absolute;
  left: 0;
  transition: 0.5s;
  transform-origin: right;
}

.login {
  background-image: linear-gradient(to right, #48c6ef 0%, #6f86d6 200%);
  border-radius: 5px 0 0 5px;
}

.register {
  background-image: linear-gradient(to right, #48c6ef -100%, #6f86d6 100%);
  border-radius: 0 5px 5px 0;
  transform: rotateY(180deg);
}

.disappear {
  display: none;
}

a {
  color: #fff;
  margin: 0 5px;
  letter-spacing: 1px;
  cursor: pointer;
}

a:hover {
  color: #000;
}
</style>
