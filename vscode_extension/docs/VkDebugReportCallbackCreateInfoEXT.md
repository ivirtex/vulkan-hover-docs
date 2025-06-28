# VkDebugReportCallbackCreateInfoEXT(3) Manual Page

## Name

VkDebugReportCallbackCreateInfoEXT - Structure specifying parameters of a newly created debug report callback



## [](#_c_specification)C Specification

The definition of [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html) is:

```c++
// Provided by VK_EXT_debug_report
typedef struct VkDebugReportCallbackCreateInfoEXT {
    VkStructureType                 sType;
    const void*                     pNext;
    VkDebugReportFlagsEXT           flags;
    PFN_vkDebugReportCallbackEXT    pfnCallback;
    void*                           pUserData;
} VkDebugReportCallbackCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagBitsEXT.html) specifying which event(s) will cause this callback to be called.
- `pfnCallback` is the application callback function to call.
- `pUserData` is user data to be passed to the callback.

## [](#_description)Description

For each `VkDebugReportCallbackEXT` that is created the `VkDebugReportCallbackCreateInfoEXT`::`flags` determine when that `VkDebugReportCallbackCreateInfoEXT`::`pfnCallback` is called. When an event happens, the implementation will do a bitwise AND of the event’s [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagBitsEXT.html) flags to each `VkDebugReportCallbackEXT` object’s flags. For each non-zero result the corresponding callback will be called. The callback will come directly from the component that detected the event, unless some other layer intercepts the calls for its own purposes (filter them in a different way, log to a system error log, etc.).

An application **may** receive multiple callbacks if multiple `VkDebugReportCallbackEXT` objects were created. A callback will always be executed in the same thread as the originating Vulkan call.

A callback may be called from multiple threads simultaneously (if the application is making Vulkan calls from multiple threads).

Valid Usage (Implicit)

- [](#VUID-VkDebugReportCallbackCreateInfoEXT-sType-sType)VUID-VkDebugReportCallbackCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEBUG_REPORT_CALLBACK_CREATE_INFO_EXT`
- [](#VUID-VkDebugReportCallbackCreateInfoEXT-flags-parameter)VUID-VkDebugReportCallbackCreateInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of [VkDebugReportFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagBitsEXT.html) values
- [](#VUID-VkDebugReportCallbackCreateInfoEXT-pfnCallback-parameter)VUID-VkDebugReportCallbackCreateInfoEXT-pfnCallback-parameter  
  `pfnCallback` **must** be a valid [PFN\_vkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkDebugReportCallbackEXT.html) value

## [](#_see_also)See Also

[PFN\_vkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkDebugReportCallbackEXT.html), [VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html), [VkDebugReportFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDebugReportCallbackEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDebugReportCallbackCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0