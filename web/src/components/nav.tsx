import { Icons } from "@/components/icons";
import { Link } from "react-router-dom";

export function MainNav() {
  return (
    <div className="flex gap-6 md:gap-10">
      <Link to="/" className="items-center space-x-2 flex">
        <Icons.watch />
        <span className="font-bold sm:inline-block text-lg">
          tls watch
          <span className="text-xs ml-1 text-muted-foreground space-x-1 hidden md:inline-block">as a service</span>
        </span>
      </Link>
    </div>
  );
}
