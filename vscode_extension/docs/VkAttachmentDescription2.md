# VkAttachmentDescription2(3) Manual Page

## Name

VkAttachmentDescription2 - Structure specifying an attachment description



## [](#_c_specification)C Specification

The `VkAttachmentDescription2` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
typedef struct VkAttachmentDescription2 {
    VkStructureType                 sType;
    const void*                     pNext;
    VkAttachmentDescriptionFlags    flags;
    VkFormat                        format;
    VkSampleCountFlagBits           samples;
    VkAttachmentLoadOp              loadOp;
    VkAttachmentStoreOp             storeOp;
    VkAttachmentLoadOp              stencilLoadOp;
    VkAttachmentStoreOp             stencilStoreOp;
    VkImageLayout                   initialLayout;
    VkImageLayout                   finalLayout;
} VkAttachmentDescription2;
```

or the equivalent

```c++
// Provided by VK_KHR_create_renderpass2
typedef VkAttachmentDescription2 VkAttachmentDescription2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkAttachmentDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionFlagBits.html) specifying additional properties of the attachment.
- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value specifying the format of the image that will be used for the attachment.
- `samples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value specifying the number of samples of the image.
- `loadOp` is a [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html) value specifying how the contents of color and depth components of the attachment are treated at the beginning of the subpass where it is first used.
- `storeOp` is a [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) value specifying how the contents of color and depth components of the attachment are treated at the end of the subpass where it is last used.
- `stencilLoadOp` is a [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html) value specifying how the contents of stencil components of the attachment are treated at the beginning of the subpass where it is first used.
- `stencilStoreOp` is a [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) value specifying how the contents of stencil components of the attachment are treated at the end of the last subpass where it is used.
- `initialLayout` is the layout the attachment image subresource will be in when a render pass instance begins.
- `finalLayout` is the layout the attachment image subresource will be transitioned to when a render pass instance ends.

## [](#_description)Description

Parameters defined by this structure with the same name as those in [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription.html) have the identical effect to those parameters.

If the [`separateDepthStencilLayouts`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-separateDepthStencilLayouts) feature is enabled, and `format` is a depth/stencil format, `initialLayout` and `finalLayout` **can** be set to a layout that only specifies the layout of the depth aspect.

If the `pNext` chain includes a [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html) structure, then the `stencilInitialLayout` and `stencilFinalLayout` members specify the initial and final layouts of the stencil aspect of a depth/stencil format, and `initialLayout` and `finalLayout` only apply to the depth aspect. For depth-only formats, the [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html) structure is ignored. For stencil-only formats, the initial and final layouts of the stencil aspect are taken from the [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html) structure if present, or `initialLayout` and `finalLayout` if not present.

If `format` is a depth/stencil format, and either `initialLayout` or `finalLayout` does not specify a layout for the stencil aspect, then the application **must** specify the initial and final layouts of the stencil aspect by including a [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html) structure in the `pNext` chain.

`loadOp` and `storeOp` are ignored for fragment shading rate attachments. No access to the shading rate attachment is performed in `loadOp` and `storeOp`. Instead, access to `VK_ACCESS_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR` is performed as fragments are rasterized.

Valid Usage

- [](#VUID-VkAttachmentDescription2-format-06699)VUID-VkAttachmentDescription2-format-06699  
  If `format` includes a color or depth component and `loadOp` is `VK_ATTACHMENT_LOAD_OP_LOAD`, then `initialLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`
- [](#VUID-VkAttachmentDescription2-finalLayout-00843)VUID-VkAttachmentDescription2-finalLayout-00843  
  `finalLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED` or `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT` or `VK_IMAGE_LAYOUT_PREINITIALIZED`
- [](#VUID-VkAttachmentDescription2-format-03280)VUID-VkAttachmentDescription2-format-03280  
  If `format` is a color format, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-03281)VUID-VkAttachmentDescription2-format-03281  
  If `format` is a depth/stencil format, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-03282)VUID-VkAttachmentDescription2-format-03282  
  If `format` is a color format, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-03283)VUID-VkAttachmentDescription2-format-03283  
  If `format` is a depth/stencil format, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-06487)VUID-VkAttachmentDescription2-format-06487  
  If `format` is a color format, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-06488)VUID-VkAttachmentDescription2-format-06488  
  If `format` is a color format, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-separateDepthStencilLayouts-03284)VUID-VkAttachmentDescription2-separateDepthStencilLayouts-03284  
  If the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is not enabled, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,
- [](#VUID-VkAttachmentDescription2-separateDepthStencilLayouts-03285)VUID-VkAttachmentDescription2-separateDepthStencilLayouts-03285  
  If the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is not enabled, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,
- [](#VUID-VkAttachmentDescription2-format-03286)VUID-VkAttachmentDescription2-format-03286  
  If `format` is a color format, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-03287)VUID-VkAttachmentDescription2-format-03287  
  If `format` is a color format, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-06906)VUID-VkAttachmentDescription2-format-06906  
  If `format` is a depth/stencil format which includes both depth and stencil components, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-06907)VUID-VkAttachmentDescription2-format-06907  
  If `format` is a depth/stencil format which includes both depth and stencil components, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-03290)VUID-VkAttachmentDescription2-format-03290  
  If `format` is a depth/stencil format which includes only the depth component, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-03291)VUID-VkAttachmentDescription2-format-03291  
  If `format` is a depth/stencil format which includes only the depth component, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-synchronization2-06908)VUID-VkAttachmentDescription2-synchronization2-06908  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkAttachmentDescription2-synchronization2-06909)VUID-VkAttachmentDescription2-synchronization2-06909  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkAttachmentDescription2-attachmentFeedbackLoopLayout-07309)VUID-VkAttachmentDescription2-attachmentFeedbackLoopLayout-07309  
  If the [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout) feature is not enabled, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`
