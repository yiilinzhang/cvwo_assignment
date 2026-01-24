import { Outlet } from "react-router";
import { useRequireAuth } from "~/hooks/useAuth";

export default function ProtectedLayout() {
  const { userID, isLoading } = useRequireAuth();
  if (isLoading || !userID) return <div>Loading...</div>;
  return <Outlet />;
}
