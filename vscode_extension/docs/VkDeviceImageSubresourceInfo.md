# VkDeviceImageSubresourceInfo(3) Manual Page

## Name

VkDeviceImageSubresourceInfo - Image creation information for querying subresource layout



## [](#_c_specification)C Specification

The `VkDeviceImageSubresourceInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkDeviceImageSubresourceInfo {
    VkStructureType               sType;
    const void*                   pNext;
    const VkImageCreateInfo*      pCreateInfo;
    const VkImageSubresource2*    pSubresource;
} VkDeviceImageSubresourceInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance5
typedef VkDeviceImageSubresourceInfo VkDeviceImageSubresourceInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pCreateInfo` is a pointer to a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure containing parameters affecting creation of the image to query.
- `pSubresource` pSubresource is a pointer to a [VkImageSubresource2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2.html) structure selecting a specific image subresource for the query.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDeviceImageSubresourceInfo-aspectMask-00997)VUID-VkDeviceImageSubresourceInfo-aspectMask-00997  
  The `aspectMask` member of `pSubresource` **must** only have a single bit set
- [](#VUID-VkDeviceImageSubresourceInfo-mipLevel-01716)VUID-VkDeviceImageSubresourceInfo-mipLevel-01716  
  The `mipLevel` member of `pSubresource` **must** be less than the `mipLevels` specified in `pCreateInfo`
- [](#VUID-VkDeviceImageSubresourceInfo-arrayLayer-01717)VUID-VkDeviceImageSubresourceInfo-arrayLayer-01717  
  The `arrayLayer` member of `pSubresource` **must** be less than the `arrayLayers` specified in `pCreateInfo`
- [](#VUID-VkDeviceImageSubresourceInfo-format-08886)VUID-VkDeviceImageSubresourceInfo-format-08886  
  If `format` of the `image` is a color format that is not a [multi-planar format](#formats-multiplanar), and `tiling` of the `pCreateInfo` is `VK_IMAGE_TILING_LINEAR` or `VK_IMAGE_TILING_OPTIMAL`, the `aspectMask` member of `pSubresource` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkDeviceImageSubresourceInfo-format-04462)VUID-VkDeviceImageSubresourceInfo-format-04462  
  If `format` of the `pCreateInfo` has a depth component, the `aspectMask` member of `pSubresource` **must** contain `VK_IMAGE_ASPECT_DEPTH_BIT`
- [](#VUID-VkDeviceImageSubresourceInfo-format-04463)VUID-VkDeviceImageSubresourceInfo-format-04463  
  If `format` of the `pCreateInfo` has a stencil component, the `aspectMask` member of `pSubresource` **must** contain `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkDeviceImageSubresourceInfo-format-04464)VUID-VkDeviceImageSubresourceInfo-format-04464  
  If `format` of the `pCreateInfo` does not contain a stencil or depth component, the `aspectMask` member of `pSubresource` **must** not contain `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkDeviceImageSubresourceInfo-tiling-08717)VUID-VkDeviceImageSubresourceInfo-tiling-08717  
  If the `tiling` of the `pCreateInfo` is `VK_IMAGE_TILING_LINEAR` and has a [multi-planar format](#formats-multiplanar), then the `aspectMask` member of `pSubresource` **must** be a single valid [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit

Valid Usage (Implicit)

- [](#VUID-VkDeviceImageSubresourceInfo-sType-sType)VUID-VkDeviceImageSubresourceInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_IMAGE_SUBRESOURCE_INFO`
- [](#VUID-VkDeviceImageSubresourceInfo-pNext-pNext)VUID-VkDeviceImageSubresourceInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkDeviceImageSubresourceInfo-pCreateInfo-parameter)VUID-VkDeviceImageSubresourceInfo-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure
- [](#VUID-VkDeviceImageSubresourceInfo-pSubresource-parameter)VUID-VkDeviceImageSubresourceInfo-pSubresource-parameter  
  `pSubresource` **must** be a valid pointer to a valid [VkImageSubresource2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2.html) structure

## [](#_see_also)See Also

[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkImageSubresource2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresource2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDeviceImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSubresourceLayout.html), [vkGetDeviceImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceImageSubresourceLayout.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceImageSubresourceInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0