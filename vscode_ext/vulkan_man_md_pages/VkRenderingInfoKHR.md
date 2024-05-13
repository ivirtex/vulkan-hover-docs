# VkRenderingInfo(3) Manual Page

## Name

VkRenderingInfo - Structure specifying render pass instance begin info



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderingInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkRenderingInfo {
    VkStructureType                     sType;
    const void*                         pNext;
    VkRenderingFlags                    flags;
    VkRect2D                            renderArea;
    uint32_t                            layerCount;
    uint32_t                            viewMask;
    uint32_t                            colorAttachmentCount;
    const VkRenderingAttachmentInfo*    pColorAttachments;
    const VkRenderingAttachmentInfo*    pDepthAttachment;
    const VkRenderingAttachmentInfo*    pStencilAttachment;
} VkRenderingInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_dynamic_rendering, VK_QCOM_tile_properties with VK_KHR_dynamic_rendering or VK_VERSION_1_3
typedef VkRenderingInfo VkRenderingInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFlagBits.html).

- `renderArea` is the render area that is affected by the render pass
  instance.

- `layerCount` is the number of layers rendered to in each attachment
  when `viewMask` is `0`.

- `viewMask` is the view mask indicating the indices of attachment
  layers that will be rendered when it is not `0`.

- `colorAttachmentCount` is the number of elements in
  `pColorAttachments`.

- `pColorAttachments` is a pointer to an array of `colorAttachmentCount`
  [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html) structures
  describing any color attachments used.

- `pDepthAttachment` is a pointer to a
  [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html) structure
  describing a depth attachment.

- `pStencilAttachment` is a pointer to a
  [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html) structure
  describing a stencil attachment.

## <a href="#_description" class="anchor"></a>Description

If `viewMask` is not `0`, multiview is enabled.

If there is an instance of
[VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
included in the `pNext` chain and its `deviceRenderAreaCount` member is
not `0`, then `renderArea` is ignored, and the render area is defined
per-device by that structure.

If multiview is enabled, and the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-per-view-render-areas"
target="_blank"
rel="noopener"><code>multiviewPerViewRenderAreas</code></a> feature is
enabled, and there is an instance of
[VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html)
included in the `pNext` chain with `perViewRenderAreaCount` not equal to
`0`, then the elements of
[VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html)::`pPerViewRenderAreas`
override `renderArea` and define a render area for each view. In this
case, `renderArea` **must** be set to an area at least as large as the
union of all the per-view render areas.

Each element of the `pColorAttachments` array corresponds to an output
location in the shader, i.e. if the shader declares an output variable
decorated with a `Location` value of **X**, then it uses the attachment
provided in `pColorAttachments`\[**X**\]. If the `imageView` member of
any element of `pColorAttachments` is
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and `resolveMode` is not
`VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`, writes to the
corresponding location by a fragment are discarded.

Valid Usage

- <a href="#VUID-VkRenderingInfo-viewMask-06069"
  id="VUID-VkRenderingInfo-viewMask-06069"></a>
  VUID-VkRenderingInfo-viewMask-06069  
  If `viewMask` is `0`, `layerCount` **must** not be `0`

- <a href="#VUID-VkRenderingInfo-multisampledRenderToSingleSampled-06857"
  id="VUID-VkRenderingInfo-multisampledRenderToSingleSampled-06857"></a>
  VUID-VkRenderingInfo-multisampledRenderToSingleSampled-06857  
  `imageView` members of `pDepthAttachment`, `pStencilAttachment`, and
  elements of `pColorAttachments` that are not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **must** have been created with
  the same `sampleCount` , if none of the following are enabled:

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

- <a href="#VUID-VkRenderingInfo-imageView-09429"
  id="VUID-VkRenderingInfo-imageView-09429"></a>
  VUID-VkRenderingInfo-imageView-09429  
  `imageView` members of elements of `pColorAttachments` that are not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **must** have been created with
  the same `sampleCount` , if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multisampledRenderToSingleSampled"
  target="_blank"
  rel="noopener"><code>multisampledRenderToSingleSampled</code></a>
  feature is not enabled

- <a href="#VUID-VkRenderingInfo-None-08994"
  id="VUID-VkRenderingInfo-None-08994"></a>
  VUID-VkRenderingInfo-None-08994  
  If
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)::`deviceRenderAreaCount`
  is 0, `renderArea.extent.width` **must** be greater than 0

- <a href="#VUID-VkRenderingInfo-None-08995"
  id="VUID-VkRenderingInfo-None-08995"></a>
  VUID-VkRenderingInfo-None-08995  
  If
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)::`deviceRenderAreaCount`
  is 0, `renderArea.extent.height` **must** be greater than 0

