# VkMemoryRequirements(3) Manual Page

## Name

VkMemoryRequirements - Structure specifying memory requirements



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryRequirements` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkMemoryRequirements {
    VkDeviceSize    size;
    VkDeviceSize    alignment;
    uint32_t        memoryTypeBits;
} VkMemoryRequirements;
```

## <a href="#_members" class="anchor"></a>Members

- `size` is the size, in bytes, of the memory allocation **required**
  for the resource.

- `alignment` is the alignment, in bytes, of the offset within the
  allocation **required** for the resource.

- `memoryTypeBits` is a bitmask and contains one bit set for every
  supported memory type for the resource. Bit `i` is set if and only if
  the memory type `i` in the `VkPhysicalDeviceMemoryProperties`
  structure for the physical device is supported for the resource.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html),
[VkVideoSessionMemoryRequirementsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionMemoryRequirementsKHR.html),
[vkGetBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements.html),
[vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryRequirements"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
