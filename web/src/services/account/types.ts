export type AuthenticationToken = {
  token: string
  expiry: Date
}

export type InputLogin = {
  email: string
  password: string
}
