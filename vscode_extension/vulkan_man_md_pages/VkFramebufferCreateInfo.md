# VkFramebufferCreateInfo(3) Manual Page

## Name

VkFramebufferCreateInfo - Structure specifying parameters of a newly
created framebuffer



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkFramebufferCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkFramebufferCreateInfo {
    VkStructureType             sType;
    const void*                 pNext;
    VkFramebufferCreateFlags    flags;
    VkRenderPass                renderPass;
    uint32_t                    attachmentCount;
    const VkImageView*          pAttachments;
    uint32_t                    width;
    uint32_t                    height;
    uint32_t                    layers;
} VkFramebufferCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkFramebufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateFlagBits.html)

- `renderPass` is a render pass defining what render passes the
  framebuffer will be compatible with. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-compatibility"
  target="_blank" rel="noopener">Render Pass Compatibility</a> for
  details.

- `attachmentCount` is the number of attachments.

- `pAttachments` is a pointer to an array of
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handles, each of which will be used as
  the corresponding attachment in a render pass instance. If `flags`
  includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, this parameter is
  ignored.

- `width`, `height` and `layers` define the dimensions of the
  framebuffer. If the render pass uses multiview, then `layers` **must**
  be one and each attachment requires a number of layers that is greater
  than the maximum bit index set in the view mask in the subpasses in
  which it is used.

## <a href="#_description" class="anchor"></a>Description

It is legal for a subpass to use no color or depth/stencil attachments,
either because it has no attachment references or because all of them
are `VK_ATTACHMENT_UNUSED`. This kind of subpass **can** use shader side
effects such as image stores and atomics to produce an output. In this
case, the subpass continues to use the `width`, `height`, and `layers`
of the framebuffer to define the dimensions of the rendering area, and
the `rasterizationSamples` from each pipeline’s
[VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)
to define the number of samples used in rasterization; however, if
[VkPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures.html)::`variableMultisampleRate`
is `VK_FALSE`, then all pipelines to be bound with the subpass **must**
have the same value for
[VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)::`rasterizationSamples`.
In all such cases, `rasterizationSamples` **must** be a valid
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value that is set in
[VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`framebufferNoAttachmentsSampleCounts`.

Valid Usage

- <a href="#VUID-VkFramebufferCreateInfo-attachmentCount-00876"
  id="VUID-VkFramebufferCreateInfo-attachmentCount-00876"></a>
  VUID-VkFramebufferCreateInfo-attachmentCount-00876  
  `attachmentCount` **must** be equal to the attachment count specified
  in `renderPass`

- <a href="#VUID-VkFramebufferCreateInfo-flags-02778"
  id="VUID-VkFramebufferCreateInfo-flags-02778"></a>
  VUID-VkFramebufferCreateInfo-flags-02778  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT` and
  `attachmentCount` is not `0`, `pAttachments` **must** be a valid
  pointer to an array of `attachmentCount` valid
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handles

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-00877"
  id="VUID-VkFramebufferCreateInfo-pAttachments-00877"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-00877  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` that is used as a color attachment or
  resolve attachment by `renderPass` **must** have been created with a
  `usage` value including `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-02633"
  id="VUID-VkFramebufferCreateInfo-pAttachments-02633"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-02633  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` that is used as a depth/stencil
  attachment by `renderPass` **must** have been created with a `usage`
  value including `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-02634"
  id="VUID-VkFramebufferCreateInfo-pAttachments-02634"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-02634  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` that is used as a depth/stencil resolve
  attachment by `renderPass` **must** have been created with a `usage`
  value including `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-00879"
  id="VUID-VkFramebufferCreateInfo-pAttachments-00879"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-00879  
  If `renderpass` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `flags`
  does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element
  of `pAttachments` that is used as an input attachment by `renderPass`
  **must** have been created with a `usage` value including
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-02552"
  id="VUID-VkFramebufferCreateInfo-pAttachments-02552"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-02552  
  Each element of `pAttachments` that is used as a fragment density map
  attachment by `renderPass` **must** not have been created with a
  `flags` value including `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

- <a href="#VUID-VkFramebufferCreateInfo-renderPass-02553"
  id="VUID-VkFramebufferCreateInfo-renderPass-02553"></a>
  VUID-VkFramebufferCreateInfo-renderPass-02553  
  If `renderPass` has a fragment density map attachment and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-fragmentDensityMapNonSubsampledImages"
  target="_blank"
  rel="noopener"><code>fragmentDensityMapNonSubsampledImages</code></a>
  feature is not enabled, each element of `pAttachments` **must** have
  been created with a `flags` value including
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT` unless that element is the
  fragment density map attachment

- <a href="#VUID-VkFramebufferCreateInfo-renderPass-06502"
  id="VUID-VkFramebufferCreateInfo-renderPass-06502"></a>
  VUID-VkFramebufferCreateInfo-renderPass-06502  
  If `renderPass` was created with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-fragmentdensitymapoffsets"
  target="_blank" rel="noopener">fragment density map offsets</a> other
  than (0,0), each element of `pAttachments` **must** have been created
  with a `flags` value including
  `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_QCOM`

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-00880"
  id="VUID-VkFramebufferCreateInfo-pAttachments-00880"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-00880  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` **must** have been created with a
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value that matches the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) specified by the corresponding
  `VkAttachmentDescription` in `renderPass`

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-00881"
  id="VUID-VkFramebufferCreateInfo-pAttachments-00881"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-00881  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` **must** have been created with a
  `samples` value that matches the `samples` value specified by the
  corresponding `VkAttachmentDescription` in `renderPass`

- <a href="#VUID-VkFramebufferCreateInfo-flags-04533"
  id="VUID-VkFramebufferCreateInfo-flags-04533"></a>
  VUID-VkFramebufferCreateInfo-flags-04533  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` that is used as an input, color,
  resolve, or depth/stencil attachment by `renderPass` **must** have
  been created with a
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`extent.width` greater
  than or equal to `width`

