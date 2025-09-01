# VkFragmentShadingRateAttachmentInfoKHR(3) Manual Page

## Name

VkFragmentShadingRateAttachmentInfoKHR - Structure specifying a fragment shading rate attachment for a subpass



## [](#_c_specification)C Specification

The `VkFragmentShadingRateAttachmentInfoKHR` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_KHR_fragment_shading_rate
typedef struct VkFragmentShadingRateAttachmentInfoKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    const VkAttachmentReference2*    pFragmentShadingRateAttachment;
    VkExtent2D                       shadingRateAttachmentTexelSize;
} VkFragmentShadingRateAttachmentInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pFragmentShadingRateAttachment` is `NULL` or a pointer to a [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structure defining the fragment shading rate attachment for this subpass.
- `shadingRateAttachmentTexelSize` specifies the size of the portion of the framebuffer corresponding to each texel in `pFragmentShadingRateAttachment`.

## [](#_description)Description

If no shading rate attachment is specified, or if this structure is not specified, the implementation behaves as if a valid shading rate attachment was specified with all texels specifying a single pixel per fragment.

Valid Usage

- [](#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04524)VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04524  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment` member is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** be equal to `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_FRAGMENT_SHADING_RATE_ATTACHMENT_OPTIMAL_KHR`
- [](#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04525)VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04525  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment` member is not `VK_ATTACHMENT_UNUSED`, `shadingRateAttachmentTexelSize.width` **must** be a power of two value
- [](#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04526)VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04526  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment` member is not `VK_ATTACHMENT_UNUSED`, `shadingRateAttachmentTexelSize.width` **must** be less than or equal to [`maxFragmentShadingRateAttachmentTexelSize.width`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSize)
- [](#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04527)VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04527  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment` member is not `VK_ATTACHMENT_UNUSED`, `shadingRateAttachmentTexelSize.width` **must** be greater than or equal to [`minFragmentShadingRateAttachmentTexelSize.width`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-minFragmentShadingRateAttachmentTexelSize)
- [](#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04528)VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04528  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment` member is not `VK_ATTACHMENT_UNUSED`, `shadingRateAttachmentTexelSize.height` **must** be a power of two value
- [](#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04529)VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04529  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment` member is not `VK_ATTACHMENT_UNUSED`, `shadingRateAttachmentTexelSize.height` **must** be less than or equal to [`maxFragmentShadingRateAttachmentTexelSize.height`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSize)
- [](#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04530)VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04530  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment` member is not `VK_ATTACHMENT_UNUSED`, `shadingRateAttachmentTexelSize.height` **must** be greater than or equal to [`minFragmentShadingRateAttachmentTexelSize.height`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-minFragmentShadingRateAttachmentTexelSize)
- [](#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04531)VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04531  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment` member is not `VK_ATTACHMENT_UNUSED`, the quotient of `shadingRateAttachmentTexelSize.width` and `shadingRateAttachmentTexelSize.height` **must** be less than or equal to [`maxFragmentShadingRateAttachmentTexelSizeAspectRatio`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSizeAspectRatio)
- [](#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04532)VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-04532  
  If `pFragmentShadingRateAttachment` is not `NULL` and its `attachment` member is not `VK_ATTACHMENT_UNUSED`, the quotient of `shadingRateAttachmentTexelSize.height` and `shadingRateAttachmentTexelSize.width` **must** be less than or equal to [`maxFragmentShadingRateAttachmentTexelSizeAspectRatio`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentShadingRateAttachmentTexelSizeAspectRatio)

Valid Usage (Implicit)

- [](#VUID-VkFragmentShadingRateAttachmentInfoKHR-sType-sType)VUID-VkFragmentShadingRateAttachmentInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FRAGMENT_SHADING_RATE_ATTACHMENT_INFO_KHR`
- [](#VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-parameter)VUID-VkFragmentShadingRateAttachmentInfoKHR-pFragmentShadingRateAttachment-parameter  
  If `pFragmentShadingRateAttachment` is not `NULL`, `pFragmentShadingRateAttachment` **must** be a valid pointer to a valid [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structure

## [](#_see_also)See Also

[VK\_KHR\_fragment\_shading\_rate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shading_rate.html), [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFragmentShadingRateAttachmentInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0