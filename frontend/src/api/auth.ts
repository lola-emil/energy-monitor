import { axiosInstance } from "./axios";


interface SignInForm {
    username: string
    password: string
};

export async function signIn(form: SignInForm) {
    return axiosInstance.post("/auth/login", form)
}