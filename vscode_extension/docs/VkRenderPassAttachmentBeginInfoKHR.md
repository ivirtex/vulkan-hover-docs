# VkRenderPassAttachmentBeginInfo(3) Manual Page

## Name

VkRenderPassAttachmentBeginInfo - Structure specifying images to be used as framebuffer attachments



## [](#_c_specification)C Specification

The `VkRenderPassAttachmentBeginInfo` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
typedef struct VkRenderPassAttachmentBeginInfo {
    VkStructureType       sType;
    const void*           pNext;
    uint32_t              attachmentCount;
    const VkImageView*    pAttachments;
} VkRenderPassAttachmentBeginInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_imageless_framebuffer
typedef VkRenderPassAttachmentBeginInfo VkRenderPassAttachmentBeginInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `attachmentCount` is the number of attachments.
- `pAttachments` is a pointer to an array of `VkImageView` handles, each of which will be used as the corresponding attachment in the render pass instance.

## [](#_description)Description

Valid Usage

- [](#VUID-VkRenderPassAttachmentBeginInfo-pAttachments-03218)VUID-VkRenderPassAttachmentBeginInfo-pAttachments-03218  
  Each element of `pAttachments` **must** only specify a single mip level
- [](#VUID-VkRenderPassAttachmentBeginInfo-pAttachments-03219)VUID-VkRenderPassAttachmentBeginInfo-pAttachments-03219  
  Each element of `pAttachments` **must** have been created with the [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings)
- [](#VUID-VkRenderPassAttachmentBeginInfo-pAttachments-04114)VUID-VkRenderPassAttachmentBeginInfo-pAttachments-04114  
  Each element of `pAttachments` **must** have been created with [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html)::`viewType` not equal to `VK_IMAGE_VIEW_TYPE_3D`
- [](#VUID-VkRenderPassAttachmentBeginInfo-pAttachments-07010)VUID-VkRenderPassAttachmentBeginInfo-pAttachments-07010  
  If [multisampled-render-to-single-sampled](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#subpass-multisampledrendertosinglesampled) is enabled for any subpass, all element of `pAttachments` which have a sample count equal to `VK_SAMPLE_COUNT_1_BIT` **must** have a format that supports the sample count specified in [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)::`rasterizationSamples`

Valid Usage (Implicit)

- [](#VUID-VkRenderPassAttachmentBeginInfo-sType-sType)VUID-VkRenderPassAttachmentBeginInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO`
- [](#VUID-VkRenderPassAttachmentBeginInfo-pAttachments-parameter)VUID-VkRenderPassAttachmentBeginInfo-pAttachments-parameter  
  If `attachmentCount` is not `0`, `pAttachments` **must** be a valid pointer to an array of `attachmentCount` valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handles

## [](#_see_also)See Also

[VK\_KHR\_imageless\_framebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_imageless_framebuffer.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassAttachmentBeginInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0