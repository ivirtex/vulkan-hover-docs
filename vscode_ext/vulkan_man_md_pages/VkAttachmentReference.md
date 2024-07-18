# VkAttachmentReference(3) Manual Page

## Name

VkAttachmentReference - Structure specifying an attachment reference



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAttachmentReference` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkAttachmentReference {
    uint32_t         attachment;
    VkImageLayout    layout;
} VkAttachmentReference;
```

## <a href="#_members" class="anchor"></a>Members

- `attachment` is either an integer value identifying an attachment at
  the corresponding index in
  [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html)::`pAttachments`,
  or `VK_ATTACHMENT_UNUSED` to signify that this attachment is not used.

- `layout` is a [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value specifying the
  layout the attachment uses during the subpass.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkAttachmentReference-layout-03077"
  id="VUID-VkAttachmentReference-layout-03077"></a>
  VUID-VkAttachmentReference-layout-03077  
  If `attachment` is not `VK_ATTACHMENT_UNUSED`, `layout` **must** not
  be `VK_IMAGE_LAYOUT_UNDEFINED`, `VK_IMAGE_LAYOUT_PREINITIALIZED`, or
  `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR`

- <a href="#VUID-VkAttachmentReference-separateDepthStencilLayouts-03313"
  id="VUID-VkAttachmentReference-separateDepthStencilLayouts-03313"></a>
  VUID-VkAttachmentReference-separateDepthStencilLayouts-03313  
  If the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is not enabled, and `attachment` is not
  `VK_ATTACHMENT_UNUSED`, `layout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,

- <a href="#VUID-VkAttachmentReference-synchronization2-06910"
  id="VUID-VkAttachmentReference-synchronization2-06910"></a>
  VUID-VkAttachmentReference-synchronization2-06910  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `layout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a href="#VUID-VkAttachmentReference-attachmentFeedbackLoopLayout-07311"
  id="VUID-VkAttachmentReference-attachmentFeedbackLoopLayout-07311"></a>
  VUID-VkAttachmentReference-attachmentFeedbackLoopLayout-07311  
  If the
  [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout)
  feature is not enabled, `layout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`

- <a href="#VUID-VkAttachmentReference-dynamicRenderingLocalRead-09546"
  id="VUID-VkAttachmentReference-dynamicRenderingLocalRead-09546"></a>
  VUID-VkAttachmentReference-dynamicRenderingLocalRead-09546  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `layout` **must** not be
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`

Valid Usage (Implicit)

- <a href="#VUID-VkAttachmentReference-layout-parameter"
  id="VUID-VkAttachmentReference-layout-parameter"></a>
  VUID-VkAttachmentReference-layout-parameter  
  `layout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html),
[VkSubpassDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAttachmentReference"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
