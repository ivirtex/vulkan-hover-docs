# VkRenderPassBeginInfo(3) Manual Page

## Name

VkRenderPassBeginInfo - Structure specifying render pass begin information



## [](#_c_specification)C Specification

The `VkRenderPassBeginInfo` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_VERSION_1_0
typedef struct VkRenderPassBeginInfo {
    VkStructureType        sType;
    const void*            pNext;
    VkRenderPass           renderPass;
    VkFramebuffer          framebuffer;
    VkRect2D               renderArea;
    uint32_t               clearValueCount;
    const VkClearValue*    pClearValues;
} VkRenderPassBeginInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `renderPass` is the render pass to begin an instance of.
- `framebuffer` is the framebuffer containing the attachments that are used with the render pass.
- `renderArea` is the render area that is affected by the render pass instance, and is described in more detail below.
- `clearValueCount` is the number of elements in `pClearValues`.
- `pClearValues` is a pointer to an array of `clearValueCount` [VkClearValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearValue.html) structures containing clear values for each attachment, if the attachment uses a `loadOp` value of `VK_ATTACHMENT_LOAD_OP_CLEAR` or if the attachment has a depth/stencil format and uses a `stencilLoadOp` value of `VK_ATTACHMENT_LOAD_OP_CLEAR`. The array is indexed by attachment number. Only elements corresponding to cleared attachments are used. Other elements of `pClearValues` are ignored.

## [](#_description)Description

`renderArea` is the render area that is affected by the render pass instance. The effects of attachment load, store and multisample resolve operations are restricted to the pixels whose x and y coordinates fall within the render area on all attachments. The render area extends to all layers of `framebuffer`. The application **must** ensure (using scissor if necessary) that all rendering is contained within the render area. The render area, after any transform specified by [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTransformBeginInfoQCOM.html)::`transform` is applied, **must** be contained within the framebuffer dimensions.

If [render pass transform](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-renderpass-transform) is enabled, then `renderArea` **must** equal the framebuffer pre-transformed dimensions. After `renderArea` has been transformed by [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTransformBeginInfoQCOM.html)::`transform`, the resulting render area **must** be equal to the framebuffer dimensions.

If multiview is enabled in `renderPass`, and [`multiviewPerViewRenderAreas`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiviewPerViewRenderAreas) feature is enabled, and there is an instance of [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html) included in the `pNext` chain with `perViewRenderAreaCount` not equal to `0`, then the elements of [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html)::`pPerViewRenderAreas` override `renderArea` and define a render area for each view. In this case, `renderArea` **must** be an area at least as large as the union of all the per-view render areas.

If the [`subpassShading`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-subpassShading) feature is enabled, then `renderArea` **must** equal the framebuffer dimensions.

Note

There **may** be a performance cost for using a render area smaller than the framebuffer, unless it matches the render area granularity for the render pass.

Valid Usage

