<template>
  <div class="user-activation">
    <div class="container">
      <div class="title">Enter activation token</div>
      <InputGroup class="input">
        <InputText placeholder="Token" />
        <Button label="Send" />
      </InputGroup>
      <Button
        label="Resend token"
        severity="success"
        variant="text"
        @click="resendToken"
      />
    </div>
  </div>
</template>
<script lang="ts" setup>
import { InputGroup, Button, InputText } from 'primevue'
import { services } from '@/services'

defineOptions({
  name: 'UserActivation',
})

async function resendToken() {
  // services.account.resendActivationToken()
}
</script>
<style lang="scss" scoped>
.user-activation {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  > .container {
    padding: 48px;
    width: 600px;
    background-color: #fff;
    border-radius: 4px;
  }
  > .container > .back {
    margin-bottom: 16px;
  }
  > .container > .title {
    margin-bottom: 8px;
    font-size: 18px;
    font-weight: 600;
  }
  > .container > .input {
    margin-bottom: 8px;
  }
}
</style>
