# VkAttachmentReference2(3) Manual Page

## Name

VkAttachmentReference2 - Structure specifying an attachment reference



## [](#_c_specification)C Specification

The `VkAttachmentReference2` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
typedef struct VkAttachmentReference2 {
    VkStructureType       sType;
    const void*           pNext;
    uint32_t              attachment;
    VkImageLayout         layout;
    VkImageAspectFlags    aspectMask;
} VkAttachmentReference2;
```

or the equivalent

```c++
// Provided by VK_KHR_create_renderpass2
typedef VkAttachmentReference2 VkAttachmentReference2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `attachment` is either an integer value identifying an attachment at the corresponding index in [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html)::`pAttachments`, or `VK_ATTACHMENT_UNUSED` to signify that this attachment is not used.
- `layout` is a [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value specifying the layout the attachment uses during the subpass.
- `aspectMask` is a mask of which aspect(s) **can** be accessed within the specified subpass as an input attachment.

## [](#_description)Description

Parameters defined by this structure with the same name as those in [VkAttachmentReference](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference.html) have the identical effect to those parameters.

`aspectMask` is ignored when this structure is used to describe anything other than an input attachment reference.

If the [`separateDepthStencilLayouts`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-separateDepthStencilLayouts) feature is enabled, and `attachment` has a depth/stencil format, `layout` **can** be set to a layout that only specifies the layout of the depth aspect.

If `layout` only specifies the layout of the depth aspect of the attachment, the layout of the stencil aspect is specified by the `stencilLayout` member of a [VkAttachmentReferenceStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReferenceStencilLayout.html) structure included in the `pNext` chain. Otherwise, `layout` describes the layout for all relevant image aspects.

Valid Usage

- [](#VUID-VkAttachmentReference2-layout-03077)VUID-VkAttachmentReference2-layout-03077  
  If `attachment` is not `VK_ATTACHMENT_UNUSED`, `layout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`, `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`, `VK_IMAGE_LAYOUT_PREINITIALIZED`, or `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR`
- [](#VUID-VkAttachmentReference2-separateDepthStencilLayouts-03313)VUID-VkAttachmentReference2-separateDepthStencilLayouts-03313  
  If the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is not enabled, and `attachment` is not `VK_ATTACHMENT_UNUSED`, `layout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,
- [](#VUID-VkAttachmentReference2-synchronization2-06910)VUID-VkAttachmentReference2-synchronization2-06910  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `layout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkAttachmentReference2-attachmentFeedbackLoopLayout-07311)VUID-VkAttachmentReference2-attachmentFeedbackLoopLayout-07311  
  If the [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout) feature is not enabled, `layout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`
- [](#VUID-VkAttachmentReference2-dynamicRenderingLocalRead-09546)VUID-VkAttachmentReference2-dynamicRenderingLocalRead-09546  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `layout` **must** not be `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ`

Valid Usage (Implicit)

- [](#VUID-VkAttachmentReference2-sType-sType)VUID-VkAttachmentReference2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2`
- [](#VUID-VkAttachmentReference2-pNext-pNext)VUID-VkAttachmentReference2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkAttachmentReferenceStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReferenceStencilLayout.html)
- [](#VUID-VkAttachmentReference2-sType-unique)VUID-VkAttachmentReference2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkAttachmentReference2-layout-parameter)VUID-VkAttachmentReference2-layout-parameter  
  `layout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value

## [](#_see_also)See Also

[VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateAttachmentInfoKHR.html), [VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlags.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html), [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionDepthStencilResolve.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAttachmentReference2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0