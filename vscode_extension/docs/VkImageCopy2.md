# VkImageCopy2(3) Manual Page

## Name

VkImageCopy2 - Structure specifying an image copy operation



## [](#_c_specification)C Specification

The `VkImageCopy2` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkImageCopy2 {
    VkStructureType             sType;
    const void*                 pNext;
    VkImageSubresourceLayers    srcSubresource;
    VkOffset3D                  srcOffset;
    VkImageSubresourceLayers    dstSubresource;
    VkOffset3D                  dstOffset;
    VkExtent3D                  extent;
} VkImageCopy2;
```

or the equivalent

```c++
// Provided by VK_KHR_copy_commands2
typedef VkImageCopy2 VkImageCopy2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcSubresource` and `dstSubresource` are [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structures specifying the image subresources of the images used for the source and destination image data, respectively.
- `srcOffset` and `dstOffset` select the initial `x`, `y`, and `z` offsets in texels of the sub-regions of the source and destination image data.
- `extent` is the size in texels of the image to copy in `width`, `height` and `depth`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkImageCopy2-apiVersion-07940)VUID-VkImageCopy2-apiVersion-07940  
  If the [VK\_KHR\_sampler\_ycbcr\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_sampler_ycbcr_conversion.html) extension is not enabled, and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, the `aspectMask` member of `srcSubresource` and `dstSubresource` **must** match
- [](#VUID-VkImageCopy2-apiVersion-07941)VUID-VkImageCopy2-apiVersion-07941  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled, and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, the `layerCount` member of `srcSubresource` and `dstSubresource` **must** match
- [](#VUID-VkImageCopy2-extent-06668)VUID-VkImageCopy2-extent-06668  
  `extent.width` **must** not be 0
- [](#VUID-VkImageCopy2-extent-06669)VUID-VkImageCopy2-extent-06669  
  `extent.height` **must** not be 0
- [](#VUID-VkImageCopy2-extent-06670)VUID-VkImageCopy2-extent-06670  
  `extent.depth` **must** not be 0

Valid Usage (Implicit)

- [](#VUID-VkImageCopy2-sType-sType)VUID-VkImageCopy2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMAGE_COPY_2`
- [](#VUID-VkImageCopy2-pNext-pNext)VUID-VkImageCopy2-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImageCopy2-srcSubresource-parameter)VUID-VkImageCopy2-srcSubresource-parameter  
  `srcSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure
- [](#VUID-VkImageCopy2-dstSubresource-parameter)VUID-VkImageCopy2-dstSubresource-parameter  
  `dstSubresource` **must** be a valid [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html) structure

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkCopyImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageInfo2.html), [VkCopyImageToImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyImageToImageInfo.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent3D.html), [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceLayers.html), [VkOffset3D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset3D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageCopy2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0