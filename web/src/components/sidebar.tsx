import { cn } from "@/lib/utils";
import { useLocation } from "react-router";
import { Link } from "react-router-dom";
import { Icons } from "@/components/icons";

type SidebarNavItem = {
  title: string
  icon?: keyof typeof Icons
} & (
  | {
      href: string
      items?: never
    }
)

const sidebarItems: SidebarNavItem[] = [
  {
    title: "domains",
    href: "/",
    icon: "globe",
  },
  {
    title: "settings",
    href: "/settings",
    icon: "settings",
  },
];

export function Sidebar() {
  const { pathname: path } = useLocation();

  return (
    <nav className="grid items-start gap-2">
      {sidebarItems.map((item, index) => {
        const Icon = Icons[item.icon || "arrowRight"];
        return (
          item.href && (
            <Link key={index} to={item.href}>
              <span
                className={cn(
                  "group flex items-center rounded-md px-3 py-2 text-sm font-medium hover:bg-accent hover:text-accent-foreground",
                  path === item.href ? "bg-accent" : "transparent"
                )}
              >
                <Icon className="mr-2 h-4 w-4" />
                <span>{item.title}</span>
              </span>
            </Link>
          )
        );
      })}
    </nav>
  );
}