- <a href="#VUID-VkRenderingInfo-imageView-06858"
  id="VUID-VkRenderingInfo-imageView-06858"></a>
  VUID-VkRenderingInfo-imageView-06858  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#subpass-multisampledrendertosinglesampled"
  target="_blank" rel="noopener">multisampled-render-to-single-sampled</a>
  is enabled, then all attachments referenced by `imageView` members of
  `pDepthAttachment`, `pStencilAttachment`, and elements of
  `pColorAttachments` that are not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  **must** have a sample count that is either `VK_SAMPLE_COUNT_1_BIT` or
  equal to
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)::`rasterizationSamples`

- <a href="#VUID-VkRenderingInfo-imageView-06859"
  id="VUID-VkRenderingInfo-imageView-06859"></a>
  VUID-VkRenderingInfo-imageView-06859  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#subpass-multisampledrendertosinglesampled"
  target="_blank" rel="noopener">multisampled-render-to-single-sampled</a>
  is enabled, then all attachments referenced by `imageView` members of
  `pDepthAttachment`, `pStencilAttachment`, and elements of
  `pColorAttachments` that are not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  and have a sample count of `VK_SAMPLE_COUNT_1_BIT` **must** have been
  created with
  `VK_IMAGE_CREATE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_BIT_EXT` in
  their [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags`

- <a href="#VUID-VkRenderingInfo-pNext-06077"
  id="VUID-VkRenderingInfo-pNext-06077"></a>
  VUID-VkRenderingInfo-pNext-06077  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0,
  `renderArea.offset.x` **must** be greater than or equal to 0

- <a href="#VUID-VkRenderingInfo-pNext-06078"
  id="VUID-VkRenderingInfo-pNext-06078"></a>
  VUID-VkRenderingInfo-pNext-06078  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0,
  `renderArea.offset.y` **must** be greater than or equal to 0

- <a href="#VUID-VkRenderingInfo-pNext-07815"
  id="VUID-VkRenderingInfo-pNext-07815"></a>
  VUID-VkRenderingInfo-pNext-07815  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0, the sum of
  `renderArea.extent.width` and `renderArea.offset.x` **must** be less
  than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFramebufferWidth"
  target="_blank" rel="noopener"><code>maxFramebufferWidth</code></a>

- <a href="#VUID-VkRenderingInfo-pNext-07816"
  id="VUID-VkRenderingInfo-pNext-07816"></a>
  VUID-VkRenderingInfo-pNext-07816  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0, the sum of
  `renderArea.extent.height` and `renderArea.offset.y` **must** be less
  than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFramebufferWidth"
  target="_blank" rel="noopener"><code>maxFramebufferHeight</code></a>

- <a href="#VUID-VkRenderingInfo-pNext-06079"
  id="VUID-VkRenderingInfo-pNext-06079"></a>
  VUID-VkRenderingInfo-pNext-06079  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0, the width of the
  `imageView` member of any element of `pColorAttachments`,
  `pDepthAttachment`, or `pStencilAttachment` that is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **must** be greater than or
  equal to `renderArea.offset.x` + `renderArea.extent.width`

- <a href="#VUID-VkRenderingInfo-pNext-06080"
  id="VUID-VkRenderingInfo-pNext-06080"></a>
  VUID-VkRenderingInfo-pNext-06080  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0, the height of the
  `imageView` member of any element of `pColorAttachments`,
  `pDepthAttachment`, or `pStencilAttachment` that is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **must** be greater than or
  equal to `renderArea.offset.y` + `renderArea.extent.height`

- <a href="#VUID-VkRenderingInfo-pNext-06083"
  id="VUID-VkRenderingInfo-pNext-06083"></a>
  VUID-VkRenderingInfo-pNext-06083  
  If the `pNext` chain contains
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html),
  the width of the `imageView` member of any element of
  `pColorAttachments`, `pDepthAttachment`, or `pStencilAttachment` that
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **must** be greater than
  or equal to the sum of the `offset.x` and `extent.width` members of
  each element of `pDeviceRenderAreas`

- <a href="#VUID-VkRenderingInfo-pNext-06084"
  id="VUID-VkRenderingInfo-pNext-06084"></a>
  VUID-VkRenderingInfo-pNext-06084  
  If the `pNext` chain contains
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html),
  the height of the `imageView` member of any element of
  `pColorAttachments`, `pDepthAttachment`, or `pStencilAttachment` that
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **must** be greater than
  or equal to the sum of the `offset.y` and `extent.height` members of
  each element of `pDeviceRenderAreas`

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-06085"
  id="VUID-VkRenderingInfo-pDepthAttachment-06085"></a>
  VUID-VkRenderingInfo-pDepthAttachment-06085  
  If neither `pDepthAttachment` or `pStencilAttachment` are `NULL` and
  the `imageView` member of either structure is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `imageView` member of each
  structure **must** be the same

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-06086"
  id="VUID-VkRenderingInfo-pDepthAttachment-06086"></a>
  VUID-VkRenderingInfo-pDepthAttachment-06086  
  If neither `pDepthAttachment` or `pStencilAttachment` are `NULL`, and
  the `resolveMode` member of each is not `VK_RESOLVE_MODE_NONE`, the
  `resolveImageView` member of each structure **must** be the same

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-06087"
  id="VUID-VkRenderingInfo-colorAttachmentCount-06087"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-06087  
  If `colorAttachmentCount` is not `0` and the `imageView` member of an
  element of `pColorAttachments` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), that `imageView` **must** have
  been created with `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-09476"
  id="VUID-VkRenderingInfo-colorAttachmentCount-09476"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-09476  
  If `colorAttachmentCount` is not `0` and there is an element of
  `pColorAttachments` with either its `resolveMode` member set to
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`, or its
  `imageView` member not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and its
  `resolveMode` member not set to `VK_RESOLVE_MODE_NONE`, the
  `resolveImageView` member of that element of `pColorAttachments`
  **must** have been created with `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-06547"
  id="VUID-VkRenderingInfo-pDepthAttachment-06547"></a>
  VUID-VkRenderingInfo-pDepthAttachment-06547  
  If `pDepthAttachment` is not `NULL` and `pDepthAttachment->imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pDepthAttachment->imageView` **must** have been created with a format
  that includes a depth component

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-06088"
  id="VUID-VkRenderingInfo-pDepthAttachment-06088"></a>
  VUID-VkRenderingInfo-pDepthAttachment-06088  
  If `pDepthAttachment` is not `NULL` and `pDepthAttachment->imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pDepthAttachment->imageView` **must** have been created with
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-09477"
  id="VUID-VkRenderingInfo-pDepthAttachment-09477"></a>
  VUID-VkRenderingInfo-pDepthAttachment-09477  
  If `pDepthAttachment` is not `NULL` and
  `pDepthAttachment->resolveMode` is not `VK_RESOLVE_MODE_NONE`,
  `pDepthAttachment->resolveImageView` **must** have been created with
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-06548"
  id="VUID-VkRenderingInfo-pStencilAttachment-06548"></a>
  VUID-VkRenderingInfo-pStencilAttachment-06548  
  If `pStencilAttachment` is not `NULL` and
  `pStencilAttachment->imageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pStencilAttachment->imageView`
  **must** have been created with a format that includes a stencil
  aspect

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-06089"
  id="VUID-VkRenderingInfo-pStencilAttachment-06089"></a>
  VUID-VkRenderingInfo-pStencilAttachment-06089  
  If `pStencilAttachment` is not `NULL` and
  `pStencilAttachment->imageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pStencilAttachment->imageView`
  **must** have been created with a stencil usage including
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-09478"
  id="VUID-VkRenderingInfo-pStencilAttachment-09478"></a>
  VUID-VkRenderingInfo-pStencilAttachment-09478  
  If `pStencilAttachment` is not `NULL` and
  `pStencilAttachment->resolveMode` is not `VK_RESOLVE_MODE_NONE`,
  `pStencilAttachment->resolveImageView` **must** have been created with
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-06090"
  id="VUID-VkRenderingInfo-colorAttachmentCount-06090"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-06090  
  If `colorAttachmentCount` is not `0` and the `imageView` member of an
  element of `pColorAttachments` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `layout` member of that
  element of `pColorAttachments` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-06091"
  id="VUID-VkRenderingInfo-colorAttachmentCount-06091"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-06091  
  If `colorAttachmentCount` is not `0` and the `imageView` member of an
  element of `pColorAttachments` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), if the `resolveMode` member of
  that element of `pColorAttachments` is not `VK_RESOLVE_MODE_NONE`, its
  `resolveImageLayout` member **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-06092"
  id="VUID-VkRenderingInfo-pDepthAttachment-06092"></a>
  VUID-VkRenderingInfo-pDepthAttachment-06092  
  If `pDepthAttachment` is not `NULL` and `pDepthAttachment->imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pDepthAttachment->layout` **must** not be
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-06093"
  id="VUID-VkRenderingInfo-pDepthAttachment-06093"></a>
  VUID-VkRenderingInfo-pDepthAttachment-06093  
  If `pDepthAttachment` is not `NULL`, `pDepthAttachment->imageView` is
  not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  `pDepthAttachment->resolveMode` is not `VK_RESOLVE_MODE_NONE`,
  `pDepthAttachment->resolveImageLayout` **must** not be
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-06094"
  id="VUID-VkRenderingInfo-pStencilAttachment-06094"></a>
  VUID-VkRenderingInfo-pStencilAttachment-06094  
  If `pStencilAttachment` is not `NULL` and
  `pStencilAttachment->imageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pStencilAttachment->layout`
  **must** not be `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-06095"
  id="VUID-VkRenderingInfo-pStencilAttachment-06095"></a>
  VUID-VkRenderingInfo-pStencilAttachment-06095  
  If `pStencilAttachment` is not `NULL`, `pStencilAttachment->imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  `pStencilAttachment->resolveMode` is not `VK_RESOLVE_MODE_NONE`,
  `pStencilAttachment->resolveImageLayout` **must** not be
  `VK_IMAGE_LAYOUT_COLOR_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-06096"
  id="VUID-VkRenderingInfo-colorAttachmentCount-06096"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-06096  
  If `colorAttachmentCount` is not `0` and the `imageView` member of an
  element of `pColorAttachments` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `layout` member of that
  element of `pColorAttachments` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-06097"
  id="VUID-VkRenderingInfo-colorAttachmentCount-06097"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-06097  
  If `colorAttachmentCount` is not `0` and the `imageView` member of an
  element of `pColorAttachments` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), if the `resolveMode` member of
  that element of `pColorAttachments` is not `VK_RESOLVE_MODE_NONE`, its
  `resolveImageLayout` member **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-06098"
  id="VUID-VkRenderingInfo-pDepthAttachment-06098"></a>
  VUID-VkRenderingInfo-pDepthAttachment-06098  
  If `pDepthAttachment` is not `NULL`, `pDepthAttachment->imageView` is
  not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  `pDepthAttachment->resolveMode` is not `VK_RESOLVE_MODE_NONE`,
  `pDepthAttachment->resolveImageLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-06099"
  id="VUID-VkRenderingInfo-pStencilAttachment-06099"></a>
  VUID-VkRenderingInfo-pStencilAttachment-06099  
  If `pStencilAttachment` is not `NULL`, `pStencilAttachment->imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  `pStencilAttachment->resolveMode` is not `VK_RESOLVE_MODE_NONE`,
  `pStencilAttachment->resolveImageLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-06100"
  id="VUID-VkRenderingInfo-colorAttachmentCount-06100"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-06100  
  If `colorAttachmentCount` is not `0` and the `imageView` member of an
  element of `pColorAttachments` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `layout` member of that
  element of `pColorAttachments` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-06101"
  id="VUID-VkRenderingInfo-colorAttachmentCount-06101"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-06101  
  If `colorAttachmentCount` is not `0` and the `imageView` member of an
  element of `pColorAttachments` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), if the `resolveMode` member of
  that element of `pColorAttachments` is not `VK_RESOLVE_MODE_NONE`, its
  `resolveImageLayout` member **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL`,
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`,
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL`, or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-07732"
  id="VUID-VkRenderingInfo-pDepthAttachment-07732"></a>
  VUID-VkRenderingInfo-pDepthAttachment-07732  
  If `pDepthAttachment` is not `NULL` and `pDepthAttachment->imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pDepthAttachment->layout` **must** not be
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-07733"
  id="VUID-VkRenderingInfo-pDepthAttachment-07733"></a>
  VUID-VkRenderingInfo-pDepthAttachment-07733  
  If `pDepthAttachment` is not `NULL`, `pDepthAttachment->imageView` is
  not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  `pDepthAttachment->resolveMode` is not `VK_RESOLVE_MODE_NONE`,
  `pDepthAttachment->resolveImageLayout` **must** not be
  `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-07734"
  id="VUID-VkRenderingInfo-pStencilAttachment-07734"></a>
  VUID-VkRenderingInfo-pStencilAttachment-07734  
  If `pStencilAttachment` is not `NULL` and
  `pStencilAttachment->imageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pStencilAttachment->layout`
  **must** not be `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-07735"
  id="VUID-VkRenderingInfo-pStencilAttachment-07735"></a>
  VUID-VkRenderingInfo-pStencilAttachment-07735  
  If `pStencilAttachment` is not `NULL`, `pStencilAttachment->imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  `pStencilAttachment->resolveMode` is not `VK_RESOLVE_MODE_NONE`,
  `pStencilAttachment->resolveImageLayout` **must** not be
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-06102"
  id="VUID-VkRenderingInfo-pDepthAttachment-06102"></a>
  VUID-VkRenderingInfo-pDepthAttachment-06102  
  If `pDepthAttachment` is not `NULL` and `pDepthAttachment->imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pDepthAttachment->resolveMode` **must** be one of the bits set in
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`supportedDepthResolveModes`

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-06103"
  id="VUID-VkRenderingInfo-pStencilAttachment-06103"></a>
  VUID-VkRenderingInfo-pStencilAttachment-06103  
  If `pStencilAttachment` is not `NULL` and
  `pStencilAttachment->imageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pStencilAttachment->resolveMode` **must** be one of the bits set in
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`supportedStencilResolveModes`

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-06104"
  id="VUID-VkRenderingInfo-pDepthAttachment-06104"></a>
  VUID-VkRenderingInfo-pDepthAttachment-06104  
  If `pDepthAttachment` or `pStencilAttachment` are both not `NULL`,
  `pDepthAttachment->imageView` and `pStencilAttachment->imageView` are
  both not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolveNone`
  is `VK_FALSE`, the `resolveMode` of both structures **must** be the
  same value

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-06105"
  id="VUID-VkRenderingInfo-pDepthAttachment-06105"></a>
  VUID-VkRenderingInfo-pDepthAttachment-06105  
  If `pDepthAttachment` or `pStencilAttachment` are both not `NULL`,
  `pDepthAttachment->imageView` and `pStencilAttachment->imageView` are
  both not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html)::`independentResolve`
  is `VK_FALSE`, and the `resolveMode` of neither structure is
  `VK_RESOLVE_MODE_NONE`, the `resolveMode` of both structures **must**
  be the same value

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-06106"
  id="VUID-VkRenderingInfo-colorAttachmentCount-06106"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-06106  
  `colorAttachmentCount` **must** be less than or equal to
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxColorAttachments`