- [](#VUID-VkRenderPassBeginInfo-clearValueCount-00902)VUID-VkRenderPassBeginInfo-clearValueCount-00902  
  `clearValueCount` **must** be greater than the largest attachment index in `renderPass` specifying a `loadOp` (or `stencilLoadOp`, if the attachment has a depth/stencil format) of `VK_ATTACHMENT_LOAD_OP_CLEAR`
- [](#VUID-VkRenderPassBeginInfo-clearValueCount-04962)VUID-VkRenderPassBeginInfo-clearValueCount-04962  
  If `clearValueCount` is not `0`, `pClearValues` **must** be a valid pointer to an array of `clearValueCount` [VkClearValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearValue.html) unions
- [](#VUID-VkRenderPassBeginInfo-renderPass-00904)VUID-VkRenderPassBeginInfo-renderPass-00904  
  `renderPass` **must** be [compatible](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-compatibility) with the `renderPass` member of the [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html) structure specified when creating `framebuffer`
- [](#VUID-VkRenderPassBeginInfo-None-08996)VUID-VkRenderPassBeginInfo-None-08996  
  If the `pNext` chain does not contain [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfo.html) or its `deviceRenderAreaCount` member is equal to 0, `renderArea.extent.width` **must** be greater than 0
- [](#VUID-VkRenderPassBeginInfo-None-08997)VUID-VkRenderPassBeginInfo-None-08997  
  If the `pNext` chain does not contain [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfo.html) or its `deviceRenderAreaCount` member is equal to 0, `renderArea.extent.height` **must** be greater than 0
- [](#VUID-VkRenderPassBeginInfo-pNext-02850)VUID-VkRenderPassBeginInfo-pNext-02850  
  If the `pNext` chain does not contain [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfo.html) or its `deviceRenderAreaCount` member is equal to 0, `renderArea.offset.x` **must** be greater than or equal to 0
- [](#VUID-VkRenderPassBeginInfo-pNext-02851)VUID-VkRenderPassBeginInfo-pNext-02851  
  If the `pNext` chain does not contain [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfo.html) or its `deviceRenderAreaCount` member is equal to 0, `renderArea.offset.y` **must** be greater than or equal to 0
- [](#VUID-VkRenderPassBeginInfo-pNext-02852)VUID-VkRenderPassBeginInfo-pNext-02852  
  If the `pNext` chain does not contain [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfo.html) or its `deviceRenderAreaCount` member is equal to 0, `renderArea.offset.x` + `renderArea.extent.width` **must** be less than or equal to [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`width` the `framebuffer` was created with
- [](#VUID-VkRenderPassBeginInfo-pNext-02853)VUID-VkRenderPassBeginInfo-pNext-02853  
  If the `pNext` chain does not contain [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfo.html) or its `deviceRenderAreaCount` member is equal to 0, `renderArea.offset.y` + `renderArea.extent.height` **must** be less than or equal to [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`height` the `framebuffer` was created with
- [](#VUID-VkRenderPassBeginInfo-pNext-02856)VUID-VkRenderPassBeginInfo-pNext-02856  
  If the `pNext` chain contains [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfo.html), `offset.x` + `extent.width` of each element of `pDeviceRenderAreas` **must** be less than or equal to [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`width` the `framebuffer` was created with
- [](#VUID-VkRenderPassBeginInfo-pNext-02857)VUID-VkRenderPassBeginInfo-pNext-02857  
  If the `pNext` chain contains [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfo.html), `offset.y` + `extent.height` of each element of `pDeviceRenderAreas` **must** be less than or equal to [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`height` the `framebuffer` was created with
- [](#VUID-VkRenderPassBeginInfo-framebuffer-03207)VUID-VkRenderPassBeginInfo-framebuffer-03207  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that did not include `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, and the `pNext` chain includes a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure, its `attachmentCount` **must** be zero
- [](#VUID-VkRenderPassBeginInfo-framebuffer-03208)VUID-VkRenderPassBeginInfo-framebuffer-03208  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, the `attachmentCount` of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** be equal to the value of [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html)::`attachmentImageInfoCount` used to create `framebuffer`
- [](#VUID-VkRenderPassBeginInfo-framebuffer-02780)VUID-VkRenderPassBeginInfo-framebuffer-02780  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of the `pAttachments` member of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** have been created on the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) as `framebuffer` and `renderPass`
- [](#VUID-VkRenderPassBeginInfo-framebuffer-03209)VUID-VkRenderPassBeginInfo-framebuffer-03209  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of the `pAttachments` member of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) of an image created with a value of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags` equal to the `flags` member of the corresponding element of [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos` used to create `framebuffer`
- [](#VUID-VkRenderPassBeginInfo-framebuffer-04627)VUID-VkRenderPassBeginInfo-framebuffer-04627  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of the `pAttachments` member of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) with [an inherited usage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-inherited-usage) equal to the `usage` member of the corresponding element of [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos` used to create `framebuffer`
- [](#VUID-VkRenderPassBeginInfo-framebuffer-03211)VUID-VkRenderPassBeginInfo-framebuffer-03211  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of the `pAttachments` member of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) with a width equal to the `width` member of the corresponding element of [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos` used to create `framebuffer`
- [](#VUID-VkRenderPassBeginInfo-framebuffer-03212)VUID-VkRenderPassBeginInfo-framebuffer-03212  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of the `pAttachments` member of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) with a height equal to the `height` member of the corresponding element of [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos` used to create `framebuffer`
- [](#VUID-VkRenderPassBeginInfo-framebuffer-03213)VUID-VkRenderPassBeginInfo-framebuffer-03213  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of the `pAttachments` member of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) of an image created with a value of [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html)::`subresourceRange.layerCount` equal to the `layerCount` member of the corresponding element of [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos` used to create `framebuffer`
- [](#VUID-VkRenderPassBeginInfo-framebuffer-03214)VUID-VkRenderPassBeginInfo-framebuffer-03214  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of the `pAttachments` member of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) of an image created with a value of [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html)::`viewFormatCount` equal to the `viewFormatCount` member of the corresponding element of [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos` used to create `framebuffer`
- [](#VUID-VkRenderPassBeginInfo-framebuffer-03215)VUID-VkRenderPassBeginInfo-framebuffer-03215  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of the `pAttachments` member of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) of an image created with a set of elements in [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html)::`pViewFormats` equal to the set of elements in the `pViewFormats` member of the corresponding element of [VkFramebufferAttachmentsCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfo.html)::`pAttachmentImageInfos` used to create `framebuffer`
- [](#VUID-VkRenderPassBeginInfo-framebuffer-03216)VUID-VkRenderPassBeginInfo-framebuffer-03216  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of the `pAttachments` member of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) of an image created with a value of [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html)::`format` equal to the corresponding value of [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription.html)::`format` in `renderPass`
- [](#VUID-VkRenderPassBeginInfo-framebuffer-09353)VUID-VkRenderPassBeginInfo-framebuffer-09353  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, and the [`nullColorAttachmentWithExternalFormatResolve`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-nullColorAttachmentWithExternalFormatResolve) is `VK_FALSE`, the format of the color attachment for each subpass that includes an external format image as a resolve attachment **must** have a format equal to the value of [VkAndroidHardwareBufferFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatResolvePropertiesANDROID.html)::`colorAttachmentFormat` as returned by a call to [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html) for the Android hardware buffer that was used to create the image view use as its resolve attachment
- [](#VUID-VkRenderPassBeginInfo-framebuffer-09354)VUID-VkRenderPassBeginInfo-framebuffer-09354  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of the `pAttachments` member of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) of an image created with a value of [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html)::`externalFormat` equal to [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html)::`externalFormat` in the `pNext` chain of the corresponding [VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription2.html) structure used to create `renderPass`
- [](#VUID-VkRenderPassBeginInfo-framebuffer-09047)VUID-VkRenderPassBeginInfo-framebuffer-09047  
  If `framebuffer` was created with a [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html)::`flags` value that included `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT`, each element of the `pAttachments` member of a [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html) structure included in the `pNext` chain **must** be a [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html) of an image created with a value of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`samples` equal to the corresponding value of [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription.html)::`samples` in `renderPass` , or `VK_SAMPLE_COUNT_1_BIT` if `renderPass` was created with [VkMultisampledRenderToSingleSampledInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultisampledRenderToSingleSampledInfoEXT.html) structure in the `pNext` chain with `multisampledRenderToSingleSampledEnable` equal to `VK_TRUE`
- [](#VUID-VkRenderPassBeginInfo-pNext-02869)VUID-VkRenderPassBeginInfo-pNext-02869  
  If the `pNext` chain includes [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTransformBeginInfoQCOM.html), `renderArea.offset` **must** equal (0,0)
- [](#VUID-VkRenderPassBeginInfo-pNext-02870)VUID-VkRenderPassBeginInfo-pNext-02870  
  If the `pNext` chain includes [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTransformBeginInfoQCOM.html), `renderArea.extent` transformed by [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTransformBeginInfoQCOM.html)::`transform` **must** equal the `framebuffer` dimensions
- [](#VUID-VkRenderPassBeginInfo-perViewRenderAreaCount-07859)VUID-VkRenderPassBeginInfo-perViewRenderAreaCount-07859  
  If the `perViewRenderAreaCount` member of a [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html) structure included in the `pNext` chain is not `0`, then the [`multiviewPerViewRenderAreas`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiviewPerViewRenderAreas) feature **must** be enabled
- [](#VUID-VkRenderPassBeginInfo-perViewRenderAreaCount-07860)VUID-VkRenderPassBeginInfo-perViewRenderAreaCount-07860  
  If the `perViewRenderAreaCount` member of a [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html) structure included in the `pNext` chain is not `0`, then `renderArea` **must** specify a render area that includes the union of all per view render areas
- [](#VUID-VkRenderPassBeginInfo-pNext-09539)VUID-VkRenderPassBeginInfo-pNext-09539  
  If the `pNext` chain contains a [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeBeginInfoARM.html) structure, the union of stripe areas defined by the elements of [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeBeginInfoARM.html)::`pStripeInfos` **must** cover the `renderArea`

Valid Usage (Implicit)

- [](#VUID-VkRenderPassBeginInfo-sType-sType)VUID-VkRenderPassBeginInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_BEGIN_INFO`
- [](#VUID-VkRenderPassBeginInfo-pNext-pNext)VUID-VkRenderPassBeginInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkDeviceGroupRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceGroupRenderPassBeginInfo.html), [VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultiviewPerViewRenderAreasRenderPassBeginInfoQCOM.html), [VkRenderPassAttachmentBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfo.html), [VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html), [VkRenderPassStripeBeginInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeBeginInfoARM.html), or [VkRenderPassTransformBeginInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTransformBeginInfoQCOM.html)
- [](#VUID-VkRenderPassBeginInfo-sType-unique)VUID-VkRenderPassBeginInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkRenderPassBeginInfo-renderPass-parameter)VUID-VkRenderPassBeginInfo-renderPass-parameter  
  `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) handle
- [](#VUID-VkRenderPassBeginInfo-framebuffer-parameter)VUID-VkRenderPassBeginInfo-framebuffer-parameter  
  `framebuffer` **must** be a valid [VkFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebuffer.html) handle
- [](#VUID-VkRenderPassBeginInfo-commonparent)VUID-VkRenderPassBeginInfo-commonparent  
  Both of `framebuffer`, and `renderPass` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkClearValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearValue.html), [VkFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebuffer.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdBeginRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderPass.html), [vkCmdBeginRenderPass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderPass2.html), [vkCmdBeginRenderPass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRenderPass2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassBeginInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0