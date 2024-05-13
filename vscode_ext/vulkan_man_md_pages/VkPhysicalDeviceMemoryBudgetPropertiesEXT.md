# VkPhysicalDeviceMemoryBudgetPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceMemoryBudgetPropertiesEXT - Structure specifying
physical device memory budget and usage



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `VkPhysicalDeviceMemoryBudgetPropertiesEXT` structure is included
in the `pNext` chain of
[VkPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties2.html),
it is filled with the current memory budgets and usages.

The `VkPhysicalDeviceMemoryBudgetPropertiesEXT` structure is defined as:

``` c
// Provided by VK_EXT_memory_budget
typedef struct VkPhysicalDeviceMemoryBudgetPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       heapBudget[VK_MAX_MEMORY_HEAPS];
    VkDeviceSize       heapUsage[VK_MAX_MEMORY_HEAPS];
} VkPhysicalDeviceMemoryBudgetPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `heapBudget` is an array of `VK_MAX_MEMORY_HEAPS`
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values in which memory budgets are
  returned, with one element for each memory heap. A heap’s budget is a
  rough estimate of how much memory the process **can** allocate from
  that heap before allocations **may** fail or cause performance
  degradation. The budget includes any currently allocated device
  memory.

- `heapUsage` is an array of `VK_MAX_MEMORY_HEAPS`
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values in which memory usages are
  returned, with one element for each memory heap. A heap’s usage is an
  estimate of how much memory the process is currently using in that
  heap.

## <a href="#_description" class="anchor"></a>Description

The values returned in this structure are not invariant. The
`heapBudget` and `heapUsage` values **must** be zero for array elements
greater than or equal to
[VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)::`memoryHeapCount`.
The `heapBudget` value **must** be non-zero for array elements less than
[VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)::`memoryHeapCount`.
The `heapBudget` value **must** be less than or equal to
[VkMemoryHeap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeap.html)::`size` for each heap.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceMemoryBudgetPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceMemoryBudgetPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceMemoryBudgetPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_BUDGET_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_memory_budget](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_memory_budget.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceMemoryBudgetPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
