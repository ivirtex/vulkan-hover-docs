# VkAttachmentDescription2(3) Manual Page

## Name

VkAttachmentDescription2 - Structure specifying an attachment
description



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAttachmentDescription2` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_create_renderpass2
typedef VkAttachmentDescription2 VkAttachmentDescription2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkAttachmentDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionFlagBits.html)
  specifying additional properties of the attachment.

- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value specifying the format of
  the image that will be used for the attachment.

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

Parameters defined by this structure with the same name as those in
[VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html) have the
identical effect to those parameters.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-separateDepthStencilLayouts"
target="_blank"
rel="noopener"><code>separateDepthStencilLayouts</code></a> feature is
enabled, and `format` is a depth/stencil format, `initialLayout` and
`finalLayout` **can** be set to a layout that only specifies the layout
of the depth aspect.

If the `pNext` chain includes a
[VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)
structure, then the `stencilInitialLayout` and `stencilFinalLayout`
members specify the initial and final layouts of the stencil aspect of a
depth/stencil format, and `initialLayout` and `finalLayout` only apply
to the depth aspect. For depth-only formats, the
[VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)
structure is ignored. For stencil-only formats, the initial and final
layouts of the stencil aspect are taken from the
[VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)
structure if present, or `initialLayout` and `finalLayout` if not
present.

If `format` is a depth/stencil format, and either `initialLayout` or
`finalLayout` does not specify a layout for the stencil aspect, then the
application **must** specify the initial and final layouts of the
stencil aspect by including a
[VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)
structure in the `pNext` chain.

`loadOp` and `storeOp` are ignored for fragment shading rate
attachments. No access to the shading rate attachment is performed in
`loadOp` and `storeOp`. Instead, access to
`VK_ACCESS_FRAGMENT_SHADING_RATE_ATTACHMENT_READ_BIT_KHR` is performed
as fragments are rasterized.

Valid Usage

- <a href="#VUID-VkAttachmentDescription2-format-06699"
  id="VUID-VkAttachmentDescription2-format-06699"></a>
  VUID-VkAttachmentDescription2-format-06699  
  If `format` includes a color or depth component and `loadOp` is
  `VK_ATTACHMENT_LOAD_OP_LOAD`, then `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_UNDEFINED`

- <a href="#VUID-VkAttachmentDescription2-finalLayout-00843"
  id="VUID-VkAttachmentDescription2-finalLayout-00843"></a>
  VUID-VkAttachmentDescription2-finalLayout-00843  
  `finalLayout` **must** not be `VK_IMAGE_LAYOUT_UNDEFINED` or
  `VK_IMAGE_LAYOUT_PREINITIALIZED`

- <a href="#VUID-VkAttachmentDescription2-format-03280"
  id="VUID-VkAttachmentDescription2-format-03280"></a>
  VUID-VkAttachmentDescription2-format-03280  
  If `format` is a color format, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-03281"
  id="VUID-VkAttachmentDescription2-format-03281"></a>
  VUID-VkAttachmentDescription2-format-03281  
  If `format` is a depth/stencil format, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-03282"
  id="VUID-VkAttachmentDescription2-format-03282"></a>
  VUID-VkAttachmentDescription2-format-03282  
  If `format` is a color format, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-03283"
  id="VUID-VkAttachmentDescription2-format-03283"></a>
  VUID-VkAttachmentDescription2-format-03283  
  If `format` is a depth/stencil format, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-06487"
  id="VUID-VkAttachmentDescription2-format-06487"></a>
  VUID-VkAttachmentDescription2-format-06487  
  If `format` is a color format, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-06488"
  id="VUID-VkAttachmentDescription2-format-06488"></a>
  VUID-VkAttachmentDescription2-format-06488  
  If `format` is a color format, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- <a
  href="#VUID-VkAttachmentDescription2-separateDepthStencilLayouts-03284"
  id="VUID-VkAttachmentDescription2-separateDepthStencilLayouts-03284"></a>
  VUID-VkAttachmentDescription2-separateDepthStencilLayouts-03284  
  If the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is not enabled, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,

- <a
  href="#VUID-VkAttachmentDescription2-separateDepthStencilLayouts-03285"
  id="VUID-VkAttachmentDescription2-separateDepthStencilLayouts-03285"></a>
  VUID-VkAttachmentDescription2-separateDepthStencilLayouts-03285  
  If the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is not enabled, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`,

