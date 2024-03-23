import { Link } from "react-router-dom";

export function MainNav() {
  return (
    <div className="flex gap-6 md:gap-10">
      <Link to="/" className="items-center space-x-2 flex">
        <span className="font-bold sm:inline-block">
          tls watch
        </span>
      </Link>
    </div>
  );
}
