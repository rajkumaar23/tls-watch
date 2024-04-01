import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Checkbox } from "@/components/ui/checkbox";
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
  Form,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useToast } from "@/components/ui/use-toast";
import { NotificationSetting } from "@/lib/types";
import { API } from "@/lib/utils";
import { zodResolver } from "@hookform/resolvers/zod";
import { useCallback, useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";

const telegramFormSchema = z.object({
  user_id: z
    .string()
    .regex(/^[a-z0-9_]+$/, "should only contain a-z, 0-9 and underscores"),
  enabled: z.boolean(),
});

export function Settings() {
  const { toast } = useToast();
  const [telegramSettings, setTelegramSettings] =
    useState<NotificationSetting | null>(null);

  const fetchNotificationSettings = useCallback(async () => {
    try {
      const { data } = await API.get("/notifications/settings/");
      setTelegramSettings(
        data.settings.find(
          (it: NotificationSetting) => it.provider === "telegram"
        )
      );
    } catch (error) {
      toast({
        description: "error fetching telegram settings",
        variant: "destructive",
      });
      console.error(error);
    }
  }, []);

  useEffect(() => {
    fetchNotificationSettings();
  }, []);

  const telegramForm = useForm<z.infer<typeof telegramFormSchema>>({
    resolver: zodResolver(telegramFormSchema),
    defaultValues: {
      user_id: "",
      enabled: false,
    },
  });

  useEffect(() => {
    if (telegramSettings) {
      telegramForm.setValue("user_id", telegramSettings.provider_user_id);
      telegramForm.setValue("enabled", telegramSettings.enabled);
    }
  }, [telegramSettings, telegramForm]);

  const onTelegramSettingsSubmit = async (
    values: z.infer<typeof telegramFormSchema>
  ) => {
    try {
      const { data } = await API.post("/notifications/settings/create", {
        provider_user_id: values.user_id,
        enabled: values.enabled,
        provider: "telegram",
      });

      toast({ description: data.message });
    } catch (error) {
      toast({
        description: "error updating telegram settings",
        variant: "destructive",
      });
      console.error(error);
    }
  };

  return (
    <>
      <main className="flex w-full flex-1 flex-col overflow-hidden">
        <div className="grid items-start gap-8">
          <div className="flex items-center justify-between px-2">
            <div className="grid gap-1">
              <h1 className="font-bold text-3xl md:text-4xl">settings</h1>
            </div>
          </div>
          <Card className="lg:w-2/3">
            <CardHeader>
              <CardTitle>telegram</CardTitle>
            </CardHeader>
            <CardContent>
              <Form {...telegramForm}>
                <form
                  onSubmit={telegramForm.handleSubmit(onTelegramSettingsSubmit)}
                  className="space-y-2"
                >
                  <FormField
                    control={telegramForm.control}
                    name="user_id"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>your user id</FormLabel>
                        <FormControl>
                          <Input
                            placeholder="12345678"
                            {...field}
                            defaultValue={telegramSettings?.provider_user_id}
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={telegramForm.control}
                    name="enabled"
                    render={({ field }) => (
                      <FormItem className="flex flex-row items-start space-x-3 space-y-0 pl-1 py-2">
                        <FormControl>
                          <Checkbox
                            checked={field.value}
                            onCheckedChange={field.onChange}
                          />
                        </FormControl>
                        <FormLabel>enable notifications</FormLabel>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <Button type="submit" className="w-1/3">
                    save
                  </Button>
                </form>
              </Form>
            </CardContent>
          </Card>
        </div>
      </main>
    </>
  );
}
