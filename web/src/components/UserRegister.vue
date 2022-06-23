<template>
  <form class="form" @submit.stop.prevent="submit">
    <h1>{{ title }}</h1>
    <p>{{ description }}</p>
    <input type="text" v-model="form.username" placeholder="帐号">
    <input type="password" v-model="form.password" placeholder="密码">
    <button type="submit">{{ submitText }}</button>
    <div class="control">
      <slot name="footer"/>
    </div>
  </form>
</template>

<script setup lang="ts">
import {defineEmits, defineProps, ref} from 'vue'
import {LoginUser, UserRequest} from '@/api'
import {sync} from '@/utils'
import {ElMessage} from 'element-plus'
import {debounce} from 'lodash'
const form = ref<UserRequest>({username: '', password: ''})

interface UserRegisterProps {
  title: string;
  description: string;
  submitText: string;
  api: (data: UserRequest) => Promise<LoginUser>
}

const props = defineProps<UserRegisterProps>()
const emit = defineEmits<{
  (e: 'submit', data: LoginUser): void
}>()

const error = debounce((message: string) => ElMessage.error(message), 360)

const submit = () => {
  if (!form.value.username) {
    return error('请输入账号')
  }

  if (!form.value.password) {
    return error('请输入密码')
  }

  return sync(async () => {
    const data = await props.api(form.value)
    await emit('submit', data)
  })
}

</script>

<style scoped lang="scss">
.form {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  flex-direction: column;

  p {
    font-size: 12px;
    color: #555;
  }

  h1 {
    margin: 50px 0 25px 0;
    text-transform: uppercase;
    color: #fff;
    letter-spacing: 5px;
  }

  input {
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

    &::placeholder {
      color: #fff;
    }
  }


  button[type="submit"] {
    width: 70%;
    height: 32px;
    margin: 30px auto 10px;
    font-size: 14px;
    color: #fff;
    border: none;
    border-radius: 16px;
    background-color: rgba(255, 255, 255, 0.2);
    transition: .4s;

    &:hover {
      letter-spacing: 3px;
      background-color: #fff;
      color: #000;
    }
  }

  .control {
    color: #fff;
    margin: 5px;
    font-size: 13px;

    a {
      color: #fff;
      margin: 0 5px;
      letter-spacing: 1px;
      cursor: pointer;

      &:hover {
        color: #000;
      }
    }
  }


}
</style>
