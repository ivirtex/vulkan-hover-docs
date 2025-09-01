# VkMemoryType(3) Manual Page

## Name

VkMemoryType - Structure specifying memory type



## [](#_c_specification)C Specification

The `VkMemoryType` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkMemoryType {
    VkMemoryPropertyFlags    propertyFlags;
    uint32_t                 heapIndex;
} VkMemoryType;
```

## [](#_members)Members

- `heapIndex` describes which memory heap this memory type corresponds to, and **must** be less than `memoryHeapCount` from the [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html) structure.
- `propertyFlags` is a bitmask of [VkMemoryPropertyFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryPropertyFlagBits.html) of properties for this memory type.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkMemoryPropertyFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryPropertyFlags.html), [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryType).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0