# VkMemoryRequirements(3) Manual Page

## Name

VkMemoryRequirements - Structure specifying memory requirements



## [](#_c_specification)C Specification

The `VkMemoryRequirements` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkMemoryRequirements {
    VkDeviceSize    size;
    VkDeviceSize    alignment;
    uint32_t        memoryTypeBits;
} VkMemoryRequirements;
```

## [](#_members)Members

- `size` is the size, in bytes, of the memory allocation **required** for the resource.
- `alignment` is the alignment, in bytes, of the offset within the allocation **required** for the resource.
- `memoryTypeBits` is a bitmask and contains one bit set for every supported memory type for the resource. Bit `i` is set if and only if the memory type `i` in the `VkPhysicalDeviceMemoryProperties` structure for the physical device is supported for the resource.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html), [VkVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionMemoryRequirementsKHR.html), [vkGetBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferMemoryRequirements.html), [vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageMemoryRequirements.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryRequirements)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0