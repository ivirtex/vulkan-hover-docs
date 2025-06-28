# VkFramebufferAttachmentImageInfo(3) Manual Page

## Name

VkFramebufferAttachmentImageInfo - Structure specifying parameters of an image that will be used with a framebuffer



## [](#_c_specification)C Specification

The `VkFramebufferAttachmentImageInfo` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
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

```c++
// Provided by VK_KHR_imageless_framebuffer
typedef VkFramebufferAttachmentImageInfo VkFramebufferAttachmentImageInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html), matching the value of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags` used to create an image that will be used with this framebuffer.
- `usage` is a bitmask of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html), matching the value of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` used to create an image used with this framebuffer.
- `width` is the width of the image view used for rendering.
- `height` is the height of the image view used for rendering.
- `layerCount` is the number of array layers of the image view used for rendering.
- `viewFormatCount` is the number of entries in the `pViewFormats` array, matching the value of [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount` used to create an image used with this framebuffer.
- `pViewFormats` is a pointer to an array of [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values specifying all of the formats which **can** be used when creating views of the image, matching the value of [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html)::`pViewFormats` used to create an image used with this framebuffer.

## [](#_description)Description

Images that **can** be used with the framebuffer when beginning a render pass, as specified by [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html), **must** be created with parameters that are identical to those specified here.

Valid Usage

- [](#VUID-VkFramebufferAttachmentImageInfo-viewFormatCount-09536)VUID-VkFramebufferAttachmentImageInfo-viewFormatCount-09536  
  If `viewFormatCount` is not 0, and the render pass is not being used with an external format resolve attachment, each element of `pViewFormats` **must** not be `VK_FORMAT_UNDEFINED`

Valid Usage (Implicit)

- [](#VUID-VkFramebufferAttachmentImageInfo-sType-sType)VUID-VkFramebufferAttachmentImageInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO`
- [](#VUID-VkFramebufferAttachmentImageInfo-pNext-pNext)VUID-VkFramebufferAttachmentImageInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkFramebufferAttachmentImageInfo-flags-parameter)VUID-VkFramebufferAttachmentImageInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html) values
- [](#VUID-VkFramebufferAttachmentImageInfo-usage-parameter)VUID-VkFramebufferAttachmentImageInfo-usage-parameter  
  `usage` **must** be a valid combination of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) values
- [](#VUID-VkFramebufferAttachmentImageInfo-usage-requiredbitmask)VUID-VkFramebufferAttachmentImageInfo-usage-requiredbitmask  
  `usage` **must** not be `0`
- [](#VUID-VkFramebufferAttachmentImageInfo-pViewFormats-parameter)VUID-VkFramebufferAttachmentImageInfo-pViewFormats-parameter  
  If `viewFormatCount` is not `0`, `pViewFormats` **must** be a valid pointer to an array of `viewFormatCount` valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values

## [](#_see_also)See Also

[VK\_KHR\_imageless\_framebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_imageless_framebuffer.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html), [VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlags.html), [VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFramebufferAttachmentImageInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0