- <a href="#VUID-VkRenderingInfo-imageView-06107"
  id="VUID-VkRenderingInfo-imageView-06107"></a>
  VUID-VkRenderingInfo-imageView-06107  
  If the `imageView` member of a
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMapNonSubsampledImages"
  target="_blank"
  rel="noopener"><code>fragmentDensityMapNonSubsampledImages</code></a>
  feature is not enabled, valid `imageView` and `resolveImageView`
  members of `pDepthAttachment`, `pStencilAttachment`, and each element
  of `pColorAttachments` **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)
  created with `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

- <a href="#VUID-VkRenderingInfo-imageView-06108"
  id="VUID-VkRenderingInfo-imageView-06108"></a>
  VUID-VkRenderingInfo-imageView-06108  
  If the `imageView` member of a
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and `viewMask` is not `0`,
  `imageView` **must** have a `layerCount` greater than or equal to the
  index of the most significant bit in `viewMask`

- <a href="#VUID-VkRenderingInfo-imageView-06109"
  id="VUID-VkRenderingInfo-imageView-06109"></a>
  VUID-VkRenderingInfo-imageView-06109  
  If the `imageView` member of a
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and `viewMask` is `0`,
  `imageView` **must** have a `layerCount` equal to `1`

- <a href="#VUID-VkRenderingInfo-pNext-06112"
  id="VUID-VkRenderingInfo-pNext-06112"></a>
  VUID-VkRenderingInfo-pNext-06112  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0 and the
  `imageView` member of a
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `imageView` **must** have a
  width greater than or equal to
  ⌈maxFragmentDensityTexelSizewidth​renderAreax​+renderAreawidth​​⌉

- <a href="#VUID-VkRenderingInfo-pNext-06114"
  id="VUID-VkRenderingInfo-pNext-06114"></a>
  VUID-VkRenderingInfo-pNext-06114  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0 and the
  `imageView` member of a
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `imageView` **must** have a
  height greater than or equal to
  ⌈maxFragmentDensityTexelSizeheight​renderAreay​+renderAreaheight​​⌉

- <a href="#VUID-VkRenderingInfo-pNext-06113"
  id="VUID-VkRenderingInfo-pNext-06113"></a>
  VUID-VkRenderingInfo-pNext-06113  
  If the `pNext` chain contains a
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  structure, its `deviceRenderAreaCount` member is not 0, and the
  `imageView` member of a
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `imageView` **must** have a
  width greater than or equal to
  ⌈maxFragmentDensityTexelSizewidth​pDeviceRenderAreasx​+pDeviceRenderAreaswidth​​⌉
  for each element of `pDeviceRenderAreas`

- <a href="#VUID-VkRenderingInfo-pNext-06115"
  id="VUID-VkRenderingInfo-pNext-06115"></a>
  VUID-VkRenderingInfo-pNext-06115  
  If the `pNext` chain contains a
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  structure, its `deviceRenderAreaCount` member is not 0, and the
  `imageView` member of a
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `imageView` **must** have a
  height greater than or equal to
  ⌈maxFragmentDensityTexelSizeheight​pDeviceRenderAreasy​+pDeviceRenderAreasheight​​⌉
  for each element of `pDeviceRenderAreas`

- <a href="#VUID-VkRenderingInfo-imageView-06116"
  id="VUID-VkRenderingInfo-imageView-06116"></a>
  VUID-VkRenderingInfo-imageView-06116  
  If the `imageView` member of a
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it **must** not be equal to the
  `imageView` or `resolveImageView` member of `pDepthAttachment`,
  `pStencilAttachment`, or any element of `pColorAttachments`

- <a href="#VUID-VkRenderingInfo-pNext-06119"
  id="VUID-VkRenderingInfo-pNext-06119"></a>
  VUID-VkRenderingInfo-pNext-06119  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0 and the
  `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `imageView` **must** have a
  width greater than or equal to
  ⌈shadingRateAttachmentTexelSizewidth​renderAreax​+renderAreawidth​​⌉

