import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";
import { useLocation, useNavigate } from "react-router";
import axios from "axios";

type UserResponse = { payload?: { data?: number } };
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
    refetchOnWindowFocus: false,
    staleTime: 60 * 10 * 1000,
  });

  return {
    userID: isError ? null : (data?.payload?.data ?? null),
    isLoading,
  };
}

export function useRequireAuth() {
  const { userID, isLoading } = useAuth();
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    if (!isLoading && !userID) {
      navigate("/sign-in", { replace: true, state: { from: location.pathname } });
    }
  }, [isLoading, userID, navigate, location.pathname]);

  return { userID, isLoading };
}
