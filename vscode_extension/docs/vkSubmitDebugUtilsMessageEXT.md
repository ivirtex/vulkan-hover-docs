# vkSubmitDebugUtilsMessageEXT(3) Manual Page

## Name

vkSubmitDebugUtilsMessageEXT - Inject a message into a debug stream



## [](#_c_specification)C Specification

To intentionally submit a debug message, call:

```c++
// Provided by VK_EXT_debug_utils
void vkSubmitDebugUtilsMessageEXT(
    VkInstance                                  instance,
    VkDebugUtilsMessageSeverityFlagBitsEXT      messageSeverity,
    VkDebugUtilsMessageTypeFlagsEXT             messageTypes,
    const VkDebugUtilsMessengerCallbackDataEXT* pCallbackData);
```

## [](#_parameters)Parameters

- `instance` is the debug stream’s [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html).
- `messageSeverity` is a [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html) value specifying the severity of this event/message.
- `messageTypes` is a bitmask of [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html) specifying which type of event(s) to identify with this message.
- `pCallbackData` contains all the callback related data in the [VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCallbackDataEXT.html) structure.

## [](#_description)Description

The call will propagate through the layers and generate callback(s) as indicated by the message’s flags. The parameters are passed on to the callback in addition to the `pUserData` value that was defined at the time the messenger was registered.

Valid Usage

- [](#VUID-vkSubmitDebugUtilsMessageEXT-objectType-02591)VUID-vkSubmitDebugUtilsMessageEXT-objectType-02591  
  The `objectType` member of each element of `pCallbackData->pObjects` **must** not be `VK_OBJECT_TYPE_UNKNOWN`

Valid Usage (Implicit)

- [](#VUID-vkSubmitDebugUtilsMessageEXT-instance-parameter)VUID-vkSubmitDebugUtilsMessageEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkSubmitDebugUtilsMessageEXT-messageSeverity-parameter)VUID-vkSubmitDebugUtilsMessageEXT-messageSeverity-parameter  
  `messageSeverity` **must** be a valid [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html) value
- [](#VUID-vkSubmitDebugUtilsMessageEXT-messageTypes-parameter)VUID-vkSubmitDebugUtilsMessageEXT-messageTypes-parameter  
  `messageTypes` **must** be a valid combination of [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html) values
- [](#VUID-vkSubmitDebugUtilsMessageEXT-messageTypes-requiredbitmask)VUID-vkSubmitDebugUtilsMessageEXT-messageTypes-requiredbitmask  
  `messageTypes` **must** not be `0`
- [](#VUID-vkSubmitDebugUtilsMessageEXT-pCallbackData-parameter)VUID-vkSubmitDebugUtilsMessageEXT-pCallbackData-parameter  
  `pCallbackData` **must** be a valid pointer to a valid [VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCallbackDataEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html), [VkDebugUtilsMessageTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageTypeFlagsEXT.html), [VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCallbackDataEXT.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSubmitDebugUtilsMessageEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0