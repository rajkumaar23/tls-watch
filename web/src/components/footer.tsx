import * as React from "react";

import { cn } from "@/lib/utils";
import { ModeToggle } from "@/components/mode-toggle";

export function Footer({ className }: React.HTMLAttributes<HTMLElement>) {
  return (
    <footer className={cn(className)}>
      <div className="container flex items-center justify-between gap-4 py-10 md:h-24 flex-row md:py-0">
        <div className="flex flex-col items-center md:flex-row md:px-0">
          <p className="text-sm leading-loose text-left">
            built by{" "}
            <a
              href="https://rajkumaar.co.in"
              target="_blank"
              rel="noreferrer"
              className="font-medium underline underline-offset-4"
            >
              rajkumar
            </a>
          </p>
        </div>
        <ModeToggle />
      </div>
    </footer>
  );
}