- <a href="#VUID-VkAttachmentDescription2-format-03286"
  id="VUID-VkAttachmentDescription2-format-03286"></a>
  VUID-VkAttachmentDescription2-format-03286  
  If `format` is a color format, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-03287"
  id="VUID-VkAttachmentDescription2-format-03287"></a>
  VUID-VkAttachmentDescription2-format-03287  
  If `format` is a color format, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-06906"
  id="VUID-VkAttachmentDescription2-format-06906"></a>
  VUID-VkAttachmentDescription2-format-06906  
  If `format` is a depth/stencil format which includes both depth and
  stencil components, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-06907"
  id="VUID-VkAttachmentDescription2-format-06907"></a>
  VUID-VkAttachmentDescription2-format-06907  
  If `format` is a depth/stencil format which includes both depth and
  stencil components, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-03290"
  id="VUID-VkAttachmentDescription2-format-03290"></a>
  VUID-VkAttachmentDescription2-format-03290  
  If `format` is a depth/stencil format which includes only the depth
  component, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-03291"
  id="VUID-VkAttachmentDescription2-format-03291"></a>
  VUID-VkAttachmentDescription2-format-03291  
  If `format` is a depth/stencil format which includes only the depth
  component, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-synchronization2-06908"
  id="VUID-VkAttachmentDescription2-synchronization2-06908"></a>
  VUID-VkAttachmentDescription2-synchronization2-06908  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a href="#VUID-VkAttachmentDescription2-synchronization2-06909"
  id="VUID-VkAttachmentDescription2-synchronization2-06909"></a>
  VUID-VkAttachmentDescription2-synchronization2-06909  
  If the [`synchronization2`](#features-synchronization2) feature is not
  enabled, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR` or
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a
  href="#VUID-VkAttachmentDescription2-attachmentFeedbackLoopLayout-07309"
  id="VUID-VkAttachmentDescription2-attachmentFeedbackLoopLayout-07309"></a>
  VUID-VkAttachmentDescription2-attachmentFeedbackLoopLayout-07309  
  If the
  [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout)
  feature is not enabled, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`

- <a
  href="#VUID-VkAttachmentDescription2-attachmentFeedbackLoopLayout-07310"
  id="VUID-VkAttachmentDescription2-attachmentFeedbackLoopLayout-07310"></a>
  VUID-VkAttachmentDescription2-attachmentFeedbackLoopLayout-07310  
  If the
  [`attachmentFeedbackLoopLayout`](#features-attachmentFeedbackLoopLayout)
  feature is not enabled, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT`

- <a href="#VUID-VkAttachmentDescription2-samples-08745"
  id="VUID-VkAttachmentDescription2-samples-08745"></a>
  VUID-VkAttachmentDescription2-samples-08745  
  `samples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value that is set
  in `imageCreateSampleCounts` (as defined in [Image Creation
  Limits](#resources-image-creation-limits)) for the given `format`

- <a href="#VUID-VkAttachmentDescription2-dynamicRenderingLocalRead-09544"
  id="VUID-VkAttachmentDescription2-dynamicRenderingLocalRead-09544"></a>
  VUID-VkAttachmentDescription2-dynamicRenderingLocalRead-09544  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`

- <a href="#VUID-VkAttachmentDescription2-dynamicRenderingLocalRead-09545"
  id="VUID-VkAttachmentDescription2-dynamicRenderingLocalRead-09545"></a>
  VUID-VkAttachmentDescription2-dynamicRenderingLocalRead-09545  
  If the
  [`dynamicRenderingLocalRead`](#features-dynamicRenderingLocalRead)
  feature is not enabled, `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_RENDERING_LOCAL_READ_KHR`

- <a href="#VUID-VkAttachmentDescription2-pNext-06704"
  id="VUID-VkAttachmentDescription2-pNext-06704"></a>
  VUID-VkAttachmentDescription2-pNext-06704  
  If the `pNext` chain does not include a
  [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)
  structure, `format` includes a stencil component, and `stencilLoadOp`
  is `VK_ATTACHMENT_LOAD_OP_LOAD`, then `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_UNDEFINED`

- <a href="#VUID-VkAttachmentDescription2-pNext-06705"
  id="VUID-VkAttachmentDescription2-pNext-06705"></a>
  VUID-VkAttachmentDescription2-pNext-06705  
  If the `pNext` chain includes a
  [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)
  structure, `format` includes a stencil component, and `stencilLoadOp`
  is `VK_ATTACHMENT_LOAD_OP_LOAD`, then
  [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)::`stencilInitialLayout`
  **must** not be `VK_IMAGE_LAYOUT_UNDEFINED`

- <a href="#VUID-VkAttachmentDescription2-format-06249"
  id="VUID-VkAttachmentDescription2-format-06249"></a>
  VUID-VkAttachmentDescription2-format-06249  
  If `format` is a depth/stencil format which includes both depth and
  stencil components, and `initialLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, the `pNext` chain **must**
  include a
  [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)
  structure

- <a href="#VUID-VkAttachmentDescription2-format-06250"
  id="VUID-VkAttachmentDescription2-format-06250"></a>
  VUID-VkAttachmentDescription2-format-06250  
  If `format` is a depth/stencil format which includes both depth and
  stencil components, and `finalLayout` is
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, the `pNext` chain **must**
  include a
  [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)
  structure

- <a href="#VUID-VkAttachmentDescription2-format-06247"
  id="VUID-VkAttachmentDescription2-format-06247"></a>
  VUID-VkAttachmentDescription2-format-06247  
  If the `pNext` chain does not include a
  [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)
  structure and `format` only includes a stencil component,
  `initialLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-06248"
  id="VUID-VkAttachmentDescription2-format-06248"></a>
  VUID-VkAttachmentDescription2-format-06248  
  If the `pNext` chain does not include a
  [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)
  structure and `format` only includes a stencil component,
  `finalLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkAttachmentDescription2-format-09332"
  id="VUID-VkAttachmentDescription2-format-09332"></a>
  VUID-VkAttachmentDescription2-format-09332  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  is not enabled, `format` **must** not be `VK_FORMAT_UNDEFINED`

- <a href="#VUID-VkAttachmentDescription2-format-09334"
  id="VUID-VkAttachmentDescription2-format-09334"></a>
  VUID-VkAttachmentDescription2-format-09334  
  If `format` is `VK_FORMAT_UNDEFINED`, there **must** be a
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html) structure in
  the `pNext` chain with a `externalFormat` that is not equal to `0`

Valid Usage (Implicit)

- <a href="#VUID-VkAttachmentDescription2-sType-sType"
  id="VUID-VkAttachmentDescription2-sType-sType"></a>
  VUID-VkAttachmentDescription2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_2`

- <a href="#VUID-VkAttachmentDescription2-pNext-pNext"
  id="VUID-VkAttachmentDescription2-pNext-pNext"></a>
  VUID-VkAttachmentDescription2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkAttachmentDescriptionStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionStencilLayout.html)
  or [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)

- <a href="#VUID-VkAttachmentDescription2-sType-unique"
  id="VUID-VkAttachmentDescription2-sType-unique"></a>
  VUID-VkAttachmentDescription2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkAttachmentDescription2-flags-parameter"
  id="VUID-VkAttachmentDescription2-flags-parameter"></a>
  VUID-VkAttachmentDescription2-flags-parameter  
  `flags` **must** be a valid combination of
  [VkAttachmentDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionFlagBits.html)
  values

- <a href="#VUID-VkAttachmentDescription2-format-parameter"
  id="VUID-VkAttachmentDescription2-format-parameter"></a>
  VUID-VkAttachmentDescription2-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkAttachmentDescription2-samples-parameter"
  id="VUID-VkAttachmentDescription2-samples-parameter"></a>
  VUID-VkAttachmentDescription2-samples-parameter  
  `samples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

- <a href="#VUID-VkAttachmentDescription2-loadOp-parameter"
  id="VUID-VkAttachmentDescription2-loadOp-parameter"></a>
  VUID-VkAttachmentDescription2-loadOp-parameter  
  `loadOp` **must** be a valid
  [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html) value

- <a href="#VUID-VkAttachmentDescription2-storeOp-parameter"
  id="VUID-VkAttachmentDescription2-storeOp-parameter"></a>
  VUID-VkAttachmentDescription2-storeOp-parameter  
  `storeOp` **must** be a valid
  [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html) value

- <a href="#VUID-VkAttachmentDescription2-stencilLoadOp-parameter"
  id="VUID-VkAttachmentDescription2-stencilLoadOp-parameter"></a>
  VUID-VkAttachmentDescription2-stencilLoadOp-parameter  
  `stencilLoadOp` **must** be a valid
  [VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html) value

- <a href="#VUID-VkAttachmentDescription2-stencilStoreOp-parameter"
  id="VUID-VkAttachmentDescription2-stencilStoreOp-parameter"></a>
  VUID-VkAttachmentDescription2-stencilStoreOp-parameter  
  `stencilStoreOp` **must** be a valid
  [VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html) value

- <a href="#VUID-VkAttachmentDescription2-initialLayout-parameter"
  id="VUID-VkAttachmentDescription2-initialLayout-parameter"></a>
  VUID-VkAttachmentDescription2-initialLayout-parameter  
  `initialLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkAttachmentDescription2-finalLayout-parameter"
  id="VUID-VkAttachmentDescription2-finalLayout-parameter"></a>
  VUID-VkAttachmentDescription2-finalLayout-parameter  
  `finalLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkAttachmentDescriptionFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionFlags.html),
[VkAttachmentLoadOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentLoadOp.html),
[VkAttachmentStoreOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentStoreOp.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAttachmentDescription2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