- <a href="#VUID-VkRenderingInfo-pNext-06121"
  id="VUID-VkRenderingInfo-pNext-06121"></a>
  VUID-VkRenderingInfo-pNext-06121  
  If the `pNext` chain does not contain
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  or its `deviceRenderAreaCount` member is equal to 0 and the
  `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `imageView` **must** have a
  height greater than or equal to
  ⌈shadingRateAttachmentTexelSizeheight​renderAreay​+renderAreaheight​​⌉

- <a href="#VUID-VkRenderingInfo-pNext-06120"
  id="VUID-VkRenderingInfo-pNext-06120"></a>
  VUID-VkRenderingInfo-pNext-06120  
  If the `pNext` chain contains a
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  structure, its `deviceRenderAreaCount` member is not 0, and the
  `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `imageView` **must** have a
  width greater than or equal to
  ⌈shadingRateAttachmentTexelSizewidth​pDeviceRenderAreasx​+pDeviceRenderAreaswidth​​⌉
  for each element of `pDeviceRenderAreas`

- <a href="#VUID-VkRenderingInfo-pNext-06122"
  id="VUID-VkRenderingInfo-pNext-06122"></a>
  VUID-VkRenderingInfo-pNext-06122  
  If the `pNext` chain contains a
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html)
  structure, its `deviceRenderAreaCount` member is not 0, and the
  `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `imageView` **must** have a
  height greater than or equal to
  ⌈shadingRateAttachmentTexelSizeheight​pDeviceRenderAreasy​+pDeviceRenderAreasheight​​⌉
  for each element of `pDeviceRenderAreas`

- <a href="#VUID-VkRenderingInfo-layerCount-07817"
  id="VUID-VkRenderingInfo-layerCount-07817"></a>
  VUID-VkRenderingInfo-layerCount-07817  
  `layerCount` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFramebufferLayers"
  target="_blank" rel="noopener"><code>maxFramebufferLayers</code></a>

- <a href="#VUID-VkRenderingInfo-imageView-06123"
  id="VUID-VkRenderingInfo-imageView-06123"></a>
  VUID-VkRenderingInfo-imageView-06123  
  If the `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and `viewMask` is `0`,
  `imageView` **must** have a `layerCount` that is either equal to `1`
  or greater than or equal to `layerCount`

