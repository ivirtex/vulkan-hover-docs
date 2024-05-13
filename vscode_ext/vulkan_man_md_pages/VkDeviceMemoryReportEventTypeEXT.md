# VkDeviceMemoryReportEventTypeEXT(3) Manual Page

## Name

VkDeviceMemoryReportEventTypeEXT - Events that can occur on a device
memory object



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[VkDeviceMemoryReportCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportCallbackDataEXT.html)::`type`,
specifying event types which cause the device driver to call the
callback, are:

``` c
// Provided by VK_EXT_device_memory_report
typedef enum VkDeviceMemoryReportEventTypeEXT {
    VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT = 0,
    VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXT = 1,
    VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT = 2,
    VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXT = 3,
    VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT = 4,
} VkDeviceMemoryReportEventTypeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT` specifies this event
  corresponds to the allocation of an internal device memory object or a
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html).

- `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXT` specifies this event
  corresponds to the deallocation of an internally-allocated device
  memory object or a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html).

- `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT` specifies this event
  corresponds to the import of an external memory object.

- `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXT` specifies this event
  is the release of an imported external memory object.

- `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT` specifies
  this event corresponds to the failed allocation of an internal device
  memory object or a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_memory_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_memory_report.html),
[VkDeviceMemoryReportCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportCallbackDataEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceMemoryReportEventTypeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
