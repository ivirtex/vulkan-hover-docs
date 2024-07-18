# VkFramebufferAttachmentsCreateInfo(3) Manual Page

## Name

VkFramebufferAttachmentsCreateInfo - Structure specifying parameters of
images that will be used with a framebuffer



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFramebufferAttachmentsCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkFramebufferAttachmentsCreateInfo {
    VkStructureType                            sType;
    const void*                                pNext;
    uint32_t                                   attachmentImageInfoCount;
    const VkFramebufferAttachmentImageInfo*    pAttachmentImageInfos;
} VkFramebufferAttachmentsCreateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_imageless_framebuffer
typedef VkFramebufferAttachmentsCreateInfo VkFramebufferAttachmentsCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `attachmentImageInfoCount` is the number of attachments being
  described.

- `pAttachmentImageInfos` is a pointer to an array of
  [VkFramebufferAttachmentImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentImageInfo.html)
  structures, each structure describing a number of parameters of the
  corresponding attachment in a render pass instance.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkFramebufferAttachmentsCreateInfo-sType-sType"
  id="VUID-VkFramebufferAttachmentsCreateInfo-sType-sType"></a>
  VUID-VkFramebufferAttachmentsCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO`

- <a
  href="#VUID-VkFramebufferAttachmentsCreateInfo-pAttachmentImageInfos-parameter"
  id="VUID-VkFramebufferAttachmentsCreateInfo-pAttachmentImageInfos-parameter"></a>
  VUID-VkFramebufferAttachmentsCreateInfo-pAttachmentImageInfos-parameter  
  If `attachmentImageInfoCount` is not `0`, `pAttachmentImageInfos`
  **must** be a valid pointer to an array of `attachmentImageInfoCount`
  valid
  [VkFramebufferAttachmentImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentImageInfo.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkFramebufferAttachmentImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentImageInfo.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFramebufferAttachmentsCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
