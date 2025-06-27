# VkSubpassDescription2(3) Manual Page

## Name

VkSubpassDescription2 - Structure specifying a subpass description



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSubpassDescription2` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_create_renderpass2
typedef VkSubpassDescription2 VkSubpassDescription2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionFlagBits.html)
  specifying usage of the subpass.

- `pipelineBindPoint` is a
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value specifying the
  pipeline type supported for this subpass.

- `viewMask` is a bitfield of view indices describing which views
  rendering is broadcast to in this subpass, when multiview is enabled.

- `inputAttachmentCount` is the number of input attachments.

- `pInputAttachments` is a pointer to an array of
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structures
  defining the input attachments for this subpass and their layouts.

- `colorAttachmentCount` is the number of color attachments.

- `pColorAttachments` is a pointer to an array of `colorAttachmentCount`
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structures
  defining the color attachments for this subpass and their layouts.

- `pResolveAttachments` is `NULL` or a pointer to an array of
  `colorAttachmentCount`
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structures
  defining the resolve attachments for this subpass and their layouts.

- `pDepthStencilAttachment` is a pointer to a
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structure
  specifying the depth/stencil attachment for this subpass and its
  layout.

- `preserveAttachmentCount` is the number of preserved attachments.

- `pPreserveAttachments` is a pointer to an array of
  `preserveAttachmentCount` render pass attachment indices identifying
  attachments that are not used by this subpass, but whose contents
  **must** be preserved throughout the subpass.

## <a href="#_description" class="anchor"></a>Description

Parameters defined by this structure with the same name as those in
[VkSubpassDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription.html) have the identical
effect to those parameters.

`viewMask` has the same effect for the described subpass as
[VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)::`pViewMasks`
has on each corresponding subpass.

If a
[VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateAttachmentInfoKHR.html)
structure is included in the `pNext` chain,
`pFragmentShadingRateAttachment` is not `NULL`, and its `attachment`
member is not `VK_ATTACHMENT_UNUSED`, the identified attachment defines
a fragment shading rate attachment for that subpass.

If any element of `pResolveAttachments` is an image specified with an
[VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html), values in the
corresponding color attachment will be resolved to the resolve
attachment in the same manner as specified for
<a href="VkResolveModeFlagBits.html" target="_blank"
rel="noopener"><code>VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID</code></a>.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
target="_blank"
rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
limit is `VK_TRUE`, values in the color attachment will be loaded from
the resolve attachment at the start of rendering, and **may** also be
reloaded any time after a resolve occurs or the resolve attachment is
written to; if this occurs it **must** happen-before any writes to the
color attachment are performed which happen-after the resolve that
triggers this. If any color component in the external format is
subsampled, values will be read from the nearest sample in the image
when they are loaded. If the color attachment is also used as an input
attachment, the same behavior applies.

Setting the color attachment to `VK_ATTACHMENT_UNUSED` when an external
resolve attachment is used and the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
target="_blank"
rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
limit is `VK_TRUE` will not result in color attachment writes to be
discarded for that attachment.

When <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
target="_blank"
rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
is `VK_TRUE`, the color output from the subpass can still be read via an
input attachment; but the application cannot bind an image view for the
color attachment as there is no such image view bound. Instead to access
the data as an input attachment applications **can** use the resolve
attachment in its place - using the resolve attachment image for the
descriptor, and setting the corresponding element of `pInputAttachments`
to the index of the resolve attachment.

Loads or input attachment reads from the resolve attachment are
performed as if using a
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)
with the following parameters:

``` c
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

where `properties` is equal to
[VkPhysicalDeviceExternalFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFormatResolvePropertiesANDROID.html)
returned by the device and `forceExplicitReconstruction` is effectively
ignored as the `VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY` model is
used. The applied swizzle is the same effective swizzle that would be
applied by the `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY` model,
but no range expansion is applied.

Valid Usage

