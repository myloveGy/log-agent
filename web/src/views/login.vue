<template>
  <div class="body">
    <div
      class="container"
      :class="{'right-panel-active': action === 'register'}"
    >
      <div class="container-form container-signup">
        <user-form
          title="日志收集系统"
          description="请输入账号密码注册"
          :api="registerApi"
          submit-text="注册"
          @submit="handler"
        >
          <template #footer>
            <span>已有帐号？<a @click="action = 'login'">去登录</a></span>
          </template>
        </user-form>
      </div>
      <div class="container-form container-sign-in">
        <UserForm
          title="日志收集系统"
          description="请输入账号密码登录"
          :api="loginApi"
          submit-text="登录"
          @submit="handler"
        >
          <template #footer v-if="allow === 'Y'">
            <span>没有帐号？<a @click="action = 'register'">去注册</a></span>
          </template>
        </UserForm>
      </div>

      <div class="container-overlay">
        <div class="overlay">
          <div class="overlay-panel overlay-left">
            <el-button class="btn" @click="action = 'login'">
              去登录
            </el-button>
          </div>
          <div class="overlay-panel overlay-right">
            <el-button
              class="btn"
              @click="action = 'register'"
              v-if="allow === 'Y'"
            >
              去注册
            </el-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import {loginApi, registerApi} from '@/api'
import UserForm from '@/components/UserForm.vue'
import {useLogin} from '@/hooks'
const {action, handler, allow} = useLogin()
</script>
<style scoped lang="scss">
$white: #fff;
$lightblue: #008997;
$button-radius: 5px;
$max-width: 600px;
$max-height: 400px;

.body {
  align-items: center;
  background: $white url('@/assets/background.jpg') no-repeat fixed center;
  width: 100%;
  background-size: cover;
  display: grid;
  height: 100vh;
  place-items: center;
}

.container {
  background: $white;
  border-radius: $button-radius;
  box-shadow: 0 0.9rem 1.7rem rgba(0, 0, 0, 0.25),
    0 0.7rem 0.7rem rgba(0, 0, 0, 0.22);
  height: $max-height;
  max-width: $max-width;
  overflow: hidden;
  position: relative;
  width: 100%;
}

.container-form {
  height: 100%;
  position: absolute;
  top: 0;
  transition: all 0.6s ease-in-out;
}

.container-sign-in {
  left: 0;
  width: 50%;
  z-index: 2;
}

.container.right-panel-active .container-sign-in {
  transform: translateX(100%);
}

.container-signup {
  left: 0;
  opacity: 0;
  width: 50%;
  z-index: 1;
}

.container.right-panel-active .container-signup {
  animation: show 0.6s;
  opacity: 1;
  transform: translateX(100%);
  z-index: 5;
}

.container-overlay {
  height: 100%;
  left: 50%;
  overflow: hidden;
  position: absolute;
  top: 0;
  transition: transform 0.6s ease-in-out;
  width: 50%;
  z-index: 100;
}

.container.right-panel-active .container-overlay {
  transform: translateX(-100%);
}

.overlay {
  background: $lightblue url('@/assets/background.jpg') no-repeat fixed center;
  background-size: cover;
  height: 100%;
  left: -100%;
  position: relative;
  transform: translateX(0);
  transition: transform 0.6s ease-in-out;
  width: 200%;
}

.container.right-panel-active .overlay {
  transform: translateX(50%);
}

.overlay-panel {
  align-items: center;
  display: flex;
  flex-direction: column;
  height: 100%;
  justify-content: center;
  position: absolute;
  text-align: center;
  top: 0;
  transform: translateX(0);
  transition: transform 0.6s ease-in-out;
  width: 50%;
}

.overlay-left {
  transform: translateX(-20%);
}

.container.right-panel-active .overlay-left {
  transform: translateX(0);
}

.overlay-right {
  right: 0;
  transform: translateX(0);
}

.container.right-panel-active .overlay-right {
  transform: translateX(20%);
}

.container {
  .btn {
    width: 50%;
    height: 32px;
    margin: 30px auto 10px;
    font-size: 14px;
    color: #fff;
    border: none;
    border-radius: 16px;
    background-color: rgba(255, 255, 255, 0.5);
    transition: 0.4s;
    letter-spacing: 10px;
  }

  a {
    color: #409eff;
    margin: 0 5px;
    letter-spacing: 1px;
    cursor: pointer;

    &:hover {
      color: #000;
    }
  }
}

@keyframes show {
  0%,
  49.99% {
    opacity: 0;
    z-index: 1;
  }

  50%,
  100% {
    opacity: 1;
    z-index: 5;
  }
}
</style>
