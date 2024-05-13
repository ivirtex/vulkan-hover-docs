# VkAttachmentReference2(3) Manual Page

## Name

VkAttachmentReference2 - Structure specifying an attachment reference



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAttachmentReference2` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_create_renderpass2
typedef VkAttachmentReference2 VkAttachmentReference2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `attachment` is either an integer value identifying an attachment at
  the corresponding index in
  [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html)::`pAttachments`,
  or `VK_ATTACHMENT_UNUSED` to signify that this attachment is not used.

- `layout` is a [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value specifying the
  layout the attachment uses during the subpass.

- `aspectMask` is a mask of which aspect(s) **can** be accessed within
  the specified subpass as an input attachment.

## <a href="#_description" class="anchor"></a>Description

Parameters defined by this structure with the same name as those in
[VkAttachmentReference](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference.html) have the identical
effect to those parameters.

`aspectMask` is ignored when this structure is used to describe anything
other than an input attachment reference.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-separateDepthStencilLayouts"
target="_blank"
rel="noopener"><code>separateDepthStencilLayouts</code></a> feature is
enabled, and `attachment` has a depth/stencil format, `layout` **can**
be set to a layout that only specifies the layout of the depth aspect.

If `layout` only specifies the layout of the depth aspect of the
attachment, the layout of the stencil aspect is specified by the
`stencilLayout` member of a
[VkAttachmentReferenceStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReferenceStencilLayout.html)
structure included in the `pNext` chain. Otherwise, `layout` describes
the layout for all relevant image aspects.

Valid Usage

- <a href="#VUID-VkAttachmentReference2-layout-03077"
  id="VUID-VkAttachmentReference2-layout-03077"></a>
  VUID-VkAttachmentReference2-layout-03077  
  If `attachment` is not `VK_ATTACHMENT_UNUSED`, `layout` **must** not
  be `VK_IMAGE_LAYOUT_UNDEFINED`, `VK_IMAGE_LAYOUT_PREINITIALIZED`, or
  `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR`

- <a href="#VUID-VkAttachmentReference2-separateDepthStencilLayouts-03313"
  id="VUID-VkAttachmentReference2-separateDepthStencilLayouts-03313"></a>
  VUID-VkAttachmentReference2-separateDepthStencilLayouts-03313  
  If the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is not enabled, and `attachment` is not
  `VK_ATTACHMENT_UNUSED`, `layout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,

- <a href="#VUID-VkAttachmentReference2-synchronization2-06910"
  id="VUID-VkAttachmentReference2-synchronization2-06910"></a>
  VUID-VkAttachmentReference2-synchronization2-06910  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `layout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a
  href="#VUID-VkAttachmentReference2-attachmentFeedbackLoopLayout-07311"
  id="VUID-VkAttachmentReference2-attachmentFeedbackLoopLayout-07311"></a>
  VUID-VkAttachmentReference2-attachmentFeedbackLoopLayout-07311  
  If the
  [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout)
  feature is not enabled, `layout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`

- <a href="#VUID-VkAttachmentReference2-dynamicRenderingLocalRead-09546"
  id="VUID-VkAttachmentReference2-dynamicRenderingLocalRead-09546"></a>
  VUID-VkAttachmentReference2-dynamicRenderingLocalRead-09546  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `layout` **must** not be
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`

Valid Usage (Implicit)

- <a href="#VUID-VkAttachmentReference2-sType-sType"
  id="VUID-VkAttachmentReference2-sType-sType"></a>
  VUID-VkAttachmentReference2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_2`

- <a href="#VUID-VkAttachmentReference2-pNext-pNext"
  id="VUID-VkAttachmentReference2-pNext-pNext"></a>
  VUID-VkAttachmentReference2-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkAttachmentReferenceStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReferenceStencilLayout.html)

- <a href="#VUID-VkAttachmentReference2-sType-unique"
  id="VUID-VkAttachmentReference2-sType-unique"></a>
  VUID-VkAttachmentReference2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkAttachmentReference2-layout-parameter"
  id="VUID-VkAttachmentReference2-layout-parameter"></a>
  VUID-VkAttachmentReference2-layout-parameter  
  `layout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateAttachmentInfoKHR.html),
[VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlags.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html),
[VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionDepthStencilResolve.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAttachmentReference2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
