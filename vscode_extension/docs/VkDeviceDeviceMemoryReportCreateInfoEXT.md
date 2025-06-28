# VkDeviceDeviceMemoryReportCreateInfoEXT(3) Manual Page

## Name

VkDeviceDeviceMemoryReportCreateInfoEXT - Register device memory report callbacks for a Vulkan device



## [](#_c_specification)C Specification

To register callbacks for underlying device memory events of type [VkDeviceMemoryReportEventTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportEventTypeEXT.html), add one or multiple [VkDeviceDeviceMemoryReportCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceDeviceMemoryReportCreateInfoEXT.html) structures to the `pNext` chain of the [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) structure.

```c++
// Provided by VK_EXT_device_memory_report
typedef struct VkDeviceDeviceMemoryReportCreateInfoEXT {
    VkStructureType                        sType;
    const void*                            pNext;
    VkDeviceMemoryReportFlagsEXT           flags;
    PFN_vkDeviceMemoryReportCallbackEXT    pfnUserCallback;
    void*                                  pUserData;
} VkDeviceDeviceMemoryReportCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is 0 and reserved for future use.
- `pfnUserCallback` is the application callback function to call.
- `pUserData` is user data to be passed to the callback.

## [](#_description)Description

The callback **may** be called from multiple threads simultaneously.

The callback **must** be called only once by the implementation when a [VkDeviceMemoryReportEventTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportEventTypeEXT.html) event occurs.

Note

The callback could be called from a background thread other than the thread calling the Vulkan commands.

Valid Usage (Implicit)

- [](#VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-sType-sType)VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_DEVICE_MEMORY_REPORT_CREATE_INFO_EXT`
- [](#VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-flags-zerobitmask)VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-pfnUserCallback-parameter)VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-pfnUserCallback-parameter  
  `pfnUserCallback` **must** be a valid [PFN\_vkDeviceMemoryReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkDeviceMemoryReportCallbackEXT.html) value
- [](#VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-pUserData-parameter)VUID-VkDeviceDeviceMemoryReportCreateInfoEXT-pUserData-parameter  
  `pUserData` **must** be a pointer value

## [](#_see_also)See Also

[PFN\_vkDeviceMemoryReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkDeviceMemoryReportCallbackEXT.html), [VK\_EXT\_device\_memory\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_memory_report.html), [VkDeviceMemoryReportFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceDeviceMemoryReportCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0