- <a href="#VUID-VkRenderingInfo-imageView-06124"
  id="VUID-VkRenderingInfo-imageView-06124"></a>
  VUID-VkRenderingInfo-imageView-06124  
  If the `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and `viewMask` is not `0`,
  `imageView` **must** have a `layerCount` that either equal to `1` or
  greater than or equal to the index of the most significant bit in
  `viewMask`

- <a href="#VUID-VkRenderingInfo-imageView-06125"
  id="VUID-VkRenderingInfo-imageView-06125"></a>
  VUID-VkRenderingInfo-imageView-06125  
  If the `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it **must** not be equal to the
  `imageView` or `resolveImageView` member of `pDepthAttachment`,
  `pStencilAttachment`, or any element of `pColorAttachments`

- <a href="#VUID-VkRenderingInfo-imageView-06126"
  id="VUID-VkRenderingInfo-imageView-06126"></a>
  VUID-VkRenderingInfo-imageView-06126  
  If the `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it **must** not be equal to the
  `imageView` member of a
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)
  structure included in the `pNext` chain

- <a href="#VUID-VkRenderingInfo-multiview-06127"
  id="VUID-VkRenderingInfo-multiview-06127"></a>
  VUID-VkRenderingInfo-multiview-06127  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview"
  target="_blank" rel="noopener"><code>multiview</code></a> feature is
  not enabled, `viewMask` **must** be `0`

- <a href="#VUID-VkRenderingInfo-viewMask-06128"
  id="VUID-VkRenderingInfo-viewMask-06128"></a>
  VUID-VkRenderingInfo-viewMask-06128  
  The index of the most significant bit in `viewMask` **must** be less
  than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxMultiviewViewCount"
  target="_blank" rel="noopener"><code>maxMultiviewViewCount</code></a>

- <a href="#VUID-VkRenderingInfo-perViewRenderAreaCount-07857"
  id="VUID-VkRenderingInfo-perViewRenderAreaCount-07857"></a>
  VUID-VkRenderingInfo-perViewRenderAreaCount-07857  
  If the `perViewRenderAreaCount` member of a
  [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html)
  structure included in the `pNext` chain is not `0`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview-per-view-render-areas"
  target="_blank"
  rel="noopener"><code>multiviewPerViewRenderAreas</code></a> feature
  **must** be enabled

- <a href="#VUID-VkRenderingInfo-perViewRenderAreaCount-07858"
  id="VUID-VkRenderingInfo-perViewRenderAreaCount-07858"></a>
  VUID-VkRenderingInfo-perViewRenderAreaCount-07858  
  If the `perViewRenderAreaCount` member of a
  [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html)
  structure included in the `pNext` chain is not `0`, then `renderArea`
  **must** specify a render area that includes the union of all per view
  render areas

- <a href="#VUID-VkRenderingInfo-None-09044"
  id="VUID-VkRenderingInfo-None-09044"></a>
  VUID-VkRenderingInfo-None-09044  
  Valid attachments specified by this structure **must** not be bound to
  memory locations that are bound to any other valid attachments
  specified by this structure

- <a href="#VUID-VkRenderingInfo-flags-09381"
  id="VUID-VkRenderingInfo-flags-09381"></a>
  VUID-VkRenderingInfo-flags-09381  
  If `flags` includes `VK_RENDERING_CONTENTS_INLINE_BIT_EXT` then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nestedCommandBuffer"
  target="_blank" rel="noopener"><code>nestedCommandBuffer</code></a>
  feature **must** be enabled

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-09318"
  id="VUID-VkRenderingInfo-pDepthAttachment-09318"></a>
  VUID-VkRenderingInfo-pDepthAttachment-09318  
  `pDepthAttachment->resolveMode` **must** not be
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-09319"
  id="VUID-VkRenderingInfo-pStencilAttachment-09319"></a>
  VUID-VkRenderingInfo-pStencilAttachment-09319  
  `pStencilAttachment->resolveMode` **must** not be
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-09320"
  id="VUID-VkRenderingInfo-colorAttachmentCount-09320"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-09320  
  If `colorAttachmentCount` is not `1`, the `resolveMode` member of any
  element of `pColorAttachments` **must** not be
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`