- <a href="#VUID-VkFramebufferCreateInfo-flags-04534"
  id="VUID-VkFramebufferCreateInfo-flags-04534"></a>
  VUID-VkFramebufferCreateInfo-flags-04534  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` that is used as an input, color,
  resolve, or depth/stencil attachment by `renderPass` **must** have
  been created with a
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`extent.height` greater
  than or equal to `height`

- <a href="#VUID-VkFramebufferCreateInfo-flags-04535"
  id="VUID-VkFramebufferCreateInfo-flags-04535"></a>
  VUID-VkFramebufferCreateInfo-flags-04535  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` that is used as an input, color,
  resolve, or depth/stencil attachment by `renderPass` **must** have
  been created with a
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)::`subresourceRange.layerCount`
  greater than or equal to `layers`

- <a href="#VUID-VkFramebufferCreateInfo-renderPass-04536"
  id="VUID-VkFramebufferCreateInfo-renderPass-04536"></a>
  VUID-VkFramebufferCreateInfo-renderPass-04536  
  If `renderPass` was specified with non-zero view masks, each element
  of `pAttachments` that is used as an input, color, resolve, or
  depth/stencil attachment by `renderPass` **must** have a `layerCount`
  greater than the index of the most significant bit set in any of those
  view masks

- <a href="#VUID-VkFramebufferCreateInfo-renderPass-02746"
  id="VUID-VkFramebufferCreateInfo-renderPass-02746"></a>
  VUID-VkFramebufferCreateInfo-renderPass-02746  
  Each element of `pAttachments` that is referenced by
  `fragmentDensityMapAttachment` **must** have a `layerCount` equal to
  `1` or if `renderPass` was specified with non-zero view masks, greater
  than the index of the most significant bit set in any of those view
  masks

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-02555"
  id="VUID-VkFramebufferCreateInfo-pAttachments-02555"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-02555  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, an
  element of `pAttachments` that is referenced by
  `fragmentDensityMapAttachment` **must** have a width at least as large
  as ⌈maxFragmentDensityTexelSizewidth​width​⌉

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-02556"
  id="VUID-VkFramebufferCreateInfo-pAttachments-02556"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-02556  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, an
  element of `pAttachments` that is referenced by
  `fragmentDensityMapAttachment` **must** have a height at least as
  large as ⌈maxFragmentDensityTexelSizeheight​height​⌉

