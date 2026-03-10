import type { AxiosResponse } from "axios";
import { ref } from "vue";


export const useAPI = (apiFunc: () => Promise<AxiosResponse>) => {
    const isLoading = ref(false)
    const apiResponse = ref<unknown>();
    const apiError = ref<unknown>();

    const request = () => {
        isLoading.value = true;

        apiFunc().then(res => {
            apiResponse.value = res;
        }).catch(err => {
            apiError.value = err
        }).finally(() => {
            isLoading.value = false
        })
    }

    return {
        isLoading,
        apiResponse,
        request
    }
};