- <a href="#VUID-VkRenderingInfo-resolveMode-09321"
  id="VUID-VkRenderingInfo-resolveMode-09321"></a>
  VUID-VkRenderingInfo-resolveMode-09321  
  If the `resolveMode` of any element of `pColorAttachments` is
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`,
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)::`imageView`
  **must** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkRenderingInfo-resolveMode-09322"
  id="VUID-VkRenderingInfo-resolveMode-09322"></a>
  VUID-VkRenderingInfo-resolveMode-09322  
  If the `resolveMode` of any element of `pColorAttachments` is
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`,
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)::`imageView`
  **must** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkRenderingInfo-pNext-09535"
  id="VUID-VkRenderingInfo-pNext-09535"></a>
  VUID-VkRenderingInfo-pNext-09535  
  If the `pNext` chain contains a
  [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeBeginInfoARM.html)
  structure, the union of stripe areas defined by the elements of
  [VkRenderPassStripeInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeInfoARM.html)::`pStripeInfos`
  **must** cover the `renderArea`

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-09479"
  id="VUID-VkRenderingInfo-colorAttachmentCount-09479"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-09479  
  If `colorAttachmentCount` is not `0` and the `imageView` member of an
  element of `pColorAttachments` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), that `imageView` **must** have
  been created with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkRenderingInfo-colorAttachmentCount-09480"
  id="VUID-VkRenderingInfo-colorAttachmentCount-09480"></a>
  VUID-VkRenderingInfo-colorAttachmentCount-09480  
  If `colorAttachmentCount` is not `0`, and there is an element of
  `pColorAttachments` with either its `resolveMode` member set to
  `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID`, or its
  `imageView` member not set to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  and its `resolveMode` member not set to `VK_RESOLVE_MODE_NONE`, the
  `resolveImageView` member of that element of `pColorAttachments`
  **must** have been created with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-09481"
  id="VUID-VkRenderingInfo-pDepthAttachment-09481"></a>
  VUID-VkRenderingInfo-pDepthAttachment-09481  
  If `pDepthAttachment` is not `NULL` and `pDepthAttachment->imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pDepthAttachment->imageView` **must** have been created with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-09482"
  id="VUID-VkRenderingInfo-pDepthAttachment-09482"></a>
  VUID-VkRenderingInfo-pDepthAttachment-09482  
  If `pDepthAttachment` is not `NULL`, `pDepthAttachment->imageView` is
  not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  `pDepthAttachment->resolveMode` is not `VK_RESOLVE_MODE_NONE`,
  `pDepthAttachment->resolveImageView` **must** have been created with
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-09483"
  id="VUID-VkRenderingInfo-pStencilAttachment-09483"></a>
  VUID-VkRenderingInfo-pStencilAttachment-09483  
  If `pStencilAttachment` is not `NULL` and
  `pStencilAttachment->imageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pStencilAttachment->imageView`
  **must** have been created with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-09484"
  id="VUID-VkRenderingInfo-pStencilAttachment-09484"></a>
  VUID-VkRenderingInfo-pStencilAttachment-09484  
  If `pStencilAttachment` is not `NULL`, `pStencilAttachment->imageView`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and
  `pStencilAttachment->resolveMode` is not `VK_RESOLVE_MODE_NONE`,
  `pStencilAttachment->resolveImageView` **must** have been created with
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkRenderingInfo-imageView-09485"
  id="VUID-VkRenderingInfo-imageView-09485"></a>
  VUID-VkRenderingInfo-imageView-09485  
  If the `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it **must** have been created
  with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkRenderingInfo-imageView-09486"
  id="VUID-VkRenderingInfo-imageView-09486"></a>
  VUID-VkRenderingInfo-imageView-09486  
  If the `imageView` member of a
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html)
  structure included in the `pNext` chain is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it **must** have been created
  with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

Valid Usage (Implicit)

- <a href="#VUID-VkRenderingInfo-sType-sType"
  id="VUID-VkRenderingInfo-sType-sType"></a>
  VUID-VkRenderingInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDERING_INFO`