- <a href="#VUID-VkFramebufferCreateInfo-flags-04537"
  id="VUID-VkFramebufferCreateInfo-flags-04537"></a>
  VUID-VkFramebufferCreateInfo-flags-04537  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, and
  `renderPass` was specified with non-zero view masks, each element of
  `pAttachments` that is used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">fragment shading rate attachment</a> by
  `renderPass` **must** have a `layerCount` that is either `1`, or
  greater than the index of the most significant bit set in any of those
  view masks

- <a href="#VUID-VkFramebufferCreateInfo-flags-04538"
  id="VUID-VkFramebufferCreateInfo-flags-04538"></a>
  VUID-VkFramebufferCreateInfo-flags-04538  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, and
  `renderPass` was not specified with non-zero view masks, each element
  of `pAttachments` that is used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">fragment shading rate attachment</a> by
  `renderPass` **must** have a `layerCount` that is either `1`, or
  greater than `layers`

- <a href="#VUID-VkFramebufferCreateInfo-flags-04539"
  id="VUID-VkFramebufferCreateInfo-flags-04539"></a>
  VUID-VkFramebufferCreateInfo-flags-04539  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> is not
  enabled or the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-robustFragmentShadingRateAttachmentAccess"
  target="_blank"
  rel="noopener"><code>robustFragmentShadingRateAttachmentAccess</code></a>
  limit is `VK_FALSE` or the `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure was created with
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`baseMipLevel`
  greater than 0, `flags` does not include
  `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, an element of `pAttachments`
  that is used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">fragment shading rate attachment</a>
  **must** have a width at least as large as ⌈`width` / `texelWidth`⌉,
  where `texelWidth` is the largest value of
  `shadingRateAttachmentTexelSize.width` in a
  [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateAttachmentInfoKHR.html)
  which references that attachment

- <a href="#VUID-VkFramebufferCreateInfo-flags-04540"
  id="VUID-VkFramebufferCreateInfo-flags-04540"></a>
  VUID-VkFramebufferCreateInfo-flags-04540  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> is not
  enabled or the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-robustFragmentShadingRateAttachmentAccess"
  target="_blank"
  rel="noopener"><code>robustFragmentShadingRateAttachmentAccess</code></a>
  limit is `VK_FALSE` or the `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure was created with
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`baseMipLevel`
  greater than 0, `flags` does not include
  `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, an element of `pAttachments`
  that is used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">fragment shading rate attachment</a>
  **must** have a height at least as large as ⌈`height` /
  `texelHeight`⌉, where `texelHeight` is the largest value of
  `shadingRateAttachmentTexelSize.height` in a
  [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateAttachmentInfoKHR.html)
  which references that attachment

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-00883"
  id="VUID-VkFramebufferCreateInfo-pAttachments-00883"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-00883  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` **must** only specify a single mip
  level

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-00884"
  id="VUID-VkFramebufferCreateInfo-pAttachments-00884"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-00884  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` **must** have been created with the
  identity swizzle

- <a href="#VUID-VkFramebufferCreateInfo-width-00885"
  id="VUID-VkFramebufferCreateInfo-width-00885"></a>
  VUID-VkFramebufferCreateInfo-width-00885  
  `width` **must** be greater than `0`

- <a href="#VUID-VkFramebufferCreateInfo-width-00886"
  id="VUID-VkFramebufferCreateInfo-width-00886"></a>
  VUID-VkFramebufferCreateInfo-width-00886  
  `width` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFramebufferWidth"
  target="_blank" rel="noopener"><code>maxFramebufferWidth</code></a>

