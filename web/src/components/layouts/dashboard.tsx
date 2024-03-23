import { Outlet } from "react-router";
import { MainNav } from "../nav";
import { Sidebar } from "../sidebar";
import { Footer } from "../footer";
import { UserAccountNav } from "../user-account-nav";
import { useAuth } from "@/context/auth";

export function DashboardLayout() {
  const { user } = useAuth();

  return (
    <>
      <header className="sticky top-0 z-40 border-b bg-background">
        <div className="container flex h-16 items-center justify-between py-4">
          <MainNav />
          {user ? (
            <UserAccountNav
              user={{
                name: user.name,
                picture: user.picture,
              }}
            />
          ) : null}
        </div>
      </header>
      <div className="mt-5 container grid flex-1 gap-12 md:grid-cols-[200px_1fr]">
        <aside className="hidden w-[200px] flex-col md:flex">
          <Sidebar />
        </aside>
        <main className="flex w-full flex-1 flex-col overflow-hidden">
          <Outlet />
        </main>
      </div>
      <Footer className="border-t fixed bottom-0 left-0 z-20 w-full p-1 md:flex md:items-center md:justify-between md:p-2" />
    </>
  );
}