- <a href="#VUID-VkRenderingInfo-pNext-pNext"
  id="VUID-VkRenderingInfo-pNext-pNext"></a>
  VUID-VkRenderingInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupRenderPassBeginInfo.html),
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html),
  [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewAttributesInfoNVX.html),
  [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html),
  [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeBeginInfoARM.html),
  [VkRenderingFragmentDensityMapAttachmentInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentDensityMapAttachmentInfoEXT.html),
  or
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)

- <a href="#VUID-VkRenderingInfo-sType-unique"
  id="VUID-VkRenderingInfo-sType-unique"></a>
  VUID-VkRenderingInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkRenderingInfo-flags-parameter"
  id="VUID-VkRenderingInfo-flags-parameter"></a>
  VUID-VkRenderingInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFlagBits.html) values

- <a href="#VUID-VkRenderingInfo-pColorAttachments-parameter"
  id="VUID-VkRenderingInfo-pColorAttachments-parameter"></a>
  VUID-VkRenderingInfo-pColorAttachments-parameter  
  If `colorAttachmentCount` is not `0`, `pColorAttachments` **must** be
  a valid pointer to an array of `colorAttachmentCount` valid
  [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html) structures

- <a href="#VUID-VkRenderingInfo-pDepthAttachment-parameter"
  id="VUID-VkRenderingInfo-pDepthAttachment-parameter"></a>
  VUID-VkRenderingInfo-pDepthAttachment-parameter  
  If `pDepthAttachment` is not `NULL`, `pDepthAttachment` **must** be a
  valid pointer to a valid
  [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html) structure

- <a href="#VUID-VkRenderingInfo-pStencilAttachment-parameter"
  id="VUID-VkRenderingInfo-pStencilAttachment-parameter"></a>
  VUID-VkRenderingInfo-pStencilAttachment-parameter  
  If `pStencilAttachment` is not `NULL`, `pStencilAttachment` **must**
  be a valid pointer to a valid
  [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html),
[VK_QCOM_tile_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_tile_properties.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html),
[VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentInfo.html),
[VkRenderingFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html),
[vkCmdBeginRenderingKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRenderingKHR.html),
[vkGetDynamicRenderingTilePropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDynamicRenderingTilePropertiesQCOM.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderingInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