- <a href="#VUID-VkSubpassDescription2-attachment-06912"
  id="VUID-VkSubpassDescription2-attachment-06912"></a>
  VUID-VkSubpassDescription2-attachment-06912  
  If the `attachment` member of an element of `pInputAttachments` is not
  `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkSubpassDescription2-attachment-06913"
  id="VUID-VkSubpassDescription2-attachment-06913"></a>
  VUID-VkSubpassDescription2-attachment-06913  
  If the `attachment` member of an element of `pColorAttachments` is not
  `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkSubpassDescription2-attachment-06914"
  id="VUID-VkSubpassDescription2-attachment-06914"></a>
  VUID-VkSubpassDescription2-attachment-06914  
  If the `attachment` member of an element of `pResolveAttachments` is
  not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkSubpassDescription2-attachment-06915"
  id="VUID-VkSubpassDescription2-attachment-06915"></a>
  VUID-VkSubpassDescription2-attachment-06915  
  If the `attachment` member of `pDepthStencilAttachment` is not
  `VK_ATTACHMENT_UNUSED`, ts `layout` member **must** not be
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkSubpassDescription2-attachment-06916"
  id="VUID-VkSubpassDescription2-attachment-06916"></a>
  VUID-VkSubpassDescription2-attachment-06916  
  If the `attachment` member of an element of `pColorAttachments` is not
  `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkSubpassDescription2-attachment-06917"
  id="VUID-VkSubpassDescription2-attachment-06917"></a>
  VUID-VkSubpassDescription2-attachment-06917  
  If the `attachment` member of an element of `pResolveAttachments` is
  not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkSubpassDescription2-attachment-06918"
  id="VUID-VkSubpassDescription2-attachment-06918"></a>
  VUID-VkSubpassDescription2-attachment-06918  
  If the `attachment` member of an element of `pInputAttachments` is not
  `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkSubpassDescription2-attachment-06919"
  id="VUID-VkSubpassDescription2-attachment-06919"></a>
  VUID-VkSubpassDescription2-attachment-06919  
  If the `attachment` member of an element of `pColorAttachments` is not
  `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkSubpassDescription2-attachment-06920"
  id="VUID-VkSubpassDescription2-attachment-06920"></a>
  VUID-VkSubpassDescription2-attachment-06920  
  If the `attachment` member of an element of `pResolveAttachments` is
  not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkSubpassDescription2-attachment-06921"
  id="VUID-VkSubpassDescription2-attachment-06921"></a>
  VUID-VkSubpassDescription2-attachment-06921  
  If the `attachment` member of an element of `pInputAttachments` is not
  `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be
  `VK_IMAGE_LAYOUT_ATTACHMENT_OPTIMAL_KHR`

- <a href="#VUID-VkSubpassDescription2-attachment-06922"
  id="VUID-VkSubpassDescription2-attachment-06922"></a>
  VUID-VkSubpassDescription2-attachment-06922  
  If the `attachment` member of an element of `pColorAttachments` is not
  `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a href="#VUID-VkSubpassDescription2-attachment-06923"
  id="VUID-VkSubpassDescription2-attachment-06923"></a>
  VUID-VkSubpassDescription2-attachment-06923  
  If the `attachment` member of an element of `pResolveAttachments` is
  not `VK_ATTACHMENT_UNUSED`, its `layout` member **must** not be
  `VK_IMAGE_LAYOUT_READ_ONLY_OPTIMAL_KHR`

- <a href="#VUID-VkSubpassDescription2-attachment-06251"
  id="VUID-VkSubpassDescription2-attachment-06251"></a>
  VUID-VkSubpassDescription2-attachment-06251  
  If the `attachment` member of `pDepthStencilAttachment` is not
  `VK_ATTACHMENT_UNUSED` and its `pNext` chain includes a
  [VkAttachmentReferenceStencilLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReferenceStencilLayout.html)
  structure, the `layout` member of `pDepthStencilAttachment` **must**
  not be `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkSubpassDescription2-pipelineBindPoint-04953"
  id="VUID-VkSubpassDescription2-pipelineBindPoint-04953"></a>
  VUID-VkSubpassDescription2-pipelineBindPoint-04953  
  `pipelineBindPoint` **must** be `VK_PIPELINE_BIND_POINT_GRAPHICS` or
  `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI`

