# VkDeviceMemoryReportEventTypeEXT(3) Manual Page

## Name

VkDeviceMemoryReportEventTypeEXT - Events that can occur on a device memory object



## [](#_c_specification)C Specification

Possible values of [VkDeviceMemoryReportCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportCallbackDataEXT.html)::`type`, specifying event types which cause the device driver to call the callback, are:

```c++
// Provided by VK_EXT_device_memory_report
typedef enum VkDeviceMemoryReportEventTypeEXT {
    VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT = 0,
    VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXT = 1,
    VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT = 2,
    VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXT = 3,
    VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT = 4,
} VkDeviceMemoryReportEventTypeEXT;
```

## [](#_description)Description

- `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT` specifies this event corresponds to the allocation of an internal device memory object or a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html).
- `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXT` specifies this event corresponds to the deallocation of an internally-allocated device memory object or a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html).
- `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT` specifies this event corresponds to the import of an external memory object.
- `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXT` specifies this event is the release of an imported external memory object.
- `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT` specifies this event corresponds to the failed allocation of an internal device memory object or a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html).

## [](#_see_also)See Also

[VK\_EXT\_device\_memory\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_memory_report.html), [VkDeviceMemoryReportCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportCallbackDataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceMemoryReportEventTypeEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0