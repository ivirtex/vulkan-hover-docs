# VkAttachmentReferenceStencilLayout(3) Manual Page

## Name

VkAttachmentReferenceStencilLayout - Structure specifying an attachment
description



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAttachmentReferenceStencilLayout` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkAttachmentReferenceStencilLayout {
    VkStructureType    sType;
    void*              pNext;
    VkImageLayout      stencilLayout;
} VkAttachmentReferenceStencilLayout;
```

or the equivalent

``` c
// Provided by VK_KHR_separate_depth_stencil_layouts
typedef VkAttachmentReferenceStencilLayout VkAttachmentReferenceStencilLayoutKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stencilLayout` is a [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value
  specifying the layout the stencil aspect of the attachment uses during
  the subpass.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkAttachmentReferenceStencilLayout-stencilLayout-03318"
  id="VUID-VkAttachmentReferenceStencilLayout-stencilLayout-03318"></a>
  VUID-VkAttachmentReferenceStencilLayout-stencilLayout-03318  
  `stencilLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`,
  `VK_IMAGE_LAYOUT_PREINITIALIZED`,
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR`

Valid Usage (Implicit)

- <a href="#VUID-VkAttachmentReferenceStencilLayout-sType-sType"
  id="VUID-VkAttachmentReferenceStencilLayout-sType-sType"></a>
  VUID-VkAttachmentReferenceStencilLayout-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT`

- <a
  href="#VUID-VkAttachmentReferenceStencilLayout-stencilLayout-parameter"
  id="VUID-VkAttachmentReferenceStencilLayout-stencilLayout-parameter"></a>
  VUID-VkAttachmentReferenceStencilLayout-stencilLayout-parameter  
  `stencilLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_separate_depth_stencil_layouts](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_separate_depth_stencil_layouts.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAttachmentReferenceStencilLayout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
