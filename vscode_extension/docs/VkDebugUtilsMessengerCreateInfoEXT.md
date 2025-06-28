# VkDebugUtilsMessengerCreateInfoEXT(3) Manual Page

## Name

VkDebugUtilsMessengerCreateInfoEXT - Structure specifying parameters of a newly created debug messenger



## [](#_c_specification)C Specification

The definition of `VkDebugUtilsMessengerCreateInfoEXT` is:

```c++
// Provided by VK_EXT_debug_utils
typedef struct VkDebugUtilsMessengerCreateInfoEXT {
    VkStructureType                         sType;
    const void*                             pNext;
    VkDebugUtilsMessengerCreateFlagsEXT     flags;
    VkDebugUtilsMessageSeverityFlagsEXT     messageSeverity;
    VkDebugUtilsMessageTypeFlagsEXT         messageType;
    PFN_vkDebugUtilsMessengerCallbackEXT    pfnUserCallback;
    void*                                   pUserData;
} VkDebugUtilsMessengerCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is `0` and is reserved for future use.
- `messageSeverity` is a bitmask of [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html) specifying which severity of event(s) will cause this callback to be called.
- `messageType` is a bitmask of [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html) specifying which type of event(s) will cause this callback to be called.
- `pfnUserCallback` is the application callback function to call.
- `pUserData` is user data to be passed to the callback.

## [](#_description)Description

For each `VkDebugUtilsMessengerEXT` that is created the `VkDebugUtilsMessengerCreateInfoEXT`::`messageSeverity` and `VkDebugUtilsMessengerCreateInfoEXT`::`messageType` determine when that `VkDebugUtilsMessengerCreateInfoEXT`::`pfnUserCallback` is called. The process to determine if the user’s `pfnUserCallback` is triggered when an event occurs is as follows:

1. The implementation will perform a bitwise AND of the event’s [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html) with the `messageSeverity` provided during creation of the [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerEXT.html) object.
   
   1. If the value is 0, the message is skipped.
2. The implementation will perform bitwise AND of the event’s [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html) with the `messageType` provided during the creation of the [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerEXT.html) object.
   
   1. If the value is 0, the message is skipped.
3. The callback will trigger a debug message for the current event

The callback will come directly from the component that detected the event, unless some other layer intercepts the calls for its own purposes (filter them in a different way, log to a system error log, etc.).

An application **can** receive multiple callbacks if multiple `VkDebugUtilsMessengerEXT` objects are created. A callback will always be executed in the same thread as the originating Vulkan call.

A callback **can** be called from multiple threads simultaneously (if the application is making Vulkan calls from multiple threads).

Valid Usage (Implicit)

- [](#VUID-VkDebugUtilsMessengerCreateInfoEXT-sType-sType)VUID-VkDebugUtilsMessengerCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEBUG_UTILS_MESSENGER_CREATE_INFO_EXT`
- [](#VUID-VkDebugUtilsMessengerCreateInfoEXT-flags-zerobitmask)VUID-VkDebugUtilsMessengerCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkDebugUtilsMessengerCreateInfoEXT-messageSeverity-parameter)VUID-VkDebugUtilsMessengerCreateInfoEXT-messageSeverity-parameter  
  `messageSeverity` **must** be a valid combination of [VkDebugUtilsMessageSeverityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagBitsEXT.html) values
- [](#VUID-VkDebugUtilsMessengerCreateInfoEXT-messageSeverity-requiredbitmask)VUID-VkDebugUtilsMessengerCreateInfoEXT-messageSeverity-requiredbitmask  
  `messageSeverity` **must** not be `0`
- [](#VUID-VkDebugUtilsMessengerCreateInfoEXT-messageType-parameter)VUID-VkDebugUtilsMessengerCreateInfoEXT-messageType-parameter  
  `messageType` **must** be a valid combination of [VkDebugUtilsMessageTypeFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageTypeFlagBitsEXT.html) values
- [](#VUID-VkDebugUtilsMessengerCreateInfoEXT-messageType-requiredbitmask)VUID-VkDebugUtilsMessengerCreateInfoEXT-messageType-requiredbitmask  
  `messageType` **must** not be `0`
- [](#VUID-VkDebugUtilsMessengerCreateInfoEXT-pfnUserCallback-parameter)VUID-VkDebugUtilsMessengerCreateInfoEXT-pfnUserCallback-parameter  
  `pfnUserCallback` **must** be a valid [PFN\_vkDebugUtilsMessengerCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkDebugUtilsMessengerCallbackEXT.html) value

## [](#_see_also)See Also

[PFN\_vkDebugUtilsMessengerCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkDebugUtilsMessengerCallbackEXT.html), [VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkDebugUtilsMessageSeverityFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageSeverityFlagsEXT.html), [VkDebugUtilsMessageTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessageTypeFlagsEXT.html), [VkDebugUtilsMessengerCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCreateFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDebugUtilsMessengerEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDebugUtilsMessengerCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0