# VkRenderPassAttachmentBeginInfo(3) Manual Page

## Name

VkRenderPassAttachmentBeginInfo - Structure specifying images to be used
as framebuffer attachments



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderPassAttachmentBeginInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkRenderPassAttachmentBeginInfo {
    VkStructureType       sType;
    const void*           pNext;
    uint32_t              attachmentCount;
    const VkImageView*    pAttachments;
} VkRenderPassAttachmentBeginInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_imageless_framebuffer
typedef VkRenderPassAttachmentBeginInfo VkRenderPassAttachmentBeginInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `attachmentCount` is the number of attachments.

- `pAttachments` is a pointer to an array of `VkImageView` handles, each
  of which will be used as the corresponding attachment in the render
  pass instance.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkRenderPassAttachmentBeginInfo-pAttachments-03218"
  id="VUID-VkRenderPassAttachmentBeginInfo-pAttachments-03218"></a>
  VUID-VkRenderPassAttachmentBeginInfo-pAttachments-03218  
  Each element of `pAttachments` **must** only specify a single mip
  level

- <a href="#VUID-VkRenderPassAttachmentBeginInfo-pAttachments-03219"
  id="VUID-VkRenderPassAttachmentBeginInfo-pAttachments-03219"></a>
  VUID-VkRenderPassAttachmentBeginInfo-pAttachments-03219  
  Each element of `pAttachments` **must** have been created with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkRenderPassAttachmentBeginInfo-pAttachments-04114"
  id="VUID-VkRenderPassAttachmentBeginInfo-pAttachments-04114"></a>
  VUID-VkRenderPassAttachmentBeginInfo-pAttachments-04114  
  Each element of `pAttachments` **must** have been created with
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)::`viewType` not
  equal to `VK_IMAGE_VIEW_TYPE_3D`

- <a href="#VUID-VkRenderPassAttachmentBeginInfo-pAttachments-07010"
  id="VUID-VkRenderPassAttachmentBeginInfo-pAttachments-07010"></a>
  VUID-VkRenderPassAttachmentBeginInfo-pAttachments-07010  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#subpass-multisampledrendertosinglesampled"
  target="_blank" rel="noopener">multisampled-render-to-single-sampled</a>
  is enabled for any subpass, all element of `pAttachments` which have a
  sample count equal to `VK_SAMPLE_COUNT_1_BIT` **must** have a format
  that supports the sample count specified in
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)::`rasterizationSamples`

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassAttachmentBeginInfo-sType-sType"
  id="VUID-VkRenderPassAttachmentBeginInfo-sType-sType"></a>
  VUID-VkRenderPassAttachmentBeginInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO`

- <a href="#VUID-VkRenderPassAttachmentBeginInfo-pAttachments-parameter"
  id="VUID-VkRenderPassAttachmentBeginInfo-pAttachments-parameter"></a>
  VUID-VkRenderPassAttachmentBeginInfo-pAttachments-parameter  
  If `attachmentCount` is not `0`, `pAttachments` **must** be a valid
  pointer to an array of `attachmentCount` valid
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handles

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_imageless_framebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_imageless_framebuffer.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassAttachmentBeginInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
