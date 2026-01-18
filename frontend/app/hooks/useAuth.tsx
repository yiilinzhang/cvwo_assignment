import { useQuery } from "@tanstack/react-query";
import axios from "axios";

type UserResponse = { payload?: { data?: Number } };
//Custom hook to check if user is authenticated + username
export function useAuth() {
  const { isLoading, data, isError } = useQuery<UserResponse>({
    queryKey: ["me"],
    queryFn: async () => {
      const url = "http://localhost:8000/me";
      const response = await axios.get(url, { withCredentials: true });
      return await response.data;
    },
    retry: false,
  });

  return {
    user: isError ? null : data,
    isLoading,
  };
}