- <a href="#VUID-VkFramebufferCreateInfo-height-00887"
  id="VUID-VkFramebufferCreateInfo-height-00887"></a>
  VUID-VkFramebufferCreateInfo-height-00887  
  `height` **must** be greater than `0`

- <a href="#VUID-VkFramebufferCreateInfo-height-00888"
  id="VUID-VkFramebufferCreateInfo-height-00888"></a>
  VUID-VkFramebufferCreateInfo-height-00888  
  `height` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFramebufferHeight"
  target="_blank" rel="noopener"><code>maxFramebufferHeight</code></a>

- <a href="#VUID-VkFramebufferCreateInfo-layers-00889"
  id="VUID-VkFramebufferCreateInfo-layers-00889"></a>
  VUID-VkFramebufferCreateInfo-layers-00889  
  `layers` **must** be greater than `0`

- <a href="#VUID-VkFramebufferCreateInfo-layers-00890"
  id="VUID-VkFramebufferCreateInfo-layers-00890"></a>
  VUID-VkFramebufferCreateInfo-layers-00890  
  `layers` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxFramebufferLayers"
  target="_blank" rel="noopener"><code>maxFramebufferLayers</code></a>

- <a href="#VUID-VkFramebufferCreateInfo-renderPass-02531"
  id="VUID-VkFramebufferCreateInfo-renderPass-02531"></a>
  VUID-VkFramebufferCreateInfo-renderPass-02531  
  If `renderPass` was specified with non-zero view masks, `layers`
  **must** be `1`

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-00891"
  id="VUID-VkFramebufferCreateInfo-pAttachments-00891"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-00891  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` that is a 2D or 2D array image view
  taken from a 3D image **must** not be a depth/stencil format

- <a href="#VUID-VkFramebufferCreateInfo-flags-03189"
  id="VUID-VkFramebufferCreateInfo-flags-03189"></a>
  VUID-VkFramebufferCreateInfo-flags-03189  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-imagelessFramebuffer"
  target="_blank" rel="noopener"><code>imagelessFramebuffer</code></a>
  feature is not enabled, `flags` **must** not include
  `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`

- <a href="#VUID-VkFramebufferCreateInfo-flags-03190"
  id="VUID-VkFramebufferCreateInfo-flags-03190"></a>
  VUID-VkFramebufferCreateInfo-flags-03190  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `pNext`
  chain **must** include a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure

- <a href="#VUID-VkFramebufferCreateInfo-flags-03191"
  id="VUID-VkFramebufferCreateInfo-flags-03191"></a>
  VUID-VkFramebufferCreateInfo-flags-03191  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the
  `attachmentImageInfoCount` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure in the `pNext` chain **must** be equal to either zero or
  `attachmentCount`

- <a href="#VUID-VkFramebufferCreateInfo-flags-04541"
  id="VUID-VkFramebufferCreateInfo-flags-04541"></a>
  VUID-VkFramebufferCreateInfo-flags-04541  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `width`
  member of any element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure in the `pNext` chain that is used as an input, color,
  resolve or depth/stencil attachment in `renderPass` **must** be
  greater than or equal to `width`

- <a href="#VUID-VkFramebufferCreateInfo-flags-04542"
  id="VUID-VkFramebufferCreateInfo-flags-04542"></a>
  VUID-VkFramebufferCreateInfo-flags-04542  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the
  `height` member of any element of the `pAttachmentImageInfos` member
  of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure in the `pNext` chain that is used as an input, color,
  resolve or depth/stencil attachment in `renderPass` **must** be
  greater than or equal to `height`

- <a href="#VUID-VkFramebufferCreateInfo-flags-03196"
  id="VUID-VkFramebufferCreateInfo-flags-03196"></a>
  VUID-VkFramebufferCreateInfo-flags-03196  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `width`
  member of any element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure in the `pNext` chain that is referenced by
  [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)::`fragmentDensityMapAttachment`
  in `renderPass` **must** be greater than or equal to
  ⌈maxFragmentDensityTexelSizewidth​width​⌉

- <a href="#VUID-VkFramebufferCreateInfo-flags-03197"
  id="VUID-VkFramebufferCreateInfo-flags-03197"></a>
  VUID-VkFramebufferCreateInfo-flags-03197  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the
  `height` member of any element of the `pAttachmentImageInfos` member
  of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure included in the `pNext` chain that is referenced by
  [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)::`fragmentDensityMapAttachment`
  in `renderPass` **must** be greater than or equal to
  ⌈maxFragmentDensityTexelSizeheight​height​⌉

- <a href="#VUID-VkFramebufferCreateInfo-flags-04543"
  id="VUID-VkFramebufferCreateInfo-flags-04543"></a>
  VUID-VkFramebufferCreateInfo-flags-04543  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> is not
  enabled or the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-robustFragmentShadingRateAttachmentAccess"
  target="_blank"
  rel="noopener"><code>robustFragmentShadingRateAttachmentAccess</code></a>
  limit is `VK_FALSE` or the `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure was created with
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`baseMipLevel`
  greater than 0, and `flags` includes
  `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `width` member of any
  element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure in the `pNext` chain that is used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">fragment shading rate attachment</a>
  **must** be greater than or equal to ⌈`width` / `texelWidth`⌉, where
  `texelWidth` is the largest value of
  `shadingRateAttachmentTexelSize.width` in a
  [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateAttachmentInfoKHR.html)
  which references that attachment

- <a href="#VUID-VkFramebufferCreateInfo-flags-04544"
  id="VUID-VkFramebufferCreateInfo-flags-04544"></a>
  VUID-VkFramebufferCreateInfo-flags-04544  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance7"
  target="_blank" rel="noopener"><code>maintenance7</code></a> is not
  enabled or the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-robustFragmentShadingRateAttachmentAccess"
  target="_blank"
  rel="noopener"><code>robustFragmentShadingRateAttachmentAccess</code></a>
  limit is `VK_FALSE` or the `imageView` member of a
  [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html)
  structure was created with
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html)::`baseMipLevel`
  greater than 0, and `flags` includes
  `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `height` member of any
  element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure in the `pNext` chain that is used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">fragment shading rate attachment</a>
  **must** be greater than or equal to ⌈`height` / `texelHeight`⌉, where
  `texelHeight` is the largest value of
  `shadingRateAttachmentTexelSize.height` in a
  [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFragmentShadingRateAttachmentInfoKHR.html)
  which references that attachment

