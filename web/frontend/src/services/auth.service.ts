import api from "./api"

export const authService = {
  async login(username: string, password: string) {
    const res = await api.post("/auth/login", {
      username,
      password,
    })

    return res.data
  },
}