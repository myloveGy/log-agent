<template>
  <div class="body">
    <div class="container" :class="{'right-panel-active': action === 'login'}">
      <div class="container-form container-signup">
        <el-form class="form register" @submit.stop.prevent="register">
          <h1>日志收集系统</h1>
          <p>请输入账号密码注册</p>
          <el-input
              v-model="form.username"
              class="input"
              style="width: 80%"
              placeholder="帐号"
              :prefix-icon="User"
          />
          <el-input
              v-model="form.password"
              class="input"
              style="width: 80%"
              placeholder="密码"
              type="password"
              show-password
              :prefix-icon="Lock"
          />
          <el-button type="primary" @click="register" :loading="loading" :disabled="loading">
            注册
          </el-button>
          <div class="control">
            <span>已有帐号？<a @click="action='register'">去登录</a></span>
          </div>
        </el-form>
      </div>

      <div class="container-form container-sign-in">
        <el-form ref="formRef" :model="form" :rules="rules" class="form login" @submit.stop.prevent="login">
          <h1>日志收集系统</h1>
          <p>请输入账号密码登录</p>
          <el-form-item prop="username">
            <el-input
                v-model="form.username"
                class="input"
                placeholder="帐号"
                :prefix-icon="User"
            />
          </el-form-item>
          <el-form-item prop="password">
            <el-input
                v-model="form.password"
                class="input"
                placeholder="密码"
                type="password"
                show-password
                :prefix-icon="Lock"
            />
          </el-form-item>
          <el-button type="primary" @click="login" :loading="loading" :disabled="loading">
            登录
          </el-button>
          <div class="control">
            <span>没有帐号？<a @click="action='login'">去注册</a></span>
          </div>
        </el-form>
      </div>

      <div class="container-overlay">
        <div class="overlay">
          <!--          <div class="overlay-panel overlay-left">-->
          <!--            <button class="btn" @click="action = 'register'">登录</button>-->
          <!--          </div>-->
          <!--          <div class="overlay-panel overlay-right">-->
          <!--            <button class="btn" @click="action = 'login'">注册</button>-->
          <!--          </div>-->
        </div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import {ElMessage, FormInstance} from 'element-plus'
import {userLoginApi, userRegisterApi, UserRequest} from '@/api'
import {onMounted, ref} from 'vue'
import {userStore} from '@/utils'
import {useRouter} from 'vue-router'
import {Lock, User} from '@element-plus/icons-vue'

const formRef = ref<FormInstance>()
const rules = {
  username: [
    {required: true, message: '请输入用户名称'},
  ],
  password: [
    {required: true, message: '请输入登录密码'},
  ],
}

const loading = ref(false)
const router = useRouter()
const form = ref<UserRequest>({username: '', password: ''})
const action = ref('register')

const login = async () => {
  await formRef.value?.validate()
  loading.value = true
  userLoginApi(form.value).then(data => {
    userStore.save(data)
    router.push('/database')
  }).catch(e => {
    ElMessage.error('登录失败: ' + e?.message)
  }).finally(() => loading.value = false)
}

const register = () => {
  loading.value = true
  userRegisterApi(form.value).then(data => {
    userStore.save(data)
    router.push('/database')
  }).catch(e => {
    ElMessage.error('注册失败: ' + e?.message)
  }).finally(() => loading.value = false)
}

onMounted(() => {
  if (userStore.fetch('token')) {
    router.push('/database')
  }
})
</script>
<style scoped lang="scss">
$white: #fff;
$lightblue: #008997;
$button-radius: 5px;
$max-width: 600px;
$max-height: 400px;

.body {
  align-items: center;
  background: $white url("@/assets/background.jpg") no-repeat fixed center;
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
  background: $lightblue url("@/assets/background.jpg") no-repeat fixed center;
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

p {
  font-size: 12px;
  color: #555;
  margin: 8px auto 12px;
}

.form {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  height: 100%;
  width: 100%;
  text-align: center;
  background: $white;
}

.container {
  h1 {
    font-size: 32px;
    margin-top: 0;
    text-transform: uppercase;
    letter-spacing: 5px;
    margin-bottom: 8px;
  }

  //.input {
  //  margin: 8px 0 0 0;
  //}

  .el-form-item {
    margin-top: 8px;
    width: 80%;
  }

  button {
    width: 80%;
    height: 32px;
    margin: 8px auto 8px;
    font-size: 14px;
    transition: .4s;
    letter-spacing: 10px;
  }

  .control {
    margin: 8px auto;
    font-size: 13px;

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

