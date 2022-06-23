<template>
  <el-form ref="formRef" :model="form" :rules="rules" class="form">
    <h1>{{ title }}</h1>
    <p>{{ description }}</p>
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
    <el-button
        type="primary"
        @click="submit"
        :loading="loading"
        :disabled="loading"
    >
      {{ submitText }}
    </el-button>
    <div class="control">
      <slot name="footer"></slot>
    </div>
  </el-form>
</template>

<script setup lang="ts">
import {Lock, User} from '@element-plus/icons-vue'
import {defineProps, withDefaults, ref, defineEmits} from 'vue'
import {LoginUser, UserRequest} from '@/api'
import {FormInstance} from 'element-plus'
import {sync} from '@/utils'

const loading = ref(false)

const formRef = ref<FormInstance>()

const rules = {
  username: [{required: true, message: '请输入用户名称'}],
  password: [
    {required: true, message: '请输入登录密码'},
    {min: 6, message: '密码需要至少6位'},
  ],
}

const form = ref<UserRequest>({username: '', password: ''})

interface Props {
  title: string,
  description: string,
  submitText: string,
  api: (data: UserRequest) => Promise<LoginUser>
}

const props = withDefaults(defineProps<Props>(), {
  title: '',
  description: '',
  submitText: '登录',
})

const emit = defineEmits<{
  (e: 'submit', data: LoginUser): void,
}>()

const submit = async () => {
  await formRef.value?.validate()
  loading.value = true
  return sync(async () => {
    const data = await props.api(form.value)
    emit('submit', data)
  }).finally(() => loading.value = false)
}

</script>

<style scoped lang="scss">
.form {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  height: 100%;
  width: 100%;
  text-align: center;
  background: #fff;

  p {
    font-size: 12px;
    color: #555;
    margin: 8px auto 12px;
  }

  h1 {
    font-size: 32px;
    margin-top: 0;
    text-transform: uppercase;
    letter-spacing: 5px;
    margin-bottom: 8px;
  }

  .el-form-item {
    margin-top: 8px;
    width: 80%;
  }

  button {
    width: 80%;
    height: 32px;
    margin: 8px auto 8px;
    font-size: 14px;
    transition: 0.4s;
    letter-spacing: 10px;
  }

  .control {
    margin: 8px auto;
    font-size: 13px;
  }
}
</style>
