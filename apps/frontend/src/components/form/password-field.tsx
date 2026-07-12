import { TextField } from "./text-field";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useState, type ComponentProps } from "react";

interface PasswordFieldProps extends Omit<
  ComponentProps<typeof TextField>,
  "type"
> {}

export function PasswordField(props: PasswordFieldProps) {
  const [showPassword, setShowPassword] = useState<boolean>(false);

  return (
    <TextField
      {...props}
      type={showPassword ? "text" : "password"}
      suffix={
        <button
          className="bg-teal-800 rounded-4xl size-8 hover:bg-teal-800/50 transition-colors duration-200 text-gray-100"
          type="button"
          onClick={() => setShowPassword(!showPassword)}
        >
          <FontAwesomeIcon icon={showPassword ? "eye-slash" : "eye"} />
        </button>
      }
    />
  );
}
