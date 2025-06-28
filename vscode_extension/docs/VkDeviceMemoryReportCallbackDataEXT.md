# VkDeviceMemoryReportCallbackDataEXT(3) Manual Page

## Name

VkDeviceMemoryReportCallbackDataEXT - Structure specifying parameters returned to the callback



## [](#_c_specification)C Specification

The definition of `VkDeviceMemoryReportCallbackDataEXT` is:

```c++
// Provided by VK_EXT_device_memory_report
typedef struct VkDeviceMemoryReportCallbackDataEXT {
    VkStructureType                     sType;
    void*                               pNext;
    VkDeviceMemoryReportFlagsEXT        flags;
    VkDeviceMemoryReportEventTypeEXT    type;
    uint64_t                            memoryObjectId;
    VkDeviceSize                        size;
    VkObjectType                        objectType;
    uint64_t                            objectHandle;
    uint32_t                            heapIndex;
} VkDeviceMemoryReportCallbackDataEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is 0 and reserved for future use.
- `type` is a [VkDeviceMemoryReportEventTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportEventTypeEXT.html) type specifying the type of event reported in this `VkDeviceMemoryReportCallbackDataEXT` structure.
- `memoryObjectId` is the unique id for the underlying memory object as described below.
- `size` is the size of the memory object in bytes. If `type` is `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT`, `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT` or `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT`, `size` is a valid [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) value. Otherwise, `size` is undefined.
- `objectType` is a [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html) value specifying the type of the object associated with this device memory report event. If `type` is `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT`, `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXT`, `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT`, `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXT` or `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT`, `objectType` is a valid [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html) enum. Otherwise, `objectType` is undefined.
- `objectHandle` is the object this device memory report event is attributed to. If `type` is `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT`, `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXT`, `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT` or `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXT`, `objectHandle` is a valid Vulkan handle of the type associated with `objectType` as defined in the [`VkObjectType` and Vulkan Handle Relationship](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#debugging-object-types) table. Otherwise, `objectHandle` is undefined.
- `heapIndex` describes which memory heap this device memory allocation is made from. If `type` is `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT` or `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT`, `heapIndex` corresponds to one of the valid heaps from the [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html) structure. Otherwise, `heapIndex` is undefined.

## [](#_description)Description

`memoryObjectId` is used to avoid double-counting on the same memory object.

If an internally-allocated device memory object or a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) **cannot** be exported, `memoryObjectId` **must** be unique in the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

If an internally-allocated device memory object or a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) supports being exported, `memoryObjectId` **must** be unique system wide.

If an internal device memory object or a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) is backed by an imported external memory object, `memoryObjectId` **must** be unique system wide.

Implementor’s Note

If the heap backing an internally-allocated device memory **cannot** be used to back [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), implementations **can** advertise that heap with no types.

Note

This structure should only be considered valid during the lifetime of the triggered callback.

For `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT` and `VK_DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT` events, `objectHandle` usually will not yet exist when the application or tool receives the callback. `objectHandle` will only exist when the create or allocate call that triggered the event returns, and if the allocation or import ends up failing `objectHandle` will not ever exist.

Valid Usage (Implicit)

- [](#VUID-VkDeviceMemoryReportCallbackDataEXT-sType-sType)VUID-VkDeviceMemoryReportCallbackDataEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_MEMORY_REPORT_CALLBACK_DATA_EXT`
- [](#VUID-VkDeviceMemoryReportCallbackDataEXT-pNext-pNext)VUID-VkDeviceMemoryReportCallbackDataEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[PFN\_vkDeviceMemoryReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkDeviceMemoryReportCallbackEXT.html), [VK\_EXT\_device\_memory\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_memory_report.html), [VkDeviceMemoryReportEventTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportEventTypeEXT.html), [VkDeviceMemoryReportFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemoryReportFlagsEXT.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceMemoryReportCallbackDataEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0