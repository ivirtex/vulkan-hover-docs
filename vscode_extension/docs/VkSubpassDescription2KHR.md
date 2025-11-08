# VkSubpassDescription2(3) Manual Page

## Name

VkSubpassDescription2 - Structure specifying a subpass description



## [](#_c_specification)C Specification

The `VkSubpassDescription2` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_2
typedef struct VkSubpassDescription2 {
    VkStructureType                  sType;
    const void*                      pNext;
    VkSubpassDescriptionFlags        flags;
    VkPipelineBindPoint              pipelineBindPoint;
    uint32_t                         viewMask;
    uint32_t                         inputAttachmentCount;
    const VkAttachmentReference2*    pInputAttachments;
    uint32_t                         colorAttachmentCount;
    const VkAttachmentReference2*    pColorAttachments;
    const VkAttachmentReference2*    pResolveAttachments;
    const VkAttachmentReference2*    pDepthStencilAttachment;
    uint32_t                         preserveAttachmentCount;
    const uint32_t*                  pPreserveAttachments;
} VkSubpassDescription2;
```

or the equivalent

```c++
// Provided by VK_KHR_create_renderpass2
typedef VkSubpassDescription2 VkSubpassDescription2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionFlagBits.html) specifying usage of the subpass.
- `pipelineBindPoint` is a [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value specifying the pipeline type supported for this subpass.
- `viewMask` is a bitfield of view indices describing which views rendering is broadcast to in this subpass, when multiview is enabled.
- `inputAttachmentCount` is the number of input attachments.
- `pInputAttachments` is a pointer to an array of [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structures defining the input attachments for this subpass and their layouts.
- `colorAttachmentCount` is the number of color attachments.
- `pColorAttachments` is a pointer to an array of `colorAttachmentCount` [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structures defining the color attachments for this subpass and their layouts.
- `pResolveAttachments` is `NULL` or a pointer to an array of `colorAttachmentCount` [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structures defining the resolve attachments for this subpass and their layouts.
- `pDepthStencilAttachment` is a pointer to a [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structure specifying the depth/stencil attachment for this subpass and its layout.
- `preserveAttachmentCount` is the number of preserved attachments.
- `pPreserveAttachments` is a pointer to an array of `preserveAttachmentCount` render pass attachment indices identifying attachments that are not used by this subpass, but whose contents **must** be preserved throughout the subpass.

## [](#_description)Description

Parameters defined by this structure with the same name as those in [VkSubpassDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription.html) have the identical effect to those parameters.

`viewMask` has the same effect for the described subpass as [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html)::`pViewMasks` has on each corresponding subpass.

If a [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateAttachmentInfoKHR.html) structure is included in the `pNext` chain, `pFragmentShadingRateAttachment` is not `NULL`, and its `attachment` member is not `VK_ATTACHMENT_UNUSED`, the identified attachment defines a fragment shading rate attachment for that subpass.

If any element of `pResolveAttachments` is an image specified with an [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html), values in the corresponding color attachment will be resolved to the resolve attachment in the same manner as specified for [`VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html).

If the [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) limit is `VK_TRUE`, values in the color attachment will be loaded from the resolve attachment at the start of rendering, and **may** also be reloaded any time after a resolve occurs or the resolve attachment is written to; if this occurs it **must** happen-before any writes to the color attachment are performed which happen-after the resolve that triggers this. If any color component in the external format is subsampled, values will be read from the nearest sample in the image when they are loaded. If the color attachment is also used as an input attachment, the same behavior applies.

Setting the color attachment to `VK_ATTACHMENT_UNUSED` when an external resolve attachment is used and the [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) limit is `VK_TRUE` will not result in color attachment writes to be discarded for that attachment.

When [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) is `VK_TRUE`, the color output from the subpass can still be read via an input attachment; but the application cannot bind an image view for the color attachment as there is no such image view bound. Instead to access the data as an input attachment applications **can** use the resolve attachment in its place - using the resolve attachment image for the descriptor, and setting the corresponding element of `pInputAttachments` to the index of the resolve attachment.

Loads or input attachment reads from the resolve attachment are performed as if using a [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html) with the following parameters:

```c
VkSamplerYcbcrConversionCreateInfo createInfo = {
    .sType = VK_STRUCTURE_TYPE_SAMPLER_YCBCR_CONVERSION_CREATE_INFO,
    .pNext = NULL,
    .format = VK_FORMAT_UNDEFINED,
    .ycbcrModel = VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY,
    .ycbcrRange = VK_SAMPLER_YCBCR_RANGE_ITU_FULL,
    .components = {
        .r = VK_COMPONENT_SWIZZLE_B
        .g = VK_COMPONENT_SWIZZLE_R
        .b = VK_COMPONENT_SWIZZLE_G
        .a = VK_COMPONENT_SWIZZLE_IDENTITY},
    .xChromaOffset = properties.chromaOffsetX,
    .yChromaOffset = properties.chromaOffsetY,
    .chromaFilter = VK_FILTER_NEAREST,
    .forceExplicitReconstruction = ... };
```

where `properties` is equal to [VkPhysicalDeviceExternalFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalFormatResolvePropertiesANDROID.html) returned by the device and `forceExplicitReconstruction` is effectively ignored as the `VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY` model is used. The applied swizzle is the same effective swizzle that would be applied by the `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY` model, but no range expansion is applied.

Valid Usage

- [](#VUID-VkSubpassDescription2-attachment-06912)VUID-VkSubpassDescription2-attachment-06912  
  If the `attachment` member of an element of `pInputAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkSubpassDescription2-attachment-06913)VUID-VkSubpassDescription2-attachment-06913  
  If the `attachment` member of an element of `pColorAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`
- [](#VUID-VkSubpassDescription2-attachment-06914)VUID-VkSubpassDescription2-attachment-06914  
  If the `attachment` member of an element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`
- [](#VUID-VkSubpassDescription2-attachment-06915)VUID-VkSubpassDescription2-attachment-06915  
  If the `attachment` member of `pDepthStencilAttachment` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`
- [](#VUID-VkSubpassDescription2-attachment-06916)VUID-VkSubpassDescription2-attachment-06916  
  If the `attachment` member of an element of `pColorAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkSubpassDescription2-attachment-06917)VUID-VkSubpassDescription2-attachment-06917  
  If the `attachment` member of an element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkSubpassDescription2-attachment-06918)VUID-VkSubpassDescription2-attachment-06918  
  If the `attachment` member of an element of `pInputAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkSubpassDescription2-attachment-06919)VUID-VkSubpassDescription2-attachment-06919  
  If the `attachment` member of an element of `pColorAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkSubpassDescription2-attachment-06920)VUID-VkSubpassDescription2-attachment-06920  
  If the `attachment` member of an element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`, `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`, `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkSubpassDescription2-attachment-06921)VUID-VkSubpassDescription2-attachment-06921  
  If the `attachment` member of an element of `pInputAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR`
- [](#VUID-VkSubpassDescription2-attachment-06922)VUID-VkSubpassDescription2-attachment-06922  
  If the `attachment` member of an element of `pColorAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkSubpassDescription2-attachment-06923)VUID-VkSubpassDescription2-attachment-06923  
  If the `attachment` member of an element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`
- [](#VUID-VkSubpassDescription2-attachment-10755)VUID-VkSubpassDescription2-attachment-10755  
  If the `attachment` member of an element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED`, the underlying resource must not be bound to a `VkDeviceMemory` object allocated from a `VkMemoryHeap` with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property.
- [](#VUID-VkSubpassDescription2-flags-10683)VUID-VkSubpassDescription2-flags-10683  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_TILE_SHADING_APRON_BIT_QCOM`, the render pass **must** have been created with a [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`tileApronSize` greater than `(0,0)`
- [](#VUID-VkSubpassDescription2-attachment-06251)VUID-VkSubpassDescription2-attachment-06251  
  If the `attachment` member of `pDepthStencilAttachment` is not `VK_ATTACHMENT_UNUSED` and its `pNext` chain includes a [VkAttachmentReferenceStencilLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReferenceStencilLayout.html) structure, the `layout` member of `pDepthStencilAttachment` **must** not be `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkSubpassDescription2-pipelineBindPoint-04953)VUID-VkSubpassDescription2-pipelineBindPoint-04953  
  `pipelineBindPoint` **must** be `VK_PIPELINE_BIND_POINT_GRAPHICS` or `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI`
- [](#VUID-VkSubpassDescription2-colorAttachmentCount-03063)VUID-VkSubpassDescription2-colorAttachmentCount-03063  
  `colorAttachmentCount` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxColorAttachments`
- [](#VUID-VkSubpassDescription2-loadOp-03064)VUID-VkSubpassDescription2-loadOp-03064  
  If the first use of an attachment in this render pass is as an input attachment, and the attachment is not also used as a color or depth/stencil attachment in the same subpass, then `loadOp` **must** not be `VK_ATTACHMENT_LOAD_OP_CLEAR`
- [](#VUID-VkSubpassDescription2-pResolveAttachments-03067)VUID-VkSubpassDescription2-pResolveAttachments-03067  
  If `pResolveAttachments` is not `NULL`, each resolve attachment that is not `VK_ATTACHMENT_UNUSED` **must** have a sample count of `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkSubpassDescription2-externalFormatResolve-09335)VUID-VkSubpassDescription2-externalFormatResolve-09335  
  If the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is not enabled and `pResolveAttachments` is not `NULL`, for each resolve attachment that does not have the value `VK_ATTACHMENT_UNUSED`, the corresponding color attachment **must** not have the value `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription2-nullColorAttachmentWithExternalFormatResolve-09336)VUID-VkSubpassDescription2-nullColorAttachmentWithExternalFormatResolve-09336  
  If the [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) property is `VK_FALSE` and `pResolveAttachments` is not `NULL`, for each resolve attachment that has a format of `VK_FORMAT_UNDEFINED`, the corresponding color attachment **must** not have the value `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription2-nullColorAttachmentWithExternalFormatResolve-09337)VUID-VkSubpassDescription2-nullColorAttachmentWithExternalFormatResolve-09337  
  If the [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) property is `VK_TRUE` and `pResolveAttachments` is not `NULL`, for each resolve attachment that has a format of `VK_FORMAT_UNDEFINED`, the corresponding color attachment **must** have the value `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription2-externalFormatResolve-09338)VUID-VkSubpassDescription2-externalFormatResolve-09338  
  If the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is not enabled and `pResolveAttachments` is not `NULL`, for each resolve attachment that is not `VK_ATTACHMENT_UNUSED`, the corresponding color attachment **must** not have a sample count of `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkSubpassDescription2-externalFormatResolve-09339)VUID-VkSubpassDescription2-externalFormatResolve-09339  
  If the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is not enabled, each element of `pResolveAttachments` **must** have the same [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) as its corresponding color attachment
- [](#VUID-VkSubpassDescription2-multisampledRenderToSingleSampled-06869)VUID-VkSubpassDescription2-multisampledRenderToSingleSampled-06869  
  If the [`multisampledRenderToSingleSampled`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multisampledRenderToSingleSampled) feature is not enabled, all attachments in `pColorAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have the same sample count
- [](#VUID-VkSubpassDescription2-pInputAttachments-02897)VUID-VkSubpassDescription2-pInputAttachments-02897  
  All attachments in `pInputAttachments` that are not `VK_ATTACHMENT_UNUSED` and any of the following is true:
  
  - the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is not enabled
  - the [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) property is `VK_FALSE`
  - does not have a non-zero value of [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  
  **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) contain at least `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` or `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkSubpassDescription2-pColorAttachments-02898)VUID-VkSubpassDescription2-pColorAttachments-02898  
  All attachments in `pColorAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) contain `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT`
- [](#VUID-VkSubpassDescription2-pResolveAttachments-09343)VUID-VkSubpassDescription2-pResolveAttachments-09343  
  All attachments in `pResolveAttachments` that are not `VK_ATTACHMENT_UNUSED` and do not have an image format of `VK_FORMAT_UNDEFINED` **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) contain `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT`
- [](#VUID-VkSubpassDescription2-pDepthStencilAttachment-02900)VUID-VkSubpassDescription2-pDepthStencilAttachment-02900  
  If `pDepthStencilAttachment` is not `NULL` and the attachment is not `VK_ATTACHMENT_UNUSED` then it **must** have an image format whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) contain `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkSubpassDescription2-linearColorAttachment-06499)VUID-VkSubpassDescription2-linearColorAttachment-06499  
  If the [`linearColorAttachment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-linearColorAttachment) feature is enabled and the image is created with `VK_IMAGE_TILING_LINEAR`, all attachments in `pInputAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) **must** contain `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`
- [](#VUID-VkSubpassDescription2-linearColorAttachment-06500)VUID-VkSubpassDescription2-linearColorAttachment-06500  
  If the [`linearColorAttachment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-linearColorAttachment) feature is enabled and the image is created with `VK_IMAGE_TILING_LINEAR`, all attachments in `pColorAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) **must** contain `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`
- [](#VUID-VkSubpassDescription2-linearColorAttachment-06501)VUID-VkSubpassDescription2-linearColorAttachment-06501  
  If the [`linearColorAttachment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-linearColorAttachment) feature is enabled and the image is created with `VK_IMAGE_TILING_LINEAR`, all attachments in `pResolveAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) **must** contain `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`
- [](#VUID-VkSubpassDescription2-None-09456)VUID-VkSubpassDescription2-None-09456  
  If either of the following is enabled:
  
  - The `VK_AMD_mixed_attachment_samples` extension
  - The `VK_NV_framebuffer_mixed_samples` extension
  
  all attachments in `pColorAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have a sample count that is smaller than or equal to the sample count of `pDepthStencilAttachment` if it is not `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription2-pNext-06870)VUID-VkSubpassDescription2-pNext-06870  
  If the `pNext` chain includes a [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html) structure with `multisampledRenderToSingleSampledEnable` equal to `VK_TRUE`, then all attachments in `pColorAttachments` and `pDepthStencilAttachment` that are not `VK_ATTACHMENT_UNUSED` **must** have a sample count that is either `VK_SAMPLE_COUNT_1_BIT` or equal to [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)::`rasterizationSamples`
- [](#VUID-VkSubpassDescription2-pNext-06871)VUID-VkSubpassDescription2-pNext-06871  
  If the `pNext` chain includes a [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html) structure with `multisampledRenderToSingleSampledEnable` equal to `VK_TRUE`, and `pDepthStencilAttachment` is not `NULL`, does not have the value `VK_ATTACHMENT_UNUSED`, and has a sample count of `VK_SAMPLE_COUNT_1_BIT`, the `pNext` chain **must** also include a [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionDepthStencilResolve.html) structure with `pDepthStencilResolveAttachment` that is either `NULL` or has the value `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription2-multisampledRenderToSingleSampled-06872)VUID-VkSubpassDescription2-multisampledRenderToSingleSampled-06872  
  All attachments in `pDepthStencilAttachment` or `pColorAttachments` that are not `VK_ATTACHMENT_UNUSED` **must** have the same sample count , if none of the following are enabled:
  
  - The `VK_AMD_mixed_attachment_samples` extension
  - The `VK_NV_framebuffer_mixed_samples` extension
  - The [`multisampledRenderToSingleSampled`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multisampledRenderToSingleSampled) feature,
- [](#VUID-VkSubpassDescription2-attachment-03073)VUID-VkSubpassDescription2-attachment-03073  
  Each element of `pPreserveAttachments` **must** not be `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription2-pPreserveAttachments-03074)VUID-VkSubpassDescription2-pPreserveAttachments-03074  
  Each element of `pPreserveAttachments` **must** not also be an element of any other member of the subpass description
- [](#VUID-VkSubpassDescription2-layout-02528)VUID-VkSubpassDescription2-layout-02528  
  If any attachment is used by more than one [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) member, then each use **must** use the same `layout`
- [](#VUID-VkSubpassDescription2-flags-03076)VUID-VkSubpassDescription2-flags-03076  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX`, it **must** also include `VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX`
- [](#VUID-VkSubpassDescription2-attachment-02799)VUID-VkSubpassDescription2-attachment-02799  
  If the `attachment` member of any element of `pInputAttachments` is not `VK_ATTACHMENT_UNUSED`, then the `aspectMask` member **must** be a valid combination of [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html)
- [](#VUID-VkSubpassDescription2-attachment-02800)VUID-VkSubpassDescription2-attachment-02800  
  If the `attachment` member of any element of `pInputAttachments` is not `VK_ATTACHMENT_UNUSED`, then the `aspectMask` member **must** not be `0`
- [](#VUID-VkSubpassDescription2-attachment-02801)VUID-VkSubpassDescription2-attachment-02801  
  If the `attachment` member of any element of `pInputAttachments` is not `VK_ATTACHMENT_UNUSED`, then the `aspectMask` member **must** not include `VK_IMAGE_ASPECT_METADATA_BIT`
- [](#VUID-VkSubpassDescription2-attachment-04563)VUID-VkSubpassDescription2-attachment-04563  
  If the `attachment` member of any element of `pInputAttachments` is not `VK_ATTACHMENT_UNUSED`, then the `aspectMask` member **must** not include `VK_IMAGE_ASPECT_MEMORY_PLANE_i_BIT_EXT` for any index *i*
- [](#VUID-VkSubpassDescription2-pDepthStencilAttachment-04440)VUID-VkSubpassDescription2-pDepthStencilAttachment-04440  
  An attachment **must** not be used in both `pDepthStencilAttachment` and `pColorAttachments`
- [](#VUID-VkSubpassDescription2-multiview-06558)VUID-VkSubpassDescription2-multiview-06558  
  If the [`multiview`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiview) feature is not enabled, `viewMask` **must** be `0`
- [](#VUID-VkSubpassDescription2-viewMask-06706)VUID-VkSubpassDescription2-viewMask-06706  
  The index of the most significant bit in `viewMask` **must** be less than [`maxMultiviewViewCount`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxMultiviewViewCount)
- [](#VUID-VkSubpassDescription2-externalFormatResolve-09344)VUID-VkSubpassDescription2-externalFormatResolve-09344  
  If the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is enabled, `pResolveAttachments` is not `NULL`, and `colorAttachmentCount` is not `1`, any element of `pResolveAttachments` that is not `VK_ATTACHMENT_UNUSED`, **must** not have a format of `VK_FORMAT_UNDEFINED`
- [](#VUID-VkSubpassDescription2-externalFormatResolve-09345)VUID-VkSubpassDescription2-externalFormatResolve-09345  
  If the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is enabled, `pResolveAttachments` is not `NULL`, any element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED` and has a format of `VK_FORMAT_UNDEFINED`, and the corresponding element of `pColorAttachments` is not `VK_ATTACHMENT_UNUSED`, the color attachment **must** have a `samples` value of `1`
- [](#VUID-VkSubpassDescription2-externalFormatResolve-09346)VUID-VkSubpassDescription2-externalFormatResolve-09346  
  If the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is enabled, `pResolveAttachments` is not `NULL`, and any element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED` and has a format of `VK_FORMAT_UNDEFINED`, `viewMask` **must** be `0`
- [](#VUID-VkSubpassDescription2-externalFormatResolve-09347)VUID-VkSubpassDescription2-externalFormatResolve-09347  
  If the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is enabled, `pResolveAttachments` is not `NULL`, and any element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED` and has a format of `VK_FORMAT_UNDEFINED`, [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateAttachmentInfoKHR.html)::`pFragmentShadingRateAttachment` **must** either be `NULL` or a [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structure with an `attachment` value of `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription2-externalFormatResolve-09348)VUID-VkSubpassDescription2-externalFormatResolve-09348  
  If the [`externalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-externalFormatResolve) feature is enabled, `pResolveAttachments` is not `NULL`, and any element of `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED` and has a format of `VK_FORMAT_UNDEFINED`, elements of `pInputAttachments` referencing either a color attachment or resolve attachment used in this subpass **must** not include `VK_IMAGE_ASPECT_PLANE_i_BIT` for any index *i* in its `aspectMask`
- [](#VUID-VkSubpassDescription2-flags-04907)VUID-VkSubpassDescription2-flags-04907  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`, and if `pResolveAttachments` is not `NULL`, then each resolve attachment **must** be `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription2-flags-04908)VUID-VkSubpassDescription2-flags-04908  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`, and if `pDepthStencilResolveAttachment` is not `NULL`, then the depth/stencil resolve attachment **must** be `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkSubpassDescription2-flags-04909)VUID-VkSubpassDescription2-flags-04909  
  If `flags` includes `VK_SUBPASS_DESCRIPTION_SHADER_RESOLVE_BIT_QCOM`, then the subpass **must** be the last subpass in a subpass dependency chain

Valid Usage (Implicit)

- [](#VUID-VkSubpassDescription2-sType-sType)VUID-VkSubpassDescription2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2`
- [](#VUID-VkSubpassDescription2-pNext-pNext)VUID-VkSubpassDescription2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateAttachmentInfoKHR.html), [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html), [VkRenderPassCreationControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreationControlEXT.html), [VkRenderPassSubpassFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSubpassFeedbackCreateInfoEXT.html), or [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionDepthStencilResolve.html)
- [](#VUID-VkSubpassDescription2-sType-unique)VUID-VkSubpassDescription2-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkSubpassDescription2-flags-parameter)VUID-VkSubpassDescription2-flags-parameter  
  `flags` **must** be a valid combination of [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionFlagBits.html) values
- [](#VUID-VkSubpassDescription2-pipelineBindPoint-parameter)VUID-VkSubpassDescription2-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html) value
- [](#VUID-VkSubpassDescription2-pInputAttachments-parameter)VUID-VkSubpassDescription2-pInputAttachments-parameter  
  If `inputAttachmentCount` is not `0`, `pInputAttachments` **must** be a valid pointer to an array of `inputAttachmentCount` valid [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structures
- [](#VUID-VkSubpassDescription2-pColorAttachments-parameter)VUID-VkSubpassDescription2-pColorAttachments-parameter  
  If `colorAttachmentCount` is not `0`, `pColorAttachments` **must** be a valid pointer to an array of `colorAttachmentCount` valid [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structures
- [](#VUID-VkSubpassDescription2-pResolveAttachments-parameter)VUID-VkSubpassDescription2-pResolveAttachments-parameter  
  If `colorAttachmentCount` is not `0`, and `pResolveAttachments` is not `NULL`, `pResolveAttachments` **must** be a valid pointer to an array of `colorAttachmentCount` valid [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structures
- [](#VUID-VkSubpassDescription2-pDepthStencilAttachment-parameter)VUID-VkSubpassDescription2-pDepthStencilAttachment-parameter  
  If `pDepthStencilAttachment` is not `NULL`, `pDepthStencilAttachment` **must** be a valid pointer to a valid [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html) structure
- [](#VUID-VkSubpassDescription2-pPreserveAttachments-parameter)VUID-VkSubpassDescription2-pPreserveAttachments-parameter  
  If `preserveAttachmentCount` is not `0`, `pPreserveAttachments` **must** be a valid pointer to an array of `preserveAttachmentCount` `uint32_t` values

## [](#_see_also)See Also

[VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html), [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBindPoint.html), [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubpassDescriptionFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubpassDescription2).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0