- <a href="#VUID-VkSubpassDescription2-colorAttachmentCount-03063"
  id="VUID-VkSubpassDescription2-colorAttachmentCount-03063"></a>
  VUID-VkSubpassDescription2-colorAttachmentCount-03063  
  `colorAttachmentCount` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxColorAttachments`

- <a href="#VUID-VkSubpassDescription2-loadOp-03064"
  id="VUID-VkSubpassDescription2-loadOp-03064"></a>
  VUID-VkSubpassDescription2-loadOp-03064  
  If the first use of an attachment in this render pass is as an input
  attachment, and the attachment is not also used as a color or
  depth/stencil attachment in the same subpass, then `loadOp` **must**
  not be `VK_ATTACHMENT_LOAD_OP_CLEAR`

- <a href="#VUID-VkSubpassDescription2-pResolveAttachments-03067"
  id="VUID-VkSubpassDescription2-pResolveAttachments-03067"></a>
  VUID-VkSubpassDescription2-pResolveAttachments-03067  
  If `pResolveAttachments` is not `NULL`, each resolve attachment that
  is not `VK_ATTACHMENT_UNUSED` **must** have a sample count of
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkSubpassDescription2-externalFormatResolve-09335"
  id="VUID-VkSubpassDescription2-externalFormatResolve-09335"></a>
  VUID-VkSubpassDescription2-externalFormatResolve-09335  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  is not enabled and `pResolveAttachments` is not `NULL`, for each
  resolve attachment that does not have the value
  `VK_ATTACHMENT_UNUSED`, the corresponding color attachment **must**
  not have the value `VK_ATTACHMENT_UNUSED`

- <a
  href="#VUID-VkSubpassDescription2-nullColorAttachmentWithExternalFormatResolve-09336"
  id="VUID-VkSubpassDescription2-nullColorAttachmentWithExternalFormatResolve-09336"></a>
  VUID-VkSubpassDescription2-nullColorAttachmentWithExternalFormatResolve-09336  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
  target="_blank"
  rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
  property is `VK_FALSE` and `pResolveAttachments` is not `NULL`, for
  each resolve attachment that has a format of `VK_FORMAT_UNDEFINED`,
  the corresponding color attachment **must** not have the value
  `VK_ATTACHMENT_UNUSED`

- <a
  href="#VUID-VkSubpassDescription2-nullColorAttachmentWithExternalFormatResolve-09337"
  id="VUID-VkSubpassDescription2-nullColorAttachmentWithExternalFormatResolve-09337"></a>
  VUID-VkSubpassDescription2-nullColorAttachmentWithExternalFormatResolve-09337  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
  target="_blank"
  rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
  property is `VK_TRUE` and `pResolveAttachments` is not `NULL`, for
  each resolve attachment that has a format of `VK_FORMAT_UNDEFINED`,
  the corresponding color attachment **must** have the value
  `VK_ATTACHMENT_UNUSED`

- <a href="#VUID-VkSubpassDescription2-externalFormatResolve-09338"
  id="VUID-VkSubpassDescription2-externalFormatResolve-09338"></a>
  VUID-VkSubpassDescription2-externalFormatResolve-09338  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  is not enabled and `pResolveAttachments` is not `NULL`, for each
  resolve attachment that is not `VK_ATTACHMENT_UNUSED`, the
  corresponding color attachment **must** not have a sample count of
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkSubpassDescription2-externalFormatResolve-09339"
  id="VUID-VkSubpassDescription2-externalFormatResolve-09339"></a>
  VUID-VkSubpassDescription2-externalFormatResolve-09339  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  is not enabled, each element of `pResolveAttachments` **must** have
  the same [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) as its corresponding color
  attachment

- <a
  href="#VUID-VkSubpassDescription2-multisampledRenderToSingleSampled-06869"
  id="VUID-VkSubpassDescription2-multisampledRenderToSingleSampled-06869"></a>
  VUID-VkSubpassDescription2-multisampledRenderToSingleSampled-06869  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multisampledRenderToSingleSampled"
  target="_blank"
  rel="noopener"><code>multisampledRenderToSingleSampled</code></a>
  feature is not enabled, all attachments in `pColorAttachments` that
  are not `VK_ATTACHMENT_UNUSED` **must** have the same sample count

- <a href="#VUID-VkSubpassDescription2-pInputAttachments-02897"
  id="VUID-VkSubpassDescription2-pInputAttachments-02897"></a>
  VUID-VkSubpassDescription2-pInputAttachments-02897  
  All attachments in `pInputAttachments` that are not
  `VK_ATTACHMENT_UNUSED` and any of the following is true:

  - the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
    target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
    feature is not enabled

  - the <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
    target="_blank"
    rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
    property is `VK_FALSE`

  - does not have a non-zero value of
    [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`

  **must** have image formats whose <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> contain
  at least `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` or
  `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkSubpassDescription2-pColorAttachments-02898"
  id="VUID-VkSubpassDescription2-pColorAttachments-02898"></a>
  VUID-VkSubpassDescription2-pColorAttachments-02898  
  All attachments in `pColorAttachments` that are not
  `VK_ATTACHMENT_UNUSED` **must** have image formats whose <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> contain
  `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-VkSubpassDescription2-pResolveAttachments-09343"
  id="VUID-VkSubpassDescription2-pResolveAttachments-09343"></a>
  VUID-VkSubpassDescription2-pResolveAttachments-09343  
  All attachments in `pResolveAttachments` that are not
  `VK_ATTACHMENT_UNUSED` and do not have an image format of
  `VK_FORMAT_UNDEFINED` **must** have image formats whose <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> contain
  `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-VkSubpassDescription2-pDepthStencilAttachment-02900"
  id="VUID-VkSubpassDescription2-pDepthStencilAttachment-02900"></a>
  VUID-VkSubpassDescription2-pDepthStencilAttachment-02900  
  If `pDepthStencilAttachment` is not `NULL` and the attachment is not
  `VK_ATTACHMENT_UNUSED` then it **must** have an image format whose <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> contain
  `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkSubpassDescription2-linearColorAttachment-06499"
  id="VUID-VkSubpassDescription2-linearColorAttachment-06499"></a>
  VUID-VkSubpassDescription2-linearColorAttachment-06499  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-linearColorAttachment"
  target="_blank" rel="noopener"><code>linearColorAttachment</code></a>
  feature is enabled and the image is created with
  `VK_IMAGE_TILING_LINEAR`, all attachments in `pInputAttachments` that
  are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> **must**
  contain `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`