- [](#VUID-VkAttachmentDescription2-attachmentFeedbackLoopLayout-07310)VUID-VkAttachmentDescription2-attachmentFeedbackLoopLayout-07310  
  If the [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout) feature is not enabled, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`
- [](#VUID-VkAttachmentDescription2-samples-08745)VUID-VkAttachmentDescription2-samples-08745  
  `samples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value that is set in `imageCreateSampleCounts` (as defined in [Image Creation Limits](#resources-image-creation-limits)) for the given `format`
- [](#VUID-VkAttachmentDescription2-dynamicRenderingLocalRead-09544)VUID-VkAttachmentDescription2-dynamicRenderingLocalRead-09544  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ`
- [](#VUID-VkAttachmentDescription2-dynamicRenderingLocalRead-09545)VUID-VkAttachmentDescription2-dynamicRenderingLocalRead-09545  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ`
- [](#VUID-VkAttachmentDescription2-pNext-06704)VUID-VkAttachmentDescription2-pNext-06704  
  If the `pNext` chain does not include a [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html) structure, `format` includes a stencil component, and `stencilLoadOp` is `VK_ATTACHMENT_LOAD_OP_LOAD`, then `initialLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`
- [](#VUID-VkAttachmentDescription2-pNext-06705)VUID-VkAttachmentDescription2-pNext-06705  
  If the `pNext` chain includes a [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html) structure, `format` includes a stencil component, and `stencilLoadOp` is `VK_ATTACHMENT_LOAD_OP_LOAD`, then [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html)::`stencilInitialLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`
- [](#VUID-VkAttachmentDescription2-format-06249)VUID-VkAttachmentDescription2-format-06249  
  If `format` is a depth/stencil format which includes both depth and stencil components, and `initialLayout` is `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, the `pNext` chain **must** include a [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html) structure
- [](#VUID-VkAttachmentDescription2-format-06250)VUID-VkAttachmentDescription2-format-06250  
  If `format` is a depth/stencil format which includes both depth and stencil components, and `finalLayout` is `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, the `pNext` chain **must** include a [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html) structure
- [](#VUID-VkAttachmentDescription2-format-06247)VUID-VkAttachmentDescription2-format-06247  
  If the `pNext` chain does not include a [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html) structure and `format` only includes a stencil component, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-06248)VUID-VkAttachmentDescription2-format-06248  
  If the `pNext` chain does not include a [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html) structure and `format` only includes a stencil component, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription2-format-09332)VUID-VkAttachmentDescription2-format-09332  
  If the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is not enabled, `format` **must** not be `VK_FORMAT_UNDEFINED`
- [](#VUID-VkAttachmentDescription2-format-09334)VUID-VkAttachmentDescription2-format-09334  
  If `format` is `VK_FORMAT_UNDEFINED`, there **must** be a [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html) structure in the `pNext` chain with a `externalFormat` that is not equal to `0`

Valid Usage (Implicit)

- [](#VUID-VkAttachmentDescription2-sType-sType)VUID-VkAttachmentDescription2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2`
- [](#VUID-VkAttachmentDescription2-pNext-pNext)VUID-VkAttachmentDescription2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayout.html) or [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html)
- [](#VUID-VkAttachmentDescription2-sType-unique)VUID-VkAttachmentDescription2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkAttachmentDescription2-flags-parameter)VUID-VkAttachmentDescription2-flags-parameter  
  `flags` **must** be a valid combination of [VkAttachmentDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionFlagBits.html) values
- [](#VUID-VkAttachmentDescription2-format-parameter)VUID-VkAttachmentDescription2-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkAttachmentDescription2-samples-parameter)VUID-VkAttachmentDescription2-samples-parameter  
  `samples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value
- [](#VUID-VkAttachmentDescription2-loadOp-parameter)VUID-VkAttachmentDescription2-loadOp-parameter  
  `loadOp` **must** be a valid [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html) value
- [](#VUID-VkAttachmentDescription2-storeOp-parameter)VUID-VkAttachmentDescription2-storeOp-parameter  
  `storeOp` **must** be a valid [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) value
- [](#VUID-VkAttachmentDescription2-stencilLoadOp-parameter)VUID-VkAttachmentDescription2-stencilLoadOp-parameter  
  `stencilLoadOp` **must** be a valid [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html) value
- [](#VUID-VkAttachmentDescription2-stencilStoreOp-parameter)VUID-VkAttachmentDescription2-stencilStoreOp-parameter  
  `stencilStoreOp` **must** be a valid [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) value
- [](#VUID-VkAttachmentDescription2-initialLayout-parameter)VUID-VkAttachmentDescription2-initialLayout-parameter  
  `initialLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkAttachmentDescription2-finalLayout-parameter)VUID-VkAttachmentDescription2-finalLayout-parameter  
  `finalLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value

## [](#_see_also)See Also

[VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkAttachmentDescriptionFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionFlags.html), [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html), [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html), [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAttachmentDescription2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0