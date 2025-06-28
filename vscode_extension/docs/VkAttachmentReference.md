# VkAttachmentReference(3) Manual Page

## Name

VkAttachmentReference - Structure specifying an attachment reference



## [](#_c_specification)C Specification

The `VkAttachmentReference` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.2](#versions-1.2). See [Deprecated Functionality](#deprecation-renderpass2) for more information.

```c++
// Provided by VK_VERSION_1_0
typedef struct VkAttachmentReference {
    uint32_t         attachment;
    VkImageLayout    layout;
} VkAttachmentReference;
```

## [](#_members)Members

- `attachment` is either an integer value identifying an attachment at the corresponding index in [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html)::`pAttachments`, or `VK_ATTACHMENT_UNUSED` to signify that this attachment is not used.
- `layout` is a [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value specifying the layout the attachment uses during the subpass.

## [](#_description)Description

Valid Usage

- [](#VUID-VkAttachmentReference-layout-03077)VUID-VkAttachmentReference-layout-03077  
  If `attachment` is not `VK_ATTACHMENT_UNUSED`, `layout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`, `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`, `VK_IMAGE_LAYOUT_PREINITIALIZED`, or `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR`
- [](#VUID-VkAttachmentReference-separateDepthStencilLayouts-03313)VUID-VkAttachmentReference-separateDepthStencilLayouts-03313  
  If the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is not enabled, and `attachment` is not `VK_ATTACHMENT_UNUSED`, `layout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,
- [](#VUID-VkAttachmentReference-synchronization2-06910)VUID-VkAttachmentReference-synchronization2-06910  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `layout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkAttachmentReference-attachmentFeedbackLoopLayout-07311)VUID-VkAttachmentReference-attachmentFeedbackLoopLayout-07311  
  If the [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout) feature is not enabled, `layout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`
- [](#VUID-VkAttachmentReference-dynamicRenderingLocalRead-09546)VUID-VkAttachmentReference-dynamicRenderingLocalRead-09546  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `layout` **must** not be `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ`

Valid Usage (Implicit)

- [](#VUID-VkAttachmentReference-layout-parameter)VUID-VkAttachmentReference-layout-parameter  
  `layout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html), [VkSubpassDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAttachmentReference)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0