import { Icons } from "@/components/icons";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import { Domain } from "@/lib/types";
import { API } from "@/lib/utils";
import { useCallback, useEffect, useState } from "react";
import { z } from "zod";
import { Link } from "react-router-dom";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form";

export function Domains() {
  const [domains, setDomains] = useState<Domain[] | null>(null);

  const fetchDomains = useCallback(async () => {
    const { data } = await API.get("/domains/");
    setDomains(data.domains);
  }, []);

  useEffect(() => {
    if (!domains) {
      fetchDomains();
    }
  });

  const newDomainFormSchema = z.object({
    domain: z.string().min(2),
  });
  const newDomainForm = useForm<z.infer<typeof newDomainFormSchema>>({
    resolver: zodResolver(newDomainFormSchema),
    defaultValues: {
      domain: "",
    },
  });

  const onNewDomainSubmit = async (
    values: z.infer<typeof newDomainFormSchema>
  ) => {
    try {
      await API.post("/domains/create", {
        domain: values.domain,
      });
      newDomainForm.reset();
      fetchDomains();
    } catch {
      console.log("error adding a new domain");
    }
  };

  return (
    <>
      <main className="flex w-full flex-1 flex-col overflow-hidden">
        <div className="grid items-start gap-8">
          <div className="flex items-center justify-between px-2">
            <div className="grid gap-1">
              <h1 className="font-bold text-3xl md:text-4xl">domains</h1>
            </div>
            <Dialog>
              <DialogTrigger asChild>
                <Button className="px-2 pr-3">
                  <Icons.add className="mr-1" />
                  add new
                </Button>
              </DialogTrigger>
              <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                  <DialogTitle>add a new domain</DialogTitle>
                </DialogHeader>
                <Form {...newDomainForm}>
                  <form
                    onSubmit={newDomainForm.handleSubmit(onNewDomainSubmit)}
                    className="space-y-4"
                  >
                    <FormField
                      control={newDomainForm.control}
                      name="domain"
                      render={({ field }) => (
                        <FormItem>
                          <FormControl>
                            <Input placeholder="x.com" {...field} />
                          </FormControl>
                          <FormMessage />
                        </FormItem>
                      )}
                    />
                    <DialogClose asChild>
                      <Button className="w-full" type="submit">
                        Submit
                      </Button>
                    </DialogClose>
                  </form>
                </Form>
              </DialogContent>
            </Dialog>
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
