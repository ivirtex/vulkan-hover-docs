# VkAttachmentDescription(3) Manual Page

## Name

VkAttachmentDescription - Structure specifying an attachment description



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAttachmentDescription` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `flags` is a bitmask of
  [VkAttachmentDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionFlagBits.html)
  specifying additional properties of the attachment.

- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value specifying the format of
  the image view that will be used for the attachment.

- `samples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html)
  value specifying the number of samples of the image.

- `loadOp` is a [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html) value
  specifying how the contents of color and depth components of the
  attachment are treated at the beginning of the subpass where it is
  first used.

- `storeOp` is a [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html) value
  specifying how the contents of color and depth components of the
  attachment are treated at the end of the subpass where it is last
  used.

- `stencilLoadOp` is a [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html)
  value specifying how the contents of stencil components of the
  attachment are treated at the beginning of the subpass where it is
  first used.

- `stencilStoreOp` is a [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html)
  value specifying how the contents of stencil components of the
  attachment are treated at the end of the last subpass where it is
  used.

- `initialLayout` is the layout the attachment image subresource will be
  in when a render pass instance begins.

- `finalLayout` is the layout the attachment image subresource will be
  transitioned to when a render pass instance ends.

## <a href="#_description" class="anchor"></a>Description

If the attachment uses a color format, then `loadOp` and `storeOp` are
used, and `stencilLoadOp` and `stencilStoreOp` are ignored. If the
format has depth and/or stencil components, `loadOp` and `storeOp` apply
only to the depth data, while `stencilLoadOp` and `stencilStoreOp`
define how the stencil data is handled. `loadOp` and `stencilLoadOp`
define the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-load-operations"
target="_blank" rel="noopener">load operations</a> for the attachment.
`storeOp` and `stencilStoreOp` define the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-store-operations"
target="_blank" rel="noopener">store operations</a> for the attachment.
If an attachment is not used by any subpass, `loadOp`, `storeOp`,
`stencilStoreOp`, and `stencilLoadOp` will be ignored for that
attachment, and no load or store ops will be performed. However, any
transition specified by `initialLayout` and `finalLayout` will still be
executed.

If `flags` includes `VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT`, then the
attachment is treated as if it shares physical memory with another
attachment in the same render pass. This information limits the ability
of the implementation to reorder certain operations (like layout
transitions and the `loadOp`) such that it is not improperly reordered
against other uses of the same physical memory via a different
attachment. This is described in more detail below.

If a render pass uses multiple attachments that alias the same device
memory, those attachments **must** each include the
`VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT` bit in their attachment
description flags. Attachments aliasing the same memory occurs in
multiple ways:

- Multiple attachments being assigned the same image view as part of
  framebuffer creation.

- Attachments using distinct image views that correspond to the same
  image subresource of an image.

- Attachments using views of distinct image subresources which are bound
  to overlapping memory ranges.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Render passes <strong>must</strong> include subpass dependencies
(either directly or via a subpass dependency chain) between any two
subpasses that operate on the same attachment or aliasing attachments
and those subpass dependencies <strong>must</strong> include execution
and memory dependencies separating uses of the aliases, if at least one
of those subpasses writes to one of the aliases. These dependencies
<strong>must</strong> not include the
<code>VK_DEPENDENCY_BY_REGION_BIT</code> if the aliases are views of
distinct image subresources which overlap in memory.</p></td>
</tr>
</tbody>
</table>

Multiple attachments that alias the same memory **must** not be used in
a single subpass. A given attachment index **must** not be used multiple
times in a single subpass, with one exception: two subpass attachments
**can** use the same attachment index if at least one use is as an input
attachment and neither use is as a resolve or preserve attachment. In
other words, the same view **can** be used simultaneously as an input
and color or depth/stencil attachment, but **must** not be used as
multiple color or depth/stencil attachments nor as resolve or preserve
attachments.

If a set of attachments alias each other, then all except the first to
be used in the render pass **must** use an `initialLayout` of
`VK_IMAGE_LAYOUT_UNDEFINED`, since the earlier uses of the other aliases
make their contents undefined. Once an alias has been used and a
different alias has been used after it, the first alias **must** not be
used in any later subpasses. However, an application **can** assign the
same image view to multiple aliasing attachment indices, which allows
that image view to be used multiple times even if other aliases are used
in between.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Once an attachment needs the
<code>VK_ATTACHMENT_DESCRIPTION_MAY_ALIAS_BIT</code> bit, there
<strong>should</strong> be no additional cost of introducing additional
aliases, and using these additional aliases <strong>may</strong> allow
more efficient clearing of the attachments on multiple uses via
<code>VK_ATTACHMENT_LOAD_OP_CLEAR</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkAttachmentDescription-format-06699"
  id="VUID-VkAttachmentDescription-format-06699"></a>
  VUID-VkAttachmentDescription-format-06699  
  If `format` includes a color or depth component and `loadOp` is
  `VK_ATTACHMENT_LOAD_OP_LOAD`, then `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_UNDEFINED`

