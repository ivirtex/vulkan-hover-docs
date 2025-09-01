# VkImageCopy(3) Manual Page

## Name

VkImageCopy - Structure specifying an image copy operation



## [](#_c_specification)C Specification

The `VkImageCopy` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkImageCopy {
    VkImageSubresourceLayers    srcSubresource;
    VkOffset3D                  srcOffset;
    VkImageSubresourceLayers    dstSubresource;
    VkOffset3D                  dstOffset;
    VkExtent3D                  extent;
} VkImageCopy;
```

## [](#_members)Members

- `srcSubresource` and `dstSubresource` are [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structures specifying the image subresources of the images used for the source and destination image data, respectively.
- `srcOffset` and `dstOffset` select the initial `x`, `y`, and `z` offsets in texels of the sub-regions of the source and destination image data.
- `extent` is the size in texels of the image to copy in `width`, `height` and `depth`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImageCopy-apiVersion-07940)VUID-VkImageCopy-apiVersion-07940  
  If the [VK\_KHR\_sampler\_ycbcr\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_sampler_ycbcr_conversion.html) extension is not enabled, and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, the `aspectMask` member of `srcSubresource` and `dstSubresource` **must** match
- [](#VUID-VkImageCopy-apiVersion-07941)VUID-VkImageCopy-apiVersion-07941  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled, and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, the `layerCount` member of `srcSubresource` and `dstSubresource` **must** match
- [](#VUID-VkImageCopy-extent-06668)VUID-VkImageCopy-extent-06668  
  `extent.width` **must** not be 0
- [](#VUID-VkImageCopy-extent-06669)VUID-VkImageCopy-extent-06669  
  `extent.height` **must** not be 0
- [](#VUID-VkImageCopy-extent-06670)VUID-VkImageCopy-extent-06670  
  `extent.depth` **must** not be 0

Valid Usage (Implicit)

- [](#VUID-VkImageCopy-srcSubresource-parameter)VUID-VkImageCopy-srcSubresource-parameter  
  `srcSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure
- [](#VUID-VkImageCopy-dstSubresource-parameter)VUID-VkImageCopy-dstSubresource-parameter  
  `dstSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html), [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html), [vkCmdCopyImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImage.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageCopy).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0