- <a href="#VUID-VkFramebufferCreateInfo-flags-04545"
  id="VUID-VkFramebufferCreateInfo-flags-04545"></a>
  VUID-VkFramebufferCreateInfo-flags-04545  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the
  `layerCount` member of any element of the `pAttachmentImageInfos`
  member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure in the `pNext` chain that is used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">fragment shading rate attachment</a>
  **must** be either `1`, or greater than or equal to `layers`

- <a href="#VUID-VkFramebufferCreateInfo-flags-04587"
  id="VUID-VkFramebufferCreateInfo-flags-04587"></a>
  VUID-VkFramebufferCreateInfo-flags-04587  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT` and
  `renderPass` was specified with non-zero view masks, the `layerCount`
  member of any element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure in the `pNext` chain that is used as a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-fragment-shading-rate-attachment"
  target="_blank" rel="noopener">fragment shading rate attachment</a>
  **must** be either `1`, or greater than the index of the most
  significant bit set in any of those view masks

- <a href="#VUID-VkFramebufferCreateInfo-renderPass-03198"
  id="VUID-VkFramebufferCreateInfo-renderPass-03198"></a>
  VUID-VkFramebufferCreateInfo-renderPass-03198  
  If multiview is enabled for `renderPass` and `flags` includes
  `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `layerCount` member of any
  element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure included in the `pNext` chain used as an input, color,
  resolve, or depth/stencil attachment in `renderPass` **must** be
  greater than the maximum bit index set in the view mask in the
  subpasses in which it is used in `renderPass`