- <a href="#VUID-VkAttachmentDescription-finalLayout-00843"
  id="VUID-VkAttachmentDescription-finalLayout-00843"></a>
  VUID-VkAttachmentDescription-finalLayout-00843  
  `finalLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED` or
  `VK_IMAGE_LAYOUT_PREINITIALIZED`

- <a href="#VUID-VkAttachmentDescription-format-03280"
  id="VUID-VkAttachmentDescription-format-03280"></a>
  VUID-VkAttachmentDescription-format-03280  
  If `format` is a color format, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-03281"
  id="VUID-VkAttachmentDescription-format-03281"></a>
  VUID-VkAttachmentDescription-format-03281  
  If `format` is a depth/stencil format, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-03282"
  id="VUID-VkAttachmentDescription-format-03282"></a>
  VUID-VkAttachmentDescription-format-03282  
  If `format` is a color format, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-03283"
  id="VUID-VkAttachmentDescription-format-03283"></a>
  VUID-VkAttachmentDescription-format-03283  
  If `format` is a depth/stencil format, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-06487"
  id="VUID-VkAttachmentDescription-format-06487"></a>
  VUID-VkAttachmentDescription-format-06487  
  If `format` is a color format, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-06488"
  id="VUID-VkAttachmentDescription-format-06488"></a>
  VUID-VkAttachmentDescription-format-06488  
  If `format` is a color format, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- <a
  href="#VUID-VkAttachmentDescription-separateDepthStencilLayouts-03284"
  id="VUID-VkAttachmentDescription-separateDepthStencilLayouts-03284"></a>
  VUID-VkAttachmentDescription-separateDepthStencilLayouts-03284  
  If the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is not enabled, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,

- <a
  href="#VUID-VkAttachmentDescription-separateDepthStencilLayouts-03285"
  id="VUID-VkAttachmentDescription-separateDepthStencilLayouts-03285"></a>
  VUID-VkAttachmentDescription-separateDepthStencilLayouts-03285  
  If the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is not enabled, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,

- <a href="#VUID-VkAttachmentDescription-format-03286"
  id="VUID-VkAttachmentDescription-format-03286"></a>
  VUID-VkAttachmentDescription-format-03286  
  If `format` is a color format, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-03287"
  id="VUID-VkAttachmentDescription-format-03287"></a>
  VUID-VkAttachmentDescription-format-03287  
  If `format` is a color format, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-06906"
  id="VUID-VkAttachmentDescription-format-06906"></a>
  VUID-VkAttachmentDescription-format-06906  
  If `format` is a depth/stencil format which includes both depth and
  stencil components, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-06907"
  id="VUID-VkAttachmentDescription-format-06907"></a>
  VUID-VkAttachmentDescription-format-06907  
  If `format` is a depth/stencil format which includes both depth and
  stencil components, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-03290"
  id="VUID-VkAttachmentDescription-format-03290"></a>
  VUID-VkAttachmentDescription-format-03290  
  If `format` is a depth/stencil format which includes only the depth
  component, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-03291"
  id="VUID-VkAttachmentDescription-format-03291"></a>
  VUID-VkAttachmentDescription-format-03291  
  If `format` is a depth/stencil format which includes only the depth
  component, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-synchronization2-06908"
  id="VUID-VkAttachmentDescription-synchronization2-06908"></a>
  VUID-VkAttachmentDescription-synchronization2-06908  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a href="#VUID-VkAttachmentDescription-synchronization2-06909"
  id="VUID-VkAttachmentDescription-synchronization2-06909"></a>
  VUID-VkAttachmentDescription-synchronization2-06909  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a
  href="#VUID-VkAttachmentDescription-attachmentFeedbackLoopLayout-07309"
  id="VUID-VkAttachmentDescription-attachmentFeedbackLoopLayout-07309"></a>
  VUID-VkAttachmentDescription-attachmentFeedbackLoopLayout-07309  
  If the
  [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout)
  feature is not enabled, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`

- <a
  href="#VUID-VkAttachmentDescription-attachmentFeedbackLoopLayout-07310"
  id="VUID-VkAttachmentDescription-attachmentFeedbackLoopLayout-07310"></a>
  VUID-VkAttachmentDescription-attachmentFeedbackLoopLayout-07310  
  If the
  [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout)
  feature is not enabled, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`

