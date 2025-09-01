# VkRenderingFragmentDensityMapAttachmentInfoEXT(3) Manual Page

## Name

VkRenderingFragmentDensityMapAttachmentInfoEXT - Structure specifying fragment shading rate attachment information



## [](#_c_specification)C Specification

The `VkRenderingFragmentDensityMapAttachmentInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_fragment_density_map with VK_VERSION_1_3 or VK_KHR_dynamic_rendering
typedef struct VkRenderingFragmentDensityMapAttachmentInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkImageView        imageView;
    VkImageLayout      imageLayout;
} VkRenderingFragmentDensityMapAttachmentInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `imageView` is the image view that will be used as a fragment density map attachment.
- `imageLayout` is the layout that `imageView` will be in during rendering.

## [](#_description)Description

This structure can be included in the `pNext` chain of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) to define a fragment density map. If this structure is not included in the `pNext` chain, `imageView` is treated as [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html).

Valid Usage

- [](#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06157)VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06157  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageLayout` **must** be `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT`
- [](#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06158)VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06158  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** have been created with `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`
- [](#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06159)VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06159  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** not have been created with `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`
- [](#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-apiVersion-07908)VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-apiVersion-07908  
  If the [`multiview`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiview) feature is not enabled, [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, and `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** have a `layerCount` equal to `1`

Valid Usage (Implicit)

- [](#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-sType-sType)VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_INFO_EXT`
- [](#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-parameter)VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-parameter  
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageLayout-parameter)VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value

## [](#_see_also)See Also

[VK\_EXT\_fragment\_density\_map](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map.html), [VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderingFragmentDensityMapAttachmentInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0