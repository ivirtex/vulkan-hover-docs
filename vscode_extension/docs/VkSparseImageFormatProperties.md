# VkSparseImageFormatProperties(3) Manual Page

## Name

VkSparseImageFormatProperties - Structure specifying sparse image format properties



## [](#_c_specification)C Specification

The `VkSparseImageFormatProperties` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkSparseImageFormatProperties {
    VkImageAspectFlags          aspectMask;
    VkExtent3D                  imageGranularity;
    VkSparseImageFormatFlags    flags;
} VkSparseImageFormatProperties;
```

## [](#_members)Members

- `aspectMask` is a bitmask [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) specifying which aspects of the image the properties apply to.
- `imageGranularity` is the width, height, and depth of the sparse image block in texels or compressed texel blocks.
- `flags` is a bitmask of [VkSparseImageFormatFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageFormatFlagBits.html) specifying additional information about the sparse resource.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlags.html), [VkSparseImageFormatFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageFormatFlags.html), [VkSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageFormatProperties2.html), [VkSparseImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSparseImageMemoryRequirements.html), [vkGetPhysicalDeviceSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSparseImageFormatProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSparseImageFormatProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0