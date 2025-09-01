# PFN\_vkDeviceMemoryReportCallbackEXT(3) Manual Page

## Name

PFN\_vkDeviceMemoryReportCallbackEXT - Application-defined device memory report callback function



## [](#_c_specification)C Specification

The prototype for the [VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html)::`pfnUserCallback` function implemented by the application is:

```c++
// Provided by VK_EXT_device_memory_report
typedef void (VKAPI_PTR *PFN_vkDeviceMemoryReportCallbackEXT)(
    const VkDeviceMemoryReportCallbackDataEXT*  pCallbackData,
    void*                                       pUserData);
```

## [](#_parameters)Parameters

- `pCallbackData` contains all the callback related data in the [VkDeviceMemoryReportCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportCallbackDataEXT.html) structure.
- `pUserData` is the user data provided when the [VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html) was created.

## [](#_description)Description

The callback **must** not make calls to any Vulkan commands.

## [](#_see_also)See Also

[VK\_EXT\_device\_memory\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_memory_report.html), [VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html), [VkDeviceMemoryReportCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportCallbackDataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#PFN_vkDeviceMemoryReportCallbackEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0