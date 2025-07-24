import axios from "axios";

export const apiWrapper = async (
  endpoint: string,
  method: "GET" | "POST" | "PUT" | "DELETE" = "GET",
  body?: Record<string, string | number>,
  customHeaders?: Record<string, string>
) => {
  const baseUrl = process.env.NEXT_PUBLIC_BASE_URL;

  const headers: Record<string, string> = {
    "Content-Type": "application/json",
    ...customHeaders,
  };

  try {
    const response = await axios({
      url: `${baseUrl}/${endpoint}`,
      method,
      headers,
      data: body,
    });

    return response.data;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      const serverError =
        error.response?.data?.message || error.response?.data?.error;

      if (serverError) {
        throw new Error(serverError);
      } else {
        throw new Error(error.message || "An unexpected error occurred");
      }
    }
    throw error;
  }
};
