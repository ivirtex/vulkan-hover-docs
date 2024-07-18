# VkMemoryPriorityAllocateInfoEXT(3) Manual Page

## Name

VkMemoryPriorityAllocateInfoEXT - Specify a memory allocation priority



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain includes a `VkMemoryPriorityAllocateInfoEXT`
structure, then that structure includes a priority for the memory.

The `VkMemoryPriorityAllocateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_memory_priority
typedef struct VkMemoryPriorityAllocateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    float              priority;
} VkMemoryPriorityAllocateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `priority` is a floating-point value between `0` and `1`, indicating
  the priority of the allocation relative to other memory allocations.
  Larger values are higher priority. The granularity of the priorities
  is implementation-dependent.

## <a href="#_description" class="anchor"></a>Description

Memory allocations with higher priority **may** be more likely to stay
in device-local memory when the system is under memory pressure.

If this structure is not included, it is as if the `priority` value were
`0.5`.

Valid Usage

- <a href="#VUID-VkMemoryPriorityAllocateInfoEXT-priority-02602"
  id="VUID-VkMemoryPriorityAllocateInfoEXT-priority-02602"></a>
  VUID-VkMemoryPriorityAllocateInfoEXT-priority-02602  
  `priority` **must** be between `0` and `1`, inclusive

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryPriorityAllocateInfoEXT-sType-sType"
  id="VUID-VkMemoryPriorityAllocateInfoEXT-sType-sType"></a>
  VUID-VkMemoryPriorityAllocateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MEMORY_PRIORITY_ALLOCATE_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_memory_priority](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_memory_priority.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryPriorityAllocateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
