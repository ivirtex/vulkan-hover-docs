# vkCreateDebugReportCallbackEXT(3) Manual Page

## Name

vkCreateDebugReportCallbackEXT - Create a debug report callback object



## [](#_c_specification)C Specification

Debug report callbacks give more detailed feedback on the application’s use of Vulkan when events of interest occur.

To register a debug report callback, an application uses [vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDebugReportCallbackEXT.html).

```c++
// Provided by VK_EXT_debug_report
VkResult vkCreateDebugReportCallbackEXT(
    VkInstance                                  instance,
    const VkDebugReportCallbackCreateInfoEXT*   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDebugReportCallbackEXT*                   pCallback);
```

## [](#_parameters)Parameters

- `instance` is the instance the callback will be logged on.
- `pCreateInfo` is a pointer to a [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html) structure defining the conditions under which this callback will be called.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pCallback` is a pointer to a [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackEXT.html) handle in which the created object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateDebugReportCallbackEXT-instance-parameter)VUID-vkCreateDebugReportCallbackEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkCreateDebugReportCallbackEXT-pCreateInfo-parameter)VUID-vkCreateDebugReportCallbackEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html) structure
- [](#VUID-vkCreateDebugReportCallbackEXT-pAllocator-parameter)VUID-vkCreateDebugReportCallbackEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateDebugReportCallbackEXT-pCallback-parameter)VUID-vkCreateDebugReportCallbackEXT-pCallback-parameter  
  `pCallback` **must** be a valid pointer to a [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackEXT.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDebugReportCallbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackCreateInfoEXT.html), [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackEXT.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateDebugReportCallbackEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0