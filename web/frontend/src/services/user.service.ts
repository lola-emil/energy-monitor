import api from "./api";
import type { User } from "@/types/user";

export const userService = {

    async getProfile(): Promise<User> {
        const response = await api.get("/user/profile");
        return response.data;
    }
};
