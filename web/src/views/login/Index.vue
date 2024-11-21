<template>
  <div class="login-page">
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
          v-model="email"
          id="email"
          type="text"
          placeholder="Email address"
          class="input"
        />

        <label for="password" class="label">Password</label>
        <InputText
          v-model="password"
          id="password"
          type="password"
          placeholder="Password"
          class="input"
        />

        <div class="wrapper">
          <div class="register">
            <span class="button" @click="$router.push('/signup')"
              >Create new account</span
            >
          </div>
          <div class="forgot">
            <span class="button">Forgot password?</span>
          </div>
        </div>
        <Button
          :disabled="!email || !password"
          label="Sign In"
          @click="signin"
        />
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import router from '@/router'
import { useAccountStore } from '@/stores/account'
import { InputText, Button, Message } from 'primevue'
import { ref } from 'vue'

defineOptions({
  name: 'LoginPage',
})

const accountStore = useAccountStore()

const email = ref('')
const password = ref('')
const errorMessage = ref('')

async function signin() {
  const res = await accountStore.signin({
    email: email.value,
    password: password.value,
  })
  if (!res.success) {
    errorMessage.value = 'invalid email or password'
    return
  } else {
    localStorage.setItem('authToken', res.data.authentication_token.token)
    localStorage.setItem(
      'authTokenExpiry',
      res.data.authentication_token.expiry
    )
    router.push({ name: 'HomePage' })
  }
}
</script>
<style lang="scss" scoped>
.login-page {
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
    margin-top: 1.5rem;
  }
  > .label {
    margin-bottom: 0.2rem;
    font-weight: 500;
    color: #0f172a;
  }
  > .input {
    margin-bottom: 1.5rem;
    width: 30rem;
  }
  > .wrapper {
    display: flex;
    justify-content: space-between;
    margin-top: 1rem;
    margin-bottom: 2rem;
  }
  > .wrapper > .register > .button,
  > .wrapper > .forgot > .button {
    color: #10b981;
    cursor: pointer;
  }
}
</style>
