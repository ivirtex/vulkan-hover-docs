# VkMemoryHeap(3) Manual Page

## Name

VkMemoryHeap - Structure specifying a memory heap



## [](#_c_specification)C Specification

The `VkMemoryHeap` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkMemoryHeap {
    VkDeviceSize         size;
    VkMemoryHeapFlags    flags;
} VkMemoryHeap;
```

## [](#_members)Members

- `size` is the total memory size in bytes in the heap.
- `flags` is a bitmask of [VkMemoryHeapFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeapFlagBits.html) specifying attribute flags for the heap.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkMemoryHeapFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeapFlags.html), [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryHeap)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0