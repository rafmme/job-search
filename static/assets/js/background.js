const MY_CUSTOM_RULE_ID = 1;



chrome.declarativeNetRequest.updateDynamicRules({
  removeRuleIds: [MY_CUSTOM_RULE_ID],
  addRules: [
    {
      id: MY_CUSTOM_RULE_ID,
      priority: 1,
      action: {
        type: "modifyHeaders",
        requestHeaders: [
          {
            operation: "set",
            header: "User-Agent",
            value: "PostmanRuntime/7.28.0",
          },
          {
            operation: "set",
            header: "Content-Type",
            value: "application/json",
          },
        ],
      },
      condition: {
        resourceTypes: ["main_frame", "sub_frame"],
      },
    },
  ],
});
