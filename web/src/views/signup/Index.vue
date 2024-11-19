<template>
  <div class="signup-page">
    <div class="wrapper">
      <div class="form-container">
        <div class="logo">
          <img src="/images/logoG.png" alt="" class="image" />
        </div>
        <Message v-if="!!errorMessage" severity="error" class="message">{{
          errorMessage
        }}</Message>
        <label for="email" class="label">Email</label>
        <InputText
          v-model="state.email"
          id="email"
          type="text"
          placeholder="Email address"
          class="input"
        />

        <label for="password" class="label">Password</label>
        <InputText
          v-model="state.password"
          id="password"
          type="password"
          placeholder="Password"
          class="input"
        />

        <label for="confirmPassword" class="label">Confirm password</label>
        <InputText
          v-model="state.confirmPassword"
          id="confirmPassword"
          type="password"
          placeholder="Confirm password"
          class="input"
        />
        <Button
          :disabled="!!v$.$silentErrors.length"
          label="Register"
          class="button"
          @click="register"
        />
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import router from '@/router'
import { useAccountStore } from '@/stores/account'
import useVuelidate from '@vuelidate/core'
import { required } from '@vuelidate/validators'
import { InputText, Button, Message } from 'primevue'
import { reactive, ref } from 'vue'

defineOptions({
  name: 'SignupPage',
})

const accountStore = useAccountStore()

const state = reactive({
  email: '',
  password: '',
  confirmPassword: '',
})

const errorMessage = ref('')

const rules = {
  email: { required },
  password: { required },
}

const v$ = useVuelidate(rules, state)

async function register() {
  if (state.password !== state.confirmPassword) {
    errorMessage.value = 'passwords are not matched'
    return
  }
  if (state.password.length < 8) {
    errorMessage.value = 'password must be at least 8 charaters long'
    return
  }
  const res = await accountStore.signup({
    email: state.email,
    password: state.password,
  })
  if (!res.success) {
    errorMessage.value = res.message
    return
  } else {
    router.push('/login')
  }
}
</script>
<style lang="scss" scoped>
.signup-page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  > .wrapper {
    border-radius: 56px;
    padding: 0.3rem;
    background: linear-gradient(
      180deg,
      var(--primary-color) 10%,
      rgba(33, 150, 243, 0) 30%
    );
  }
}

.form-container {
  display: flex;
  flex-direction: column;
  padding: 5rem;
  border-radius: 53px;
  background-color: #fff;
  > .logo {
    display: flex;
    justify-content: center;
    margin-bottom: 1.5rem;
  }
  > .logo > .image {
    width: 4rem;
  }
  > .message {
    margin-bottom: 1.5rem;
  }
  > .label {
    margin-bottom: 0.2rem;
    font-weight: 500;
    color: #0f172a;
  }
  > .input {
    width: 30rem;
    margin-bottom: 1.5rem;
  }
}
</style>
