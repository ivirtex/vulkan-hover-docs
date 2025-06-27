# VkAttachmentDescriptionStencilLayout(3) Manual Page

## Name

VkAttachmentDescriptionStencilLayout - Structure specifying an
attachment description



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAttachmentDescriptionStencilLayout` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkAttachmentDescriptionStencilLayout {
    VkStructureType    sType;
    void*              pNext;
    VkImageLayout      stencilInitialLayout;
    VkImageLayout      stencilFinalLayout;
} VkAttachmentDescriptionStencilLayout;
```

or the equivalent

``` c
// Provided by VK_KHR_separate_depth_stencil_layouts
typedef VkAttachmentDescriptionStencilLayout VkAttachmentDescriptionStencilLayoutKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stencilInitialLayout` is the layout the stencil aspect of the
  attachment image subresource will be in when a render pass instance
  begins.

- `stencilFinalLayout` is the layout the stencil aspect of the
  attachment image subresource will be transitioned to when a render
  pass instance ends.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkAttachmentDescriptionStencilLayout-stencilInitialLayout-03308"
  id="VUID-VkAttachmentDescriptionStencilLayout-stencilInitialLayout-03308"></a>
  VUID-VkAttachmentDescriptionStencilLayout-stencilInitialLayout-03308  
  `stencilInitialLayout` **must** not be
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- <a
  href="#VUID-VkAttachmentDescriptionStencilLayout-stencilFinalLayout-03309"
  id="VUID-VkAttachmentDescriptionStencilLayout-stencilFinalLayout-03309"></a>
  VUID-VkAttachmentDescriptionStencilLayout-stencilFinalLayout-03309  
  `stencilFinalLayout` **must** not be
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- <a
  href="#VUID-VkAttachmentDescriptionStencilLayout-stencilFinalLayout-03310"
  id="VUID-VkAttachmentDescriptionStencilLayout-stencilFinalLayout-03310"></a>
  VUID-VkAttachmentDescriptionStencilLayout-stencilFinalLayout-03310  
  `stencilFinalLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED` or
  `VK_IMAGE_LAYOUT_PREINITIALIZED`

Valid Usage (Implicit)

- <a href="#VUID-VkAttachmentDescriptionStencilLayout-sType-sType"
  id="VUID-VkAttachmentDescriptionStencilLayout-sType-sType"></a>
  VUID-VkAttachmentDescriptionStencilLayout-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT`

- <a
  href="#VUID-VkAttachmentDescriptionStencilLayout-stencilInitialLayout-parameter"
  id="VUID-VkAttachmentDescriptionStencilLayout-stencilInitialLayout-parameter"></a>
  VUID-VkAttachmentDescriptionStencilLayout-stencilInitialLayout-parameter  
  `stencilInitialLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a
  href="#VUID-VkAttachmentDescriptionStencilLayout-stencilFinalLayout-parameter"
  id="VUID-VkAttachmentDescriptionStencilLayout-stencilFinalLayout-parameter"></a>
  VUID-VkAttachmentDescriptionStencilLayout-stencilFinalLayout-parameter  
  `stencilFinalLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_separate_depth_stencil_layouts](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_separate_depth_stencil_layouts.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAttachmentDescriptionStencilLayout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
