# VkRenderingFragmentShadingRateAttachmentInfoKHR(3) Manual Page

## Name

VkRenderingFragmentShadingRateAttachmentInfoKHR - Structure specifying fragment shading rate attachment information



## [](#_c_specification)C Specification

The `VkRenderingFragmentShadingRateAttachmentInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_fragment_shading_rate with VK_VERSION_1_3 or VK_KHR_dynamic_rendering
typedef struct VkRenderingFragmentShadingRateAttachmentInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkImageView        imageView;
    VkImageLayout      imageLayout;
    VkExtent2D         shadingRateAttachmentTexelSize;
} VkRenderingFragmentShadingRateAttachmentInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `imageView` is the image view that will be used as a fragment shading rate attachment.
- `imageLayout` is the layout that `imageView` will be in during rendering.
- `shadingRateAttachmentTexelSize` specifies the number of pixels corresponding to each texel in `imageView`.

## [](#_description)Description

This structure can be included in the `pNext` chain of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) to define a [fragment shading rate attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment). If `imageView` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), or if this structure is not specified, the implementation behaves as if a valid shading rate attachment was specified with all texels specifying a single pixel per fragment.

Valid Usage

- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06147)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06147  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `layout` **must** be `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR`
- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06148)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06148  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), it **must** have been created with `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06149)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06149  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `shadingRateAttachmentTexelSize.width` **must** be a power of two value
- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06150)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06150  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `shadingRateAttachmentTexelSize.width` **must** be less than or equal to [`maxFragmentShadingRateAttachmentTexelSize.width`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSize)
- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06151)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06151  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `shadingRateAttachmentTexelSize.width` **must** be greater than or equal to [`minFragmentShadingRateAttachmentTexelSize.width`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-minFragmentShadingRateAttachmentTexelSize)
- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06152)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06152  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `shadingRateAttachmentTexelSize.height` **must** be a power of two value
- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06153)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06153  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `shadingRateAttachmentTexelSize.height` **must** be less than or equal to [`maxFragmentShadingRateAttachmentTexelSize.height`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSize)
- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06154)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06154  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `shadingRateAttachmentTexelSize.height` **must** be greater than or equal to [`minFragmentShadingRateAttachmentTexelSize.height`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-minFragmentShadingRateAttachmentTexelSize)
- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06155)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06155  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the quotient of `shadingRateAttachmentTexelSize.width` and `shadingRateAttachmentTexelSize.height` **must** be less than or equal to [`maxFragmentShadingRateAttachmentTexelSizeAspectRatio`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSizeAspectRatio)
- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06156)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-06156  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the quotient of `shadingRateAttachmentTexelSize.height` and `shadingRateAttachmentTexelSize.width` **must** be less than or equal to [`maxFragmentShadingRateAttachmentTexelSizeAspectRatio`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSizeAspectRatio)

Valid Usage (Implicit)

- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-sType-sType)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDERING_FRAGMENT_SHADING_RATE_ATTACHMENT_INFO_KHR`
- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-parameter)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageView-parameter  
  If `imageView` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handle
- [](#VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageLayout-parameter)VUID-VkRenderingFragmentShadingRateAttachmentInfoKHR-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value

## [](#_see_also)See Also

[VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html), [VK\_KHR\_fragment\_shading\_rate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shading_rate.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderingFragmentShadingRateAttachmentInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0