- <a href="#VUID-VkSubpassDescription2-linearColorAttachment-06500"
  id="VUID-VkSubpassDescription2-linearColorAttachment-06500"></a>
  VUID-VkSubpassDescription2-linearColorAttachment-06500  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-linearColorAttachment"
  target="_blank" rel="noopener"><code>linearColorAttachment</code></a>
  feature is enabled and the image is created with
  `VK_IMAGE_TILING_LINEAR`, all attachments in `pColorAttachments` that
  are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> **must**
  contain `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`

- <a href="#VUID-VkSubpassDescription2-linearColorAttachment-06501"
  id="VUID-VkSubpassDescription2-linearColorAttachment-06501"></a>
  VUID-VkSubpassDescription2-linearColorAttachment-06501  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-linearColorAttachment"
  target="_blank" rel="noopener"><code>linearColorAttachment</code></a>
  feature is enabled and the image is created with
  `VK_IMAGE_TILING_LINEAR`, all attachments in `pResolveAttachments`
  that are not `VK_ATTACHMENT_UNUSED` **must** have image formats whose
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> **must**
  contain `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV`

- <a href="#VUID-VkSubpassDescription2-None-09456"
  id="VUID-VkSubpassDescription2-None-09456"></a>
  VUID-VkSubpassDescription2-None-09456  
  If either of the following is enabled:

  - The
    [`VK_AMD_mixed_attachment_samples`](VK_AMD_mixed_attachment_samples.html)
    extension

  - The
    [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
    extension

  all attachments in `pColorAttachments` that are not
  `VK_ATTACHMENT_UNUSED` **must** have a sample count that is smaller
  than or equal to the sample count of `pDepthStencilAttachment` if it
  is not `VK_ATTACHMENT_UNUSED`

- <a href="#VUID-VkSubpassDescription2-pNext-06870"
  id="VUID-VkSubpassDescription2-pNext-06870"></a>
  VUID-VkSubpassDescription2-pNext-06870  
  If the `pNext` chain includes a
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)
  structure with `multisampledRenderToSingleSampledEnable` equal to
  `VK_TRUE`, then all attachments in `pColorAttachments` and
  `pDepthStencilAttachment` that are not `VK_ATTACHMENT_UNUSED` **must**
  have a sample count that is either `VK_SAMPLE_COUNT_1_BIT` or equal to
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)::`rasterizationSamples`

