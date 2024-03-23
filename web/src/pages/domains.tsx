import { Icons } from "@/components/icons";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Domain } from "@/lib/types";
import { API } from "@/lib/utils";
import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

export function Domains() {
  const [domains, setDomains] = useState<Domain[] | null>(null);

  useEffect(() => {
    const fetchDomains = async () => {
      const { data } = await API.get("/domains/");
      setDomains(data.domains);
    };

    if (!domains) {
      fetchDomains();
    }
  });

  return (
    <>
      <main className="flex w-full flex-1 flex-col overflow-hidden">
        <div className="grid items-start gap-8">
          <div className="flex items-center justify-between px-2">
            <div className="grid gap-1">
              <h1 className="font-bold text-3xl md:text-4xl">domains</h1>
              <p className="text-lg text-muted-foreground">
                manage your domains for monitoring tls certificates
              </p>
            </div>
            <button className="inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:opacity-50 disabled:pointer-events-none ring-offset-background bg-primary text-primary-foreground hover:bg-primary/90 h-10 py-2 px-2">
              <Icons.add className="mr-1" />
              add domain
            </button>
          </div>
          <div className="divide-y divide-border rounded-md border">
            {domains
              ? domains.map((domain, idx) => (
                  <div
                    key={idx}
                    className="flex items-center justify-between p-4"
                  >
                    <div className="grid gap-1">
                      <div className="flex gap-6 md:gap-10">
                        <Link
                          to={`https://${domain.domain}`}
                          className="items-center space-x-1 flex"
                          target="_blank"
                        >
                          <span className="font-semibold sm:inline-block">
                            {domain.domain}
                          </span>
                          <Icons.externalLink className="h-4 w-4" />
                        </Link>
                      </div>
                      <div>
                        <p className="text-sm text-muted-foreground">
                          {domain.created_at}
                        </p>
                      </div>
                    </div>
                    <DropdownMenu>
                      <DropdownMenuTrigger className="flex h-8 w-8 items-center justify-center rounded-md border transition-colors hover:bg-muted">
                        <Icons.ellipsis className="h-4 w-4" />
                        <span className="sr-only">Open</span>
                      </DropdownMenuTrigger>
                      <DropdownMenuContent align="end">
                        <DropdownMenuItem>Edit</DropdownMenuItem>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem
                          className="flex cursor-pointer items-center text-destructive focus:text-destructive"
                          onSelect={() => console.log("deleting domain ")}
                        >
                          Delete
                        </DropdownMenuItem>
                      </DropdownMenuContent>
                    </DropdownMenu>
                  </div>
                ))
              : null}
          </div>
        </div>
      </main>
    </>
  );
}