- <a href="#VUID-VkFramebufferCreateInfo-renderPass-04546"
  id="VUID-VkFramebufferCreateInfo-renderPass-04546"></a>
  VUID-VkFramebufferCreateInfo-renderPass-04546  
  If multiview is not enabled for `renderPass` and `flags` includes
  `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `layerCount` member of any
  element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure included in the `pNext` chain used as an input, color,
  resolve, or depth/stencil attachment in `renderPass` **must** be
  greater than or equal to `layers`

- <a href="#VUID-VkFramebufferCreateInfo-flags-03201"
  id="VUID-VkFramebufferCreateInfo-flags-03201"></a>
  VUID-VkFramebufferCreateInfo-flags-03201  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `usage`
  member of any element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure included in the `pNext` chain that refers to an attachment
  used as a color attachment or resolve attachment by `renderPass`
  **must** include `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`

- <a href="#VUID-VkFramebufferCreateInfo-flags-03202"
  id="VUID-VkFramebufferCreateInfo-flags-03202"></a>
  VUID-VkFramebufferCreateInfo-flags-03202  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `usage`
  member of any element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure included in the `pNext` chain that refers to an attachment
  used as a depth/stencil attachment by `renderPass` **must** include
  `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkFramebufferCreateInfo-flags-03203"
  id="VUID-VkFramebufferCreateInfo-flags-03203"></a>
  VUID-VkFramebufferCreateInfo-flags-03203  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `usage`
  member of any element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure included in the `pNext` chain that refers to an attachment
  used as a depth/stencil resolve attachment by `renderPass` **must**
  include `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a href="#VUID-VkFramebufferCreateInfo-flags-03204"
  id="VUID-VkFramebufferCreateInfo-flags-03204"></a>
  VUID-VkFramebufferCreateInfo-flags-03204  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `usage`
  member of any element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure included in the `pNext` chain that refers to an attachment
  used as an input attachment by `renderPass` **must** include
  `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`

- <a href="#VUID-VkFramebufferCreateInfo-flags-03205"
  id="VUID-VkFramebufferCreateInfo-flags-03205"></a>
  VUID-VkFramebufferCreateInfo-flags-03205  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, at least
  one element of the `pViewFormats` member of any element of the
  `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure included in the `pNext` chain **must** be equal to the
  corresponding value of
  [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html)::`format` used
  to create `renderPass`

- <a href="#VUID-VkFramebufferCreateInfo-flags-04113"
  id="VUID-VkFramebufferCreateInfo-flags-04113"></a>
  VUID-VkFramebufferCreateInfo-flags-04113  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` **must** have been created with
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)::`viewType` not
  equal to `VK_IMAGE_VIEW_TYPE_3D`

- <a href="#VUID-VkFramebufferCreateInfo-flags-04548"
  id="VUID-VkFramebufferCreateInfo-flags-04548"></a>
  VUID-VkFramebufferCreateInfo-flags-04548  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  each element of `pAttachments` that is used as a fragment shading rate
  attachment by `renderPass` **must** have been created with a `usage`
  value including
  `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkFramebufferCreateInfo-flags-04549"
  id="VUID-VkFramebufferCreateInfo-flags-04549"></a>
  VUID-VkFramebufferCreateInfo-flags-04549  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `usage`
  member of any element of the `pAttachmentImageInfos` member of a
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)
  structure included in the `pNext` chain that refers to an attachment
  used as a fragment shading rate attachment by `renderPass` **must**
  include `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkFramebufferCreateInfo-samples-06881"
  id="VUID-VkFramebufferCreateInfo-samples-06881"></a>
  VUID-VkFramebufferCreateInfo-samples-06881  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#subpass-multisampledrendertosinglesampled"
  target="_blank" rel="noopener">multisampled-render-to-single-sampled</a>
  is enabled for any subpass, all color, depth/stencil and input
  attachments used in that subpass which have
  `VkAttachmentDescription`::`samples` or
  `VkAttachmentDescription2`::`samples` equal to `VK_SAMPLE_COUNT_1_BIT`
  **must** have been created with
  `VK_IMAGE_CREATE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_BIT_EXT` in
  their [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags`