- <a href="#VUID-VkSubpassDescription2-pNext-06871"
  id="VUID-VkSubpassDescription2-pNext-06871"></a>
  VUID-VkSubpassDescription2-pNext-06871  
  If the `pNext` chain includes a
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)
  structure with `multisampledRenderToSingleSampledEnable` equal to
  `VK_TRUE`, and `pDepthStencilAttachment` is not `NULL`, does not have
  the value `VK_ATTACHMENT_UNUSED`, and has a sample count of
  `VK_SAMPLE_COUNT_1_BIT`, the `pNext` chain **must** also include a
  [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionDepthStencilResolve.html)
  structure with `pDepthStencilResolveAttachment` that is either `NULL`
  or has the value `VK_ATTACHMENT_UNUSED`

- <a
  href="#VUID-VkSubpassDescription2-multisampledRenderToSingleSampled-06872"
  id="VUID-VkSubpassDescription2-multisampledRenderToSingleSampled-06872"></a>
  VUID-VkSubpassDescription2-multisampledRenderToSingleSampled-06872  
  All attachments in `pDepthStencilAttachment` or `pColorAttachments`
  that are not `VK_ATTACHMENT_UNUSED` **must** have the same sample
  count , if none of the following are enabled:

  - The
    [`VK_AMD_mixed_attachment_samples`](VK_AMD_mixed_attachment_samples.html)
    extension

  - The
    [`VK_NV_framebuffer_mixed_samples`](VK_NV_framebuffer_mixed_samples.html)
    extension

  - The <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multisampledRenderToSingleSampled"
    target="_blank"
    rel="noopener"><code>multisampledRenderToSingleSampled</code></a>
    feature,

- <a href="#VUID-VkSubpassDescription2-attachment-03073"
  id="VUID-VkSubpassDescription2-attachment-03073"></a>
  VUID-VkSubpassDescription2-attachment-03073  
  Each element of `pPreserveAttachments` **must** not be
  `VK_ATTACHMENT_UNUSED`

- <a href="#VUID-VkSubpassDescription2-pPreserveAttachments-03074"
  id="VUID-VkSubpassDescription2-pPreserveAttachments-03074"></a>
  VUID-VkSubpassDescription2-pPreserveAttachments-03074  
  Each element of `pPreserveAttachments` **must** not also be an element
  of any other member of the subpass description

- <a href="#VUID-VkSubpassDescription2-layout-02528"
  id="VUID-VkSubpassDescription2-layout-02528"></a>
  VUID-VkSubpassDescription2-layout-02528  
  If any attachment is used by more than one
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) member, then
  each use **must** use the same `layout`

- <a href="#VUID-VkSubpassDescription2-flags-03076"
  id="VUID-VkSubpassDescription2-flags-03076"></a>
  VUID-VkSubpassDescription2-flags-03076  
  If `flags` includes
  `VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX`, it **must**
  also include `VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX`

- <a href="#VUID-VkSubpassDescription2-attachment-02799"
  id="VUID-VkSubpassDescription2-attachment-02799"></a>
  VUID-VkSubpassDescription2-attachment-02799  
  If the `attachment` member of any element of `pInputAttachments` is
  not `VK_ATTACHMENT_UNUSED`, then the `aspectMask` member **must** be a
  valid combination of
  [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlagBits.html)

- <a href="#VUID-VkSubpassDescription2-attachment-02800"
  id="VUID-VkSubpassDescription2-attachment-02800"></a>
  VUID-VkSubpassDescription2-attachment-02800  
  If the `attachment` member of any element of `pInputAttachments` is
  not `VK_ATTACHMENT_UNUSED`, then the `aspectMask` member **must** not
  be `0`

