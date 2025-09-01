# VkMemoryPriorityAllocateInfoEXT(3) Manual Page

## Name

VkMemoryPriorityAllocateInfoEXT - Specify a memory allocation priority



## [](#_c_specification)C Specification

If the `pNext` chain includes a `VkMemoryPriorityAllocateInfoEXT` structure, then that structure includes a priority for the memory.

The `VkMemoryPriorityAllocateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_memory_priority
typedef struct VkMemoryPriorityAllocateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    float              priority;
} VkMemoryPriorityAllocateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `priority` is a floating-point value between `0` and `1`, indicating the priority of the allocation relative to other memory allocations. Larger values are higher priority. The granularity of the priorities is implementation-dependent.

## [](#_description)Description

Memory allocations with higher priority **may** be more likely to stay in device-local memory when the system is under memory pressure.

If this structure is not included, it is as if the `priority` value were `0.5`.

Valid Usage

- [](#VUID-VkMemoryPriorityAllocateInfoEXT-priority-02602)VUID-VkMemoryPriorityAllocateInfoEXT-priority-02602  
  `priority` **must** be between `0` and `1`, inclusive

Valid Usage (Implicit)

- [](#VUID-VkMemoryPriorityAllocateInfoEXT-sType-sType)VUID-VkMemoryPriorityAllocateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_PRIORITY_ALLOCATE_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_memory\_priority](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_memory_priority.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryPriorityAllocateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0