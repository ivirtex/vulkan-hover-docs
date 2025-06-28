# VkFramebufferCreateInfo(3) Manual Page

## Name

VkFramebufferCreateInfo - Structure specifying parameters of a newly created framebuffer



## [](#_c_specification)C Specification

The `VkFramebufferCreateInfo` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkFramebufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateFlagBits.html)
- `renderPass` is a render pass defining what render passes the framebuffer will be compatible with. See [Render Pass Compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-compatibility) for details. The implementation **must** not access this object outside of the duration of the command this structure is passed to.
- `attachmentCount` is the number of attachments.
- `pAttachments` is a pointer to an array of [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handles, each of which will be used as the corresponding attachment in a render pass instance. If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, this parameter is ignored.
- `width`, `height` and `layers` define the dimensions of the framebuffer. If the render pass uses multiview, then `layers` **must** be one and each attachment requires a number of layers that is greater than the maximum bit index set in the view mask in the subpasses in which it is used.

## [](#_description)Description

It is legal for a subpass to use no color or depth/stencil attachments, either because it has no attachment references or because all of them are `VK_ATTACHMENT_UNUSED`. This kind of subpass **can** use shader side effects such as image stores and atomics to produce an output. In this case, the subpass continues to use the `width`, `height`, and `layers` of the framebuffer to define the dimensions of the rendering area, and the `rasterizationSamples` from each pipeline’s [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineMultisampleStateCreateInfo.html) to define the number of samples used in rasterization; however, if [VkPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures.html)::`variableMultisampleRate` is `VK_FALSE`, then all pipelines to be bound with the subpass **must** have the same value for [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineMultisampleStateCreateInfo.html)::`rasterizationSamples`. In all such cases, `rasterizationSamples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value that is set in [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`framebufferNoAttachmentsSampleCounts`.

Valid Usage

- [](#VUID-VkFramebufferCreateInfo-attachmentCount-00876)VUID-VkFramebufferCreateInfo-attachmentCount-00876  
  `attachmentCount` **must** be equal to the attachment count specified in `renderPass`
- [](#VUID-VkFramebufferCreateInfo-flags-02778)VUID-VkFramebufferCreateInfo-flags-02778  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT` and `attachmentCount` is not `0`, `pAttachments` **must** be a valid pointer to an array of `attachmentCount` valid [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) handles
- [](#VUID-VkFramebufferCreateInfo-pAttachments-00877)VUID-VkFramebufferCreateInfo-pAttachments-00877  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` that is used as a color attachment or resolve attachment by `renderPass` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`
- [](#VUID-VkFramebufferCreateInfo-pAttachments-02633)VUID-VkFramebufferCreateInfo-pAttachments-02633  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` that is used as a depth/stencil attachment by `renderPass` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkFramebufferCreateInfo-pAttachments-02634)VUID-VkFramebufferCreateInfo-pAttachments-02634  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` that is used as a depth/stencil resolve attachment by `renderPass` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkFramebufferCreateInfo-pAttachments-00879)VUID-VkFramebufferCreateInfo-pAttachments-00879  
  If `renderpass` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` that is used as an input attachment by `renderPass` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-VkFramebufferCreateInfo-pAttachments-02552)VUID-VkFramebufferCreateInfo-pAttachments-02552  
  Each element of `pAttachments` that is used as a fragment density map attachment by `renderPass` **must** not have been created with a `flags` value including `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`
- [](#VUID-VkFramebufferCreateInfo-renderPass-02553)VUID-VkFramebufferCreateInfo-renderPass-02553  
  If `renderPass` has a fragment density map attachment and the [`fragmentDensityMapNonSubsampledImages`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-fragmentDensityMapNonSubsampledImages) feature is not enabled, each element of `pAttachments` **must** have been created with a `flags` value including `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT` unless that element is the fragment density map attachment
- [](#VUID-VkFramebufferCreateInfo-renderPass-10830)VUID-VkFramebufferCreateInfo-renderPass-10830  
  If `renderPass` was created with `VK_RENDER_PASS_CREATE_PER_LAYER_FRAGMENT_DENSITY_BIT_VALVE`, then `layers` **must** be less than or equal to [`maxFragmentDensityMapLayers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentDensityMapLayers)
- [](#VUID-VkFramebufferCreateInfo-pAttachments-00880)VUID-VkFramebufferCreateInfo-pAttachments-00880  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` **must** have been created with a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value that matches the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) specified by the corresponding `VkAttachmentDescription` in `renderPass`
- [](#VUID-VkFramebufferCreateInfo-pAttachments-00881)VUID-VkFramebufferCreateInfo-pAttachments-00881  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` **must** have been created with a `samples` value that matches the `samples` value specified by the corresponding `VkAttachmentDescription` in `renderPass`
- [](#VUID-VkFramebufferCreateInfo-flags-04533)VUID-VkFramebufferCreateInfo-flags-04533  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` that is used as an input, color, resolve, or depth/stencil attachment by `renderPass` **must** have been created with a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`extent.width` greater than or equal to `width`
- [](#VUID-VkFramebufferCreateInfo-flags-04534)VUID-VkFramebufferCreateInfo-flags-04534  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` that is used as an input, color, resolve, or depth/stencil attachment by `renderPass` **must** have been created with a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`extent.height` greater than or equal to `height`
- [](#VUID-VkFramebufferCreateInfo-flags-04535)VUID-VkFramebufferCreateInfo-flags-04535  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` that is used as an input, color, resolve, or depth/stencil attachment by `renderPass` **must** have been created with a [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html)::`subresourceRange.layerCount` greater than or equal to `layers`
- [](#VUID-VkFramebufferCreateInfo-renderPass-04536)VUID-VkFramebufferCreateInfo-renderPass-04536  
  If `renderPass` was specified with non-zero view masks, each element of `pAttachments` that is used as an input, color, resolve, or depth/stencil attachment by `renderPass` **must** have a `layerCount` greater than the index of the most significant bit set in any of those view masks
- [](#VUID-VkFramebufferCreateInfo-renderPass-02746)VUID-VkFramebufferCreateInfo-renderPass-02746  
  Each element of `pAttachments` that is referenced by `fragmentDensityMapAttachment` **must** have a `layerCount` equal to `1` or if `renderPass` was specified with non-zero view masks, greater than the index of the most significant bit set in any of those view masks
- [](#VUID-VkFramebufferCreateInfo-pAttachments-02555)VUID-VkFramebufferCreateInfo-pAttachments-02555  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, an element of `pAttachments` that is referenced by `fragmentDensityMapAttachment` **must** have a width at least as large as ⌈maxFragmentDensityTexelSizewidth​width​⌉
- [](#VUID-VkFramebufferCreateInfo-pAttachments-02556)VUID-VkFramebufferCreateInfo-pAttachments-02556  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, an element of `pAttachments` that is referenced by `fragmentDensityMapAttachment` **must** have a height at least as large as ⌈maxFragmentDensityTexelSizeheight​height​⌉
- [](#VUID-VkFramebufferCreateInfo-flags-04537)VUID-VkFramebufferCreateInfo-flags-04537  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, and `renderPass` was specified with non-zero view masks, each element of `pAttachments` that is used as a [fragment shading rate attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment) by `renderPass` **must** have a `layerCount` that is either `1`, or greater than the index of the most significant bit set in any of those view masks
- [](#VUID-VkFramebufferCreateInfo-flags-04538)VUID-VkFramebufferCreateInfo-flags-04538  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, and `renderPass` was not specified with non-zero view masks, each element of `pAttachments` that is used as a [fragment shading rate attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment) by `renderPass` **must** have a `layerCount` that is either `1`, or greater than `layers`
- [](#VUID-VkFramebufferCreateInfo-flags-04539)VUID-VkFramebufferCreateInfo-flags-04539  
  If the [`maintenance7`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance7) feature is not enabled or the [`robustFragmentShadingRateAttachmentAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-robustFragmentShadingRateAttachmentAccess) limit is `VK_FALSE` or the `imageView` member of a [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html) structure was created with [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`baseMipLevel` greater than 0, `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, an element of `pAttachments` that is used as a [fragment shading rate attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment) **must** have a width at least as large as ⌈`width` / `texelWidth`⌉, where `texelWidth` is the largest value of `shadingRateAttachmentTexelSize.width` in a [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateAttachmentInfoKHR.html) which references that attachment
- [](#VUID-VkFramebufferCreateInfo-flags-04540)VUID-VkFramebufferCreateInfo-flags-04540  
  If the [`maintenance7`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance7) feature is not enabled or the [`robustFragmentShadingRateAttachmentAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-robustFragmentShadingRateAttachmentAccess) limit is `VK_FALSE` or the `imageView` member of a [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html) structure was created with [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`baseMipLevel` greater than 0, `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, an element of `pAttachments` that is used as a [fragment shading rate attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment) **must** have a height at least as large as ⌈`height` / `texelHeight`⌉, where `texelHeight` is the largest value of `shadingRateAttachmentTexelSize.height` in a [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateAttachmentInfoKHR.html) which references that attachment
- [](#VUID-VkFramebufferCreateInfo-pAttachments-00883)VUID-VkFramebufferCreateInfo-pAttachments-00883  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` **must** only specify a single mip level
- [](#VUID-VkFramebufferCreateInfo-pAttachments-00884)VUID-VkFramebufferCreateInfo-pAttachments-00884  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` **must** have been created with the identity swizzle
- [](#VUID-VkFramebufferCreateInfo-width-00885)VUID-VkFramebufferCreateInfo-width-00885  
  `width` **must** be greater than `0`
- [](#VUID-VkFramebufferCreateInfo-width-00886)VUID-VkFramebufferCreateInfo-width-00886  
  `width` **must** be less than or equal to [`maxFramebufferWidth`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFramebufferWidth)
- [](#VUID-VkFramebufferCreateInfo-height-00887)VUID-VkFramebufferCreateInfo-height-00887  
  `height` **must** be greater than `0`
- [](#VUID-VkFramebufferCreateInfo-height-00888)VUID-VkFramebufferCreateInfo-height-00888  
  `height` **must** be less than or equal to [`maxFramebufferHeight`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFramebufferHeight)
- [](#VUID-VkFramebufferCreateInfo-layers-00889)VUID-VkFramebufferCreateInfo-layers-00889  
  `layers` **must** be greater than `0`
- [](#VUID-VkFramebufferCreateInfo-layers-00890)VUID-VkFramebufferCreateInfo-layers-00890  
  `layers` **must** be less than or equal to [`maxFramebufferLayers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFramebufferLayers)
- [](#VUID-VkFramebufferCreateInfo-renderPass-02531)VUID-VkFramebufferCreateInfo-renderPass-02531  
  If `renderPass` was specified with non-zero view masks, `layers` **must** be `1`
- [](#VUID-VkFramebufferCreateInfo-pAttachments-00891)VUID-VkFramebufferCreateInfo-pAttachments-00891  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` that is a 2D or 2D array image view taken from a 3D image **must** not be a depth/stencil format
- [](#VUID-VkFramebufferCreateInfo-flags-03189)VUID-VkFramebufferCreateInfo-flags-03189  
  If the [`imagelessFramebuffer`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-imagelessFramebuffer) feature is not enabled, `flags` **must** not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`
- [](#VUID-VkFramebufferCreateInfo-flags-03190)VUID-VkFramebufferCreateInfo-flags-03190  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `pNext` chain **must** include a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure
- [](#VUID-VkFramebufferCreateInfo-flags-03191)VUID-VkFramebufferCreateInfo-flags-03191  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `attachmentImageInfoCount` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure in the `pNext` chain **must** be equal to either zero or `attachmentCount`
- [](#VUID-VkFramebufferCreateInfo-flags-04541)VUID-VkFramebufferCreateInfo-flags-04541  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `width` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure in the `pNext` chain that is used as an input, color, resolve or depth/stencil attachment in `renderPass` **must** be greater than or equal to `width`
- [](#VUID-VkFramebufferCreateInfo-flags-04542)VUID-VkFramebufferCreateInfo-flags-04542  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `height` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure in the `pNext` chain that is used as an input, color, resolve or depth/stencil attachment in `renderPass` **must** be greater than or equal to `height`
- [](#VUID-VkFramebufferCreateInfo-flags-03196)VUID-VkFramebufferCreateInfo-flags-03196  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `width` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure in the `pNext` chain that is referenced by [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)::`fragmentDensityMapAttachment` in `renderPass` **must** be greater than or equal to ⌈maxFragmentDensityTexelSizewidth​width​⌉
- [](#VUID-VkFramebufferCreateInfo-flags-03197)VUID-VkFramebufferCreateInfo-flags-03197  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `height` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure included in the `pNext` chain that is referenced by [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)::`fragmentDensityMapAttachment` in `renderPass` **must** be greater than or equal to ⌈maxFragmentDensityTexelSizeheight​height​⌉
- [](#VUID-VkFramebufferCreateInfo-flags-04543)VUID-VkFramebufferCreateInfo-flags-04543  
  If the [`maintenance7`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance7) feature is not enabled or the [`robustFragmentShadingRateAttachmentAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-robustFragmentShadingRateAttachmentAccess) limit is `VK_FALSE` or the `imageView` member of a [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html) structure was created with [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`baseMipLevel` greater than 0, and `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `width` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure in the `pNext` chain that is used as a [fragment shading rate attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment) **must** be greater than or equal to ⌈`width` / `texelWidth`⌉, where `texelWidth` is the largest value of `shadingRateAttachmentTexelSize.width` in a [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateAttachmentInfoKHR.html) which references that attachment
- [](#VUID-VkFramebufferCreateInfo-flags-04544)VUID-VkFramebufferCreateInfo-flags-04544  
  If the [`maintenance7`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance7) feature is not enabled or the [`robustFragmentShadingRateAttachmentAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-robustFragmentShadingRateAttachmentAccess) limit is `VK_FALSE` or the `imageView` member of a [VkRenderingFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFragmentShadingRateAttachmentInfoKHR.html) structure was created with [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html)::`baseMipLevel` greater than 0, and `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `height` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure in the `pNext` chain that is used as a [fragment shading rate attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment) **must** be greater than or equal to ⌈`height` / `texelHeight`⌉, where `texelHeight` is the largest value of `shadingRateAttachmentTexelSize.height` in a [VkFragmentShadingRateAttachmentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFragmentShadingRateAttachmentInfoKHR.html) which references that attachment
- [](#VUID-VkFramebufferCreateInfo-flags-04545)VUID-VkFramebufferCreateInfo-flags-04545  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `layerCount` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure in the `pNext` chain that is used as a [fragment shading rate attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment) **must** be either `1`, or greater than or equal to `layers`
- [](#VUID-VkFramebufferCreateInfo-flags-04587)VUID-VkFramebufferCreateInfo-flags-04587  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT` and `renderPass` was specified with non-zero view masks, the `layerCount` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure in the `pNext` chain that is used as a [fragment shading rate attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-fragment-shading-rate-attachment) **must** be either `1`, or greater than the index of the most significant bit set in any of those view masks
- [](#VUID-VkFramebufferCreateInfo-renderPass-03198)VUID-VkFramebufferCreateInfo-renderPass-03198  
  If multiview is enabled for `renderPass` and `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `layerCount` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure included in the `pNext` chain used as an input, color, resolve, or depth/stencil attachment in `renderPass` **must** be greater than the maximum bit index set in the view mask in the subpasses in which it is used in `renderPass`
- [](#VUID-VkFramebufferCreateInfo-renderPass-04546)VUID-VkFramebufferCreateInfo-renderPass-04546  
  If multiview is not enabled for `renderPass` and `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `layerCount` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure included in the `pNext` chain used as an input, color, resolve, or depth/stencil attachment in `renderPass` **must** be greater than or equal to `layers`
- [](#VUID-VkFramebufferCreateInfo-flags-03201)VUID-VkFramebufferCreateInfo-flags-03201  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `usage` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure included in the `pNext` chain that refers to an attachment used as a color attachment or resolve attachment by `renderPass` **must** include `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT`
- [](#VUID-VkFramebufferCreateInfo-flags-03202)VUID-VkFramebufferCreateInfo-flags-03202  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `usage` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure included in the `pNext` chain that refers to an attachment used as a depth/stencil attachment by `renderPass` **must** include `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkFramebufferCreateInfo-flags-03203)VUID-VkFramebufferCreateInfo-flags-03203  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `usage` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure included in the `pNext` chain that refers to an attachment used as a depth/stencil resolve attachment by `renderPass` **must** include `VK_IMAGE_USAGE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkFramebufferCreateInfo-flags-03204)VUID-VkFramebufferCreateInfo-flags-03204  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `usage` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure included in the `pNext` chain that refers to an attachment used as an input attachment by `renderPass` **must** include `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT`
- [](#VUID-VkFramebufferCreateInfo-flags-03205)VUID-VkFramebufferCreateInfo-flags-03205  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, at least one element of the `pViewFormats` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure included in the `pNext` chain **must** be equal to the corresponding value of [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription.html)::`format` used to create `renderPass`
- [](#VUID-VkFramebufferCreateInfo-flags-04113)VUID-VkFramebufferCreateInfo-flags-04113  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` **must** have been created with [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html)::`viewType` not equal to `VK_IMAGE_VIEW_TYPE_3D`
- [](#VUID-VkFramebufferCreateInfo-flags-04548)VUID-VkFramebufferCreateInfo-flags-04548  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of `pAttachments` that is used as a fragment shading rate attachment by `renderPass` **must** have been created with a `usage` value including `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkFramebufferCreateInfo-flags-04549)VUID-VkFramebufferCreateInfo-flags-04549  
  If `flags` includes `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `usage` member of any element of the `pAttachmentImageInfos` member of a [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html) structure included in the `pNext` chain that refers to an attachment used as a fragment shading rate attachment by `renderPass` **must** include `VK_IMAGE_USAGE_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkFramebufferCreateInfo-samples-06881)VUID-VkFramebufferCreateInfo-samples-06881  
  If [multisampled-render-to-single-sampled](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#subpass-multisampledrendertosinglesampled) is enabled for any subpass, all color, depth/stencil and input attachments used in that subpass which have `VkAttachmentDescription`::`samples` or `VkAttachmentDescription2`::`samples` equal to `VK_SAMPLE_COUNT_1_BIT` **must** have been created with `VK_IMAGE_CREATE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_BIT_EXT` in their [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags`
- [](#VUID-VkFramebufferCreateInfo-samples-07009)VUID-VkFramebufferCreateInfo-samples-07009  
  If [multisampled-render-to-single-sampled](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#subpass-multisampledrendertosinglesampled) is enabled for any subpass, all color, depth/stencil and input attachments used in that subpass which have `VkAttachmentDescription`::`samples` or `VkAttachmentDescription2`::`samples` equal to `VK_SAMPLE_COUNT_1_BIT` **must** have a format that supports the sample count specified in [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html)::`rasterizationSamples`
- [](#VUID-VkFramebufferCreateInfo-nullColorAttachmentWithExternalFormatResolve-09349)VUID-VkFramebufferCreateInfo-nullColorAttachmentWithExternalFormatResolve-09349  
  If the [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) is `VK_FALSE`, and `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the format of the color attachment for each subpass in `renderPass` that includes an external format image as a resolve attachment **must** have a format equal to the value of [VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html)::`colorAttachmentFormat` as returned by a call to [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html) for the Android hardware buffer that was used to create the image view use as its resolve attachment
- [](#VUID-VkFramebufferCreateInfo-pAttachments-09350)VUID-VkFramebufferCreateInfo-pAttachments-09350  
  If `flags` does not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, then if an element of `pAttachments` has a format of `VK_FORMAT_UNDEFINED`, it **must** have been created with a [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html)::`externalFormat` value identical to that provided in the [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html)::`externalFormat` specified by the corresponding [VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription2.html) in `renderPass`

Valid Usage (Implicit)

- [](#VUID-VkFramebufferCreateInfo-sType-sType)VUID-VkFramebufferCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_FRAMEBUFFER_CREATE_INFO`
- [](#VUID-VkFramebufferCreateInfo-pNext-pNext)VUID-VkFramebufferCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html)
- [](#VUID-VkFramebufferCreateInfo-sType-unique)VUID-VkFramebufferCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkFramebufferCreateInfo-flags-parameter)VUID-VkFramebufferCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkFramebufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateFlagBits.html) values
- [](#VUID-VkFramebufferCreateInfo-renderPass-parameter)VUID-VkFramebufferCreateInfo-renderPass-parameter  
  `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) handle
- [](#VUID-VkFramebufferCreateInfo-commonparent)VUID-VkFramebufferCreateInfo-commonparent  
  Both of `renderPass`, and the elements of `pAttachments` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkFramebufferCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateFlags.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateFramebuffer.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkFramebufferCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0