import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { UserAvatar } from "@/components/user-avatar";
import { API_URL, USER_SESSION_COOKIE } from "@/lib/constants";
import { User } from "@/lib/types";
import Cookies from "js-cookie";
import { HTMLAttributes } from "react";
import { Link } from "react-router-dom";

interface UserAccountNavProps extends HTMLAttributes<HTMLDivElement> {
  user: Pick<User, "name" | "picture">;
}

export function UserAccountNav({ user }: UserAccountNavProps) {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger>
        <UserAvatar
          user={{ name: user.name || "", picture: user.picture }}
          className="h-8 w-8"
        />
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <div className="flex items-center justify-start gap-2 p-2">
          <div className="flex flex-col space-y-1 leading-none">
            {user.name && (
              <p>
                hola,{" "}
                <span className="font-medium">{user.name.toLowerCase()}</span>
              </p>
            )}
            {/* {user.email && (
              <p className="w-[200px] truncate text-sm text-muted-foreground">
                {user.email}
              </p>
            )} */}
          </div>
        </div>
        <DropdownMenuSeparator />
        <DropdownMenuItem asChild className="cursor-pointer">
          <Link
            to={`${API_URL}/auth/logout`}
            onClick={() =>
              Cookies.remove(USER_SESSION_COOKIE, {
                domain: window.location.hostname,
                path: "/",
                secure: API_URL.startsWith("https"),
              })
            }
          >
            logout
          </Link>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
