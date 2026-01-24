import { createContext, useContext, useState } from "react";
import { Alert, Snackbar } from "@mui/material";

type ToastApi = {
  show: (message: string) => void;
};

type Severity = "success" | "error" | "info";

const ToastContext = createContext<(msg: string, s?: Severity) => void>(
  () => {}
);


export function ToastProvider({ children }: { children: React.ReactNode }) {
  const [open, setOpen] = useState(false);
  const [message, setMessage] = useState("");
  const [severity, setSeverity] = useState<Severity>("info");

  const notify = (msg: string, s: Severity = "info") => {
    setMessage(msg);
    setSeverity(s);
    setOpen(true);
  };

  return (
    <ToastContext.Provider value={ notify }>
      {children}
      <Snackbar
        open={open}
        autoHideDuration={3000}
        onClose={() => setOpen(false)}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert
          severity={severity}
          onClose={() => setOpen(false)}
        >
          {message}
        </Alert>
      </Snackbar>
    </ToastContext.Provider>
  );
}

export const useToast = () => useContext(ToastContext);