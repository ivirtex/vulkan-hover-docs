# VkAttachmentDescription(3) Manual Page

## Name

VkAttachmentDescription - Structure specifying an attachment description



## [](#_c_specification)C Specification

The `VkAttachmentDescription` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.2](#versions-1.2). See [Deprecated Functionality](#deprecation-renderpass2) for more information.

```c++
// Provided by VK_VERSION_1_0
typedef struct VkAttachmentDescription {
    VkAttachmentDescriptionFlags    flags;
    VkFormat                        format;
    VkSampleCountFlagBits           samples;
    VkAttachmentLoadOp              loadOp;
    VkAttachmentStoreOp             storeOp;
    VkAttachmentLoadOp              stencilLoadOp;
    VkAttachmentStoreOp             stencilStoreOp;
    VkImageLayout                   initialLayout;
    VkImageLayout                   finalLayout;
} VkAttachmentDescription;
```

## [](#_members)Members

- `flags` is a bitmask of [VkAttachmentDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionFlagBits.html) specifying additional properties of the attachment.
- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value specifying the format of the image view that will be used for the attachment.
- `samples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value specifying the number of samples of the image.
- `loadOp` is a [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html) value specifying how the contents of color and depth components of the attachment are treated at the beginning of the subpass where it is first used.
- `storeOp` is a [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) value specifying how the contents of color and depth components of the attachment are treated at the end of the subpass where it is last used.
- `stencilLoadOp` is a [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html) value specifying how the contents of stencil components of the attachment are treated at the beginning of the subpass where it is first used.
- `stencilStoreOp` is a [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) value specifying how the contents of stencil components of the attachment are treated at the end of the last subpass where it is used.
- `initialLayout` is the layout the attachment image subresource will be in when a render pass instance begins.
- `finalLayout` is the layout the attachment image subresource will be transitioned to when a render pass instance ends.

## [](#_description)Description

If the attachment uses a color format, then `loadOp` and `storeOp` are used, and `stencilLoadOp` and `stencilStoreOp` are ignored. If the format has depth and/or stencil components, `loadOp` and `storeOp` apply only to the depth data, while `stencilLoadOp` and `stencilStoreOp` define how the stencil data is handled. `loadOp` and `stencilLoadOp` define the [load operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-load-operations) for the attachment. `storeOp` and `stencilStoreOp` define the [store operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-store-operations) for the attachment. If an attachment is not used by any subpass, `loadOp`, `storeOp`, `stencilStoreOp`, and `stencilLoadOp` will be ignored for that attachment, and no load or store ops will be performed. However, any transition specified by `initialLayout` and `finalLayout` will still be executed.

If `flags` includes `VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT`, then the attachment is treated as if it shares physical memory with another attachment in the same render pass. This information limits the ability of the implementation to reorder certain operations (like layout transitions and the `loadOp`) such that it is not improperly reordered against other uses of the same physical memory via a different attachment. This is described in more detail below.

If a render pass uses multiple attachments that alias the same device memory, those attachments **must** each include the `VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT` bit in their attachment description flags. Attachments aliasing the same memory occurs in multiple ways:

- Multiple attachments being assigned the same image view as part of framebuffer creation.
- Attachments using distinct image views that correspond to the same image subresource of an image.
- Attachments using views of distinct image subresources which are bound to overlapping memory ranges.

Note

Render passes **must** include subpass dependencies (either directly or via a subpass dependency chain) between any two subpasses that operate on the same attachment or aliasing attachments and those subpass dependencies **must** include execution and memory dependencies separating uses of the aliases, if at least one of those subpasses writes to one of the aliases. These dependencies **must** not include the `VK_DEPENDENCY_BY_REGION_BIT` if the aliases are views of distinct image subresources which overlap in memory.

Multiple attachments that alias the same memory **must** not be used in a single subpass. A given attachment index **must** not be used multiple times in a single subpass, with one exception: two subpass attachments **can** use the same attachment index if at least one use is as an input attachment and neither use is as a resolve or preserve attachment. In other words, the same view **can** be used simultaneously as an input and color or depth/stencil attachment, but **must** not be used as multiple color or depth/stencil attachments nor as resolve or preserve attachments.

If a set of attachments alias each other, then all except the first to be used in the render pass **must** use an `initialLayout` of `VK_IMAGE_LAYOUT_UNDEFINED`, since the earlier uses of the other aliases make their contents undefined. Once an alias has been used and a different alias has been used after it, the first alias **must** not be used in any later subpasses. However, an application **can** assign the same image view to multiple aliasing attachment indices, which allows that image view to be used multiple times even if other aliases are used in between.

Note

Once an attachment needs the `VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT` bit, there **should** be no additional cost of introducing additional aliases, and using these additional aliases **may** allow more efficient clearing of the attachments on multiple uses via `VK_ATTACHMENT_LOAD_OP_CLEAR`.

Valid Usage

- [](#VUID-VkAttachmentDescription-format-06699)VUID-VkAttachmentDescription-format-06699  
  If `format` includes a color or depth component and `loadOp` is `VK_ATTACHMENT_LOAD_OP_LOAD`, then `initialLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`
- [](#VUID-VkAttachmentDescription-finalLayout-00843)VUID-VkAttachmentDescription-finalLayout-00843  
  `finalLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED` or `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT` or `VK_IMAGE_LAYOUT_PREINITIALIZED`
- [](#VUID-VkAttachmentDescription-format-03280)VUID-VkAttachmentDescription-format-03280  
  If `format` is a color format, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-03281)VUID-VkAttachmentDescription-format-03281  
  If `format` is a depth/stencil format, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-03282)VUID-VkAttachmentDescription-format-03282  
  If `format` is a color format, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-03283)VUID-VkAttachmentDescription-format-03283  
  If `format` is a depth/stencil format, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-06487)VUID-VkAttachmentDescription-format-06487  
  If `format` is a color format, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-06488)VUID-VkAttachmentDescription-format-06488  
  If `format` is a color format, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkAttachmentDescription-separateDepthStencilLayouts-03284)VUID-VkAttachmentDescription-separateDepthStencilLayouts-03284  
  If the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is not enabled, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,
- [](#VUID-VkAttachmentDescription-separateDepthStencilLayouts-03285)VUID-VkAttachmentDescription-separateDepthStencilLayouts-03285  
  If the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is not enabled, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,
- [](#VUID-VkAttachmentDescription-format-03286)VUID-VkAttachmentDescription-format-03286  
  If `format` is a color format, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-03287)VUID-VkAttachmentDescription-format-03287  
  If `format` is a color format, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-06906)VUID-VkAttachmentDescription-format-06906  
  If `format` is a depth/stencil format which includes both depth and stencil components, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-06907)VUID-VkAttachmentDescription-format-06907  
  If `format` is a depth/stencil format which includes both depth and stencil components, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-03290)VUID-VkAttachmentDescription-format-03290  
  If `format` is a depth/stencil format which includes only the depth component, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-03291)VUID-VkAttachmentDescription-format-03291  
  If `format` is a depth/stencil format which includes only the depth component, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription-synchronization2-06908)VUID-VkAttachmentDescription-synchronization2-06908  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkAttachmentDescription-synchronization2-06909)VUID-VkAttachmentDescription-synchronization2-06909  
  If the [`synchronization2`](#features-synchronization2) feature is not enabled, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkAttachmentDescription-attachmentFeedbackLoopLayout-07309)VUID-VkAttachmentDescription-attachmentFeedbackLoopLayout-07309  
  If the [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout) feature is not enabled, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`
- [](#VUID-VkAttachmentDescription-attachmentFeedbackLoopLayout-07310)VUID-VkAttachmentDescription-attachmentFeedbackLoopLayout-07310  
  If the [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout) feature is not enabled, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`
- [](#VUID-VkAttachmentDescription-samples-08745)VUID-VkAttachmentDescription-samples-08745  
  `samples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value that is set in `imageCreateSampleCounts` (as defined in [Image Creation Limits](#resources-image-creation-limits)) for the given `format`
- [](#VUID-VkAttachmentDescription-dynamicRenderingLocalRead-09544)VUID-VkAttachmentDescription-dynamicRenderingLocalRead-09544  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ`
- [](#VUID-VkAttachmentDescription-dynamicRenderingLocalRead-09545)VUID-VkAttachmentDescription-dynamicRenderingLocalRead-09545  
  If the [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead) feature is not enabled, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ`
- [](#VUID-VkAttachmentDescription-format-06698)VUID-VkAttachmentDescription-format-06698  
  `format` **must** not be VK\_FORMAT\_UNDEFINED
- [](#VUID-VkAttachmentDescription-format-06700)VUID-VkAttachmentDescription-format-06700  
  If `format` includes a stencil component and `stencilLoadOp` is `VK_ATTACHMENT_LOAD_OP_LOAD`, then `initialLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`
- [](#VUID-VkAttachmentDescription-format-03292)VUID-VkAttachmentDescription-format-03292  
  If `format` is a depth/stencil format which includes only the stencil component, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-03293)VUID-VkAttachmentDescription-format-03293  
  If `format` is a depth/stencil format which includes only the stencil component, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-06242)VUID-VkAttachmentDescription-format-06242  
  If `format` is a depth/stencil format which includes both depth and stencil components, `initialLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`
- [](#VUID-VkAttachmentDescription-format-06243)VUID-VkAttachmentDescription-format-06243  
  If `format` is a depth/stencil format which includes both depth and stencil components, `finalLayout` **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

Valid Usage (Implicit)

- [](#VUID-VkAttachmentDescription-flags-parameter)VUID-VkAttachmentDescription-flags-parameter  
  `flags` **must** be a valid combination of [VkAttachmentDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionFlagBits.html) values
- [](#VUID-VkAttachmentDescription-format-parameter)VUID-VkAttachmentDescription-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkAttachmentDescription-samples-parameter)VUID-VkAttachmentDescription-samples-parameter  
  `samples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value
- [](#VUID-VkAttachmentDescription-loadOp-parameter)VUID-VkAttachmentDescription-loadOp-parameter  
  `loadOp` **must** be a valid [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html) value
- [](#VUID-VkAttachmentDescription-storeOp-parameter)VUID-VkAttachmentDescription-storeOp-parameter  
  `storeOp` **must** be a valid [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) value
- [](#VUID-VkAttachmentDescription-stencilLoadOp-parameter)VUID-VkAttachmentDescription-stencilLoadOp-parameter  
  `stencilLoadOp` **must** be a valid [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html) value
- [](#VUID-VkAttachmentDescription-stencilStoreOp-parameter)VUID-VkAttachmentDescription-stencilStoreOp-parameter  
  `stencilStoreOp` **must** be a valid [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html) value
- [](#VUID-VkAttachmentDescription-initialLayout-parameter)VUID-VkAttachmentDescription-initialLayout-parameter  
  `initialLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkAttachmentDescription-finalLayout-parameter)VUID-VkAttachmentDescription-finalLayout-parameter  
  `finalLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAttachmentDescriptionFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionFlags.html), [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentLoadOp.html), [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentStoreOp.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo.html), [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAttachmentDescription).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0