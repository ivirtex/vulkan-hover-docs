# VkFramebufferAttachmentImageInfo(3) Manual Page

## Name

VkFramebufferAttachmentImageInfo - Structure specifying parameters of an
image that will be used with a framebuffer



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFramebufferAttachmentImageInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkFramebufferAttachmentImageInfo {
    VkStructureType       sType;
    const void*           pNext;
    VkImageCreateFlags    flags;
    VkImageUsageFlags     usage;
    uint32_t              width;
    uint32_t              height;
    uint32_t              layerCount;
    uint32_t              viewFormatCount;
    const VkFormat*       pViewFormats;
} VkFramebufferAttachmentImageInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_imageless_framebuffer
typedef VkFramebufferAttachmentImageInfo VkFramebufferAttachmentImageInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html), matching the
  value of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags` used to
  create an image that will be used with this framebuffer.

- `usage` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html), matching the value
  of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` used to create
  an image used with this framebuffer.

- `width` is the width of the image view used for rendering.

- `height` is the height of the image view used for rendering.

- `layerCount` is the number of array layers of the image view used for
  rendering.

- `viewFormatCount` is the number of entries in the `pViewFormats`
  array, matching the value of
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount`
  used to create an image used with this framebuffer.

- `pViewFormats` is a pointer to an array of [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)
  values specifying all of the formats which **can** be used when
  creating views of the image, matching the value of
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)::`pViewFormats`
  used to create an image used with this framebuffer.

## <a href="#_description" class="anchor"></a>Description

Images that **can** be used with the framebuffer when beginning a render
pass, as specified by
[VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassAttachmentBeginInfo.html),
**must** be created with parameters that are identical to those
specified here.

Valid Usage

- <a href="#VUID-VkFramebufferAttachmentImageInfo-viewFormatCount-09536"
  id="VUID-VkFramebufferAttachmentImageInfo-viewFormatCount-09536"></a>
  VUID-VkFramebufferAttachmentImageInfo-viewFormatCount-09536  
  If `viewFormatCount` is not 0, and the render pass is not being used
  with an external format resolve attachment, each element of
  `pViewFormats` **must** not be `VK_FORMAT_UNDEFINED`

Valid Usage (Implicit)

- <a href="#VUID-VkFramebufferAttachmentImageInfo-sType-sType"
  id="VUID-VkFramebufferAttachmentImageInfo-sType-sType"></a>
  VUID-VkFramebufferAttachmentImageInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO`

- <a href="#VUID-VkFramebufferAttachmentImageInfo-pNext-pNext"
  id="VUID-VkFramebufferAttachmentImageInfo-pNext-pNext"></a>
  VUID-VkFramebufferAttachmentImageInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkFramebufferAttachmentImageInfo-flags-parameter"
  id="VUID-VkFramebufferAttachmentImageInfo-flags-parameter"></a>
  VUID-VkFramebufferAttachmentImageInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html) values

- <a href="#VUID-VkFramebufferAttachmentImageInfo-usage-parameter"
  id="VUID-VkFramebufferAttachmentImageInfo-usage-parameter"></a>
  VUID-VkFramebufferAttachmentImageInfo-usage-parameter  
  `usage` **must** be a valid combination of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) values

- <a href="#VUID-VkFramebufferAttachmentImageInfo-usage-requiredbitmask"
  id="VUID-VkFramebufferAttachmentImageInfo-usage-requiredbitmask"></a>
  VUID-VkFramebufferAttachmentImageInfo-usage-requiredbitmask  
  `usage` **must** not be `0`

- <a href="#VUID-VkFramebufferAttachmentImageInfo-pViewFormats-parameter"
  id="VUID-VkFramebufferAttachmentImageInfo-pViewFormats-parameter"></a>
  VUID-VkFramebufferAttachmentImageInfo-pViewFormats-parameter  
  If `viewFormatCount` is not `0`, `pViewFormats` **must** be a valid
  pointer to an array of `viewFormatCount` valid
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_imageless_framebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_imageless_framebuffer.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html),
[VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlags.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFramebufferAttachmentImageInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
