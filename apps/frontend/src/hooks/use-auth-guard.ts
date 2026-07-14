import { useEffect, useMemo } from "react";
import { useCurrentUser } from "./use-current-user";

interface UseAuthGuardProps {
  onAuthenticated?: () => void;
  onUnauthenticated?: () => void;
}

export function useAuthGuard({
  onAuthenticated,
  onUnauthenticated,
}: UseAuthGuardProps) {
  const { data: currentUser } = useCurrentUser();

  useEffect(() => {
    if (currentUser) {
      onAuthenticated?.();
    } else {
      onUnauthenticated?.();
    }
  }, [currentUser, onAuthenticated, onUnauthenticated]);

  return useMemo(() => !!currentUser, [currentUser]);
}
