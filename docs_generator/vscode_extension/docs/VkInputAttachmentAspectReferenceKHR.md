# VkInputAttachmentAspectReference(3) Manual Page

## Name

VkInputAttachmentAspectReference - Structure specifying a subpass/input attachment pair and an aspect mask that can: be read.



## [](#_c_specification)C Specification

The `VkInputAttachmentAspectReference` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.2](#versions-1.2). See [Deprecated Functionality](#deprecation-renderpass2) for more information.

```c++
// Provided by VK_VERSION_1_1
typedef struct VkInputAttachmentAspectReference {
    uint32_t              subpass;
    uint32_t              inputAttachmentIndex;
    VkImageAspectFlags    aspectMask;
} VkInputAttachmentAspectReference;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance2
typedef VkInputAttachmentAspectReference VkInputAttachmentAspectReferenceKHR;
```

## [](#_members)Members

- `subpass` is an index into the `pSubpasses` array of the parent `VkRenderPassCreateInfo` structure.
- `inputAttachmentIndex` is an index into the `pInputAttachments` of the specified subpass.
- `aspectMask` is a mask of which aspect(s) **can** be accessed within the specified subpass.

## [](#_description)Description

This structure specifies an aspect mask for a specific input attachment of a specific subpass in the render pass.

`subpass` and `inputAttachmentIndex` index into the render pass as:

```c
pCreateInfo->pSubpasses[subpass].pInputAttachments[inputAttachmentIndex]
```

Valid Usage

- [](#VUID-VkInputAttachmentAspectReference-aspectMask-01964)VUID-VkInputAttachmentAspectReference-aspectMask-01964  
  `aspectMask` **must** not include `VK_IMAGE_ASPECT_METADATA_BIT`
- [](#VUID-VkInputAttachmentAspectReference-aspectMask-02250)VUID-VkInputAttachmentAspectReference-aspectMask-02250  
  `aspectMask` **must** not include `VK_IMAGE_ASPECT_MEMORY_PLANE_i_BIT_EXT` for any index *i*

Valid Usage (Implicit)

- [](#VUID-VkInputAttachmentAspectReference-aspectMask-parameter)VUID-VkInputAttachmentAspectReference-aspectMask-parameter  
  `aspectMask` **must** be a valid combination of [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) values
- [](#VUID-VkInputAttachmentAspectReference-aspectMask-requiredbitmask)VUID-VkInputAttachmentAspectReference-aspectMask-requiredbitmask  
  `aspectMask` **must** not be `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlags.html), [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkInputAttachmentAspectReference)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0