- <a href="#VUID-VkFramebufferCreateInfo-samples-07009"
  id="VUID-VkFramebufferCreateInfo-samples-07009"></a>
  VUID-VkFramebufferCreateInfo-samples-07009  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#subpass-multisampledrendertosinglesampled"
  target="_blank" rel="noopener">multisampled-render-to-single-sampled</a>
  is enabled for any subpass, all color, depth/stencil and input
  attachments used in that subpass which have
  `VkAttachmentDescription`::`samples` or
  `VkAttachmentDescription2`::`samples` equal to `VK_SAMPLE_COUNT_1_BIT`
  **must** have a format that supports the sample count specified in
  [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)::`rasterizationSamples`

- <a
  href="#VUID-VkFramebufferCreateInfo-nullColorAttachmentWithExternalFormatResolve-09349"
  id="VUID-VkFramebufferCreateInfo-nullColorAttachmentWithExternalFormatResolve-09349"></a>
  VUID-VkFramebufferCreateInfo-nullColorAttachmentWithExternalFormatResolve-09349  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve"
  target="_blank"
  rel="noopener"><code>nullColorAttachmentWithExternalFormatResolve</code></a>
  is `VK_FALSE`, and `flags` does not include
  `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the format of the color
  attachment for each subpass in `renderPass` that includes an external
  format image as a resolve attachment **must** have a format equal to
  the value of
  [VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html)::`colorAttachmentFormat`
  as returned by a call to
  [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)
  for the Android hardware buffer that was used to create the image view
  use as its resolve attachment

- <a href="#VUID-VkFramebufferCreateInfo-pAttachments-09350"
  id="VUID-VkFramebufferCreateInfo-pAttachments-09350"></a>
  VUID-VkFramebufferCreateInfo-pAttachments-09350  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`,
  then if an element of `pAttachments` has a format of
  `VK_FORMAT_UNDEFINED`, it **must** have been created with a
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  value identical to that provided in the
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html)::`externalFormat`
  specified by the corresponding
  [VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription2.html) in
  `renderPass`

Valid Usage (Implicit)

- <a href="#VUID-VkFramebufferCreateInfo-sType-sType"
  id="VUID-VkFramebufferCreateInfo-sType-sType"></a>
  VUID-VkFramebufferCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO`

- <a href="#VUID-VkFramebufferCreateInfo-pNext-pNext"
  id="VUID-VkFramebufferCreateInfo-pNext-pNext"></a>
  VUID-VkFramebufferCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferAttachmentsCreateInfo.html)

- <a href="#VUID-VkFramebufferCreateInfo-sType-unique"
  id="VUID-VkFramebufferCreateInfo-sType-unique"></a>
  VUID-VkFramebufferCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkFramebufferCreateInfo-flags-parameter"
  id="VUID-VkFramebufferCreateInfo-flags-parameter"></a>
  VUID-VkFramebufferCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkFramebufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateFlagBits.html) values

- <a href="#VUID-VkFramebufferCreateInfo-renderPass-parameter"
  id="VUID-VkFramebufferCreateInfo-renderPass-parameter"></a>
  VUID-VkFramebufferCreateInfo-renderPass-parameter  
  `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)
  handle

- <a href="#VUID-VkFramebufferCreateInfo-commonparent"
  id="VUID-VkFramebufferCreateInfo-commonparent"></a>
  VUID-VkFramebufferCreateInfo-commonparent  
  Both of `renderPass`, and the elements of `pAttachments` that are
  valid handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkFramebufferCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateFlags.html),
[VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateFramebuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFramebufferCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
