# PFN\_vkDebugUtilsMessengerCallbackEXT(3) Manual Page

## Name

PFN\_vkDebugUtilsMessengerCallbackEXT - Application-defined debug messenger callback function



## [](#_c_specification)C Specification

The prototype for the [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCreateInfoEXT.html)::`pfnUserCallback` function implemented by the application is:

```c++
// Provided by VK_EXT_debug_utils
typedef VkBool32 (VKAPI_PTR *PFN_vkDebugUtilsMessengerCallbackEXT)(
    VkDebugUtilsMessageSeverityFlagBitsEXT           messageSeverity,
    VkDebugUtilsMessageTypeFlagsEXT                  messageTypes,
    const VkDebugUtilsMessengerCallbackDataEXT*      pCallbackData,
    void*                                            pUserData);
```

## [](#_parameters)Parameters

- `messageSeverity` specifies the [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html) that triggered this callback.
- `messageTypes` is a bitmask of [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html) specifying which type of event(s) triggered this callback.
- `pCallbackData` contains all the callback related data in the [VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCallbackDataEXT.html) structure.
- `pUserData` is the user data provided when the [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerEXT.html) was created.

## [](#_description)Description

The callback returns a [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), which is interpreted in a layer-specified manner. The application **should** always return `VK_FALSE`. The `VK_TRUE` value is reserved for use in layer development.

Valid Usage

- [](#VUID-PFN_vkDebugUtilsMessengerCallbackEXT-None-04769)VUID-PFN\_vkDebugUtilsMessengerCallbackEXT-None-04769  
  The callback **must** not make calls to any Vulkan commands

## [](#_see_also)See Also

[VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html), [VkDebugUtilsMessageTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageTypeFlagsEXT.html), [VkDebugUtilsMessengerCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCallbackDataEXT.html), [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCreateInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PFN_vkDebugUtilsMessengerCallbackEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0