- <a href="#VUID-VkSubpassDescription2-attachment-02801"
  id="VUID-VkSubpassDescription2-attachment-02801"></a>
  VUID-VkSubpassDescription2-attachment-02801  
  If the `attachment` member of any element of `pInputAttachments` is
  not `VK_ATTACHMENT_UNUSED`, then the `aspectMask` member **must** not
  include `VK_IMAGE_ASPECT_METADATA_BIT`

- <a href="#VUID-VkSubpassDescription2-attachment-04563"
  id="VUID-VkSubpassDescription2-attachment-04563"></a>
  VUID-VkSubpassDescription2-attachment-04563  
  If the `attachment` member of any element of `pInputAttachments` is
  not `VK_ATTACHMENT_UNUSED`, then the `aspectMask` member **must** not
  include `VK_IMAGE_ASPECT_MEMORY_PLANE`*`_i_`*`BIT_EXT` for any index
  *i*

- <a href="#VUID-VkSubpassDescription2-pDepthStencilAttachment-04440"
  id="VUID-VkSubpassDescription2-pDepthStencilAttachment-04440"></a>
  VUID-VkSubpassDescription2-pDepthStencilAttachment-04440  
  An attachment **must** not be used in both `pDepthStencilAttachment`
  and `pColorAttachments`

- <a href="#VUID-VkSubpassDescription2-multiview-06558"
  id="VUID-VkSubpassDescription2-multiview-06558"></a>
  VUID-VkSubpassDescription2-multiview-06558  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview"
  target="_blank" rel="noopener"><code>multiview</code></a> feature is
  not enabled, `viewMask` **must** be `0`

- <a href="#VUID-VkSubpassDescription2-viewMask-06706"
  id="VUID-VkSubpassDescription2-viewMask-06706"></a>
  VUID-VkSubpassDescription2-viewMask-06706  
  The index of the most significant bit in `viewMask` **must** be less
  than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxMultiviewViewCount"
  target="_blank" rel="noopener"><code>maxMultiviewViewCount</code></a>

- <a href="#VUID-VkSubpassDescription2-externalFormatResolve-09344"
  id="VUID-VkSubpassDescription2-externalFormatResolve-09344"></a>
  VUID-VkSubpassDescription2-externalFormatResolve-09344  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  is enabled, `pResolveAttachments` is not `NULL`, and
  `colorAttachmentCount` is not `1`, any element of
  `pResolveAttachments` that is not `VK_ATTACHMENT_UNUSED`, **must** not
  have a format of `VK_FORMAT_UNDEFINED`

- <a href="#VUID-VkSubpassDescription2-externalFormatResolve-09345"
  id="VUID-VkSubpassDescription2-externalFormatResolve-09345"></a>
  VUID-VkSubpassDescription2-externalFormatResolve-09345  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  is enabled, `pResolveAttachments` is not `NULL`, any element of
  `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED` and has a format
  of `VK_FORMAT_UNDEFINED`, and the corresponding element of
  `pColorAttachments` is not `VK_ATTACHMENT_UNUSED`, the color
  attachment **must** have a `samples` value of `1`

- <a href="#VUID-VkSubpassDescription2-externalFormatResolve-09346"
  id="VUID-VkSubpassDescription2-externalFormatResolve-09346"></a>
  VUID-VkSubpassDescription2-externalFormatResolve-09346  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  is enabled, `pResolveAttachments` is not `NULL`, and any element of
  `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED` and has a format
  of `VK_FORMAT_UNDEFINED`, `viewMask` **must** be `0`

- <a href="#VUID-VkSubpassDescription2-externalFormatResolve-09347"
  id="VUID-VkSubpassDescription2-externalFormatResolve-09347"></a>
  VUID-VkSubpassDescription2-externalFormatResolve-09347  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  is enabled, `pResolveAttachments` is not `NULL`, and any element of
  `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED` and has a format
  of `VK_FORMAT_UNDEFINED`,
  [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateAttachmentInfoKHR.html)::`pFragmentShadingRateAttachment`
  **must** either be `NULL` or a
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structure with
  an `attachment` value of `VK_ATTACHMENT_UNUSED`

