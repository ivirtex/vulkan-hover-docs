# VkPhysicalDeviceMemoryBudgetPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceMemoryBudgetPropertiesEXT - Structure specifying physical device memory budget and usage



## [](#_c_specification)C Specification

If the `VkPhysicalDeviceMemoryBudgetPropertiesEXT` structure is included in the `pNext` chain of [VkPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties2.html), it is filled with the current memory budgets and usages.

The `VkPhysicalDeviceMemoryBudgetPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_memory_budget
typedef struct VkPhysicalDeviceMemoryBudgetPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       heapBudget[VK_MAX_MEMORY_HEAPS];
    VkDeviceSize       heapUsage[VK_MAX_MEMORY_HEAPS];
} VkPhysicalDeviceMemoryBudgetPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `heapBudget` is an array of `VK_MAX_MEMORY_HEAPS` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) values in which memory budgets are returned, with one element for each memory heap. A heap’s budget is a rough estimate of how much memory the process **can** allocate from that heap before allocations **may** fail or cause performance degradation. The budget includes any currently allocated device memory.
- `heapUsage` is an array of `VK_MAX_MEMORY_HEAPS` [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html) values in which memory usages are returned, with one element for each memory heap. A heap’s usage is an estimate of how much memory the process is currently using in that heap.

## [](#_description)Description

The values returned in this structure are not invariant. The `heapBudget` and `heapUsage` values **must** be zero for array elements greater than or equal to [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html)::`memoryHeapCount`. The `heapBudget` value **must** be non-zero for array elements less than [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html)::`memoryHeapCount`. The `heapBudget` value **must** be less than or equal to [VkMemoryHeap](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeap.html)::`size` for each heap.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceMemoryBudgetPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceMemoryBudgetPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_MEMORY_BUDGET_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_memory\_budget](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_memory_budget.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceMemoryBudgetPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0