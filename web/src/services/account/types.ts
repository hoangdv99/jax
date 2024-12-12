export type InputSignup = {
  email: string
  password: string
}

export type InputSignin = {
  email: string
  password: string
}

export type UserActivationPayload = {
  email: string
}

export type User = {
  id: number
  email: string
  status: number
}