- <a href="#VUID-VkSubpassDescription2-externalFormatResolve-09348"
  id="VUID-VkSubpassDescription2-externalFormatResolve-09348"></a>
  VUID-VkSubpassDescription2-externalFormatResolve-09348  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-externalFormatResolve"
  target="_blank" rel="noopener"><code>externalFormatResolve</code></a>
  is enabled, `pResolveAttachments` is not `NULL`, and any element of
  `pResolveAttachments` is not `VK_ATTACHMENT_UNUSED` and has a format
  of `VK_FORMAT_UNDEFINED`, elements of `pInputAttachments` referencing
  either a color attachment or resolve attachment used in this subpass
  **must** not include `VK_IMAGE_ASPECT_PLANE`*`_i_`*`BIT` for any index
  *i* in its `aspectMask`

Valid Usage (Implicit)

- <a href="#VUID-VkSubpassDescription2-sType-sType"
  id="VUID-VkSubpassDescription2-sType-sType"></a>
  VUID-VkSubpassDescription2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_2`

- <a href="#VUID-VkSubpassDescription2-pNext-pNext"
  id="VUID-VkSubpassDescription2-pNext-pNext"></a>
  VUID-VkSubpassDescription2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateAttachmentInfoKHR.html),
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html),
  [VkRenderPassCreationControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationControlEXT.html),
  [VkRenderPassSubpassFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSubpassFeedbackCreateInfoEXT.html),
  or
  [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionDepthStencilResolve.html)

- <a href="#VUID-VkSubpassDescription2-sType-unique"
  id="VUID-VkSubpassDescription2-sType-unique"></a>
  VUID-VkSubpassDescription2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkSubpassDescription2-flags-parameter"
  id="VUID-VkSubpassDescription2-flags-parameter"></a>
  VUID-VkSubpassDescription2-flags-parameter  
  `flags` **must** be a valid combination of
  [VkSubpassDescriptionFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionFlagBits.html)
  values

- <a href="#VUID-VkSubpassDescription2-pipelineBindPoint-parameter"
  id="VUID-VkSubpassDescription2-pipelineBindPoint-parameter"></a>
  VUID-VkSubpassDescription2-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a href="#VUID-VkSubpassDescription2-pInputAttachments-parameter"
  id="VUID-VkSubpassDescription2-pInputAttachments-parameter"></a>
  VUID-VkSubpassDescription2-pInputAttachments-parameter  
  If `inputAttachmentCount` is not `0`, `pInputAttachments` **must** be
  a valid pointer to an array of `inputAttachmentCount` valid
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structures

- <a href="#VUID-VkSubpassDescription2-pColorAttachments-parameter"
  id="VUID-VkSubpassDescription2-pColorAttachments-parameter"></a>
  VUID-VkSubpassDescription2-pColorAttachments-parameter  
  If `colorAttachmentCount` is not `0`, `pColorAttachments` **must** be
  a valid pointer to an array of `colorAttachmentCount` valid
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structures

- <a href="#VUID-VkSubpassDescription2-pResolveAttachments-parameter"
  id="VUID-VkSubpassDescription2-pResolveAttachments-parameter"></a>
  VUID-VkSubpassDescription2-pResolveAttachments-parameter  
  If `colorAttachmentCount` is not `0`, and `pResolveAttachments` is not
  `NULL`, `pResolveAttachments` **must** be a valid pointer to an array
  of `colorAttachmentCount` valid
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structures

- <a href="#VUID-VkSubpassDescription2-pDepthStencilAttachment-parameter"
  id="VUID-VkSubpassDescription2-pDepthStencilAttachment-parameter"></a>
  VUID-VkSubpassDescription2-pDepthStencilAttachment-parameter  
  If `pDepthStencilAttachment` is not `NULL`, `pDepthStencilAttachment`
  **must** be a valid pointer to a valid
  [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html) structure

- <a href="#VUID-VkSubpassDescription2-pPreserveAttachments-parameter"
  id="VUID-VkSubpassDescription2-pPreserveAttachments-parameter"></a>
  VUID-VkSubpassDescription2-pPreserveAttachments-parameter  
  If `preserveAttachmentCount` is not `0`, `pPreserveAttachments`
  **must** be a valid pointer to an array of `preserveAttachmentCount`
  `uint32_t` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentReference2.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html),
[VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubpassDescriptionFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassDescription2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
