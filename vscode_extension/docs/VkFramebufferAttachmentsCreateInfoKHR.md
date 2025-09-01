# VkFramebufferAttachmentsCreateInfo(3) Manual Page

## Name

VkFramebufferAttachmentsCreateInfo - Structure specifying parameters of images that will be used with a framebuffer



## [](#_c_specification)C Specification

The `VkFramebufferAttachmentsCreateInfo` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
typedef struct VkFramebufferAttachmentsCreateInfo {
    VkStructureType                            sType;
    const void*                                pNext;
    uint32_t                                   attachmentImageInfoCount;
    const VkFramebufferAttachmentImageInfo*    pAttachmentImageInfos;
} VkFramebufferAttachmentsCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_imageless_framebuffer
typedef VkFramebufferAttachmentsCreateInfo VkFramebufferAttachmentsCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `attachmentImageInfoCount` is the number of attachments being described.
- `pAttachmentImageInfos` is a pointer to an array of [VkFramebufferAttachmentImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentImageInfo.html) structures, each structure describing a number of parameters of the corresponding attachment in a render pass instance.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkFramebufferAttachmentsCreateInfo-sType-sType)VUID-VkFramebufferAttachmentsCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO`
- [](#VUID-VkFramebufferAttachmentsCreateInfo-pAttachmentImageInfos-parameter)VUID-VkFramebufferAttachmentsCreateInfo-pAttachmentImageInfos-parameter  
  If `attachmentImageInfoCount` is not `0`, `pAttachmentImageInfos` **must** be a valid pointer to an array of `attachmentImageInfoCount` valid [VkFramebufferAttachmentImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentImageInfo.html) structures

## [](#_see_also)See Also

[VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkFramebufferAttachmentImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentImageInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFramebufferAttachmentsCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0