- <a href="#VUID-VkAttachmentDescription-samples-08745"
  id="VUID-VkAttachmentDescription-samples-08745"></a>
  VUID-VkAttachmentDescription-samples-08745  
  `samples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value that is set
  in `imageCreateSampleCounts` (as defined in [Image Creation
  Limits](#resources-image-creation-limits)) for the given `format`

- <a href="#VUID-VkAttachmentDescription-dynamicRenderingLocalRead-09544"
  id="VUID-VkAttachmentDescription-dynamicRenderingLocalRead-09544"></a>
  VUID-VkAttachmentDescription-dynamicRenderingLocalRead-09544  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`

- <a href="#VUID-VkAttachmentDescription-dynamicRenderingLocalRead-09545"
  id="VUID-VkAttachmentDescription-dynamicRenderingLocalRead-09545"></a>
  VUID-VkAttachmentDescription-dynamicRenderingLocalRead-09545  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`

- <a href="#VUID-VkAttachmentDescription-format-06698"
  id="VUID-VkAttachmentDescription-format-06698"></a>
  VUID-VkAttachmentDescription-format-06698  
  `format` **must** not be VK_FORMAT_UNDEFINED

- <a href="#VUID-VkAttachmentDescription-format-06700"
  id="VUID-VkAttachmentDescription-format-06700"></a>
  VUID-VkAttachmentDescription-format-06700  
  If `format` includes a stencil component and `stencilLoadOp` is
  `VK_ATTACHMENT_LOAD_OP_LOAD`, then `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_UNDEFINED`

- <a href="#VUID-VkAttachmentDescription-format-03292"
  id="VUID-VkAttachmentDescription-format-03292"></a>
  VUID-VkAttachmentDescription-format-03292  
  If `format` is a depth/stencil format which includes only the stencil
  component, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-03293"
  id="VUID-VkAttachmentDescription-format-03293"></a>
  VUID-VkAttachmentDescription-format-03293  
  If `format` is a depth/stencil format which includes only the stencil
  component, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-06242"
  id="VUID-VkAttachmentDescription-format-06242"></a>
  VUID-VkAttachmentDescription-format-06242  
  If `format` is a depth/stencil format which includes both depth and
  stencil components, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription-format-06243"
  id="VUID-VkAttachmentDescription-format-06243"></a>
  VUID-VkAttachmentDescription-format-06243  
  If `format` is a depth/stencil format which includes both depth and
  stencil components, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

Valid Usage (Implicit)

- <a href="#VUID-VkAttachmentDescription-flags-parameter"
  id="VUID-VkAttachmentDescription-flags-parameter"></a>
  VUID-VkAttachmentDescription-flags-parameter  
  `flags` **must** be a valid combination of
  [VkAttachmentDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionFlagBits.html)
  values

- <a href="#VUID-VkAttachmentDescription-format-parameter"
  id="VUID-VkAttachmentDescription-format-parameter"></a>
  VUID-VkAttachmentDescription-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkAttachmentDescription-samples-parameter"
  id="VUID-VkAttachmentDescription-samples-parameter"></a>
  VUID-VkAttachmentDescription-samples-parameter  
  `samples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

- <a href="#VUID-VkAttachmentDescription-loadOp-parameter"
  id="VUID-VkAttachmentDescription-loadOp-parameter"></a>
  VUID-VkAttachmentDescription-loadOp-parameter  
  `loadOp` **must** be a valid
  [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html) value

- <a href="#VUID-VkAttachmentDescription-storeOp-parameter"
  id="VUID-VkAttachmentDescription-storeOp-parameter"></a>
  VUID-VkAttachmentDescription-storeOp-parameter  
  `storeOp` **must** be a valid
  [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html) value

- <a href="#VUID-VkAttachmentDescription-stencilLoadOp-parameter"
  id="VUID-VkAttachmentDescription-stencilLoadOp-parameter"></a>
  VUID-VkAttachmentDescription-stencilLoadOp-parameter  
  `stencilLoadOp` **must** be a valid
  [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html) value

- <a href="#VUID-VkAttachmentDescription-stencilStoreOp-parameter"
  id="VUID-VkAttachmentDescription-stencilStoreOp-parameter"></a>
  VUID-VkAttachmentDescription-stencilStoreOp-parameter  
  `stencilStoreOp` **must** be a valid
  [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html) value

- <a href="#VUID-VkAttachmentDescription-initialLayout-parameter"
  id="VUID-VkAttachmentDescription-initialLayout-parameter"></a>
  VUID-VkAttachmentDescription-initialLayout-parameter  
  `initialLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkAttachmentDescription-finalLayout-parameter"
  id="VUID-VkAttachmentDescription-finalLayout-parameter"></a>
  VUID-VkAttachmentDescription-finalLayout-parameter  
  `finalLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAttachmentDescriptionFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionFlags.html),
[VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html),
[VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAttachmentDescription"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
