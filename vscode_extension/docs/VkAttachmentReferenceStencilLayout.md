# VkAttachmentReferenceStencilLayout(3) Manual Page

## Name

VkAttachmentReferenceStencilLayout - Structure specifying an attachment description



## [](#_c_specification)C Specification

The `VkAttachmentReferenceStencilLayout` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkAttachmentReferenceStencilLayout {
    VkStructureType    sType;
    void*              pNext;
    VkImageLayout      stencilLayout;
} VkAttachmentReferenceStencilLayout;
```

or the equivalent

```c++
// Provided by VK_KHR_separate_depth_stencil_layouts
typedef VkAttachmentReferenceStencilLayout VkAttachmentReferenceStencilLayoutKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stencilLayout` is a [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value specifying the layout the stencil aspect of the attachment uses during the subpass.

## [](#_description)Description

Valid Usage

- [](#VUID-VkAttachmentReferenceStencilLayout-stencilLayout-03318)VUID-VkAttachmentReferenceStencilLayout-stencilLayout-03318  
  `stencilLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`, `VK_IMAGE_LAYOUT_PREINITIALIZED`, `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR`

Valid Usage (Implicit)

- [](#VUID-VkAttachmentReferenceStencilLayout-sType-sType)VUID-VkAttachmentReferenceStencilLayout-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT`
- [](#VUID-VkAttachmentReferenceStencilLayout-stencilLayout-parameter)VUID-VkAttachmentReferenceStencilLayout-stencilLayout-parameter  
  `stencilLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value

## [](#_see_also)See Also

[VK\_KHR\_separate\_depth\_stencil\_layouts](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_separate_depth_stencil_layouts.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAttachmentReferenceStencilLayout)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0